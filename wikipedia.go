// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package perigon

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/goperigon/perigon-go-sdk/v2/internal/apijson"
	"github.com/goperigon/perigon-go-sdk/v2/internal/apiquery"
	"github.com/goperigon/perigon-go-sdk/v2/internal/requestconfig"
	"github.com/goperigon/perigon-go-sdk/v2/option"
	"github.com/goperigon/perigon-go-sdk/v2/packages/param"
	"github.com/goperigon/perigon-go-sdk/v2/packages/respjson"
)

// WikipediaService contains methods and other services that help with interacting
// with the perigon API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWikipediaService] method instead.
type WikipediaService struct {
	Options []option.RequestOption
}

// NewWikipediaService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewWikipediaService(opts ...option.RequestOption) (r WikipediaService) {
	r = WikipediaService{}
	r.Options = opts
	return
}

// Search and filter all Wikipedia pages available via the Perigon API. The result
// includes a list of individual pages that were matched to your specific criteria.
func (r *WikipediaService) Search(ctx context.Context, query WikipediaSearchParams, opts ...option.RequestOption) (res *WikipediaSearchResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/wikipedia/all"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Perform a natural language search over Wikipedia pages using semantic relevance.
// The result includes a list of page sections most closely matched to your query
// intent.
func (r *WikipediaService) VectorSearch(ctx context.Context, body WikipediaVectorSearchParams, opts ...option.RequestOption) (res *WikipediaVectorSearchResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/vector/wikipedia/all"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Complex filter structure for Wikipedia page searches that supports nested
// logical operations (AND, OR, NOT) and multiple filtering criteria.
type WikipediaSearchFilterParam struct {
	// Adds additional AND filter objects. These objects must be of the same type as
	// the original filter object and will be combined with the existing filter using
	// the AND logical operator.
	And []WikipediaSearchFilterParam `json:"AND,omitzero"`
	// A filter object for logical NOT operations
	Not []WikipediaSearchFilterParam `json:"NOT,omitzero"`
	// Adds additional OR filter objects. These objects must be of the same type as the
	// original filter object and will be combined with the existing filter using the
	// OR logical operator.
	Or []WikipediaSearchFilterParam `json:"OR,omitzero"`
	// Filter by specific Perigon page identifiers. Accepts either a single ID or an
	// array of IDs. Returns only pages matching these IDs.
	PageID []string `json:"pageId,omitzero"`
	// Filter by specific section identifiers. Accepts either a single ID or an array
	// of IDs. Returns only pages containing these sections.
	SectionID []string `json:"sectionId,omitzero"`
	// Filter by specific Wikipedia project codes. Returns only pages matching these
	// projects.
	WikiCode []string `json:"wikiCode,omitzero"`
	// Filter by specific Wikidata entity IDs. Returns only pages whose Wikidata
	// entities match those ids.
	WikidataID []string `json:"wikidataId,omitzero"`
	// Filter by specific Wikidata entity IDs. Returns only pages whose Wikidata
	// entities are instances of provided ids.
	WikidataInstanceOfID []string `json:"wikidataInstanceOfId,omitzero"`
	// Filter by specific Wikidata entity labels. Returns only pages whose Wikidata
	// entities are instances of these labels.
	WikidataInstanceOfLabel []string `json:"wikidataInstanceOfLabel,omitzero"`
	// Filter by specific Wikipedia namespaces. Returns only pages matching these
	// namespaces.
	WikiNamespace []int64 `json:"wikiNamespace,omitzero"`
	// Filter by specific Wikipedia page identifiers. Accepts either a single ID or an
	// array of IDs. Returns only pages matching these IDs.
	WikiPageID []int64 `json:"wikiPageId,omitzero"`
	// Filter by specific Perigon page revision identifiers. Accepts either a single ID
	// or an array of IDs. Returns only pages matching these IDs.
	WikiRevisionID []int64 `json:"wikiRevisionId,omitzero"`
	paramObj
}

func (r WikipediaSearchFilterParam) MarshalJSON() (data []byte, err error) {
	type shadow WikipediaSearchFilterParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WikipediaSearchFilterParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Wikipedia search result
type WikipediaSearchResponse struct {
	NumResults int64                           `json:"numResults,required"`
	Results    []WikipediaSearchResponseResult `json:"results,required"`
	Status     int64                           `json:"status,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		NumResults  respjson.Field
		Results     respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WikipediaSearchResponse) RawJSON() string { return r.JSON.raw }
func (r *WikipediaSearchResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WikipediaSearchResponseResult struct {
	ID                 string                                 `json:"id,nullable"`
	Categories         []string                               `json:"categories,nullable"`
	ExternalLinks      []string                               `json:"externalLinks,nullable"`
	Pageviews          int64                                  `json:"pageviews,nullable"`
	RedirectTitles     []string                               `json:"redirectTitles,nullable"`
	References         []string                               `json:"references,nullable"`
	ScrapedAt          string                                 `json:"scrapedAt,nullable"`
	Sections           []WikipediaSearchResponseResultSection `json:"sections,nullable"`
	Summary            string                                 `json:"summary,nullable"`
	TopImage           string                                 `json:"topImage,nullable"`
	URL                string                                 `json:"url,nullable"`
	WikiCode           string                                 `json:"wikiCode,nullable"`
	WikidataID         string                                 `json:"wikidataId,nullable"`
	WikidataInstanceOf []WikidataLabelHolder                  `json:"wikidataInstanceOf,nullable"`
	WikiNamespace      int64                                  `json:"wikiNamespace,nullable"`
	WikiPageID         int64                                  `json:"wikiPageId,nullable"`
	WikiRevisionID     int64                                  `json:"wikiRevisionId,nullable"`
	WikiRevisionTs     string                                 `json:"wikiRevisionTs,nullable"`
	WikiTitle          string                                 `json:"wikiTitle,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		Categories         respjson.Field
		ExternalLinks      respjson.Field
		Pageviews          respjson.Field
		RedirectTitles     respjson.Field
		References         respjson.Field
		ScrapedAt          respjson.Field
		Sections           respjson.Field
		Summary            respjson.Field
		TopImage           respjson.Field
		URL                respjson.Field
		WikiCode           respjson.Field
		WikidataID         respjson.Field
		WikidataInstanceOf respjson.Field
		WikiNamespace      respjson.Field
		WikiPageID         respjson.Field
		WikiRevisionID     respjson.Field
		WikiRevisionTs     respjson.Field
		WikiTitle          respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WikipediaSearchResponseResult) RawJSON() string { return r.JSON.raw }
func (r *WikipediaSearchResponseResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WikipediaSearchResponseResultSection struct {
	ID         string  `json:"id,nullable"`
	Loc        []int64 `json:"loc,nullable"`
	StyleLevel int64   `json:"styleLevel,nullable"`
	TextRaw    string  `json:"textRaw,nullable"`
	TextRich   string  `json:"textRich,nullable"`
	Title      string  `json:"title,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Loc         respjson.Field
		StyleLevel  respjson.Field
		TextRaw     respjson.Field
		TextRich    respjson.Field
		Title       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WikipediaSearchResponseResultSection) RawJSON() string { return r.JSON.raw }
func (r *WikipediaSearchResponseResultSection) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Wikipedia vector search result
type WikipediaVectorSearchResponse struct {
	Results []WikipediaVectorSearchResponseResult `json:"results,required"`
	Status  int64                                 `json:"status,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Results     respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WikipediaVectorSearchResponse) RawJSON() string { return r.JSON.raw }
func (r *WikipediaVectorSearchResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WikipediaVectorSearchResponseResult struct {
	Data  WikipediaVectorSearchResponseResultData `json:"data"`
	Score float64                                 `json:"score,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Score       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WikipediaVectorSearchResponseResult) RawJSON() string { return r.JSON.raw }
func (r *WikipediaVectorSearchResponseResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WikipediaVectorSearchResponseResultData struct {
	Content            string                `json:"content,nullable"`
	PageID             string                `json:"pageId,nullable"`
	Pageviews          int64                 `json:"pageviews,nullable"`
	RedirectTitles     []string              `json:"redirectTitles,nullable"`
	SectionID          string                `json:"sectionId,nullable"`
	StyleLevel         int64                 `json:"styleLevel,nullable"`
	Title              string                `json:"title,nullable"`
	WikiCode           string                `json:"wikiCode,nullable"`
	WikidataID         string                `json:"wikidataId,nullable"`
	WikidataInstanceOf []WikidataLabelHolder `json:"wikidataInstanceOf,nullable"`
	WikiNamespace      int64                 `json:"wikiNamespace,nullable"`
	WikiPageID         int64                 `json:"wikiPageId,nullable"`
	WikiRevisionID     int64                 `json:"wikiRevisionId,nullable"`
	WikiRevisionTs     string                `json:"wikiRevisionTs,nullable"`
	WikiTitle          string                `json:"wikiTitle,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content            respjson.Field
		PageID             respjson.Field
		Pageviews          respjson.Field
		RedirectTitles     respjson.Field
		SectionID          respjson.Field
		StyleLevel         respjson.Field
		Title              respjson.Field
		WikiCode           respjson.Field
		WikidataID         respjson.Field
		WikidataInstanceOf respjson.Field
		WikiNamespace      respjson.Field
		WikiPageID         respjson.Field
		WikiRevisionID     respjson.Field
		WikiRevisionTs     respjson.Field
		WikiTitle          respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WikipediaVectorSearchResponseResultData) RawJSON() string { return r.JSON.raw }
func (r *WikipediaVectorSearchResponseResultData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WikipediaSearchParams struct {
	// The specific page of results to retrieve in the paginated response. Starts at 0.
	Page param.Opt[int64] `query:"page,omitzero" json:"-"`
	// Retrieve pages with the average number of views per day higher than the provided
	// value.
	PageviewsFrom param.Opt[int64] `query:"pageviewsFrom,omitzero" json:"-"`
	// Retrieve pages with the average number of views per day lower than the provided
	// value.
	PageviewsTo param.Opt[int64] `query:"pageviewsTo,omitzero" json:"-"`
	// Primary search query for filtering pages based on their title, summary, and
	// content. Supports Boolean operators (AND, OR, NOT), exact phrases with quotes,
	// and wildcards (\* and ?) for flexible searching.
	Q param.Opt[string] `query:"q,omitzero" json:"-"`
	// Search specifically across page's references. Supports Boolean operators, exact
	// phrases with quotes, and wildcards for matching title variations.
	Reference param.Opt[string] `query:"reference,omitzero" json:"-"`
	// Retrieve pages scraped after this date. Accepts ISO 8601 format (e.g.,
	// 2023-03-01T00:00:00) or yyyy-mm-dd format.
	ScrapedAtFrom param.Opt[time.Time] `query:"scrapedAtFrom,omitzero" format:"date-time" json:"-"`
	// Retrieve pages scraped before this date. Accepts ISO 8601 format (e.g.,
	// 2023-03-01T00:00:00) or yyyy-mm-dd format.
	ScrapedAtTo param.Opt[time.Time] `query:"scrapedAtTo,omitzero" format:"date-time" json:"-"`
	// Whether to show the total number of all matched pages. Default value is false
	// which makes queries a bit more efficient but also counts up to 10000 pages.
	ShowNumResults param.Opt[bool] `query:"showNumResults,omitzero" json:"-"`
	// The number of articles to return per page in the paginated response.
	Size param.Opt[int64] `query:"size,omitzero" json:"-"`
	// Search specifically within page summary. Supports Boolean operators, exact
	// phrases with quotes, and wildcards for matching title variations.
	Summary param.Opt[string] `query:"summary,omitzero" json:"-"`
	// Search specifically within the page's content (across all sections). Supports
	// Boolean operators, exact phrases with quotes, and wildcards for matching title
	// variations.
	Text param.Opt[string] `query:"text,omitzero" json:"-"`
	// Search specifically within page titles. Supports Boolean operators, exact
	// phrases with quotes, and wildcards for matching title variations.
	Title param.Opt[string] `query:"title,omitzero" json:"-"`
	// Retrieve pages modified after this date. Accepts ISO 8601 format (e.g.,
	// 2023-03-01T00:00:00) or yyyy-mm-dd format.
	WikiRevisionFrom param.Opt[time.Time] `query:"wikiRevisionFrom,omitzero" format:"date-time" json:"-"`
	// Retrieve pages modified before this date. Accepts ISO 8601 format (e.g.,
	// 2023-03-01T00:00:00) or yyyy-mm-dd format.
	WikiRevisionTo param.Opt[time.Time] `query:"wikiRevisionTo,omitzero" format:"date-time" json:"-"`
	// Retrieve pages that have any viewership statistics available for them. If
	// `false` (the default) - return all pages.
	WithPageviews param.Opt[bool] `query:"withPageviews,omitzero" json:"-"`
	// Retrieve specific pages by their unique Perigon identifiers. Multiple IDs can be
	// provided to return a collection of specific pages.
	ID []string `query:"id,omitzero" json:"-"`
	// Retrieve all pages for specified categories.
	Category []string `query:"category,omitzero" json:"-"`
	// Retrieve pages containing provided section ids. Each section ID is unique.
	SectionID []string `query:"sectionId,omitzero" json:"-"`
	// Determines the Wikipedia page sorting order. Options include relevance
	// (default), revisionTsDesc (recently edited first), revisionTsAsc (recently
	// edited last), pageViewsDesc (highest viewership first), pageViewsAsc (highest
	// viewership last), scrapedAtDesc (recently scraped first), scrapedAtAsc (recently
	// scraped last).
	//
	// Any of "createdAt", "updatedAt", "relevance", "count", "totalCount".
	SortBy SortBy `query:"sortBy,omitzero" json:"-"`
	// Retrieve pages only from specified wiki projects. Currently, the only accepted
	// value is `enwiki`.
	WikiCode []string `query:"wikiCode,omitzero" json:"-"`
	// Retrieve pages by the ids corresponding to their Wikidata entities.
	WikidataID []string `query:"wikidataId,omitzero" json:"-"`
	// Retrieve all pages whose Wikidata entities are instances of these provided ids.
	WikidataInstanceOfID []string `query:"wikidataInstanceOfId,omitzero" json:"-"`
	// Retrieve all pages whose Wikidata entities are instances of these ids (provided
	// as labels).
	WikidataInstanceOfLabel []string `query:"wikidataInstanceOfLabel,omitzero" json:"-"`
	// Retrieve pages only from specified wiki namespace. Currently, only the main
	// namespace (0) is available.
	WikiNamespace []int64 `query:"wikiNamespace,omitzero" json:"-"`
	// Retrieve specific pages by their Wikipedia identifiers. These are unique only in
	// a combination with `wikiCode` parameter. Multiple IDs can be provided to return
	// a collection of specific pages.
	WikiPageID []int64 `query:"wikiPageId,omitzero" json:"-"`
	// Retrieve specific pages by their Wikipedia revision identifiers. These are
	// unique only in a combination with `wikiCode` parameter. Multiple IDs can be
	// provided to return a collection of specific pages. This ID changes each time a
	// page is edited.
	WikiRevisionID []int64 `query:"wikiRevisionId,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [WikipediaSearchParams]'s query parameters as `url.Values`.
func (r WikipediaSearchParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type WikipediaVectorSearchParams struct {
	// Natural language query to search the Wikipedia pages database
	Prompt string `json:"prompt,required"`
	// The page number to retrieve.
	Page param.Opt[int64] `json:"page,omitzero"`
	// 'pageviewsFrom' filter, will search pages with at least the provided number of
	// views per day.
	PageviewsFrom param.Opt[int64] `json:"pageviewsFrom,omitzero"`
	// 'pageviewsFrom' filter, will search pages with at most the provided number of
	// views per day.
	PageviewsTo param.Opt[int64] `json:"pageviewsTo,omitzero"`
	// The number of items per page.
	Size param.Opt[int64] `json:"size,omitzero"`
	// 'wikiRevisionFrom' filter, will search pages modified after the specified date,
	// the date could be passed as ISO or 'yyyy-mm-dd'. Date time in ISO format, ie.
	// 2024-01-01T00:00:00.
	WikiRevisionFrom param.Opt[time.Time] `json:"wikiRevisionFrom,omitzero" format:"date-time"`
	// 'wikiRevisionFrom' filter, will search pages modified before the specified date,
	// the date could be passed as ISO or 'yyyy-mm-dd'. Date time in ISO format, ie.
	// 2024-01-01T00:00:00.
	WikiRevisionTo param.Opt[time.Time] `json:"wikiRevisionTo,omitzero" format:"date-time"`
	// Complex filter structure for Wikipedia page searches that supports nested
	// logical operations (AND, OR, NOT) and multiple filtering criteria.
	Filter WikipediaSearchFilterParam `json:"filter,omitzero"`
	paramObj
}

func (r WikipediaVectorSearchParams) MarshalJSON() (data []byte, err error) {
	type shadow WikipediaVectorSearchParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WikipediaVectorSearchParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
