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
	"github.com/goperigon/perigon-go-sdk/packages/resp"
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

// A versatile filter object to refine search results based on articles, clusters,
// sources, languages, categories, locations, companies, and people. Supports
// logical operators (AND, OR, NOT) for complex queries. Accepts single values or
// arrays, with arrays applied as OR operations.
type ArticleSearchFilterParam struct {
	// Adds additional AND filter objects. These objects must be of the same type as
	// the original filter object and will be combined with the existing filter using
	// the AND logical operator.
	And []ArticleSearchFilterParam `json:"AND,omitzero"`
	// Filter by specific article(s). Array of article IDs. For convenience, a single
	// string ID is also accepted.
	ArticleID []string `json:"articleId,omitzero"`
	// Filter by categories. Categories are general themes that the article is about.
	// Examples of categories: Tech, Politics, etc. If multiple parameters are passed,
	// they will be applied as OR operations. Use 'none' to search uncategorized
	// articles.
	Category []string `json:"category,omitzero"`
	// Filters articles where a specified city plays a central role in the content,
	// beyond mere mentions, to ensure the results are deeply relevant to the urban
	// area in question. If multiple parameters are passed, they will be applied as OR
	// operations.
	City []string `json:"city,omitzero"`
	// Filter by specific cluster(s). Accepts either a single string or an array of
	// strings.
	ClusterID []string `json:"clusterId,omitzero"`
	// Search by company domains for filtering. E.g. companyDomain=apple.com.
	CompanyDomain []string `json:"companyDomain,omitzero"`
	// List of company IDs to filter by.
	CompanyID []string `json:"companyId,omitzero"`
	// Search by company name.
	CompanyName []string `json:"companyName,omitzero"`
	// Search by company symbols.
	CompanySymbol []string `json:"companySymbol,omitzero"`
	// Country code to filter by country. If multiple parameters are passed, they will
	// be applied as OR operations. Only accepts country ISO codes.
	Country []string `json:"country,omitzero"`
	// Filters articles where a specified state plays a central role in the content,
	// beyond mere mentions, to ensure the results are deeply relevant to the state in
	// question. If multiple parameters are passed, they will be applied as OR
	// operations.
	County []string `json:"county,omitzero"`
	// A list of cities to exclude from the results. Articles that are associated with
	// any of the specified cities will be filtered out.
	ExcludeCity []string `json:"excludeCity,omitzero"`
	// A list of company domains to exclude. If an article is related to a company that
	// uses one of the specified domains (for instance, "example.com"), it will not be
	// returned in the results.
	ExcludeCompanyDomain []string `json:"excludeCompanyDomain,omitzero"`
	// A list of company identifiers. Articles associated with companies that have any
	// of these unique IDs will be filtered out from the returned results, ensuring
	// that certain companies or corporate entities are not included.
	ExcludeCompanyID []string `json:"excludeCompanyId,omitzero"`
	// A list of stock symbols (ticker symbols) that identify companies to be excluded.
	// Articles related to companies using any of these symbols will be omitted, which
	// is useful for targeting or avoiding specific public companies.
	ExcludeCompanySymbol []string `json:"excludeCompanySymbol,omitzero"`
	// Excludes articles from specific countries in the vector search results. Accepts
	// a list of country codes (e.g., 'US', 'UK', 'CA'). Use this parameter to filter
	// out articles published in countries you don't want to include in your search
	// results. See the Country and Language Support section in the documentation for a
	// full list of supported country codes.
	ExcludeCountry []string `json:"excludeCountry,omitzero"`
	// Excludes articles from specific counties or administrative divisions in the
	// vector search results. Accepts either a single county name or a list of county
	// names. County names should match the format used in article metadata (e.g., 'Los
	// Angeles County', 'Cook County'). This parameter allows for more granular
	// geographic filter
	ExcludeCounty []string `json:"excludeCounty,omitzero"`
	// Exclude results that include specific labels (Opinion, Non-news, Paid News,
	// etc.). You can filter multiple by repeating the parameter.
	ExcludeLabel []string `json:"excludeLabel,omitzero"`
	// A list of languages to be excluded. Any article published in one of the
	// languages provided in this filter will not be returned. This is useful when you
	// are interested only in news published in specific languages.
	ExcludeLanguage []string `json:"excludeLanguage,omitzero"`
	// Excludes articles where a specified country plays a central role in the content,
	// ensuring results are not deeply relevant to the country in question. If multiple
	// parameters are passed, they will be applied as AND operations, excluding
	// articles relevant to any of the specified countries.
	ExcludeLocationsCountry []string `json:"excludeLocationsCountry,omitzero"`
	// A list of person names that, when associated with the content, cause the article
	// to be excluded. This filter removes articles related to any individuals whose
	// names match those on the list.
	ExcludePersonName []string `json:"excludePersonName,omitzero"`
	// A list of Wikidata identifiers for individuals. Articles mentioning persons with
	// any of these Wikidata IDs will be filtered out. This is particularly helpful
	// when using a unique identifier to prevent ambiguity in names.
	ExcludePersonWikidataID []string `json:"excludePersonWikidataId,omitzero"`
	// The domain of the website, which should be excluded from the search. Multiple
	// parameters could be provided. Wildcards (_ and ?) are suported (e.g. _.cnn.com).
	ExcludeSource []string `json:"excludeSource,omitzero"`
	// A list of states to exclude. Articles that include, or are associated with, any
	// of the states provided here will be filtered out. This is especially useful if
	// you want to ignore news tied to certain geographical areas (e.g., US states).
	ExcludeState []string `json:"excludeState,omitzero"`
	// Filter by excluding topics. Each topic is some kind of entity that the article
	// is about. Examples of topics: Markets, Joe Biden, Green Energy, Climate Change,
	// Cryptocurrency, etc. If multiple parameters are passed, they will be applied as
	// OR operations.
	ExcludeTopic []string `json:"excludeTopic,omitzero"`
	// Labels to filter by, could be 'Opinion', 'Paid-news', 'Non-news', etc. If
	// multiple parameters are passed, they will be applied as OR operations.
	Label []string `json:"label,omitzero"`
	// Language code to filter by language. If an array parameters are passed, they
	// will be applied as OR operations. For example: ['en', 'es']. Language ISO codes
	// must be provided.
	Language []string `json:"language,omitzero"`
	// Filters articles where a specified country plays a central role in the content,
	// beyond mere mentions, to ensure the results are deeply relevant to the country
	// in question. If multiple parameters are passed, they will be applied as OR
	// operations. Only accepts country ISO codes.
	LocationsCountry []string `json:"locationsCountry,omitzero"`
	// A filter object for logical NOT operations
	Not []ArticleSearchFilterParam `json:"NOT,omitzero"`
	// Adds additional OR filter objects. These objects must be of the same type as the
	// original filter object and will be combined with the existing filter using the
	// OR logical operator.
	Or []ArticleSearchFilterParam `json:"OR,omitzero"`
	// List of person names for exact matches. Boolean and complex logic is not
	// supported on this filter.
	PersonName []string `json:"personName,omitzero"`
	// List of person Wikidata IDs for filtering.
	PersonWikidataID []string `json:"personWikidataId,omitzero"`
	// Filter by specific source(s). Accepts either a single string or an array of
	// strings.
	Source []string `json:"source,omitzero"`
	// Find articles published by sources that are located within a given city.
	SourceCity []string `json:"sourceCity,omitzero"`
	// Find articles published by sources that are located within a given country. Must
	// be 2 character country code (i.e. us, gb, etc).
	SourceCountry []string `json:"sourceCountry,omitzero"`
	// Find articles published by sources that are located within a given county.
	SourceCounty []string `json:"sourceCounty,omitzero"`
	// Filter by specific source group, for example: 'top100'. Accepts either a single
	// string or an array of strings.
	SourceGroup []string `json:"sourceGroup,omitzero"`
	// Find articles published by sources that are located within a given state.
	SourceState []string `json:"sourceState,omitzero"`
	// Filters articles where a specified state plays a central role in the content,
	// beyond mere mentions, to ensure the results are deeply relevant to the state in
	// question. If multiple parameters are passed, they will be applied as OR
	// operations. Only accepts state ISO codes.
	State []string `json:"state,omitzero"`
	// Filters by Google Content Categories. This field will accept 1 or more
	// categories, must pass the full name of the category. Example:
	// taxonomy=/Finance/Banking/Other, /Finance/Investing/Funds
	Taxonomy []string `json:"taxonomy,omitzero"`
	// Filter by topics. Each topic is some kind of entity that the article is about.
	// Examples of topics: Markets, Joe Biden, Green Energy, Climate Change,
	// Cryptocurrency, etc. If multiple parameters are passed, they will be applied as
	// OR operations.
	Topic []string `json:"topic,omitzero"`
	// Filter using sources that are located on specific coordinates using Lat, lon and
	// radius.
	Coordinates ArticleSearchFilterCoordinatesParam `json:"coordinates,omitzero"`
	// Filter using sources that are located on specific coordinates using Lat, lon and
	// radius.
	SourceCoordinates ArticleSearchFilterSourceCoordinatesParam `json:"sourceCoordinates,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f ArticleSearchFilterParam) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }
func (r ArticleSearchFilterParam) MarshalJSON() (data []byte, err error) {
	type shadow ArticleSearchFilterParam
	return param.MarshalObject(r, (*shadow)(&r))
}

// Filter using sources that are located on specific coordinates using Lat, lon and
// radius.
type ArticleSearchFilterCoordinatesParam struct {
	Lat    param.Opt[float64] `json:"lat,omitzero"`
	Lon    param.Opt[float64] `json:"lon,omitzero"`
	Radius param.Opt[float64] `json:"radius,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f ArticleSearchFilterCoordinatesParam) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r ArticleSearchFilterCoordinatesParam) MarshalJSON() (data []byte, err error) {
	type shadow ArticleSearchFilterCoordinatesParam
	return param.MarshalObject(r, (*shadow)(&r))
}

// Filter using sources that are located on specific coordinates using Lat, lon and
// radius.
type ArticleSearchFilterSourceCoordinatesParam struct {
	Lat    param.Opt[float64] `json:"lat,omitzero"`
	Lon    param.Opt[float64] `json:"lon,omitzero"`
	Radius param.Opt[float64] `json:"radius,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f ArticleSearchFilterSourceCoordinatesParam) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r ArticleSearchFilterSourceCoordinatesParam) MarshalJSON() (data []byte, err error) {
	type shadow ArticleSearchFilterSourceCoordinatesParam
	return param.MarshalObject(r, (*shadow)(&r))
}

// Vector search result
type VectorNewsSearchResponse struct {
	Results []VectorNewsSearchResponseResult `json:"results,nullable"`
	Status  int64                            `json:"status,nullable"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Results     resp.Field
		Status      resp.Field
		ExtraFields map[string]resp.Field
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
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Data        resp.Field
		Score       resp.Field
		ExtraFields map[string]resp.Field
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
	Prompt      string               `json:"prompt,required"`
	AddDateFrom param.Opt[time.Time] `json:"addDateFrom,omitzero" format:"date-time"`
	AddDateTo   param.Opt[time.Time] `json:"addDateTo,omitzero" format:"date-time"`
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
	PubDateTo      param.Opt[time.Time] `json:"pubDateTo,omitzero" format:"date-time"`
	ScoreThreshold param.Opt[float64]   `json:"scoreThreshold,omitzero"`
	// Whether to return reprints in the response or not. Reprints are usually wired
	// articles from sources like AP or Reuters that are reprinted in multiple sources
	// at the same time. By default, this parameter is 'true'.
	ShowReprints param.Opt[bool] `json:"showReprints,omitzero"`
	// The number of items per page.
	Size param.Opt[int64] `json:"size,omitzero"`
	// A versatile filter object to refine search results based on articles, clusters,
	// sources, languages, categories, locations, companies, and people. Supports
	// logical operators (AND, OR, NOT) for complex queries. Accepts single values or
	// arrays, with arrays applied as OR operations.
	Filter ArticleSearchFilterParam `json:"filter,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f VectorNewsSearchParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r VectorNewsSearchParams) MarshalJSON() (data []byte, err error) {
	type shadow VectorNewsSearchParams
	return param.MarshalObject(r, (*shadow)(&r))
}
