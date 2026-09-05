// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package perigon

import (
	"context"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/goperigon/perigon-go-sdk/v2/internal/apijson"
	"github.com/goperigon/perigon-go-sdk/v2/internal/apiquery"
	"github.com/goperigon/perigon-go-sdk/v2/internal/requestconfig"
	"github.com/goperigon/perigon-go-sdk/v2/option"
	"github.com/goperigon/perigon-go-sdk/v2/packages/param"
	"github.com/goperigon/perigon-go-sdk/v2/packages/respjson"
	"github.com/goperigon/perigon-go-sdk/v2/shared"
)

// Core endpoints for the Perigon News API v1, providing access to aggregated news
// stories, articles, and related content. These endpoints enable searching,
// filtering, and retrieving media content across multiple sources.
//
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

// Track evolving narratives and top news headlines with the '**Stories**' endpoint
// — it returns structured clusters of related articles with summaries, key points,
// sentiment scores, and metadata on top people, companies, topics, countries, and
// categories. Use it to monitor how news stories develop over time, who and what
// they impact, and where they're gaining traction globally.
//
// Each story represents a cluster of related articles grouped during processing.
// Articles are assigned to a single story, identified by the `clusterId` field.
// Once a story reaches five articles, a story title is automatically generated
// based on its content. Stories update as new coverage is processed and key
// details evolve.
//
// To fetch all articles within a specific story, query the `/articles/all`
// endpoint using the `clusterId` parameter.
func (r *StoryService) List(ctx context.Context, query StoryListParams, opts ...option.RequestOption) (res *StoryListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/stories/all"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Expanded story cluster details, when requested and available.
type NewsCluster struct {
	// Unique identifier for the story cluster.
	ID string `json:"id" api:"nullable"`
	// Categories associated with the story and their article counts.
	Categories []RecordStatHolder `json:"categories" api:"nullable"`
	// Companies associated with the story and their article counts.
	Companies []NewsClusterCompany `json:"companies" api:"nullable"`
	// Countries associated with the story and their article counts.
	Countries []RecordStatHolder `json:"countries" api:"nullable"`
	// Date and time the story was created, in ISO 8601 format.
	CreatedAt string `json:"createdAt" api:"nullable"`
	// Identifier of the original story when this story is a duplicate.
	DuplicateOf string `json:"duplicateOf" api:"nullable"`
	// Matched text fragments grouped by story field.
	Highlights map[string][]string `json:"highlights" api:"nullable"`
	// Publisher information for the selected story image.
	ImageSource NewsClusterImageSource `json:"imageSource" api:"nullable"`
	// Primary image URL selected for the story.
	ImageURL string `json:"imageUrl" api:"nullable"`
	// Date and time the story became initialized and visible, in ISO 8601 format.
	InitializedAt string `json:"initializedAt" api:"nullable"`
	// Key facts and developments identified in the story.
	KeyPoints []NewsClusterKeyPoint `json:"keyPoints" api:"nullable"`
	// Locations associated with the story and their article counts.
	Locations []NewsClusterLocation `json:"locations" api:"nullable"`
	// Generated story title.
	Name string `json:"name" api:"nullable"`
	// People associated with the story and their article counts.
	People []NewsClusterPerson `json:"people" api:"nullable"`
	// Questions and answers generated from the story content.
	Questions []NewsClusterQuestion `json:"questions" api:"nullable"`
	// Number of reprinted articles in the story.
	ReprintCount int64 `json:"reprintCount" api:"nullable"`
	// Representative articles selected for the story.
	SelectedArticles []Article `json:"selectedArticles" api:"nullable"`
	// Aggregate sentiment scores for the story.
	Sentiment NewsClusterSentiment `json:"sentiment" api:"nullable"`
	// Short generated story summary, when available.
	ShortSummary string `json:"shortSummary" api:"nullable"`
	// URL-safe identifier for the story.
	Slug string `json:"slug" api:"nullable"`
	// Ratio of unique publisher domains to total articles in the story.
	SourceDiversity float64 `json:"sourceDiversity" api:"nullable"`
	// Generated story summary.
	Summary string `json:"summary" api:"nullable"`
	// Article IDs supporting the story summary.
	SummaryReferences []string `json:"summaryReferences" api:"nullable"`
	// Google Content Categories associated with the story and their article counts.
	Taxonomies []RecordStatHolder `json:"taxonomies" api:"nullable"`
	// Categories most strongly associated with the story.
	TopCategories []shared.CategoryHolder `json:"topCategories" api:"nullable"`
	// Companies most strongly associated with the story.
	TopCompanies []NewsClusterTopCompany `json:"topCompanies" api:"nullable"`
	// Countries most strongly associated with the story.
	TopCountries []string `json:"topCountries" api:"nullable"`
	// Topics associated with the story and their article counts.
	Topics []RecordStatHolder `json:"topics" api:"nullable"`
	// Locations most strongly associated with the story.
	TopLocations []shared.LocationHolder `json:"topLocations" api:"nullable"`
	// People most strongly associated with the story.
	TopPeople []NewsClusterTopPerson `json:"topPeople" api:"nullable"`
	// Google Content Categories most strongly associated with the story.
	TopTaxonomies []shared.CategoryHolder `json:"topTaxonomies" api:"nullable"`
	// Topics most strongly associated with the story.
	TopTopics []NewsClusterTopTopic `json:"topTopics" api:"nullable"`
	// Total article count including reprints.
	TotalCount int64 `json:"totalCount" api:"nullable"`
	// Number of unique articles in the story.
	UniqueCount int64 `json:"uniqueCount" api:"nullable"`
	// Distinct publisher domains represented in the story.
	UniqueSources []string `json:"uniqueSources" api:"nullable"`
	// Date and time the story was last updated, in ISO 8601 format.
	UpdatedAt string `json:"updatedAt" api:"nullable"`
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
		SourceDiversity   respjson.Field
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

// Companies associated with the story and their article counts.
type NewsClusterCompany struct {
	// Unique Perigon identifier for the company.
	ID string `json:"id" api:"nullable"`
	// Number of story articles mentioning the company.
	Count int64 `json:"count" api:"nullable"`
	// Domains associated with the company.
	Domains []string `json:"domains" api:"nullable"`
	// Company name.
	Name string `json:"name" api:"nullable"`
	// Stock symbols associated with the company.
	Symbols []string `json:"symbols" api:"nullable"`
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

// Publisher information for the selected story image.
type NewsClusterImageSource struct {
	// Publisher domain.
	Domain string `json:"domain" api:"nullable"`
	// Geographic location of the publisher, when available.
	Location shared.SourceLocation `json:"location" api:"nullable"`
	// Whether the publisher uses a paywall, when known.
	Paywall bool `json:"paywall" api:"nullable"`
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

// Key facts and developments identified in the story.
type NewsClusterKeyPoint struct {
	// A key fact or development identified in the content.
	Point string `json:"point" api:"nullable"`
	// Article IDs supporting the key point.
	References []string `json:"references" api:"nullable"`
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

// Locations associated with the story and their article counts.
type NewsClusterLocation struct {
	// Neighborhood, borough, or district associated with the story.
	Area string `json:"area" api:"nullable"`
	// City associated with the story.
	City string `json:"city" api:"nullable"`
	// Number of story articles associated with the location.
	Count int64 `json:"count" api:"nullable"`
	// Country associated with the story.
	Country string `json:"country" api:"nullable"`
	// County associated with the story.
	County string `json:"county" api:"nullable"`
	// State or region associated with the story.
	State string `json:"state" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Area        respjson.Field
		City        respjson.Field
		Count       respjson.Field
		Country     respjson.Field
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

// People associated with the story and their article counts.
type NewsClusterPerson struct {
	// Number of story articles mentioning the person.
	Count int64 `json:"count" api:"nullable"`
	// Person name.
	Name string `json:"name" api:"nullable"`
	// Wikidata identifier for the person.
	WikidataID string `json:"wikidataId" api:"nullable"`
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

// Questions and answers generated from the story content.
type NewsClusterQuestion struct {
	// Answer generated from the story content.
	Answer string `json:"answer" api:"nullable"`
	// Question generated from the story content.
	Question string `json:"question" api:"nullable"`
	// Article IDs supporting the answer.
	References []string `json:"references" api:"nullable"`
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

// Aggregate sentiment scores for the story.
type NewsClusterSentiment struct {
	// Negative sentiment score from 0 to 1.
	Negative float64 `json:"negative" api:"nullable"`
	// Neutral sentiment score from 0 to 1.
	Neutral float64 `json:"neutral" api:"nullable"`
	// Positive sentiment score from 0 to 1.
	Positive float64 `json:"positive" api:"nullable"`
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

// Companies most strongly associated with the story.
type NewsClusterTopCompany struct {
	// Unique Perigon identifier for the company.
	ID string `json:"id" api:"nullable"`
	// Domains associated with the company.
	Domains []string `json:"domains" api:"nullable"`
	// Company name.
	Name string `json:"name" api:"nullable"`
	// Stock symbols associated with the company.
	Symbols []string `json:"symbols" api:"nullable"`
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

// People most strongly associated with the story.
type NewsClusterTopPerson struct {
	// Person name.
	Name string `json:"name" api:"nullable"`
	// Wikidata identifier for the person.
	WikidataID string `json:"wikidataId" api:"nullable"`
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

// Topics most strongly associated with the story.
type NewsClusterTopTopic struct {
	// Topic assigned to the article.
	Name string `json:"name" api:"nullable"`
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

// Google Content Categories associated with the story and their article counts.
type RecordStatHolder struct {
	// Number of matching articles for the value.
	Count int64 `json:"count" api:"nullable"`
	// Name of the aggregated value.
	Name string `json:"name" api:"nullable"`
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
	NumResults int64         `json:"numResults" api:"required"`
	Results    []NewsCluster `json:"results" api:"required"`
	Status     int64         `json:"status" api:"required"`
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
	// String. Filter stories by names of top mentioned companies. Performs an exact
	// match on company names in the topCompanies field.
	CompanyName param.Opt[string] `query:"companyName,omitzero" json:"-"`
	// Boolean. Preview 5 articles from the cluster.
	ExpandArticles param.Opt[bool] `query:"expandArticles,omitzero" json:"-"`
	// Filter for stories created on or after this date. Accepts ISO 8601 or yyyy-mm-dd
	// format.
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
	// Filter for stories initialized on or after this date. A story is initialized
	// when it becomes visible with generated story data.
	InitializedFrom param.Opt[time.Time] `query:"initializedFrom,omitzero" format:"date-time" json:"-"`
	// Filter for stories initialized on or before this date. A story is initialized
	// when it becomes visible with generated story data.
	InitializedTo param.Opt[time.Time] `query:"initializedTo,omitzero" format:"date-time" json:"-"`
	// Integer. Filter by maximum cluster size. Maximum cluster size filter applies to
	// number of unique articles in the cluster.
	MaxClusterSize param.Opt[int64] `query:"maxClusterSize,omitzero" json:"-"`
	// Integer. Filter by minimum cluster size. Minimum cluster size filter applies to
	// number of unique articles.
	MinClusterSize param.Opt[int64] `query:"minClusterSize,omitzero" json:"-"`
	// Float. Minimum ratio of unique sources to unique articles (uniqueSources /
	// uniqueCount). Filters out stories dominated by a single publisher. For example,
	// a value of 0.05 requires at least 1 unique source per 20 articles. Not applied
	// by default.
	MinSourceDiversity param.Opt[float64] `query:"minSourceDiversity,omitzero" json:"-"`
	// Integer. Specifies the minimum number of unique sources required for a story to
	// appear in results. Higher values return more significant stories covered by
	// multiple publications. Default is 3.
	MinUniqueSources param.Opt[int64] `query:"minUniqueSources,omitzero" json:"-"`
	// String. Search specifically within story names. Supports Boolean operators,
	// exact phrases with quotes, and wildcards for matching name variations.
	Name param.Opt[string] `query:"name,omitzero" json:"-"`
	// Boolean. Filter to only include stories that have been assigned names. Defaults
	// to true. Note that stories only receive names after they contain at least 5
	// unique articles.
	NameExists param.Opt[bool] `query:"nameExists,omitzero" json:"-"`
	// Filter stories with an aggregate negative sentiment score greater than or equal
	// to the specified value. Scores range from 0 to 1, with higher values indicating
	// stronger negative tone.
	NegativeSentimentFrom param.Opt[float64] `query:"negativeSentimentFrom,omitzero" json:"-"`
	// Filter stories with an aggregate negative sentiment score less than or equal to
	// the specified value. Scores range from 0 to 1, with higher values indicating
	// stronger negative tone.
	NegativeSentimentTo param.Opt[float64] `query:"negativeSentimentTo,omitzero" json:"-"`
	// Filter stories with an aggregate neutral sentiment score greater than or equal
	// to the specified value. Scores range from 0 to 1, with higher values indicating
	// stronger neutral tone.
	NeutralSentimentFrom param.Opt[float64] `query:"neutralSentimentFrom,omitzero" json:"-"`
	// Filter stories with an aggregate neutral sentiment score less than or equal to
	// the specified value. Scores range from 0 to 1, with higher values indicating
	// stronger neutral tone.
	NeutralSentimentTo param.Opt[float64] `query:"neutralSentimentTo,omitzero" json:"-"`
	// Integer. The specific page of results to retrieve in the paginated response.
	// Starts at 0.
	Page param.Opt[int64] `query:"page,omitzero" json:"-"`
	// String. Filter stories by exact name matches of top mentioned people. Does not
	// support Boolean or complex logic. For available person entities, consult the
	// /people endpoint.
	PersonName param.Opt[string] `query:"personName,omitzero" json:"-"`
	// Filter stories with an aggregate positive sentiment score greater than or equal
	// to the specified value. Scores range from 0 to 1, with higher values indicating
	// stronger positive tone.
	PositiveSentimentFrom param.Opt[float64] `query:"positiveSentimentFrom,omitzero" json:"-"`
	// Filter stories with an aggregate positive sentiment score less than or equal to
	// the specified value. Scores range from 0 to 1, with higher values indicating
	// stronger positive tone.
	PositiveSentimentTo param.Opt[float64] `query:"positiveSentimentTo,omitzero" json:"-"`
	// String. Primary search query for filtering stories based on their name, summary,
	// and key points. Supports Boolean operators (AND, OR, NOT), exact phrases with
	// quotes, and wildcards (\* and ?) for flexible searching.
	Q param.Opt[string] `query:"q,omitzero" json:"-"`
	// Boolean. Stories are deduplicated by default. If a story is deduplicated, all
	// future articles are merged into the original story. duplicateOf field contains
	// the original cluster Id. When showDuplicates=true, all stories are shown.
	ShowDuplicates param.Opt[bool] `query:"showDuplicates,omitzero" json:"-"`
	// Boolean. When set to true, enables text highlighting in search results.
	ShowHighlighting param.Opt[bool] `query:"showHighlighting,omitzero" json:"-"`
	// Boolean. Show total number of results. By default set to false, will cap result
	// count at 10000.
	ShowNumResults param.Opt[bool] `query:"showNumResults,omitzero" json:"-"`
	// Boolean.
	ShowStoryPageInfo param.Opt[bool] `query:"showStoryPageInfo,omitzero" json:"-"`
	// Integer. The number of articles to return per page in the paginated response.
	Size param.Opt[int64] `query:"size,omitzero" json:"-"`
	// Filter for stories created on or before this date. Accepts ISO 8601 or
	// yyyy-mm-dd format.
	To param.Opt[time.Time] `query:"to,omitzero" format:"date-time" json:"-"`
	// Date. Filter for stories that received new articles after this date. Useful for
	// tracking developing news events or evolving storylines.
	UpdatedFrom param.Opt[time.Time] `query:"updatedFrom,omitzero" format:"date-time" json:"-"`
	// Date. Filter for stories that received new articles before this date. Useful for
	// tracking developing news events or evolving storylines.
	UpdatedTo param.Opt[time.Time] `query:"updatedTo,omitzero" format:"date-time" json:"-"`
	// String Array. Filter local news by area. Applies only to local news, when this
	// param is passed non-local news will not be returned. If multiple parameters are
	// passed, they will be applied as OR operations.
	Area []string `query:"area,omitzero" json:"-"`
	// String Array. Filter stories by broad content categories such as Politics, Tech,
	// Sports, Business, or Finance. Use 'none' to find uncategorized stories.
	// Categories are derived from the articles within each story. Multiple values
	// create an OR filter.
	Category []string `query:"category,omitzero" json:"-"`
	// String Array. Filter local news by city. Applies only to local news, when this
	// param is passed non-local news will not be returned. If multiple parameters are
	// passed, they will be applied as OR operations.
	City []string `query:"city,omitzero" json:"-"`
	// String Array. Filter to specific stories using their unique identifiers. Each
	// clusterId represents a distinct story that groups related articles. Multiple
	// values create an OR filter.
	ClusterID []string `query:"clusterId,omitzero" json:"-"`
	// String Array. Filter stories by domains of top mentioned companies (e.g.,
	// apple.com). Returns stories where companies with these domains appear
	// prominently. For available company entities, consult the /companies endpoint.
	CompanyDomain []string `query:"companyDomain,omitzero" json:"-"`
	// String Array. Filter stories by identifiers of top mentioned companies. Returns
	// stories where these companies appear prominently. For a complete list of tracked
	// companies, refer to the /companies endpoint.
	CompanyID []string `query:"companyId,omitzero" json:"-"`
	// String Array. Filter stories by ISIN codes of top mentioned companies
	// (International Securities Identification Numbers). Returns stories where
	// companies with these ISINs appear prominently. For available company entities
	// and their ISINs, consult the /companies endpoint.
	CompanyIsin []string `query:"companyIsin,omitzero" json:"-"`
	// String Array. Filter stories by stock symbols of top mentioned companies.
	// Returns stories where companies with these symbols appear prominently. For
	// available company entities and their symbols, consult the /companies endpoint.
	CompanySymbol []string `query:"companySymbol,omitzero" json:"-"`
	// String Array. Filter stories by ISIN codes on ticker listings of top mentioned
	// companies (symbols.isin). Distinct from companyIsin, which matches the
	// company-level ISIN. Returns stories where those companies appear prominently.
	// For available company entities and their listing ISINs, consult the /companies
	// endpoint.
	CompanySymbolIsin []string `query:"companySymbolIsin,omitzero" json:"-"`
	// String Array. Country code to filter by country. If multiple parameters are
	// passed, they will be applied as OR operations.
	Country []string `query:"country,omitzero" json:"-"`
	// String Array. Excludes specific stories from the results by their unique
	// identifiers. Use this parameter to filter out unwanted or previously seen
	// stories.
	ExcludeClusterID []string `query:"excludeClusterId,omitzero" json:"-"`
	// String Array. Filter stories by Wikidata IDs of top mentioned people. Returns
	// stories where these individuals appear prominently. Refer to the /people
	// endpoint for a complete list of tracked individuals.
	PersonWikidataID []string `query:"personWikidataId,omitzero" json:"-"`
	// String. Determines the story sorting order. Options include createdAt (default,
	// when stories first emerged), updatedAt (when stories received new articles, best
	// for tracking developing events), relevance (best match to query), count (by
	// unique article count), and totalCount (by total article count including
	// reprints).
	//
	// Any of "createdAt", "updatedAt", "relevance", "count", "totalCount".
	SortBy StoryListParamsSortBy `query:"sortBy,omitzero" json:"-"`
	// String Array. Filter stories that contain articles from specific publisher
	// domains or subdomains. Supports wildcards (_ and ?) for pattern matching (e.g.,
	// _.cnn.com). A story will match if it contains at least one article from any of
	// the specified sources. Multiple values create an OR filter.
	Source []string `query:"source,omitzero" json:"-"`
	// String Array. Filter stories that contain articles from publishers in Perigon's
	// curated bundles (e.g., top100, top25crypto). A story will match if it contains
	// at least one article from any publisher in the specified bundles. Multiple
	// values create an OR filter.
	SourceGroup []string `query:"sourceGroup,omitzero" json:"-"`
	// String Array. Filter local news by state. Applies only to local news, when this
	// param is passed non-local news will not be returned. If multiple parameters are
	// passed, they will be applied as OR operations.
	State []string `query:"state,omitzero" json:"-"`
	// String Array. Filter stories by Google Content Categories. Must pass the full
	// hierarchical path of the category. Example:
	// taxonomy=/Finance/Banking/Other,/Finance/Investing/Funds. Stories are
	// categorized based on their constituent articles. Multiple values create an OR
	// filter.
	Taxonomy []string `query:"taxonomy,omitzero" json:"-"`
	// String Array. Filter stories by specific topics such as Markets, Crime,
	// Cryptocurrency, or College Sports. Topics are more granular than categories, and
	// stories can include multiple topics based on their constituent articles. Use the
	// /topics endpoint for a complete list of available topics. Multiple values create
	// an OR filter.
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

// String. Determines the story sorting order. Options include createdAt (default,
// when stories first emerged), updatedAt (when stories received new articles, best
// for tracking developing events), relevance (best match to query), count (by
// unique article count), and totalCount (by total article count including
// reprints).
type StoryListParamsSortBy string

const (
	StoryListParamsSortByCreatedAt  StoryListParamsSortBy = "createdAt"
	StoryListParamsSortByUpdatedAt  StoryListParamsSortBy = "updatedAt"
	StoryListParamsSortByRelevance  StoryListParamsSortBy = "relevance"
	StoryListParamsSortByCount      StoryListParamsSortBy = "count"
	StoryListParamsSortByTotalCount StoryListParamsSortBy = "totalCount"
)
