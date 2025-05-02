// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package perigonsdk

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/goperigon/perigon-go-sdk/internal/apijson"
	"github.com/goperigon/perigon-go-sdk/internal/apiquery"
	"github.com/goperigon/perigon-go-sdk/internal/requestconfig"
	"github.com/goperigon/perigon-go-sdk/option"
	"github.com/goperigon/perigon-go-sdk/packages/param"
	"github.com/goperigon/perigon-go-sdk/packages/resp"
)

// SummarizeService contains methods and other services that help with interacting
// with the perigon-sdk API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSummarizeService] method instead.
type SummarizeService struct {
	Options []option.RequestOption
}

// NewSummarizeService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSummarizeService(opts ...option.RequestOption) (r SummarizeService) {
	r = SummarizeService{}
	r.Options = opts
	return
}

// Produce a single, concise summary over the full corpus of articles matching your
// filters, using your prompt to guide which insights to highlight.
func (r *SummarizeService) New(ctx context.Context, params SummarizeNewParams, opts ...option.RequestOption) (res *SummarizeNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/summarize"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type SummarizeNewResponse struct {
	NumResults int64     `json:"numResults,nullable"`
	Results    []Article `json:"results,nullable"`
	Status     int64     `json:"status,nullable"`
	Summary    string    `json:"summary,nullable"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		NumResults  resp.Field
		Results     resp.Field
		Status      resp.Field
		Summary     resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SummarizeNewResponse) RawJSON() string { return r.JSON.raw }
func (r *SummarizeNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SummarizeNewParams struct {
	// The maximum number of articles to factor into the summary.
	MaxArticleCount param.Opt[int64] `json:"maxArticleCount,omitzero"`
	// Maximum number of tokens to generate in the summary.
	MaxTokens param.Opt[int64] `json:"maxTokens,omitzero"`
	// Instructions guiding how the summary should be written. Maximum length: 2 048
	// characters.
	Prompt param.Opt[string] `json:"prompt,omitzero"`
	// The maximum number of articles that should be returned in the response. This can
	// be used to return fewer than maxArticleCount results.
	ReturnedArticleCount param.Opt[int64] `json:"returnedArticleCount,omitzero"`
	// Sampling temperature for the LLM (0.0 = deterministic to 2.0 = very creative).
	Temperature param.Opt[float64] `json:"temperature,omitzero"`
	// Nucleus sampling (top-p) for the LLM (0.0 to 1.0).
	TopP param.Opt[float64] `json:"topP,omitzero"`
	// 'addDateFrom' filter, will search articles added after the specified date, the
	// date could be passed as ISO or 'yyyy-mm-dd'. Add time in ISO format, ie.
	// 2022-02-01T00:00:00
	AddDateFrom param.Opt[time.Time] `query:"addDateFrom,omitzero" format:"date-time" json:"-"`
	// 'addDateTo' filter, will search articles added before the specified date, the
	// date could be passed as ISO or 'yyyy-mm-dd'. Add time in ISO format, ie.
	// 2022-02-01T23:59:59
	AddDateTo param.Opt[time.Time] `query:"addDateTo,omitzero" format:"date-time" json:"-"`
	// Search by company name.
	CompanyName param.Opt[string] `query:"companyName,omitzero" json:"-"`
	// Search query on the article's body of content field. Semantic similar to q
	// parameter.
	Content param.Opt[string] `query:"content,omitzero" json:"-"`
	// Search query on the description field. Semantic similar to q parameter.
	Desc param.Opt[string] `query:"desc,omitzero" json:"-"`
	// 'from' filter, will search articles published after the specified date, the date
	// could be passed as ISO or 'yyyy-mm-dd'. Add time in ISO format, ie.
	// 2023-03-01T00:00:00
	From param.Opt[time.Time] `query:"from,omitzero" format:"date-time" json:"-"`
	// Latitude of the center point to search places
	Lat param.Opt[float64] `query:"lat,omitzero" json:"-"`
	// Returns only articles that point to specified links (as determined by the
	// 'links' field in the article response).
	LinkTo param.Opt[string] `query:"linkTo,omitzero" json:"-"`
	// Longitude of the center point to search places
	Lon param.Opt[float64] `query:"lon,omitzero" json:"-"`
	// Maximum distance (in km) from starting point to search articles by tagged places
	MaxDistance param.Opt[float64] `query:"maxDistance,omitzero" json:"-"`
	// Filters results with a sentiment score greater than or equal to the specified
	// value, indicating negative sentiment. See the Article Data section in Docs for
	// an explanation of scores.
	NegativeSentimentFrom param.Opt[float64] `query:"negativeSentimentFrom,omitzero" json:"-"`
	// Filters results with a sentiment score less than or equal to the specified
	// value, indicating negative sentiment. See the Article Data section in Docs for
	// an explanation of scores.
	NegativeSentimentTo param.Opt[float64] `query:"negativeSentimentTo,omitzero" json:"-"`
	// Filters results with a sentiment score greater than or equal to the specified
	// value, indicating neutral sentiment. Explanation of sentimental values can be
	// found in Docs under the Article Data section.
	NeutralSentimentFrom param.Opt[float64] `query:"neutralSentimentFrom,omitzero" json:"-"`
	// Filters results with a sentiment score less than or equal to the specified
	// value, indicating neutral sentiment. See the Article Data section in Docs for an
	// explanation of scores.
	NeutralSentimentTo param.Opt[float64] `query:"neutralSentimentTo,omitzero" json:"-"`
	// The page number to retrieve.
	Page param.Opt[int64] `query:"page,omitzero" json:"-"`
	// Filter to show only results where the source has a paywall (true) or does not
	// have a paywall (false).
	Paywall param.Opt[bool] `query:"paywall,omitzero" json:"-"`
	// Filters results with a sentiment score greater than or equal to the specified
	// value, indicating positive sentiment. See the Article Data section in Docs for
	// an explanation of scores.
	PositiveSentimentFrom param.Opt[float64] `query:"positiveSentimentFrom,omitzero" json:"-"`
	// Filters results with a sentiment score less than or equal to the specified
	// value, indicating positive sentiment. See the Article Data section in Docs for
	// an explanation of scores.
	PositiveSentimentTo param.Opt[float64] `query:"positiveSentimentTo,omitzero" json:"-"`
	// Filters by Google Content Categories. This field will filter by the category
	// prefix only. Example: prefixTaxonomy=/Finance
	PrefixTaxonomy param.Opt[string] `query:"prefixTaxonomy,omitzero" json:"-"`
	// Search query, each article will be scored and ranked against it. Articles are
	// searched on the title, description, and content fields.
	Q param.Opt[string] `query:"q,omitzero" json:"-"`
	// Will search articles that were refreshed after the specified date. The date
	// could be passed as ISO or 'yyyy-mm-dd'. Add time in ISO format, ie.
	// 2022-02-01T00:00:00
	RefreshDateFrom param.Opt[time.Time] `query:"refreshDateFrom,omitzero" format:"date-time" json:"-"`
	// Will search articles that were refreshed before the specified date. The date
	// could be passed as ISO or 'yyyy-mm-dd'. Add time in ISO format, ie.
	// 2022-02-01T23:59:59
	RefreshDateTo param.Opt[time.Time] `query:"refreshDateTo,omitzero" format:"date-time" json:"-"`
	// Shows all articles belonging to the same reprint group. A reprint group includes
	// one original article (the first one processed by the API) and all its known
	// reprints.
	ReprintGroupID param.Opt[string] `query:"reprintGroupId,omitzero" json:"-"`
	// Expand a query to search the translation, translatedTitle, and
	// translatedDescription fields for non-English articles.
	SearchTranslation param.Opt[bool] `query:"searchTranslation,omitzero" json:"-"`
	// Whether to show the total number of all matched articles. Default value is false
	// which makes queries a bit more efficient but also counts up to 10000 articles.
	ShowNumResults param.Opt[bool] `query:"showNumResults,omitzero" json:"-"`
	// Whether to return reprints in the response or not. Reprints are usually wired
	// articles from sources like AP or Reuters that are reprinted in multiple sources
	// at the same time. By default, this parameter is 'true'.
	ShowReprints param.Opt[bool] `query:"showReprints,omitzero" json:"-"`
	// The number of items per page.
	Size param.Opt[int64] `query:"size,omitzero" json:"-"`
	// Latitude of the center point to search articles created by local publications.
	SourceLat param.Opt[float64] `query:"sourceLat,omitzero" json:"-"`
	// Latitude of the center point to search articles created by local publications.
	SourceLon param.Opt[float64] `query:"sourceLon,omitzero" json:"-"`
	// Maximum distance from starting point to search articles created by local
	// publications.
	SourceMaxDistance param.Opt[float64] `query:"sourceMaxDistance,omitzero" json:"-"`
	// Search article headlines/title field. Semantic similar to q parameter.
	Title param.Opt[string] `query:"title,omitzero" json:"-"`
	// 'to' filter, will search articles published before the specified date, the date
	// could be passed as ISO or 'yyyy-mm-dd'. Add time in ISO format, ie.
	// 2022-02-01T23:59:59
	To param.Opt[time.Time] `query:"to,omitzero" format:"date-time" json:"-"`
	// Search query on the url field. Semantic similar to q parameter. E.g. could be
	// used for querying certain website sections, e.g. source=cnn.com&url=travel.
	URL param.Opt[string] `query:"url,omitzero" json:"-"`
	// Method for selecting articles: ARTICLES (include all matches) or CLUSTERS (one
	// per cluster).
	//
	// Any of "ARTICLES", "CLUSTERS".
	Method SummarizeNewParamsMethod `json:"method,omitzero"`
	// The underlying LLM model to use for generation.
	//
	// Any of "gpt-4o", "gpt-4o-mini", "gpt-4.1", "gpt-4.1-mini", "gpt-4.1-nano",
	// "llama-3.3-70b-versatile", "deepseek-r1-distill-llama-70b".
	Model SummarizeNewParamsModel `json:"model,omitzero"`
	// Which article fields to include when generating the summary. Up to three values
	// from TITLE, CONTENT, SUMMARY.
	//
	// Any of "TITLE", "CONTENT", "SUMMARY".
	SummarizeFields SummarizeNewParamsSummarizeFields `json:"summarizeFields,omitzero"`
	// Filters articles where a specified area, such as a neighborhood, borough, or
	// district, plays a central role in the content, beyond mere mentions, to ensure
	// the results are deeply relevant to the area in question. If multiple parameters
	// are passed, they will be applied as OR operations.
	Area []string `query:"area,omitzero" json:"-"`
	// Article ID will search for a news article by the ID of the article. If several
	// parameters are passed, all matched articles will be returned.
	ArticleID []string `query:"articleId,omitzero" json:"-"`
	// A list of author names to include. Only articles written by any of the specified
	// authors are returned. This is ideal when you wish to focus on content from
	// particular voices or experts.
	Author []string `query:"author,omitzero" json:"-"`
	// Author names to filter by. Article author bylines are used as a source field. If
	// multiple parameters are passed, they will be applied as OR operations.
	Byline []string `query:"byline,omitzero" json:"-"`
	// Filter by categories. Categories are general themes that the article is about.
	// Examples of categories: Tech, Politics, etc. If multiple parameters are passed,
	// they will be applied as OR operations. Use 'none' to search uncategorized
	// articles.
	Category []string `query:"category,omitzero" json:"-"`
	// Filters articles where a specified city plays a central role in the content,
	// beyond mere mentions, to ensure the results are deeply relevant to the urban
	// area in question. If multiple parameters are passed, they will be applied as OR
	// operations.
	City []string `query:"city,omitzero" json:"-"`
	// Search for related content using a cluster ID. Passing a cluster ID will filter
	// results to only the content found within the cluster.
	ClusterID []string `query:"clusterId,omitzero" json:"-"`
	// Search by company domains for filtering. E.g. companyDomain=apple.com.
	CompanyDomain []string `query:"companyDomain,omitzero" json:"-"`
	// List of company IDs to filter by.
	CompanyID []string `query:"companyId,omitzero" json:"-"`
	// Search by company symbols.
	CompanySymbol []string `query:"companySymbol,omitzero" json:"-"`
	// Country code to filter by country. If multiple parameters are passed, they will
	// be applied as OR operations.
	Country []string `query:"country,omitzero" json:"-"`
	// A list of counties to include (or specify) in the search results. This field
	// filters the returned articles based on the county associated with the event or
	// news. Only articles tagged with one of these counties will be included.
	County []string `query:"county,omitzero" json:"-"`
	// A list of author names to exclude from the search results. Any article written
	// by an author whose name matches one in this list will be omitted, which helps to
	// avoid content from certain individuals.
	ExcludeAuthor []string `query:"excludeAuthor,omitzero" json:"-"`
	// A list of article categories to be omitted. If an article is tagged with any
	// category present in this list (such as “Polotics”, “Tech”, “Sports”, etc.), it
	// will not appear in the search results.
	ExcludeCategory []string `query:"excludeCategory,omitzero" json:"-"`
	// A list of cities to exclude from the results. Articles that are associated with
	// any of the specified cities will be filtered out.
	ExcludeCity []string `query:"excludeCity,omitzero" json:"-"`
	// A list of company domains to exclude. If an article is related to a company that
	// uses one of the specified domains (for instance, “example.com”), it will not be
	// returned in the results.
	ExcludeCompanyDomain []string `query:"excludeCompanyDomain,omitzero" json:"-"`
	// A list of company identifiers. Articles associated with companies that have any
	// of these unique IDs will be filtered out from the returned results, ensuring
	// that certain companies or corporate entities are not included.
	ExcludeCompanyID []string `query:"excludeCompanyId,omitzero" json:"-"`
	// A list of stock symbols (ticker symbols) that identify companies to be excluded.
	// Articles related to companies using any of these symbols will be omitted, which
	// is useful for targeting or avoiding specific public companies.
	ExcludeCompanySymbol []string `query:"excludeCompanySymbol,omitzero" json:"-"`
	// Excludes articles from specific counties or administrative divisions in the
	// vector search results. Accepts either a single county name or a list of county
	// names. County names should match the format used in article metadata (e.g., 'Los
	// Angeles County', 'Cook County'). This parameter allows for more granular
	// geographic filter
	ExcludeCounty []string `query:"excludeCounty,omitzero" json:"-"`
	// A list of journalist (or reporter) identifiers to exclude. If an article is
	// written by a journalist whose ID matches any in this list, it will not be part
	// of the result set.
	ExcludeJournalistID []string `query:"excludeJournalistId,omitzero" json:"-"`
	// Exclude results that include specific labels (Opinion, Non-news, Paid News,
	// etc.). You can filter multiple by repeating the parameter.
	ExcludeLabel []string `query:"excludeLabel,omitzero" json:"-"`
	// A list of languages to be excluded. Any article published in one of the
	// languages provided in this filter will not be returned. This is useful when you
	// are interested only in news published in specific languages.
	ExcludeLanguage []string `query:"excludeLanguage,omitzero" json:"-"`
	// Excludes articles where a specified country plays a central role in the content,
	// ensuring results are not deeply relevant to the country in question. If multiple
	// parameters are passed, they will be applied as AND operations, excluding
	// articles relevant to any of the specified countries.
	ExcludeLocationsCountry []string `query:"excludeLocationsCountry,omitzero" json:"-"`
	// A list of person names that, when associated with the content, cause the article
	// to be excluded. This filter removes articles related to any individuals whose
	// names match those on the list.
	ExcludePersonName []string `query:"excludePersonName,omitzero" json:"-"`
	// A list of Wikidata identifiers for individuals. Articles mentioning persons with
	// any of these Wikidata IDs will be filtered out. This is particularly helpful
	// when using a unique identifier to prevent ambiguity in names.
	ExcludePersonWikidataID []string `query:"excludePersonWikidataId,omitzero" json:"-"`
	// The domain of the website, which should be excluded from the search. Multiple
	// parameters could be provided. Wildcards (_ and ?) are suported (e.g. _.cnn.com).
	ExcludeSource []string `query:"excludeSource,omitzero" json:"-"`
	// A list of built-in source group names to exclude from the results. The Perigon
	// API categorizes sources into groups (for example, “top10” or “top100”) based on
	// type or popularity. Using this filter allows you to remove articles coming from
	// any source that belongs to one or more of the specified groups.
	ExcludeSourceGroup []string `query:"excludeSourceGroup,omitzero" json:"-"`
	// A list of states to exclude. Articles that include, or are associated with, any
	// of the states provided here will be filtered out. This is especially useful if
	// you want to ignore news tied to certain geographical areas (e.g., US states).
	ExcludeState []string `query:"excludeState,omitzero" json:"-"`
	// Filter by excluding topics. Each topic is some kind of entity that the article
	// is about. Examples of topics: Markets, Joe Biden, Green Energy, Climate Change,
	// Cryptocurrency, etc. If multiple parameters are passed, they will be applied as
	// OR operations.
	ExcludeTopic []string `query:"excludeTopic,omitzero" json:"-"`
	// Filter by journalist ID. Journalist IDs are unique journalist identifiers which
	// can be found through the Journalist API, or in the matchedAuthors field.
	JournalistID []string `query:"journalistId,omitzero" json:"-"`
	// Labels to filter by, could be 'Opinion', 'Paid-news', 'Non-news', etc. If
	// multiple parameters are passed, they will be applied as OR operations.
	Label []string `query:"label,omitzero" json:"-"`
	// Language code to filter by language. If multiple parameters are passed, they
	// will be applied as OR operations.
	Language []string `query:"language,omitzero" json:"-"`
	// Return all articles that have the specified location. Location attributes are
	// delimited by ':' between key and value, and '::' between attributes. Example:
	// 'city:New York::state:NY'.
	Location []string `query:"location,omitzero" json:"-"`
	// Filters articles where a specified country plays a central role in the content,
	// beyond mere mentions, to ensure the results are deeply relevant to the country
	// in question. If multiple parameters are passed, they will be applied as OR
	// operations.
	LocationsCountry []string `query:"locationsCountry,omitzero" json:"-"`
	// Medium will filter out news articles medium, which could be 'Video' or
	// 'Article'. If several parameters are passed, all matched articles will be
	// returned.
	Medium []string `query:"medium,omitzero" json:"-"`
	// List of person names for exact matches. Boolean and complex logic is not
	// supported on this paramter.
	PersonName []string `query:"personName,omitzero" json:"-"`
	// List of person Wikidata IDs for filtering.
	PersonWikidataID []string `query:"personWikidataId,omitzero" json:"-"`
	// 'relevance' to sort by relevance to the query, 'date' to sort by the publication
	// date (desc), 'pubDate' is a synonym to 'date', 'addDate' to sort by 'addDate'
	// field (desc), 'refreshDate' to sort by 'refreshDate' field (desc). Defaults to
	// 'relevance'
	//
	// Any of "relevance", "date", "reverseDate", "reverseAddDate", "addDate",
	// "pubDate", "refreshDate".
	SortBy AllEndpointSortBy `query:"sortBy,omitzero" json:"-"`
	// Publisher's domain can include a subdomain. If multiple parameters are passed,
	// they will be applied as OR operations. Wildcards (_ and ?) are suported (e.g.
	// _.cnn.com).
	Source []string `query:"source,omitzero" json:"-"`
	// Find articles published by sources that are located within a given city.
	SourceCity []string `query:"sourceCity,omitzero" json:"-"`
	// Find articles published by sources that are located within a given country. Must
	// be 2 character country code (i.e. us, gb, etc).
	SourceCountry []string `query:"sourceCountry,omitzero" json:"-"`
	// Find articles published by sources that are located within a given county.
	SourceCounty []string `query:"sourceCounty,omitzero" json:"-"`
	// One of the supported source groups. Find Source Groups in the guided part of our
	// documentation...
	SourceGroup []string `query:"sourceGroup,omitzero" json:"-"`
	// Find articles published by sources that are located within a given state.
	SourceState []string `query:"sourceState,omitzero" json:"-"`
	// Filters articles where a specified state plays a central role in the content,
	// beyond mere mentions, to ensure the results are deeply relevant to the state in
	// question. If multiple parameters are passed, they will be applied as OR
	// operations.
	State []string `query:"state,omitzero" json:"-"`
	// Filters by Google Content Categories. This field will accept 1 or more
	// categories, must pass the full name of the category. Example:
	// taxonomy=/Finance/Banking/Other, /Finance/Investing/Funds
	Taxonomy []string `query:"taxonomy,omitzero" json:"-"`
	// Filters results to include only articles with the specified topics. Topics are
	// more specific classifications than categories, with an article potentially
	// having multiple topics assigned. Perigon uses both human and machine curation to
	// maintain an evolving list of available topics. Common examples include
	// 'Markets', 'Crime', 'Cryptocurrency', 'Social Issues', 'College Sports', etc.
	// See the Topics page in Docs for a complete list of available topics.
	Topic []string `query:"topic,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f SummarizeNewParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r SummarizeNewParams) MarshalJSON() (data []byte, err error) {
	type shadow SummarizeNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}

// URLQuery serializes [SummarizeNewParams]'s query parameters as `url.Values`.
func (r SummarizeNewParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Method for selecting articles: ARTICLES (include all matches) or CLUSTERS (one
// per cluster).
type SummarizeNewParamsMethod string

const (
	SummarizeNewParamsMethodArticles SummarizeNewParamsMethod = "ARTICLES"
	SummarizeNewParamsMethodClusters SummarizeNewParamsMethod = "CLUSTERS"
)

// The underlying LLM model to use for generation.
type SummarizeNewParamsModel string

const (
	SummarizeNewParamsModelGpt4o                     SummarizeNewParamsModel = "gpt-4o"
	SummarizeNewParamsModelGpt4oMini                 SummarizeNewParamsModel = "gpt-4o-mini"
	SummarizeNewParamsModelGpt4_1                    SummarizeNewParamsModel = "gpt-4.1"
	SummarizeNewParamsModelGpt4_1Mini                SummarizeNewParamsModel = "gpt-4.1-mini"
	SummarizeNewParamsModelGpt4_1Nano                SummarizeNewParamsModel = "gpt-4.1-nano"
	SummarizeNewParamsModelLlama3_3_70bVersatile     SummarizeNewParamsModel = "llama-3.3-70b-versatile"
	SummarizeNewParamsModelDeepseekR1DistillLlama70b SummarizeNewParamsModel = "deepseek-r1-distill-llama-70b"
)

// Which article fields to include when generating the summary. Up to three values
// from TITLE, CONTENT, SUMMARY.
type SummarizeNewParamsSummarizeFields string

const (
	SummarizeNewParamsSummarizeFieldsTitle   SummarizeNewParamsSummarizeFields = "TITLE"
	SummarizeNewParamsSummarizeFieldsContent SummarizeNewParamsSummarizeFields = "CONTENT"
	SummarizeNewParamsSummarizeFieldsSummary SummarizeNewParamsSummarizeFields = "SUMMARY"
)
