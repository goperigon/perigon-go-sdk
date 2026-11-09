// Command preprocess downloads the public Perigon OpenAPI spec and writes
// .openapi-spec/spec.local.json for oapi-codegen.
//
// Transformations applied before oapi-codegen runs:
//
//  1. Operation tags: an operation with no tags, or exactly one tag, is left
//     unchanged. When tags lists v1 and at least one other tag, tags becomes
//     ["v1"]. Several tags and no v1 stops generation.
//  2. Hidden properties from scripts/spec-filters.json are removed from the
//     schema and from its required list.
//  3. Excluded paths: every paths key that matches a listed glob is deleted.
//     A pattern with no wildcards is an exact match. * covers one segment and
//     ** covers nested paths. Brace expansion is disabled, so {uuid} stays literal.
//  4. Article/NewsCluster cycle: NewsClusterArticle is Article without cluster,
//     and ArticleNewsCluster is NewsCluster without selectedArticles. The nested
//     refs point at those clones.
//  5. Orphan schemas: a fixed-point pass deletes any components.schemas entry
//     with no remaining $ref from outside that schema.
//  6. Inline enums whose generated Go name matches a component schema get
//     x-go-type-name, so oapi-codegen can emit both types.
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

const (
	specURL    = "https://api.perigon.io/v1/openapi/public-sdk"
	userAgent  = "perigon-go-sdk-generator/1.0"
	v1Tag      = "v1"
	filtersRel = "scripts/spec-filters.json"
	outputRel  = ".openapi-spec/spec.local.json"
)

var httpMethods = map[string]struct{}{
	"get": {}, "put": {}, "post": {}, "delete": {},
	"options": {}, "head": {}, "patch": {}, "trace": {},
}

const (
	article                = "Article"
	newsCluster            = "NewsCluster"
	newsClusterArticle     = "NewsClusterArticle"
	articleNewsCluster     = "ArticleNewsCluster"
	newsClusterArticleDesc = "Article embedded in a story cluster, without the parent cluster."
	articleNewsClusterDesc = "Story cluster embedded on an article, without selected articles."
)

type specFilters struct {
	Description   string              `json:"$description"`
	HiddenFields  map[string][]string `json:"hiddenFields"`
	ExcludedPaths []string            `json:"excludedPaths"`
}

func main() {
	root, err := moduleRoot()
	if err != nil {
		fail(err)
	}
	if err := run(root, specURL, os.Stdout, os.Stderr); err != nil {
		fail(err)
	}
}

func fail(err error) {
	fmt.Fprintf(os.Stderr, "preprocess: %v\n", err)
	os.Exit(1)
}

func moduleRoot() (string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	dir := wd
	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir, nil
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			return "", fmt.Errorf("go.mod not found from %s", wd)
		}
		dir = parent
	}
}

func run(root, url string, stdout, stderr io.Writer) error {
	config, err := loadFilters(filepath.Join(root, filtersRel))
	if err != nil {
		return err
	}
	var hiddenEntries []string
	for schemaName, fields := range config.HiddenFields {
		for _, field := range fields {
			hiddenEntries = append(hiddenEntries, schemaName+"."+field)
		}
	}
	fmt.Fprintln(stdout, "Loaded spec filters from scripts/spec-filters.json")
	hiddenSummary := fmt.Sprintf("  %d hidden field(s)", len(hiddenEntries))
	if len(hiddenEntries) > 0 {
		hiddenSummary += ": " + strings.Join(hiddenEntries, ", ")
	}
	fmt.Fprintln(stdout, hiddenSummary)
	pathSummary := fmt.Sprintf("  %d excluded path pattern(s)", len(config.ExcludedPaths))
	if len(config.ExcludedPaths) > 0 {
		pathSummary += ": " + strings.Join(config.ExcludedPaths, ", ")
	}
	fmt.Fprintln(stdout, pathSummary)

	fmt.Fprintf(stdout, "Fetching OpenAPI spec from %s\n", url)
	spec, err := fetchSpec(url)
	if err != nil {
		return err
	}
	spec, err = preferV1Tag(spec, stdout)
	if err != nil {
		return err
	}
	spec, applied, missing := hideProperties(spec, config.HiddenFields, stdout)
	if len(missing) > 0 {
		warn(stderr, fmt.Sprintf("Skipped %d hidden-field entry/entries (not present in spec):", len(missing)), missing)
	}
	spec, excluded, unused := excludePaths(spec, config.ExcludedPaths, stdout)
	if len(unused) > 0 {
		warn(stderr, fmt.Sprintf("Skipped %d excluded-path pattern(s) (no matching paths):", len(unused)), unused)
	}
	spec, err = splitArticleClusterCycle(spec, stdout)
	if err != nil {
		return err
	}
	spec, pruned := pruneOrphanSchemas(spec, stdout)
	spec = disambiguateInlineTypeNames(spec, stdout)

	outputPath := filepath.Join(root, outputRel)
	if err := os.MkdirAll(filepath.Dir(outputPath), 0o755); err != nil {
		return err
	}
	payload, err := marshalSpec(spec)
	if err != nil {
		return err
	}
	if err := os.WriteFile(outputPath, payload, 0o644); err != nil {
		return err
	}
	fmt.Fprintf(stdout, "Wrote preprocessed spec to %s (%d field(s) hidden, %d path(s) excluded, %d orphan schema(s) pruned)\n", outputPath, applied, len(excluded), len(pruned))
	return nil
}

func loadFilters(path string) (specFilters, error) {
	raw, err := os.ReadFile(path)
	if err != nil {
		return specFilters{}, err
	}
	var config specFilters
	if err := json.Unmarshal(raw, &config); err != nil {
		return specFilters{}, fmt.Errorf("parse %s: %w", path, err)
	}
	if config.HiddenFields == nil {
		config.HiddenFields = map[string][]string{}
	}
	return config, nil
}

func fetchSpec(url string) (map[string]any, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", userAgent)
	req.Header.Set("Accept", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("fetch spec: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		body, _ := io.ReadAll(io.LimitReader(resp.Body, 512))
		return nil, fmt.Errorf("fetch spec: HTTP %s: %s", resp.Status, strings.TrimSpace(string(body)))
	}
	dec := json.NewDecoder(resp.Body)
	dec.UseNumber()
	var payload any
	if err := dec.Decode(&payload); err != nil {
		return nil, fmt.Errorf("fetch spec: %w", err)
	}
	spec, ok := payload.(map[string]any)
	if !ok {
		return nil, fmt.Errorf("fetch spec: response was not a JSON object")
	}
	return spec, nil
}

func marshalSpec(spec map[string]any) ([]byte, error) {
	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	enc.SetEscapeHTML(false)
	enc.SetIndent("", "  ")
	if err := enc.Encode(spec); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func warn(stderr io.Writer, title string, items []string) {
	fmt.Fprintln(stderr, title)
	for _, item := range items {
		fmt.Fprintf(stderr, "    - %s\n", item)
	}
}

func schemasOf(spec map[string]any) map[string]any {
	components, _ := spec["components"].(map[string]any)
	if components == nil {
		return nil
	}
	schemas, _ := components["schemas"].(map[string]any)
	return schemas
}

func withSchemas(spec map[string]any, schemas map[string]any) (map[string]any, error) {
	components, _ := spec["components"].(map[string]any)
	if components == nil {
		return nil, fmt.Errorf("OpenAPI spec is missing components")
	}
	nextComponents := cloneMap(components)
	nextComponents["schemas"] = schemas
	next := cloneMap(spec)
	next["components"] = nextComponents
	return next, nil
}

func cloneMap(in map[string]any) map[string]any {
	out := make(map[string]any, len(in))
	for key, value := range in {
		out[key] = value
	}
	return out
}

func deepClone(value any) any {
	switch typed := value.(type) {
	case map[string]any:
		out := make(map[string]any, len(typed))
		for key, item := range typed {
			out[key] = deepClone(item)
		}
		return out
	case []any:
		out := make([]any, len(typed))
		for i, item := range typed {
			out[i] = deepClone(item)
		}
		return out
	default:
		return value
	}
}

func without(in map[string]any, keys ...string) map[string]any {
	drop := make(map[string]struct{}, len(keys))
	for _, key := range keys {
		drop[key] = struct{}{}
	}
	out := make(map[string]any, len(in))
	for key, value := range in {
		if _, ok := drop[key]; ok {
			continue
		}
		out[key] = value
	}
	return out
}

func schemaWithoutProperty(schema map[string]any, field string) (map[string]any, bool) {
	properties, _ := schema["properties"].(map[string]any)
	if properties == nil {
		return nil, false
	}
	if _, ok := properties[field]; !ok {
		return nil, false
	}
	updated := cloneMap(schema)
	updated["properties"] = without(properties, field)
	required, ok := schema["required"].([]any)
	if !ok {
		return updated, true
	}
	var kept []any
	for _, item := range required {
		if item == field {
			continue
		}
		kept = append(kept, item)
	}
	if len(kept) > 0 {
		updated["required"] = kept
		return updated, true
	}
	return without(updated, "required"), true
}

func hideProperties(spec map[string]any, hidden map[string][]string, stdout io.Writer) (map[string]any, int, []string) {
	schemas := schemasOf(spec)
	newSchemas := cloneMap(schemas)
	applied := 0
	var missing []string
	for schemaName, fields := range hidden {
		schema, _ := newSchemas[schemaName].(map[string]any)
		if schema == nil {
			missing = append(missing, schemaName+" (schema not found)")
			continue
		}
		for _, field := range fields {
			updated, ok := schemaWithoutProperty(schema, field)
			if !ok {
				missing = append(missing, schemaName+"."+field+" (property not found)")
				continue
			}
			schema = updated
			applied++
			fmt.Fprintf(stdout, "  Hid %s.%s\n", schemaName, field)
		}
		newSchemas[schemaName] = schema
	}
	if applied == 0 {
		return spec, applied, missing
	}
	next, err := withSchemas(spec, newSchemas)
	if err != nil {
		return spec, applied, append(missing, err.Error())
	}
	return next, applied, missing
}

func pathMatches(path, pattern string) bool {
	return matchGlob(path, pattern)
}

func matchGlob(path, pattern string) bool {
	for {
		if strings.HasPrefix(pattern, "**") {
			rest := strings.TrimPrefix(pattern, "**")
			if strings.HasPrefix(rest, "/") {
				slashRest := rest[1:]
				if matchGlob(path, slashRest) {
					return true
				}
				for i := 0; i < len(path); i++ {
					if path[i] != '/' {
						continue
					}
					if matchGlob(path[i+1:], slashRest) {
						return true
					}
				}
				return false
			}
			if rest == "" {
				return true
			}
			for i := 0; i <= len(path); i++ {
				if matchGlob(path[i:], rest) {
					return true
				}
			}
			return false
		}
		if pattern == "" {
			return path == ""
		}
		if pattern[0] == '*' {
			rest := pattern[1:]
			if matchGlob(path, rest) {
				return true
			}
			for i := 0; i < len(path) && path[i] != '/'; i++ {
				if matchGlob(path[i+1:], rest) {
					return true
				}
			}
			return false
		}
		if path == "" {
			return false
		}
		if pattern[0] == '?' {
			if path[0] == '/' {
				return false
			}
			path = path[1:]
			pattern = pattern[1:]
			continue
		}
		if path[0] != pattern[0] {
			return false
		}
		path = path[1:]
		pattern = pattern[1:]
	}
}

func excludePaths(spec map[string]any, patterns []string, stdout io.Writer) (map[string]any, []string, []string) {
	paths, _ := spec["paths"].(map[string]any)
	if paths == nil {
		return spec, nil, append([]string(nil), patterns...)
	}
	unused := map[string]struct{}{}
	for _, pattern := range patterns {
		unused[pattern] = struct{}{}
	}
	var excluded []string
	kept := map[string]any{}
	for path, pathItem := range paths {
		matched := false
		for _, pattern := range patterns {
			if !pathMatches(path, pattern) {
				continue
			}
			delete(unused, pattern)
			matched = true
		}
		if !matched {
			kept[path] = pathItem
			continue
		}
		excluded = append(excluded, path)
		fmt.Fprintf(stdout, "  Excluded path %s\n", path)
	}
	if len(excluded) == 0 {
		return spec, excluded, unusedPatterns(patterns, unused)
	}
	next := cloneMap(spec)
	next["paths"] = kept
	return next, excluded, unusedPatterns(patterns, unused)
}

func unusedPatterns(patterns []string, unused map[string]struct{}) []string {
	var out []string
	for _, pattern := range patterns {
		if _, ok := unused[pattern]; ok {
			out = append(out, pattern)
		}
	}
	return out
}

func schemaRef(name string) string {
	return "#/components/schemas/" + name
}

func requireSchema(schemas map[string]any, name string) (map[string]any, error) {
	schema, _ := schemas[name].(map[string]any)
	if schema == nil {
		return nil, fmt.Errorf("cannot split Article/NewsCluster cycle: schema %s not found", name)
	}
	return schema, nil
}

func requireProperties(schema map[string]any, schemaName string) (map[string]any, error) {
	properties, _ := schema["properties"].(map[string]any)
	if properties == nil {
		return nil, fmt.Errorf("cannot split Article/NewsCluster cycle: %s.properties is missing", schemaName)
	}
	return properties, nil
}

func dropProperty(schema map[string]any, schemaName, field string) (map[string]any, error) {
	cloned, _ := deepClone(schema).(map[string]any)
	if cloned == nil {
		return nil, fmt.Errorf("cannot split Article/NewsCluster cycle: %s is not an object", schemaName)
	}
	schema = cloned
	properties, err := requireProperties(schema, schemaName)
	if err != nil {
		return nil, err
	}
	if _, ok := properties[field]; !ok {
		return nil, fmt.Errorf("cannot split Article/NewsCluster cycle: %s.%s not found", schemaName, field)
	}
	updated := cloneMap(schema)
	updated["properties"] = without(properties, field)
	required, ok := schema["required"].([]any)
	if !ok {
		return updated, nil
	}
	var kept []any
	for _, item := range required {
		if item == field {
			continue
		}
		kept = append(kept, item)
	}
	if len(kept) > 0 {
		updated["required"] = kept
		return updated, nil
	}
	return without(updated, "required"), nil
}

func retargetAllOfRef(propertySchema map[string]any, expected, replacement string) (map[string]any, error) {
	allOf, _ := propertySchema["allOf"].([]any)
	if allOf == nil {
		return nil, fmt.Errorf("cannot split Article/NewsCluster cycle: Article.cluster.allOf is missing")
	}
	updatedEntries := make([]any, 0, len(allOf))
	found := false
	for _, entry := range allOf {
		entryMap, _ := entry.(map[string]any)
		if entryMap == nil || entryMap["$ref"] != expected {
			updatedEntries = append(updatedEntries, entry)
			continue
		}
		cloned := cloneMap(entryMap)
		cloned["$ref"] = replacement
		updatedEntries = append(updatedEntries, cloned)
		found = true
	}
	if !found {
		return nil, fmt.Errorf("cannot split Article/NewsCluster cycle: Article.cluster.allOf does not reference NewsCluster")
	}
	updated := cloneMap(propertySchema)
	updated["allOf"] = updatedEntries
	return updated, nil
}

func retargetItemsRef(propertySchema map[string]any, expected, replacement string) (map[string]any, error) {
	items, _ := propertySchema["items"].(map[string]any)
	if items == nil || items["$ref"] != expected {
		return nil, fmt.Errorf("cannot split Article/NewsCluster cycle: NewsCluster.selectedArticles.items does not reference Article")
	}
	clonedItems := cloneMap(items)
	clonedItems["$ref"] = replacement
	updated := cloneMap(propertySchema)
	updated["items"] = clonedItems
	return updated, nil
}

func splitArticleClusterCycle(spec map[string]any, stdout io.Writer) (map[string]any, error) {
	schemas := schemasOf(spec)
	if len(schemas) == 0 {
		return nil, fmt.Errorf("cannot split Article/NewsCluster cycle: components.schemas is missing")
	}
	articleSchema, err := requireSchema(schemas, article)
	if err != nil {
		return nil, err
	}
	clusterSchema, err := requireSchema(schemas, newsCluster)
	if err != nil {
		return nil, err
	}
	articleProperties, err := requireProperties(articleSchema, article)
	if err != nil {
		return nil, err
	}
	clusterProperties, err := requireProperties(clusterSchema, newsCluster)
	if err != nil {
		return nil, err
	}
	clusterProperty, _ := articleProperties["cluster"].(map[string]any)
	selectedArticles, _ := clusterProperties["selectedArticles"].(map[string]any)
	if clusterProperty == nil {
		return nil, fmt.Errorf("cannot split Article/NewsCluster cycle: Article.cluster is missing")
	}
	if selectedArticles == nil {
		return nil, fmt.Errorf("cannot split Article/NewsCluster cycle: NewsCluster.selectedArticles is missing")
	}
	droppedArticle, err := dropProperty(articleSchema, article, "cluster")
	if err != nil {
		return nil, err
	}
	droppedCluster, err := dropProperty(clusterSchema, newsCluster, "selectedArticles")
	if err != nil {
		return nil, err
	}
	droppedArticle["description"] = newsClusterArticleDesc
	droppedCluster["description"] = articleNewsClusterDesc

	retargetedCluster, err := retargetAllOfRef(clusterProperty, schemaRef(newsCluster), schemaRef(articleNewsCluster))
	if err != nil {
		return nil, err
	}
	retargetedArticles, err := retargetItemsRef(selectedArticles, schemaRef(article), schemaRef(newsClusterArticle))
	if err != nil {
		return nil, err
	}
	newArticleProperties := cloneMap(articleProperties)
	newArticleProperties["cluster"] = retargetedCluster
	newArticle := cloneMap(articleSchema)
	newArticle["properties"] = newArticleProperties

	newClusterProperties := cloneMap(clusterProperties)
	newClusterProperties["selectedArticles"] = retargetedArticles
	newCluster := cloneMap(clusterSchema)
	newCluster["properties"] = newClusterProperties

	fmt.Fprintf(stdout, "  Split Article/NewsCluster cycle into %s and %s\n", newsClusterArticle, articleNewsCluster)
	nextSchemas := cloneMap(schemas)
	nextSchemas[article] = newArticle
	nextSchemas[newsCluster] = newCluster
	nextSchemas[newsClusterArticle] = droppedArticle
	nextSchemas[articleNewsCluster] = droppedCluster
	return withSchemas(spec, nextSchemas)
}

func collectRefs(node any, acc map[string]struct{}) {
	switch typed := node.(type) {
	case map[string]any:
		for key, value := range typed {
			if key == "$ref" {
				if ref, ok := value.(string); ok {
					acc[ref] = struct{}{}
				}
				continue
			}
			collectRefs(value, acc)
		}
	case []any:
		for _, item := range typed {
			collectRefs(item, acc)
		}
	}
}

func pruneOrphanSchemas(spec map[string]any, stdout io.Writer) (map[string]any, []string) {
	components, _ := spec["components"].(map[string]any)
	if components == nil {
		return spec, nil
	}
	original, _ := components["schemas"].(map[string]any)
	if len(original) == 0 {
		return spec, nil
	}
	schemas := cloneMap(original)
	var pruned []string
	for {
		view := cloneMap(spec)
		viewComponents := cloneMap(components)
		viewComponents["schemas"] = schemas
		view["components"] = viewComponents
		allRefs := map[string]struct{}{}
		collectRefs(view, allRefs)

		removedThisPass := 0
		for name := range schemas {
			target := schemaRef(name)
			if _, referenced := allRefs[target]; referenced {
				saved := schemas[name]
				delete(schemas, name)
				refsWithoutSelf := map[string]struct{}{}
				withoutView := cloneMap(spec)
				withoutComponents := cloneMap(components)
				withoutComponents["schemas"] = schemas
				withoutView["components"] = withoutComponents
				collectRefs(withoutView, refsWithoutSelf)
				if _, still := refsWithoutSelf[target]; !still {
					pruned = append(pruned, name)
					fmt.Fprintf(stdout, "  Pruned orphan schema %s\n", name)
					removedThisPass++
					continue
				}
				schemas[name] = saved
				continue
			}
			delete(schemas, name)
			pruned = append(pruned, name)
			fmt.Fprintf(stdout, "  Pruned orphan schema %s\n", name)
			removedThisPass++
		}
		if removedThisPass == 0 {
			break
		}
	}
	if len(pruned) == 0 {
		return spec, pruned
	}
	next, err := withSchemas(spec, schemas)
	if err != nil {
		return spec, pruned
	}
	return next, pruned
}

func preferV1Tag(spec map[string]any, stdout io.Writer) (map[string]any, error) {
	paths, _ := spec["paths"].(map[string]any)
	if paths == nil {
		return spec, nil
	}
	newPaths := make(map[string]any, len(paths))
	changed := false
	for path, pathItem := range paths {
		item, _ := pathItem.(map[string]any)
		if item == nil {
			newPaths[path] = pathItem
			continue
		}
		var newItem map[string]any
		for method, operationValue := range item {
			if _, ok := httpMethods[strings.ToLower(method)]; !ok {
				continue
			}
			operation, _ := operationValue.(map[string]any)
			if operation == nil {
				continue
			}
			tags, _ := operation["tags"].([]any)
			if len(tags) <= 1 {
				continue
			}
			label := strings.ToUpper(method) + " " + path
			if operationID, ok := operation["operationId"].(string); ok && operationID != "" {
				label = label + " (" + operationID + ")"
			}
			if !tagListHas(tags, v1Tag) {
				return nil, fmt.Errorf("operation has multiple tags and none of them is 'v1': %s tags=%v", label, tags)
			}
			if newItem == nil {
				newItem = cloneMap(item)
			}
			updated := cloneMap(operation)
			updated["tags"] = []any{v1Tag}
			newItem[method] = updated
			changed = true
			fmt.Fprintf(stdout, "  Kept v1 tag for %s (was %v)\n", label, tags)
		}
		if newItem == nil {
			newPaths[path] = pathItem
			continue
		}
		newPaths[path] = newItem
	}
	if !changed {
		return spec, nil
	}
	next := cloneMap(spec)
	next["paths"] = newPaths
	return next, nil
}

func tagListHas(tags []any, want string) bool {
	for _, tag := range tags {
		if tag == want {
			return true
		}
	}
	return false
}

// disambiguateInlineTypeNames sets x-go-type-name when an inline enum would
// be generated with the same Go name as a component schema. oapi-codegen names
// that enum from the schema path (FilledField.source becomes FilledFieldSource)
// and stops when the name is already taken.
func disambiguateInlineTypeNames(spec map[string]any, stdout io.Writer) map[string]any {
	schemas := schemasOf(spec)
	taken := map[string]struct{}{}
	for name := range schemas {
		taken[toCamel(name)] = struct{}{}
	}
	for _, name := range sortedKeys(schemas) {
		schema, ok := schemas[name].(map[string]any)
		if !ok {
			continue
		}
		walkInlineTypes(schema, []string{name}, taken, stdout)
	}
	return spec
}

func sortedKeys(values map[string]any) []string {
	keys := make([]string, 0, len(values))
	for key := range values {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return keys
}

func walkInlineTypes(node any, path []string, taken map[string]struct{}, stdout io.Writer) {
	switch typed := node.(type) {
	case []any:
		for i, item := range typed {
			walkInlineTypes(item, append(append([]string{}, path...), strconv.Itoa(i)), taken, stdout)
		}
	case map[string]any:
		if _, isRef := typed["$ref"]; isRef && len(typed) == 1 {
			return
		}
		if enum, ok := typed["enum"].([]any); ok && len(enum) > 0 && typed["type"] != "array" && len(path) > 1 {
			goName := inlineEnumGoName(path)
			if _, exists := taken[goName]; exists {
				alt := uniqueGoName(goName+"Kind", taken)
				typed["x-go-type-name"] = alt
				taken[alt] = struct{}{}
				fmt.Fprintf(stdout, "  Renamed inline type %s to %s\n", strings.Join(path, "."), alt)
			} else {
				taken[goName] = struct{}{}
			}
		}
		if props, ok := typed["properties"].(map[string]any); ok {
			for _, key := range sortedKeys(props) {
				walkInlineTypes(props[key], append(append([]string{}, path...), key), taken, stdout)
			}
		}
		if items, ok := typed["items"]; ok {
			walkInlineTypes(items, path, taken, stdout)
		}
		for _, key := range []string{"allOf", "oneOf", "anyOf"} {
			items, ok := typed[key].([]any)
			if !ok {
				continue
			}
			for i, item := range items {
				walkInlineTypes(item, append(append([]string{}, path...), strconv.Itoa(i)), taken, stdout)
			}
		}
	}
}

func inlineEnumGoName(path []string) string {
	parts := make([]string, len(path))
	for i, part := range path {
		parts[i] = toCamel(part)
	}
	return toCamel(strings.Join(parts, "_"))
}

func uniqueGoName(base string, taken map[string]struct{}) string {
	name := base
	for i := 2; ; i++ {
		if _, exists := taken[name]; !exists {
			return name
		}
		name = fmt.Sprintf("%s%d", base, i)
	}
}

// toCamel matches oapi-codegen v2.8.2 ToCamelCase. It trims ASCII spaces,
// copies uppercase letters and digits, capitalizes a lowercase letter after a
// separator, and drops separators. Separators are - # @ ! $ & = . + : ; _ ~ space ( ) { } [ ].
func toCamel(value string) string {
	var b strings.Builder
	capNext := true
	for _, r := range strings.Trim(value, " ") {
		switch {
		case unicode.IsUpper(r) || unicode.IsDigit(r):
			b.WriteRune(r)
		case unicode.IsLower(r):
			if capNext {
				b.WriteRune(unicode.ToUpper(r))
			} else {
				b.WriteRune(r)
			}
		}
		_, capNext = camelSeparators[r]
	}
	return b.String()
}

var camelSeparators = map[rune]struct{}{
	'-': {}, '#': {}, '@': {}, '!': {}, '$': {}, '&': {}, '=': {},
	'.': {}, '+': {}, ':': {}, ';': {}, '_': {}, '~': {},
	' ': {}, '(': {}, ')': {}, '{': {}, '}': {}, '[': {}, ']': {},
}
