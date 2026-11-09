package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestPathMatches(t *testing.T) {
	cases := []struct {
		path    string
		pattern string
		want    bool
	}{
		{path: "/v1/articles/refresh", pattern: "/v1/articles/refresh", want: true},
		{path: "/v1/articles/refresh/extra", pattern: "/v1/articles/refresh", want: false},
		{path: "/v1/articles/refresh", pattern: "/v1/articles/refresh/**", want: false},
		{path: "/v1/articles/refresh/foo", pattern: "/v1/articles/refresh/**", want: true},
		{path: "/v1/articles/refresh/foo/bar", pattern: "/v1/articles/refresh/**", want: true},
		{path: "/v1/articles/other", pattern: "/v1/articles/refresh/**", want: false},
		{path: "/v1/articles/refresh", pattern: "/v1/*/refresh", want: true},
		{path: "/v1/articles/refresh/extra", pattern: "/v1/*", want: false},
		{path: "/v1/items/{uuid}", pattern: "/v1/items/{uuid}", want: true},
		{path: "/v1/items/123", pattern: "/v1/items/{uuid}", want: false},
	}
	for _, tc := range cases {
		if got := pathMatches(tc.path, tc.pattern); got != tc.want {
			t.Errorf("pathMatches(%q, %q) = %v, want %v", tc.path, tc.pattern, got, tc.want)
		}
	}
}

func TestSplitArticleClusterCycle(t *testing.T) {
	spec := map[string]any{
		"components": map[string]any{
			"schemas": map[string]any{
				"Article": map[string]any{
					"type": "object",
					"properties": map[string]any{
						"title": map[string]any{"type": "string"},
						"cluster": map[string]any{
							"allOf": []any{
								map[string]any{"$ref": "#/components/schemas/NewsCluster"},
							},
						},
					},
				},
				"NewsCluster": map[string]any{
					"type": "object",
					"properties": map[string]any{
						"name": map[string]any{"type": "string"},
						"selectedArticles": map[string]any{
							"type": "array",
							"items": map[string]any{
								"$ref": "#/components/schemas/Article",
							},
						},
					},
				},
			},
		},
	}
	var stdout bytes.Buffer
	next, err := splitArticleClusterCycle(spec, &stdout)
	if err != nil {
		t.Fatal(err)
	}
	schemas := schemasOf(next)
	articleProps := schemas["Article"].(map[string]any)["properties"].(map[string]any)
	clusterAllOf := articleProps["cluster"].(map[string]any)["allOf"].([]any)
	if clusterAllOf[0].(map[string]any)["$ref"] != "#/components/schemas/ArticleNewsCluster" {
		t.Fatalf("Article.cluster ref = %#v", clusterAllOf[0])
	}
	clusterProps := schemas["NewsCluster"].(map[string]any)["properties"].(map[string]any)
	items := clusterProps["selectedArticles"].(map[string]any)["items"].(map[string]any)
	if items["$ref"] != "#/components/schemas/NewsClusterArticle" {
		t.Fatalf("selectedArticles items ref = %#v", items)
	}
	slice := schemas["NewsClusterArticle"].(map[string]any)
	if _, ok := slice["properties"].(map[string]any)["cluster"]; ok {
		t.Fatal("NewsClusterArticle still has cluster")
	}
	embedded := schemas["ArticleNewsCluster"].(map[string]any)
	if _, ok := embedded["properties"].(map[string]any)["selectedArticles"]; ok {
		t.Fatal("ArticleNewsCluster still has selectedArticles")
	}
	if !strings.Contains(stdout.String(), "NewsClusterArticle") {
		t.Fatalf("stdout = %q", stdout.String())
	}
	articleProps["title"].(map[string]any)["x-go-type-name"] = "Shared"
	embeddedTitle := schemas["NewsClusterArticle"].(map[string]any)["properties"].(map[string]any)["title"].(map[string]any)
	if _, ok := embeddedTitle["x-go-type-name"]; ok {
		t.Fatal("NewsClusterArticle shares Article property maps")
	}
}

func TestDisambiguateInlineTypeNames(t *testing.T) {
	spec := map[string]any{
		"components": map[string]any{
			"schemas": map[string]any{
				"FilledField": map[string]any{
					"type": "object",
					"properties": map[string]any{
						"source": map[string]any{
							"type": "string",
							"enum": []any{"MANUAL", "BACKGROUND"},
						},
					},
				},
				"FilledFieldSource": map[string]any{
					"type": "object",
				},
			},
		},
	}
	var stdout bytes.Buffer
	next := disambiguateInlineTypeNames(spec, &stdout)
	source := schemasOf(next)["FilledField"].(map[string]any)["properties"].(map[string]any)["source"].(map[string]any)
	if source["x-go-type-name"] != "FilledFieldSourceKind" {
		t.Fatalf("x-go-type-name = %#v", source["x-go-type-name"])
	}
	if !strings.Contains(stdout.String(), "FilledField.source") {
		t.Fatalf("stdout = %q", stdout.String())
	}
}

func TestDisambiguateAllOfEnum(t *testing.T) {
	spec := map[string]any{
		"components": map[string]any{
			"schemas": map[string]any{
				"Article": map[string]any{
					"type": "object",
					"properties": map[string]any{
						"cluster": map[string]any{
							"allOf": []any{
								map[string]any{
									"type": "string",
									"enum": []any{"A"},
								},
							},
						},
					},
				},
				"ArticleCluster0": map[string]any{"type": "object"},
			},
		},
	}
	var stdout bytes.Buffer
	next := disambiguateInlineTypeNames(spec, &stdout)
	cluster := schemasOf(next)["Article"].(map[string]any)["properties"].(map[string]any)["cluster"].(map[string]any)
	enumSchema := cluster["allOf"].([]any)[0].(map[string]any)
	if enumSchema["x-go-type-name"] != "ArticleCluster0Kind" {
		t.Fatalf("x-go-type-name = %#v", enumSchema["x-go-type-name"])
	}
}

func TestDisambiguateStableOrder(t *testing.T) {
	spec := map[string]any{
		"components": map[string]any{
			"schemas": map[string]any{
				"Dup": map[string]any{
					"type": "object",
					"properties": map[string]any{
						"source_kind": map[string]any{
							"type": "string",
							"enum": []any{"A"},
						},
						"sourceKind": map[string]any{
							"type": "string",
							"enum": []any{"B"},
						},
					},
				},
			},
		},
	}
	var stdout bytes.Buffer
	next := disambiguateInlineTypeNames(spec, &stdout)
	props := schemasOf(next)["Dup"].(map[string]any)["properties"].(map[string]any)
	if _, ok := props["sourceKind"].(map[string]any)["x-go-type-name"]; ok {
		t.Fatal("sourceKind should keep DupSourceKind")
	}
	if props["source_kind"].(map[string]any)["x-go-type-name"] != "DupSourceKindKind" {
		t.Fatalf("source_kind name = %#v", props["source_kind"].(map[string]any)["x-go-type-name"])
	}
}

func TestToCamel(t *testing.T) {
	if got := inlineEnumGoName([]string{"FilledField", "source"}); got != "FilledFieldSource" {
		t.Fatalf("inlineEnumGoName = %s", got)
	}
	if got := toCamel("search-articles"); got != "SearchArticles" {
		t.Fatalf("toCamel = %s", got)
	}
	if got := toCamel("foo#bar"); got != "FooBar" {
		t.Fatalf("toCamel hash = %s", got)
	}
	if got := toCamel("at&t"); got != "AtT" {
		t.Fatalf("toCamel amp = %s", got)
	}
	if got := toCamel("  price  "); got != "Price" {
		t.Fatalf("toCamel trim = %s", got)
	}
}

func TestHideAndExclude(t *testing.T) {
	spec := map[string]any{
		"paths": map[string]any{
			"/v1/articles/refresh":      map[string]any{},
			"/v1/articles/refresh/jobs": map[string]any{},
			"/v1/articles/all":          map[string]any{},
		},
		"components": map[string]any{
			"schemas": map[string]any{
				"Article": map[string]any{
					"type":     "object",
					"required": []any{"title", "vectors"},
					"properties": map[string]any{
						"title":   map[string]any{"type": "string"},
						"vectors": map[string]any{"type": "array"},
					},
				},
			},
		},
	}
	var stdout bytes.Buffer
	next, applied, missing := hideProperties(spec, map[string][]string{"Article": {"vectors"}}, &stdout)
	if applied != 1 || len(missing) != 0 {
		t.Fatalf("applied=%d missing=%v", applied, missing)
	}
	props := schemasOf(next)["Article"].(map[string]any)
	if _, ok := props["properties"].(map[string]any)["vectors"]; ok {
		t.Fatal("vectors property remains")
	}
	required := props["required"].([]any)
	if len(required) != 1 || required[0] != "title" {
		t.Fatalf("required = %#v", required)
	}

	excluded, removed, unused := excludePaths(next, []string{"/v1/articles/refresh", "/v1/articles/refresh/**"}, &stdout)
	if len(unused) != 0 {
		t.Fatalf("unused = %v", unused)
	}
	if len(removed) != 2 {
		t.Fatalf("removed = %v", removed)
	}
	paths := excluded["paths"].(map[string]any)
	if _, ok := paths["/v1/articles/all"]; !ok {
		t.Fatal("kept path was dropped")
	}
	if _, ok := paths["/v1/articles/refresh"]; ok {
		t.Fatal("refresh path remains")
	}
}
