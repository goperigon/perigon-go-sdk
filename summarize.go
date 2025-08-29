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

// SummarizeService contains methods and other services that help with interacting
// with the perigon API.
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
	NumResults int64     `json:"numResults,required"`
	Results    []Article `json:"results,required"`
	Status     int64     `json:"status,required"`
	Summary    string    `json:"summary,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		NumResults  respjson.Field
		Results     respjson.Field
		Status      respjson.Field
		Summary     respjson.Field
		ExtraFields map[string]respjson.Field
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
	// Date. Filter for articles added to Perigon's system after this date. Accepts ISO
	// 8601 format (e.g., 2022-02-01T00:00:00) or yyyy-mm-dd format.
	AddDateFrom param.Opt[time.Time] `query:"addDateFrom,omitzero" format:"date-time" json:"-"`
	// Date. Filter for articles added to Perigon's system before this date. Accepts
	// ISO 8601 format (e.g., 2022-02-01T23:59:59) or yyyy-mm-dd format.
	AddDateTo param.Opt[time.Time] `query:"addDateTo,omitzero" format:"date-time" json:"-"`
	// String. Filter articles by company name mentions. Performs an exact match on
	// company names.
	CompanyName param.Opt[string] `query:"companyName,omitzero" json:"-"`
	// String. Search within the full article body content. Supports Boolean logic,
	// exact phrase matching with quotes, and wildcards for comprehensive content
	// searching.
	Content param.Opt[string] `query:"content,omitzero" json:"-"`
	// String. Search within article description fields. Supports Boolean expressions,
	// exact phrase matching with quotes, and wildcards for flexible pattern matching.
	Desc param.Opt[string] `query:"desc,omitzero" json:"-"`
	// Date. Filter for articles published after this date. Accepts ISO 8601 format
	// (e.g., 2023-03-01T00:00:00) or yyyy-mm-dd format.
	From param.Opt[time.Time] `query:"from,omitzero" format:"date-time" json:"-"`
	// Integer. Specifies the size in characters of each highlighted text fragment.
	// Defaults to 100 if not specified.
	HighlightFragmentSize param.Opt[int64] `query:"highlightFragmentSize,omitzero" json:"-"`
	// Integer. Controls the maximum number of highlighted fragments to return per
	// field.
	HighlightNumFragments param.Opt[int64] `query:"highlightNumFragments,omitzero" json:"-"`
	// String. Defines the HTML tag that appears after highlighted text. Defaults to
	// '</em>' if not specified.
	HighlightPostTag param.Opt[string] `query:"highlightPostTag,omitzero" json:"-"`
	// String. Defines the HTML tag that appears before highlighted text. Defaults to
	// '<em>' if not specified.
	HighlightPreTag param.Opt[string] `query:"highlightPreTag,omitzero" json:"-"`
	// String. Specifies a separate query for highlighting, allowing highlights based
	// on terms different from the main search query. Example: main query 'q=climate
	// change' with 'highlightQ=renewable OR solar' will highlight terms 'renewable'
	// and 'solar' in results about climate change.
	HighlightQ param.Opt[string] `query:"highlightQ,omitzero" json:"-"`
	// Float. Latitude of the center point to search places
	Lat param.Opt[float64] `query:"lat,omitzero" json:"-"`
	// String. Returns only articles that contain links to the specified URL pattern.
	// Matches against the 'links' field in article responses.
	LinkTo param.Opt[string] `query:"linkTo,omitzero" json:"-"`
	// Float. Longitude of the center point to search places
	Lon param.Opt[float64] `query:"lon,omitzero" json:"-"`
	// Float. Maximum distance (in km) from starting point to search articles by tagged
	// places
	MaxDistance param.Opt[float64] `query:"maxDistance,omitzero" json:"-"`
	// Float. Filter articles with a negative sentiment score greater than or equal to
	// the specified value. Scores range from 0 to 1, with higher values indicating
	// stronger negative tone.
	NegativeSentimentFrom param.Opt[float64] `query:"negativeSentimentFrom,omitzero" json:"-"`
	// Float. Filter articles with a negative sentiment score less than or equal to the
	// specified value. Scores range from 0 to 1, with higher values indicating
	// stronger negative tone.
	NegativeSentimentTo param.Opt[float64] `query:"negativeSentimentTo,omitzero" json:"-"`
	// Float. Filter articles with a neutral sentiment score greater than or equal to
	// the specified value. Scores range from 0 to 1, with higher values indicating
	// stronger neutral tone.
	NeutralSentimentFrom param.Opt[float64] `query:"neutralSentimentFrom,omitzero" json:"-"`
	// Float. Filter articles with a neutral sentiment score less than or equal to the
	// specified value. Scores range from 0 to 1, with higher values indicating
	// stronger neutral tone.
	NeutralSentimentTo param.Opt[float64] `query:"neutralSentimentTo,omitzero" json:"-"`
	// Integer. The specific page of results to retrieve in the paginated response.
	// Starts at 0.
	Page param.Opt[int64] `query:"page,omitzero" json:"-"`
	// Boolean. Filter to show only results where the source has a paywall (true) or
	// does not have a paywall (false).
	Paywall param.Opt[bool] `query:"paywall,omitzero" json:"-"`
	// Float. Filter articles with a positive sentiment score greater than or equal to
	// the specified value. Scores range from 0 to 1, with higher values indicating
	// stronger positive tone.
	PositiveSentimentFrom param.Opt[float64] `query:"positiveSentimentFrom,omitzero" json:"-"`
	// Float. Filter articles with a positive sentiment score less than or equal to the
	// specified value. Scores range from 0 to 1, with higher values indicating
	// stronger positive tone.
	PositiveSentimentTo param.Opt[float64] `query:"positiveSentimentTo,omitzero" json:"-"`
	// String. Filters by Google Content Categories. This field will filter by the
	// category prefix only. Example: prefixTaxonomy=/Finance
	PrefixTaxonomy param.Opt[string] `query:"prefixTaxonomy,omitzero" json:"-"`
	// String. Primary search query for filtering articles based on their title,
	// description, and content. Supports Boolean operators (AND, OR, NOT), exact
	// phrases with quotes, and wildcards (\* and ?) for flexible searching.
	Q param.Opt[string] `query:"q,omitzero" json:"-"`
	// Date. Filter for articles refreshed/updated in Perigon's system after this date.
	// In most cases yields similar results to addDateFrom but can differ for updated
	// content. Accepts ISO 8601 format (e.g., 2022-02-01T00:00:00) or yyyy-mm-dd
	// format.
	RefreshDateFrom param.Opt[time.Time] `query:"refreshDateFrom,omitzero" format:"date-time" json:"-"`
	// Date. Filter for articles refreshed/updated in Perigon's system before this
	// date. In most cases yields similar results to addDateTo but can differ for
	// updated content. Accepts ISO 8601 format (e.g., 2022-02-01T23:59:59) or
	// yyyy-mm-dd format.
	RefreshDateTo param.Opt[time.Time] `query:"refreshDateTo,omitzero" format:"date-time" json:"-"`
	// String. Returns all articles in a specific reprint group, including the original
	// article and all its known reprints. Use when you want to see all versions of the
	// same content.
	ReprintGroupID param.Opt[string] `query:"reprintGroupId,omitzero" json:"-"`
	// Boolean. Expand search to include translated content fields for non-English
	// articles. When true, searches translated title, description, and content fields.
	SearchTranslation param.Opt[bool] `query:"searchTranslation,omitzero" json:"-"`
	// Boolean. When set to true, enables text highlighting in search results.
	ShowHighlighting param.Opt[bool] `query:"showHighlighting,omitzero" json:"-"`
	// Boolean. Whether to show the total number of all matched articles. Default value
	// is false which makes queries a bit more efficient but also counts up to 10000
	// articles.
	ShowNumResults param.Opt[bool] `query:"showNumResults,omitzero" json:"-"`
	// Boolean. Controls whether to include reprinted content in results. When true
	// (default), shows syndicated articles from wire services like AP or Reuters that
	// appear on multiple sites.
	ShowReprints param.Opt[bool] `query:"showReprints,omitzero" json:"-"`
	// Integer. The number of articles to return per page in the paginated response.
	Size param.Opt[int64] `query:"size,omitzero" json:"-"`
	// Float. Latitude of the center point to search articles created by local
	// publications.
	SourceLat param.Opt[float64] `query:"sourceLat,omitzero" json:"-"`
	// Float. Latitude of the center point to search articles created by local
	// publications.
	SourceLon param.Opt[float64] `query:"sourceLon,omitzero" json:"-"`
	// Float. Maximum distance from starting point to search articles created by local
	// publications.
	SourceMaxDistance param.Opt[float64] `query:"sourceMaxDistance,omitzero" json:"-"`
	// String. Search within article summary fields. Supports Boolean expressions,
	// exact phrase matching with quotes, and wildcards for flexible pattern matching.
	Summary param.Opt[string] `query:"summary,omitzero" json:"-"`
	// String. Search specifically within article headlines/titles. Supports Boolean
	// operators, exact phrases with quotes, and wildcards for matching title
	// variations.
	Title param.Opt[string] `query:"title,omitzero" json:"-"`
	// Date. Filter for articles published before this date. Accepts ISO 8601 format
	// (e.g., 2022-02-01T23:59:59) or yyyy-mm-dd format.
	To param.Opt[time.Time] `query:"to,omitzero" format:"date-time" json:"-"`
	// String. Search query on the url field. Semantic similar to q parameter. E.g.
	// could be used for querying certain website sections, e.g.
	// source=cnn.com&url=travel.
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
	SummarizeFields []string `json:"summarizeFields,omitzero"`
	// String Array. Filters articles where a specified area, such as a neighborhood,
	// borough, or district, plays a central role in the content, beyond mere mentions,
	// to ensure the results are deeply relevant to the area in question. If multiple
	// parameters are passed, they will be applied as OR operations.
	Area []string `query:"area,omitzero" json:"-"`
	// String Array. Retrieve specific news articles by their unique article
	// identifiers. Multiple IDs can be provided to return a collection of specific
	// articles.
	ArticleID []string `query:"articleId,omitzero" json:"-"`
	// String Array. Filter articles by specific author names. Works as an exact match
	// for each name. Multiple values create an OR filter to find articles by any of
	// the specified authors.
	Author []string `query:"author,omitzero" json:"-"`
	// String Array. Filter articles by author bylines. Works as an exact match for
	// each author name provided. Multiple values create an OR filter to find articles
	// by any of the specified authors.
	Byline []string `query:"byline,omitzero" json:"-"`
	// String Array. Filter by categories. Categories are general themes that the
	// article is about. Examples of categories: Tech, Politics, etc. If multiple
	// parameters are passed, they will be applied as OR operations. Use 'none' to
	// search uncategorized articles.
	Category []string `query:"category,omitzero" json:"-"`
	// String Array. Filters articles where a specified city plays a central role in
	// the content, beyond mere mentions, to ensure the results are deeply relevant to
	// the urban area in question. If multiple parameters are passed, they will be
	// applied as OR operations.
	City []string `query:"city,omitzero" json:"-"`
	// String Array. Filter results to only show content within a specific related
	// content cluster. Returns articles grouped together as part of Perigon Stories
	// based on topic relevance.
	ClusterID []string `query:"clusterId,omitzero" json:"-"`
	// String Array. Filter articles by company domains (e.g., apple.com). For
	// available company entities, consult the /companies endpoint.
	CompanyDomain []string `query:"companyDomain,omitzero" json:"-"`
	// String Array. Filter articles by company identifiers. For a complete list of
	// tracked companies, refer to the /companies endpoint.
	CompanyID []string `query:"companyId,omitzero" json:"-"`
	// String Array. Filter articles by company stock symbols. For available company
	// entities and their symbols, consult the /companies endpoint.
	CompanySymbol []string `query:"companySymbol,omitzero" json:"-"`
	// String Array. Country code to filter by country. If multiple parameters are
	// passed, they will be applied as OR operations.
	Country []string `query:"country,omitzero" json:"-"`
	// String Array. A list of counties to include (or specify) in the search results.
	// This field filters the returned articles based on the county associated with the
	// event or news. Only articles tagged with one of these counties will be included.
	County []string `query:"county,omitzero" json:"-"`
	// String Array. Exclude articles written by specific authors. Any article with an
	// author name matching an entry in this list will be omitted from results.
	// Multiple values create an AND-exclude filter.
	ExcludeAuthor []string `query:"excludeAuthor,omitzero" json:"-"`
	// String Array. Exclude articles with specific categories. Multiple values create
	// an AND-exclude filter, removing all content with any of these categories.
	ExcludeCategory []string `query:"excludeCategory,omitzero" json:"-"`
	// String Array. A list of cities to exclude from the results. Articles that are
	// associated with any of the specified cities will be filtered out.
	ExcludeCity []string `query:"excludeCity,omitzero" json:"-"`
	// String Array. Exclude articles related to companies with specific domains.
	// Creates an AND-exclude filter to remove content about these companies.
	ExcludeCompanyDomain []string `query:"excludeCompanyDomain,omitzero" json:"-"`
	// String Array. Exclude articles mentioning companies with specific identifiers.
	// Creates an AND-exclude filter to remove content about these corporate entities.
	ExcludeCompanyID []string `query:"excludeCompanyId,omitzero" json:"-"`
	// String Array. A list of stock symbols (ticker symbols) that identify companies
	// to be excluded. Articles related to companies using any of these symbols will be
	// omitted, which is useful for targeting or avoiding specific public companies.
	ExcludeCompanySymbol []string `query:"excludeCompanySymbol,omitzero" json:"-"`
	// String Array. Excludes articles from specific counties or administrative
	// divisions in the search results. Accepts either a single county name or a list
	// of county names. County names should match the format used in article metadata
	// (e.g., 'Los Angeles County', 'Cook County'). This parameter allows for more
	// granular geographic filter
	ExcludeCounty []string `query:"excludeCounty,omitzero" json:"-"`
	// String Array. Exclude articles written by specific journalists identified by
	// their unique IDs. Multiple values create an AND-exclude filter.
	ExcludeJournalistID []string `query:"excludeJournalistId,omitzero" json:"-"`
	// String Array. Exclude articles with specific editorial labels. Multiple values
	// create an AND-exclude filter, removing all content with any of these labels.
	ExcludeLabel []string `query:"excludeLabel,omitzero" json:"-"`
	// String Array. Exclude articles in specific languages using ISO-639 two-letter
	// codes. Multiple values create an AND-exclude filter.
	ExcludeLanguage []string `query:"excludeLanguage,omitzero" json:"-"`
	// String Array. Excludes articles where a specified country plays a central role
	// in the content, ensuring results are not deeply relevant to the country in
	// question. If multiple parameters are passed, they will be applied as AND
	// operations, excluding articles relevant to any of the specified countries.
	ExcludeLocationsCountry []string `query:"excludeLocationsCountry,omitzero" json:"-"`
	// String Array. Exclude articles mentioning specific people by name. Creates an
	// AND-exclude filter to remove content about these individuals.
	ExcludePersonName []string `query:"excludePersonName,omitzero" json:"-"`
	// String Array. Exclude articles mentioning people with specific Wikidata IDs.
	// Creates an AND-exclude filter to remove content about these individuals. Uses
	// precise identifiers to avoid name ambiguity.
	ExcludePersonWikidataID []string `query:"excludePersonWikidataId,omitzero" json:"-"`
	// String Array. Exclude articles from specific publisher domains or subdomains.
	// Supports wildcards (_ and ?) for pattern matching (e.g., _.cnn.com). Multiple
	// values create an AND-exclude filter.
	ExcludeSource []string `query:"excludeSource,omitzero" json:"-"`
	// String Array. Exclude articles from specified Perigon source groups. Multiple
	// values create an AND-exclude filter, removing content from publishers in any of
	// the specified bundles (e.g., top10, top100).
	ExcludeSourceGroup []string `query:"excludeSourceGroup,omitzero" json:"-"`
	// String Array. A list of states to exclude. Articles that include, or are
	// associated with, any of the states provided here will be filtered out. This is
	// especially useful if you want to ignore news tied to certain geographical areas
	// (e.g., US states).
	ExcludeState []string `query:"excludeState,omitzero" json:"-"`
	// String Array. Exclude articles with specific topics. Multiple values create an
	// AND-exclude filter, removing all content with any of these topics.
	ExcludeTopic []string `query:"excludeTopic,omitzero" json:"-"`
	// String Array. Exclude articles mentioning entities from specified watchlists.
	// Multiple values create an AND-exclude filter, removing content mentioning any
	// entity from the specified watchlists.
	ExcludeWatchlist []string `query:"excludeWatchlist,omitzero" json:"-"`
	// String Array. Filter by unique journalist identifiers which can be found through
	// the Journalist API or in the matchedAuthors field. Multiple values create an OR
	// filter.
	JournalistID []string `query:"journalistId,omitzero" json:"-"`
	// String Array. Filter articles by editorial labels such as Opinion, Paid-news,
	// Non-news, Fact Check, or Press Release. Multiple values create an OR filter.
	Label []string `query:"label,omitzero" json:"-"`
	// String Array. Filter articles by their language using ISO-639 two-letter codes
	// (e.g., en, es, fr). Multiple values create an OR filter.
	Language []string `query:"language,omitzero" json:"-"`
	// String Array. Return all articles that have the specified location. Location
	// attributes are delimited by ':' between key and value, and '::' between
	// attributes. Example: 'city:New York::state:NY'.
	Location []string `query:"location,omitzero" json:"-"`
	// String Array. Filters articles where a specified country plays a central role in
	// the content, beyond mere mentions, to ensure the results are deeply relevant to
	// the country in question. If multiple parameters are passed, they will be applied
	// as OR operations.
	LocationsCountry []string `query:"locationsCountry,omitzero" json:"-"`
	// String Array. Filter articles by their primary medium type. Accepts Article for
	// written content or Video for video-based stories. Multiple values create an OR
	// filter.
	Medium []string `query:"medium,omitzero" json:"-"`
	// String Array. Filter articles by exact person name matches. Does not support
	// Boolean or complex logic. For available person entities, consult the /people
	// endpoint.
	PersonName []string `query:"personName,omitzero" json:"-"`
	// String Array. Filter articles by Wikidata IDs of mentioned people. Refer to the
	// /people endpoint for a complete list of tracked individuals.
	PersonWikidataID []string `query:"personWikidataId,omitzero" json:"-"`
	// String. Determines the article sorting order. Options include relevance
	// (default), date/pubDate (newest publication date first), reverseDate (oldest
	// publication date first), addDate (newest ingestion date first), reverseAddDate
	// (oldest ingestion date first), and refreshDate (most recently updated in system
	// first, often identical to addDate).
	//
	// Any of "relevance", "date", "reverseDate", "reverseAddDate", "addDate",
	// "pubDate", "refreshDate".
	SortBy AllEndpointSortBy `query:"sortBy,omitzero" json:"-"`
	// String Array. Filter articles by specific publisher domains or subdomains.
	// Supports wildcards (_ and ?) for pattern matching (e.g., _.cnn.com). Multiple
	// values create an OR filter.
	Source []string `query:"source,omitzero" json:"-"`
	// String Array. Find articles published by sources that are located within a given
	// city.
	SourceCity []string `query:"sourceCity,omitzero" json:"-"`
	// String Array. Find articles published by sources that are located within a given
	// country. Must be 2 character country code (i.e. us, gb, etc).
	SourceCountry []string `query:"sourceCountry,omitzero" json:"-"`
	// String Array. Find articles published by sources that are located within a given
	// county.
	SourceCounty []string `query:"sourceCounty,omitzero" json:"-"`
	// String Array. Filter articles using Perigon's curated publisher bundles (e.g.,
	// top100, top25crypto). Multiple values create an OR filter to include articles
	// from any of the specified bundles.
	SourceGroup []string `query:"sourceGroup,omitzero" json:"-"`
	// String Array. Find articles published by sources that are located within a given
	// state.
	SourceState []string `query:"sourceState,omitzero" json:"-"`
	// String Array. Filters articles where a specified state plays a central role in
	// the content, beyond mere mentions, to ensure the results are deeply relevant to
	// the state in question. If multiple parameters are passed, they will be applied
	// as OR operations.
	State []string `query:"state,omitzero" json:"-"`
	// String Array. Filters by Google Content Categories. This field will accept 1 or
	// more categories, must pass the full name of the category. Example:
	// taxonomy=/Finance/Banking/Other, /Finance/Investing/Funds.
	// [Full list](https://cloud.google.com/natural-language/docs/categories)
	Taxonomy []string `query:"taxonomy,omitzero" json:"-"`
	// String Array. Filter by specific topics such as Markets, Crime, Cryptocurrency,
	// or College Sports. Topics are more granular than categories, and articles can
	// have multiple topics. Use the /topics endpoint for a complete list of available
	// topics. Multiple values create an OR filter.
	Topic []string `query:"topic,omitzero" json:"-"`
	// String Array. Filter articles using watchlists of people and companies. Multiple
	// values create an OR filter to include articles mentioning any entity from the
	// specified watchlists.
	Watchlist []string `query:"watchlist,omitzero" json:"-"`
	paramObj
}

func (r SummarizeNewParams) MarshalJSON() (data []byte, err error) {
	type shadow SummarizeNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SummarizeNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
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
