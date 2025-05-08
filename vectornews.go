// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package perigon

import (
	"context"
	"net/http"
	"time"

	"github.com/goperigon/perigon-go-sdk/internal/apijson"
	"github.com/goperigon/perigon-go-sdk/internal/requestconfig"
	"github.com/goperigon/perigon-go-sdk/option"
	"github.com/goperigon/perigon-go-sdk/packages/param"
	"github.com/goperigon/perigon-go-sdk/packages/respjson"
)

// VectorNewsService contains methods and other services that help with interacting
// with the perigon API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewVectorNewsService] method instead.
type VectorNewsService struct {
	Options []option.RequestOption
}

// NewVectorNewsService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewVectorNewsService(opts ...option.RequestOption) (r VectorNewsService) {
	r = VectorNewsService{}
	r.Options = opts
	return
}

// Perform a natural language search over news articles from the past 6 months
// using semantic relevance. The result includes a list of articles most closely
// matched to your query intent.
func (r *VectorNewsService) Search(ctx context.Context, body VectorNewsSearchParams, opts ...option.RequestOption) (res *VectorNewsSearchResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/vector/news/all"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Complex filter structure for article searches that supports nested logical
// operations (AND, OR, NOT) and multiple filtering criteria.
type ArticleSearchFilterParam struct {
	// Adds additional AND filter objects. These objects must be of the same type as
	// the original filter object and will be combined with the existing filter using
	// the AND logical operator.
	And []ArticleSearchFilterParam `json:"AND,omitzero"`
	// Filter by specific article identifiers. Accepts either a single ID or an array
	// of IDs. Returns only articles matching these IDs.
	ArticleID []string `json:"articleId,omitzero"`
	// Filter by broad content categories such as Politics, Tech, Sports, Business, or
	// Finance. Accepts either a single category or an array. Use none to find
	// uncategorized articles. Multiple values create an OR filter.
	Category []string `json:"category,omitzero"`
	// Filter articles that mention or are related to specific cities. Accepts either a
	// single city name or an array. Multiple values create an OR filter.
	City []string `json:"city,omitzero"`
	// Filter by specific story identifiers. Accepts either a single ID or an array of
	// IDs. Returns only articles belonging to these stories.
	ClusterID []string `json:"clusterId,omitzero"`
	// Filter articles by company domains (e.g., apple.com). Accepts either a single
	// domain or an array. Multiple values create an OR filter. For a complete list of
	// tracked companies and their domains, refer to the /companies endpoint.
	CompanyDomain []string `json:"companyDomain,omitzero"`
	// Filter articles by company identifiers. Accepts either a single ID or an array.
	// Multiple values create an OR filter. For a complete list of tracked companies
	// and their IDs, refer to the /companies endpoint.
	CompanyID []string `json:"companyId,omitzero"`
	// Filter articles by company name mentions. Accepts either a single name or an
	// array. Performs exact matching on company names. Multiple values create an OR
	// filter. For a complete list of tracked companies and their names, refer to the
	// /companies endpoint.
	CompanyName []string `json:"companyName,omitzero"`
	// Filter articles by company stock symbols (e.g., AAPL, MSFT). Accepts either a
	// single symbol or an array. Multiple values create an OR filter. For a complete
	// list of tracked companies and their symbols, refer to the /companies endpoint.
	CompanySymbol []string `json:"companySymbol,omitzero"`
	// Filter articles by countries they mention using two-letter country codes in
	// lowercase (e.g., us, gb, jp). Accepts either a single country code or an array.
	// Multiple values create an OR filter. See documentation for supported country
	// codes.
	Country []string `json:"country,omitzero"`
	// Filter articles that mention or are related to specific counties. Accepts either
	// a single county name or an array. Multiple values create an OR filter. County
	// names typically include the word 'County' (e.g., Los Angeles County).
	County []string `json:"county,omitzero"`
	// Exclude articles that mention or are related to specific cities. Accepts either
	// a single city name or an array. Multiple values create an AND-exclude filter.
	ExcludeCity []string `json:"excludeCity,omitzero"`
	// Exclude articles related to companies with specific domains. Accepts either a
	// single domain or an array. Multiple values create an AND-exclude filter. For a
	// complete list of tracked companies and their domains, refer to the /companies
	// endpoint.
	ExcludeCompanyDomain []string `json:"excludeCompanyDomain,omitzero"`
	// Exclude articles mentioning companies with specific identifiers. Accepts either
	// a single ID or an array. Multiple values create an AND-exclude filter. For a
	// complete list of tracked companies and their IDs, refer to the /companies
	// endpoint.
	ExcludeCompanyID []string `json:"excludeCompanyId,omitzero"`
	// Exclude articles related to companies with specific stock symbols. Accepts
	// either a single symbol or an array. Multiple values create an AND-exclude
	// filter. For a complete list of tracked companies and their symbols, refer to the
	// /companies endpoint.
	ExcludeCompanySymbol []string `json:"excludeCompanySymbol,omitzero"`
	// Exclude articles from specific countries using two-letter country codes in
	// lowercase. Accepts either a single country code or an array. Multiple values
	// create an AND-exclude filter. See documentation for supported country codes.
	ExcludeCountry []string `json:"excludeCountry,omitzero"`
	// Exclude articles that mention or are related to specific counties. Accepts
	// either a single county name or an array. Multiple values create an AND-exclude
	// filter. County names should match the format in article metadata (e.g., Los
	// Angeles County, Cook County).
	ExcludeCounty []string `json:"excludeCounty,omitzero"`
	// Exclude articles with specific editorial labels. Accepts either a single label
	// or an array. Multiple values create an AND-exclude filter, removing all content
	// with any of these labels.
	ExcludeLabel []string `json:"excludeLabel,omitzero"`
	// Exclude articles in specific languages using ISO-639 two-letter codes in
	// lowercase. Accepts either a single language code or an array. Multiple values
	// create an AND-exclude filter.
	ExcludeLanguage []string `json:"excludeLanguage,omitzero"`
	// Exclude articles where specified countries play a central role in the content.
	// Accepts either a single country code or an array. Multiple values create an
	// AND-exclude filter, removing articles focused on any of these countries. See
	// documentation for supported country codes.
	ExcludeLocationsCountry []string `json:"excludeLocationsCountry,omitzero"`
	// Exclude articles mentioning specific people by name. Accepts either a single
	// name or an array. Multiple values create an AND-exclude filter. For a complete
	// list of tracked individuals and their names, refer to the /people endpoint.
	ExcludePersonName []string `json:"excludePersonName,omitzero"`
	// Exclude articles mentioning people with specific Wikidata IDs. Accepts either a
	// single ID or an array. Multiple values create an AND-exclude filter. For a
	// complete list of tracked individuals and their Wikidata IDs, refer to the
	// /people endpoint.
	ExcludePersonWikidataID []string `json:"excludePersonWikidataId,omitzero"`
	// Exclude articles from specific publisher domains or subdomains. Accepts either a
	// single domain or an array of domains. Multiple values create an AND-exclude
	// filter.
	ExcludeSource []string `json:"excludeSource,omitzero"`
	// Exclude articles where specified states play a central role. Accepts either a
	// single state code or an array. Multiple values create an AND-exclude filter,
	// removing articles focused on any of these states. See documentation for
	// supported state codes.
	ExcludeState []string `json:"excludeState,omitzero"`
	// Exclude articles with specific topics. Accepts either a single topic or an
	// array. Multiple values create an AND-exclude filter, removing all content with
	// any of these topics.
	ExcludeTopic []string `json:"excludeTopic,omitzero"`
	// Filter articles by editorial labels such as Opinion, Paid-news, Non-news, Fact
	// Check, or Press Release. View our docs for an exhaustive list of labels. Accepts
	// either a single label or an array. Multiple values create an OR filter.
	Label []string `json:"label,omitzero"`
	// Filter articles by their language using ISO-639 two-letter codes in lowercase
	// (e.g., en, es, fr). Accepts either a single language code or an array. Multiple
	// values create an OR filter.
	Language []string `json:"language,omitzero"`
	// Filter articles where specified countries play a central role in the content,
	// not just mentioned. Uses two-letter country codes in lowercase. Accepts either a
	// single country code or an array. Multiple values create an OR filter. See
	// documentation for supported country codes.
	LocationsCountry []string `json:"locationsCountry,omitzero"`
	// A filter object for logical NOT operations
	Not []ArticleSearchFilterParam `json:"NOT,omitzero"`
	// Adds additional OR filter objects. These objects must be of the same type as the
	// original filter object and will be combined with the existing filter using the
	// OR logical operator.
	Or []ArticleSearchFilterParam `json:"OR,omitzero"`
	// Filter articles by exact person name matches. Accepts either a single name or an
	// array. Does not support Boolean operators or wildcards. Multiple values create
	// an OR filter. For a complete list of tracked individuals and their names, refer
	// to the /people endpoint.
	PersonName []string `json:"personName,omitzero"`
	// Filter articles by Wikidata IDs of mentioned people. Accepts either a single ID
	// or an array. Multiple values create an OR filter. For a complete list of tracked
	// individuals and their Wikidata IDs, refer to the /people endpoint.
	PersonWikidataID []string `json:"personWikidataId,omitzero"`
	// Filter articles by specific publisher domains or subdomains. Accepts either a
	// single domain or an array of domains. Multiple values create an OR filter.
	Source []string `json:"source,omitzero"`
	// Filter for articles from publishers based in specific cities. Accepts either a
	// single city name or an array. Multiple values create an OR filter.
	SourceCity []string `json:"sourceCity,omitzero"`
	// Filter for articles from publishers based in specific countries. Accepts either
	// a single country code or an array. Uses two-letter country codes in lowercase
	// (e.g., us, gb). See documentation for supported country codes.
	SourceCountry []string `json:"sourceCountry,omitzero"`
	// Filter for articles from publishers based in specific counties. Accepts either a
	// single county name or an array. Multiple values create an OR filter.
	SourceCounty []string `json:"sourceCounty,omitzero"`
	// Filter articles using Perigon's curated publisher bundles (e.g., top100,
	// top25tech). Accepts either a single source group or an array. Multiple values
	// create an OR filter to include articles from any of the specified bundles.
	SourceGroup []string `json:"sourceGroup,omitzero"`
	// Filter for articles from publishers based in specific states or regions. Accepts
	// either a single state code or an array. Uses two-letter state codes in
	// lowercase. See documentation for supported state codes.
	SourceState []string `json:"sourceState,omitzero"`
	// Filter articles where specified states play a central role in the content.
	// Accepts either a single state code or an array. Multiple values create an OR
	// filter. Uses two-letter state codes in lowercase. See documentation for
	// supported state codes.
	State []string `json:"state,omitzero"`
	// Filter by Google Content Categories. Must pass the full hierarchical path of the
	// category. Accepts either a single path or an array. Example:
	// taxonomy=/Finance/Banking/Other,/Finance/Investing/Funds. Multiple values create
	// an OR filter.
	Taxonomy []string `json:"taxonomy,omitzero"`
	// Filter by specific topics such as Markets, Crime, Cryptocurrency, or College
	// Sports. Accepts either a single topic or an array. Topics are more granular than
	// categories, and articles can have multiple topics. Multiple values create an OR
	// filter.
	Topic []string `json:"topic,omitzero"`
	// Filter articles from publishers based at specific geographic locations. Uses
	// latitude, longitude, and radius parameters to define a circular search area for
	// publisher locations.
	Coordinates ArticleSearchFilterCoordinatesParam `json:"coordinates,omitzero"`
	// Filter articles from publishers based at specific geographic locations. Uses
	// latitude, longitude, and radius parameters to define a circular search area for
	// publisher locations.
	SourceCoordinates ArticleSearchFilterSourceCoordinatesParam `json:"sourceCoordinates,omitzero"`
	paramObj
}

func (r ArticleSearchFilterParam) MarshalJSON() (data []byte, err error) {
	type shadow ArticleSearchFilterParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ArticleSearchFilterParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Filter articles from publishers based at specific geographic locations. Uses
// latitude, longitude, and radius parameters to define a circular search area for
// publisher locations.
type ArticleSearchFilterCoordinatesParam struct {
	Lat    param.Opt[float64] `json:"lat,omitzero"`
	Lon    param.Opt[float64] `json:"lon,omitzero"`
	Radius param.Opt[float64] `json:"radius,omitzero"`
	paramObj
}

func (r ArticleSearchFilterCoordinatesParam) MarshalJSON() (data []byte, err error) {
	type shadow ArticleSearchFilterCoordinatesParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ArticleSearchFilterCoordinatesParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Filter articles from publishers based at specific geographic locations. Uses
// latitude, longitude, and radius parameters to define a circular search area for
// publisher locations.
type ArticleSearchFilterSourceCoordinatesParam struct {
	Lat    param.Opt[float64] `json:"lat,omitzero"`
	Lon    param.Opt[float64] `json:"lon,omitzero"`
	Radius param.Opt[float64] `json:"radius,omitzero"`
	paramObj
}

func (r ArticleSearchFilterSourceCoordinatesParam) MarshalJSON() (data []byte, err error) {
	type shadow ArticleSearchFilterSourceCoordinatesParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ArticleSearchFilterSourceCoordinatesParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Vector search result
type VectorNewsSearchResponse struct {
	Results []VectorNewsSearchResponseResult `json:"results,required"`
	Status  int64                            `json:"status,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Results     respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VectorNewsSearchResponse) RawJSON() string { return r.JSON.raw }
func (r *VectorNewsSearchResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VectorNewsSearchResponseResult struct {
	Data  Article `json:"data"`
	Score float64 `json:"score,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Score       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VectorNewsSearchResponseResult) RawJSON() string { return r.JSON.raw }
func (r *VectorNewsSearchResponseResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VectorNewsSearchParams struct {
	// Natural language query to search the news articles database
	Prompt string `json:"prompt,required"`
	// The page number to retrieve.
	Page param.Opt[int64] `json:"page,omitzero"`
	// 'pubDateFrom' filter, will search articles published after the specified date,
	// the date could be passed as ISO or 'yyyy-mm-dd'. Date time in ISO format, ie.
	// 2024-01-01T00:00:00 - Default: Only articles with a pubDate within the last 30
	// days of the request
	PubDateFrom param.Opt[time.Time] `json:"pubDateFrom,omitzero" format:"date-time"`
	// 'pubDateFrom' filter, will search articles published before the specified date,
	// the date could be passed as ISO or 'yyyy-mm-dd'. Date time in ISO format, ie.
	// 2024-01-01T00:00:00
	PubDateTo param.Opt[time.Time] `json:"pubDateTo,omitzero" format:"date-time"`
	// Whether to return reprints in the response or not. Reprints are usually wired
	// articles from sources like AP or Reuters that are reprinted in multiple sources
	// at the same time. By default, this parameter is 'true'.
	ShowReprints param.Opt[bool] `json:"showReprints,omitzero"`
	// The number of items per page.
	Size param.Opt[int64] `json:"size,omitzero"`
	// Complex filter structure for article searches that supports nested logical
	// operations (AND, OR, NOT) and multiple filtering criteria.
	Filter ArticleSearchFilterParam `json:"filter,omitzero"`
	paramObj
}

func (r VectorNewsSearchParams) MarshalJSON() (data []byte, err error) {
	type shadow VectorNewsSearchParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VectorNewsSearchParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
