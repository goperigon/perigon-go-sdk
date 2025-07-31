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
	"github.com/goperigon/perigon-go-sdk/v2/shared"
)

// StoryService contains methods and other services that help with interacting with
// the perigon API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewStoryService] method instead.
type StoryService struct {
	Options []option.RequestOption
}

// NewStoryService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewStoryService(opts ...option.RequestOption) (r StoryService) {
	r = StoryService{}
	r.Options = opts
	return
}

// Search and filter all news stories available via the Perigon API. Each story
// aggregates key information across related articles, including AI-generated
// names, summaries, and key points.
func (r *StoryService) List(ctx context.Context, query StoryListParams, opts ...option.RequestOption) (res *StoryListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/stories/all"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type NewsCluster struct {
	ID                string                  `json:"id,nullable"`
	Categories        []RecordStatHolder      `json:"categories,nullable"`
	Companies         []NewsClusterCompany    `json:"companies,nullable"`
	Countries         []RecordStatHolder      `json:"countries,nullable"`
	CreatedAt         string                  `json:"createdAt,nullable"`
	DuplicateOf       string                  `json:"duplicateOf,nullable"`
	Highlights        map[string][]string     `json:"highlights,nullable"`
	ImageSource       NewsClusterImageSource  `json:"imageSource,nullable"`
	ImageURL          string                  `json:"imageUrl,nullable"`
	InitializedAt     string                  `json:"initializedAt,nullable"`
	KeyPoints         []NewsClusterKeyPoint   `json:"keyPoints,nullable"`
	Locations         []NewsClusterLocation   `json:"locations,nullable"`
	Name              string                  `json:"name,nullable"`
	People            []NewsClusterPerson     `json:"people,nullable"`
	Questions         []NewsClusterQuestion   `json:"questions,nullable"`
	ReprintCount      int64                   `json:"reprintCount,nullable"`
	SelectedArticles  []Article               `json:"selectedArticles,nullable"`
	Sentiment         NewsClusterSentiment    `json:"sentiment,nullable"`
	ShortSummary      string                  `json:"shortSummary,nullable"`
	Slug              string                  `json:"slug,nullable"`
	Summary           string                  `json:"summary,nullable"`
	SummaryReferences []string                `json:"summaryReferences,nullable"`
	Taxonomies        []RecordStatHolder      `json:"taxonomies,nullable"`
	TopCategories     []shared.CategoryHolder `json:"topCategories,nullable"`
	TopCompanies      []NewsClusterTopCompany `json:"topCompanies,nullable"`
	TopCountries      []string                `json:"topCountries,nullable"`
	Topics            []RecordStatHolder      `json:"topics,nullable"`
	TopLocations      []shared.LocationHolder `json:"topLocations,nullable"`
	TopPeople         []NewsClusterTopPerson  `json:"topPeople,nullable"`
	TopTaxonomies     []shared.CategoryHolder `json:"topTaxonomies,nullable"`
	TopTopics         []NewsClusterTopTopic   `json:"topTopics,nullable"`
	TotalCount        int64                   `json:"totalCount,nullable"`
	UniqueCount       int64                   `json:"uniqueCount,nullable"`
	UniqueSources     []string                `json:"uniqueSources,nullable"`
	UpdatedAt         string                  `json:"updatedAt,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                respjson.Field
		Categories        respjson.Field
		Companies         respjson.Field
		Countries         respjson.Field
		CreatedAt         respjson.Field
		DuplicateOf       respjson.Field
		Highlights        respjson.Field
		ImageSource       respjson.Field
		ImageURL          respjson.Field
		InitializedAt     respjson.Field
		KeyPoints         respjson.Field
		Locations         respjson.Field
		Name              respjson.Field
		People            respjson.Field
		Questions         respjson.Field
		ReprintCount      respjson.Field
		SelectedArticles  respjson.Field
		Sentiment         respjson.Field
		ShortSummary      respjson.Field
		Slug              respjson.Field
		Summary           respjson.Field
		SummaryReferences respjson.Field
		Taxonomies        respjson.Field
		TopCategories     respjson.Field
		TopCompanies      respjson.Field
		TopCountries      respjson.Field
		Topics            respjson.Field
		TopLocations      respjson.Field
		TopPeople         respjson.Field
		TopTaxonomies     respjson.Field
		TopTopics         respjson.Field
		TotalCount        respjson.Field
		UniqueCount       respjson.Field
		UniqueSources     respjson.Field
		UpdatedAt         respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NewsCluster) RawJSON() string { return r.JSON.raw }
func (r *NewsCluster) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NewsClusterCompany struct {
	ID      string   `json:"id,nullable"`
	Count   int64    `json:"count,nullable"`
	Domains []string `json:"domains,nullable"`
	Name    string   `json:"name,nullable"`
	Symbols []string `json:"symbols,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Count       respjson.Field
		Domains     respjson.Field
		Name        respjson.Field
		Symbols     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NewsClusterCompany) RawJSON() string { return r.JSON.raw }
func (r *NewsClusterCompany) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NewsClusterImageSource struct {
	Domain   string                `json:"domain,nullable"`
	Location shared.SourceLocation `json:"location,nullable"`
	Paywall  bool                  `json:"paywall,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Domain      respjson.Field
		Location    respjson.Field
		Paywall     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NewsClusterImageSource) RawJSON() string { return r.JSON.raw }
func (r *NewsClusterImageSource) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NewsClusterKeyPoint struct {
	Point      string   `json:"point,nullable"`
	References []string `json:"references,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Point       respjson.Field
		References  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NewsClusterKeyPoint) RawJSON() string { return r.JSON.raw }
func (r *NewsClusterKeyPoint) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NewsClusterLocation struct {
	Area   string `json:"area,nullable"`
	City   string `json:"city,nullable"`
	Count  int64  `json:"count,nullable"`
	County string `json:"county,nullable"`
	State  string `json:"state,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Area        respjson.Field
		City        respjson.Field
		Count       respjson.Field
		County      respjson.Field
		State       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NewsClusterLocation) RawJSON() string { return r.JSON.raw }
func (r *NewsClusterLocation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NewsClusterPerson struct {
	Count      int64  `json:"count,nullable"`
	Name       string `json:"name,nullable"`
	WikidataID string `json:"wikidataId,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		Name        respjson.Field
		WikidataID  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NewsClusterPerson) RawJSON() string { return r.JSON.raw }
func (r *NewsClusterPerson) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NewsClusterQuestion struct {
	Answer     string   `json:"answer,nullable"`
	Question   string   `json:"question,nullable"`
	References []string `json:"references,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Answer      respjson.Field
		Question    respjson.Field
		References  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NewsClusterQuestion) RawJSON() string { return r.JSON.raw }
func (r *NewsClusterQuestion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NewsClusterSentiment struct {
	Negative float64 `json:"negative,nullable"`
	Neutral  float64 `json:"neutral,nullable"`
	Positive float64 `json:"positive,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Negative    respjson.Field
		Neutral     respjson.Field
		Positive    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NewsClusterSentiment) RawJSON() string { return r.JSON.raw }
func (r *NewsClusterSentiment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NewsClusterTopCompany struct {
	ID      string   `json:"id,nullable"`
	Domains []string `json:"domains,nullable"`
	Name    string   `json:"name,nullable"`
	Symbols []string `json:"symbols,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Domains     respjson.Field
		Name        respjson.Field
		Symbols     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NewsClusterTopCompany) RawJSON() string { return r.JSON.raw }
func (r *NewsClusterTopCompany) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NewsClusterTopPerson struct {
	Name       string `json:"name,nullable"`
	WikidataID string `json:"wikidataId,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		WikidataID  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NewsClusterTopPerson) RawJSON() string { return r.JSON.raw }
func (r *NewsClusterTopPerson) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NewsClusterTopTopic struct {
	Name string `json:"name,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NewsClusterTopTopic) RawJSON() string { return r.JSON.raw }
func (r *NewsClusterTopTopic) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RecordStatHolder struct {
	Count int64  `json:"count,nullable"`
	Name  string `json:"name,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RecordStatHolder) RawJSON() string { return r.JSON.raw }
func (r *RecordStatHolder) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Story search result
type StoryListResponse struct {
	NumResults int64         `json:"numResults,required"`
	Results    []NewsCluster `json:"results,required"`
	Status     int64         `json:"status,required"`
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
func (r StoryListResponse) RawJSON() string { return r.JSON.raw }
func (r *StoryListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StoryListParams struct {
	// Filter stories by names of top mentioned companies. Performs an exact match on
	// company names in the topCompanies field.
	CompanyName param.Opt[string] `query:"companyName,omitzero" json:"-"`
	// 'from' filter, will search stories created after the specified date, the date
	// could be passed as ISO or 'yyyy-mm-dd'. Add time in ISO format, ie.
	// 2023-03-01T00:00:00
	From param.Opt[time.Time] `query:"from,omitzero" format:"date-time" json:"-"`
	// Specifies the size in characters of each highlighted text fragment. Defaults to
	// 100 if not specified.
	HighlightFragmentSize param.Opt[int64] `query:"highlightFragmentSize,omitzero" json:"-"`
	// Controls the maximum number of highlighted fragments to return per field.
	HighlightNumFragments param.Opt[int64] `query:"highlightNumFragments,omitzero" json:"-"`
	// Defines the HTML tag that appears after highlighted text. Defaults to '</em>' if
	// not specified.
	HighlightPostTag param.Opt[string] `query:"highlightPostTag,omitzero" json:"-"`
	// Defines the HTML tag that appears before highlighted text. Defaults to '<em>' if
	// not specified.
	HighlightPreTag param.Opt[string] `query:"highlightPreTag,omitzero" json:"-"`
	// Specifies a separate query for highlighting, allowing highlights based on terms
	// different from the main search query. Example: main query 'q=climate change'
	// with 'highlightQ=renewable OR solar' will highlight terms 'renewable' and
	// 'solar' in results about climate change.
	HighlightQ param.Opt[string] `query:"highlightQ,omitzero" json:"-"`
	// Filter for stories created after this date. Alternative parameter for filtering
	// by story creation date.
	InitializedFrom param.Opt[time.Time] `query:"initializedFrom,omitzero" format:"date-time" json:"-"`
	// Filter for stories created before this date. Alternative parameter for filtering
	// by story creation date.
	InitializedTo param.Opt[time.Time] `query:"initializedTo,omitzero" format:"date-time" json:"-"`
	// Filter by maximum cluster size. Maximum cluster size filter applies to number of
	// unique articles in the cluster.
	MaxClusterSize param.Opt[int64] `query:"maxClusterSize,omitzero" json:"-"`
	// Filter by minimum cluster size. Minimum cluster size filter applies to number of
	// unique articles.
	MinClusterSize param.Opt[int64] `query:"minClusterSize,omitzero" json:"-"`
	// Specifies the minimum number of unique sources required for a story to appear in
	// results. Higher values return more significant stories covered by multiple
	// publications. Default is 3.
	MinUniqueSources param.Opt[int64] `query:"minUniqueSources,omitzero" json:"-"`
	// Search specifically within story names. Supports Boolean operators, exact
	// phrases with quotes, and wildcards for matching name variations.
	Name param.Opt[string] `query:"name,omitzero" json:"-"`
	// Filter to only include stories that have been assigned names. Defaults to true.
	// Note that stories only receive names after they contain at least 5 unique
	// articles.
	NameExists param.Opt[bool] `query:"nameExists,omitzero" json:"-"`
	// Filter stories with an aggregate negative sentiment score greater than or equal
	// to the specified value. Scores range from 0 to 1, with higher values indicating
	// stronger negative tone.
	NegativeSentimentFrom param.Opt[float64] `query:"negativeSentimentFrom,omitzero" json:"-"`
	// Filter articles with an aggregate negative sentiment score less than or equal to
	// the specified value. Scores range from 0 to 1, with higher values indicating
	// stronger negative tone.
	NegativeSentimentTo param.Opt[float64] `query:"negativeSentimentTo,omitzero" json:"-"`
	// Filter articles with an aggregate neutral sentiment score greater than or equal
	// to the specified value. Scores range from 0 to 1, with higher values indicating
	// stronger neutral tone.
	NeutralSentimentFrom param.Opt[float64] `query:"neutralSentimentFrom,omitzero" json:"-"`
	// Filter articles with an aggregate neutral sentiment score less than or equal to
	// the specified value. Scores range from 0 to 1, with higher values indicating
	// stronger neutral tone.
	NeutralSentimentTo param.Opt[float64] `query:"neutralSentimentTo,omitzero" json:"-"`
	// The specific page of results to retrieve in the paginated response. Starts at 0.
	Page param.Opt[int64] `query:"page,omitzero" json:"-"`
	// Filter stories by exact name matches of top mentioned people. Does not support
	// Boolean or complex logic. For available person entities, consult the /people
	// endpoint.
	PersonName param.Opt[string] `query:"personName,omitzero" json:"-"`
	// Filter articles with an aggregate positive sentiment score greater than or equal
	// to the specified value. Scores range from 0 to 1, with higher values indicating
	// stronger positive tone.
	PositiveSentimentFrom param.Opt[float64] `query:"positiveSentimentFrom,omitzero" json:"-"`
	// Filter articles with an aggregate positive sentiment score less than or equal to
	// the specified value. Scores range from 0 to 1, with higher values indicating
	// stronger positive tone.
	PositiveSentimentTo param.Opt[float64] `query:"positiveSentimentTo,omitzero" json:"-"`
	// Primary search query for filtering stories based on their name, summary, and key
	// points. Supports Boolean operators (AND, OR, NOT), exact phrases with quotes,
	// and wildcards (\* and ?) for flexible searching.
	Q param.Opt[string] `query:"q,omitzero" json:"-"`
	// Stories are deduplicated by default. If a story is deduplicated, all future
	// articles are merged into the original story. duplicateOf field contains the
	// original cluster Id. When showDuplicates=true, all stories are shown.
	ShowDuplicates param.Opt[bool] `query:"showDuplicates,omitzero" json:"-"`
	// When set to true, enables text highlighting in search results.
	ShowHighlighting param.Opt[bool] `query:"showHighlighting,omitzero" json:"-"`
	// Show total number of results. By default set to false, will cap result count
	// at 10000.
	ShowNumResults    param.Opt[bool] `query:"showNumResults,omitzero" json:"-"`
	ShowStoryPageInfo param.Opt[bool] `query:"showStoryPageInfo,omitzero" json:"-"`
	// The number of articles to return per page in the paginated response.
	Size param.Opt[int64] `query:"size,omitzero" json:"-"`
	// 'to' filter, will search stories created before the specified date, the date
	// could be passed as ISO or 'yyyy-mm-dd'. Add time in ISO format, ie.
	// 2023-03-01T23:59:59
	To param.Opt[time.Time] `query:"to,omitzero" format:"date-time" json:"-"`
	// Filter for stories that received new articles after this date. Useful for
	// tracking developing news events or evolving storylines.
	UpdatedFrom param.Opt[time.Time] `query:"updatedFrom,omitzero" format:"date-time" json:"-"`
	// Filter for stories that received new articles before this date. Useful for
	// tracking developing news events or evolving storylines.
	UpdatedTo param.Opt[time.Time] `query:"updatedTo,omitzero" format:"date-time" json:"-"`
	// Filter local news by area. Applies only to local news, when this param is passed
	// non-local news will not be returned. If multiple parameters are passed, they
	// will be applied as OR operations.
	Area []string `query:"area,omitzero" json:"-"`
	// Filter stories by broad content categories such as Politics, Tech, Sports,
	// Business, or Finance. Use 'none' to find uncategorized stories. Categories are
	// derived from the articles within each story. Multiple values create an OR
	// filter.
	Category []string `query:"category,omitzero" json:"-"`
	// Filter local news by city. Applies only to local news, when this param is passed
	// non-local news will not be returned. If multiple parameters are passed, they
	// will be applied as OR operations.
	City []string `query:"city,omitzero" json:"-"`
	// Filter to specific stories using their unique identifiers. Each clusterId
	// represents a distinct story that groups related articles. Multiple values create
	// an OR filter.
	ClusterID []string `query:"clusterId,omitzero" json:"-"`
	// Filter stories by domains of top mentioned companies (e.g., apple.com). Returns
	// stories where companies with these domains appear prominently. For available
	// company entities, consult the /companies endpoint.
	CompanyDomain []string `query:"companyDomain,omitzero" json:"-"`
	// Filter stories by identifiers of top mentioned companies. Returns stories where
	// these companies appear prominently. For a complete list of tracked companies,
	// refer to the /companies endpoint.
	CompanyID []string `query:"companyId,omitzero" json:"-"`
	// Filter stories by stock symbols of top mentioned companies. Returns stories
	// where companies with these symbols appear prominently. For available company
	// entities and their symbols, consult the /companies endpoint.
	CompanySymbol []string `query:"companySymbol,omitzero" json:"-"`
	// Country code to filter by country. If multiple parameters are passed, they will
	// be applied as OR operations.
	Country []string `query:"country,omitzero" json:"-"`
	// Excludes specific stories from the results by their unique identifiers. Use this
	// parameter to filter out unwanted or previously seen stories.
	ExcludeClusterID []string `query:"excludeClusterId,omitzero" json:"-"`
	// Filter stories by Wikidata IDs of top mentioned people. Returns stories where
	// these individuals appear prominently. Refer to the /people endpoint for a
	// complete list of tracked individuals.
	PersonWikidataID []string `query:"personWikidataId,omitzero" json:"-"`
	// Determines the story sorting order. Options include createdAt (default, when
	// stories first emerged), updatedAt (when stories received new articles, best for
	// tracking developing events), relevance (best match to query), count (by unique
	// article count), and totalCount (by total article count including reprints).
	//
	// Any of "createdAt", "updatedAt", "relevance", "count", "totalCount".
	SortBy SortBy `query:"sortBy,omitzero" json:"-"`
	// Filter stories that contain articles from specific publisher domains or
	// subdomains. Supports wildcards (_ and ?) for pattern matching (e.g., _.cnn.com).
	// A story will match if it contains at least one article from any of the specified
	// sources. Multiple values create an OR filter.
	Source []string `query:"source,omitzero" json:"-"`
	// Filter stories that contain articles from publishers in Perigon's curated
	// bundles (e.g., top100, top25crypto). A story will match if it contains at least
	// one article from any publisher in the specified bundles. Multiple values create
	// an OR filter.
	SourceGroup []string `query:"sourceGroup,omitzero" json:"-"`
	// Filter local news by state. Applies only to local news, when this param is
	// passed non-local news will not be returned. If multiple parameters are passed,
	// they will be applied as OR operations.
	State []string `query:"state,omitzero" json:"-"`
	// Filter stories by Google Content Categories. Must pass the full hierarchical
	// path of the category. Example:
	// taxonomy=/Finance/Banking/Other,/Finance/Investing/Funds. Stories are
	// categorized based on their constituent articles. Multiple values create an OR
	// filter.
	Taxonomy []string `query:"taxonomy,omitzero" json:"-"`
	// Filter stories by specific topics such as Markets, Crime, Cryptocurrency, or
	// College Sports. Topics are more granular than categories, and stories can
	// include multiple topics based on their constituent articles. Use the /topics
	// endpoint for a complete list of available topics. Multiple values create an OR
	// filter.
	Topic []string `query:"topic,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [StoryListParams]'s query parameters as `url.Values`.
func (r StoryListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
