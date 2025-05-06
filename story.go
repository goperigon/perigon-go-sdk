// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package perigon

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
	"github.com/goperigon/perigon-go-sdk/packages/respjson"
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
	ID                string                   `json:"id,nullable"`
	Categories        []NewsClusterCategory    `json:"categories,nullable"`
	Companies         []NewsClusterCompany     `json:"companies,nullable"`
	Countries         []NewsClusterCountry     `json:"countries,nullable"`
	CreatedAt         string                   `json:"createdAt,nullable"`
	DuplicateOf       string                   `json:"duplicateOf,nullable"`
	Highlights        map[string][]string      `json:"highlights,nullable"`
	ImageSource       NewsClusterImageSource   `json:"imageSource"`
	ImageURL          string                   `json:"imageUrl,nullable"`
	InitializedAt     string                   `json:"initializedAt,nullable"`
	KeyPoints         []NewsClusterKeyPoint    `json:"keyPoints,nullable"`
	Locations         []NewsClusterLocation    `json:"locations,nullable"`
	Name              string                   `json:"name,nullable"`
	People            []NewsClusterPerson      `json:"people,nullable"`
	Questions         []NewsClusterQuestion    `json:"questions,nullable"`
	ReprintCount      int64                    `json:"reprintCount,nullable"`
	SelectedArticles  []Article                `json:"selectedArticles,nullable"`
	Sentiment         NewsClusterSentiment     `json:"sentiment"`
	ShortSummary      string                   `json:"shortSummary,nullable"`
	Slug              string                   `json:"slug,nullable"`
	Summary           string                   `json:"summary,nullable"`
	SummaryReferences []string                 `json:"summaryReferences,nullable"`
	Taxonomies        []NewsClusterTaxonomy    `json:"taxonomies,nullable"`
	TopCategories     []NewsClusterTopCategory `json:"topCategories,nullable"`
	TopCompanies      []NewsClusterTopCompany  `json:"topCompanies,nullable"`
	TopCountries      []string                 `json:"topCountries,nullable"`
	Topics            []NewsClusterTopic       `json:"topics,nullable"`
	TopLocations      []NewsClusterTopLocation `json:"topLocations,nullable"`
	TopPeople         []NewsClusterTopPerson   `json:"topPeople,nullable"`
	TopTaxonomies     []NewsClusterTopTaxonomy `json:"topTaxonomies,nullable"`
	TopTopics         []NewsClusterTopTopic    `json:"topTopics,nullable"`
	TotalCount        int64                    `json:"totalCount,nullable"`
	UniqueCount       int64                    `json:"uniqueCount,nullable"`
	UniqueSources     []string                 `json:"uniqueSources,nullable"`
	UpdatedAt         string                   `json:"updatedAt,nullable"`
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

type NewsClusterCategory struct {
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
func (r NewsClusterCategory) RawJSON() string { return r.JSON.raw }
func (r *NewsClusterCategory) UnmarshalJSON(data []byte) error {
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

type NewsClusterCountry struct {
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
func (r NewsClusterCountry) RawJSON() string { return r.JSON.raw }
func (r *NewsClusterCountry) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NewsClusterImageSource struct {
	Domain   string                         `json:"domain,nullable"`
	Location NewsClusterImageSourceLocation `json:"location"`
	Paywall  bool                           `json:"paywall,nullable"`
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

type NewsClusterImageSourceLocation struct {
	City        string                                    `json:"city,nullable"`
	Coordinates NewsClusterImageSourceLocationCoordinates `json:"coordinates"`
	Country     string                                    `json:"country,nullable"`
	County      string                                    `json:"county,nullable"`
	State       string                                    `json:"state,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		City        respjson.Field
		Coordinates respjson.Field
		Country     respjson.Field
		County      respjson.Field
		State       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NewsClusterImageSourceLocation) RawJSON() string { return r.JSON.raw }
func (r *NewsClusterImageSourceLocation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NewsClusterImageSourceLocationCoordinates struct {
	Lat float64 `json:"lat,nullable"`
	Lon float64 `json:"lon,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Lat         respjson.Field
		Lon         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NewsClusterImageSourceLocationCoordinates) RawJSON() string { return r.JSON.raw }
func (r *NewsClusterImageSourceLocationCoordinates) UnmarshalJSON(data []byte) error {
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

type NewsClusterTaxonomy struct {
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
func (r NewsClusterTaxonomy) RawJSON() string { return r.JSON.raw }
func (r *NewsClusterTaxonomy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NewsClusterTopCategory struct {
	Name string `json:"name,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NewsClusterTopCategory) RawJSON() string { return r.JSON.raw }
func (r *NewsClusterTopCategory) UnmarshalJSON(data []byte) error {
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

type NewsClusterTopic struct {
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
func (r NewsClusterTopic) RawJSON() string { return r.JSON.raw }
func (r *NewsClusterTopic) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NewsClusterTopLocation struct {
	Area    string `json:"area,nullable"`
	City    string `json:"city,nullable"`
	Country string `json:"country,nullable"`
	County  string `json:"county,nullable"`
	State   string `json:"state,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Area        respjson.Field
		City        respjson.Field
		Country     respjson.Field
		County      respjson.Field
		State       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NewsClusterTopLocation) RawJSON() string { return r.JSON.raw }
func (r *NewsClusterTopLocation) UnmarshalJSON(data []byte) error {
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

type NewsClusterTopTaxonomy struct {
	Name string `json:"name,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NewsClusterTopTaxonomy) RawJSON() string { return r.JSON.raw }
func (r *NewsClusterTopTaxonomy) UnmarshalJSON(data []byte) error {
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

// Story search result
type StoryListResponse struct {
	NumResults int64         `json:"numResults,nullable"`
	Results    []NewsCluster `json:"results,nullable"`
	Status     int64         `json:"status,nullable"`
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
	// List of company names for filtering. Filtering is applied on topCompanies field.
	CompanyName param.Opt[string] `query:"companyName,omitzero" json:"-"`
	// 'from' filter, will search stories created after the specified date, the date
	// could be passed as ISO or 'yyyy-mm-dd'. Add time in ISO format, ie.
	// 2023-03-01T00:00:00
	From param.Opt[time.Time] `query:"from,omitzero" format:"date-time" json:"-"`
	// 'initializedFrom' filter, will search stories that became available after
	// provided date
	InitializedFrom param.Opt[time.Time] `query:"initializedFrom,omitzero" format:"date-time" json:"-"`
	// 'initializedTo' filter, will search stories that became available before
	// provided date
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
	// Search story by name. Supports complex query syntax, same way as q parameter
	// from /all endpoint.
	Name param.Opt[string] `query:"name,omitzero" json:"-"`
	// Returns stories with name assigned. Defaults to true.
	NameExists param.Opt[bool] `query:"nameExists,omitzero" json:"-"`
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
	// List of people names. Filtering is applied on topPeople field.
	PersonName param.Opt[string] `query:"personName,omitzero" json:"-"`
	// Filters results with a sentiment score greater than or equal to the specified
	// value, indicating positive sentiment. See the Article Data section in Docs for
	// an explanation of scores.
	PositiveSentimentFrom param.Opt[float64] `query:"positiveSentimentFrom,omitzero" json:"-"`
	// Filters results with a sentiment score less than or equal to the specified
	// value, indicating positive sentiment. See the Article Data section in Docs for
	// an explanation of scores.
	PositiveSentimentTo param.Opt[float64] `query:"positiveSentimentTo,omitzero" json:"-"`
	// Search story by name, summary and key points. Preference is given to the name
	// field. Supports complex query syntax, same way as q parameter from /all
	// endpoint.
	Q param.Opt[string] `query:"q,omitzero" json:"-"`
	// Stories are deduplicated by default. If a story is deduplicated, all future
	// articles are merged into the original story. duplicateOf field contains the
	// original cluster Id. When showDuplicates=true, all stories are shown.
	ShowDuplicates param.Opt[bool] `query:"showDuplicates,omitzero" json:"-"`
	// Show total number of results. By default set to false, will cap result count
	// at 10000.
	ShowNumResults    param.Opt[bool] `query:"showNumResults,omitzero" json:"-"`
	ShowStoryPageInfo param.Opt[bool] `query:"showStoryPageInfo,omitzero" json:"-"`
	// The number of items per page.
	Size param.Opt[int64] `query:"size,omitzero" json:"-"`
	// 'to' filter, will search stories created before the specified date, the date
	// could be passed as ISO or 'yyyy-mm-dd'. Add time in ISO format, ie.
	// 2023-03-01T23:59:59
	To param.Opt[time.Time] `query:"to,omitzero" format:"date-time" json:"-"`
	// Will return stories with 'updatedAt' >= 'updatedFrom'.
	UpdatedFrom param.Opt[time.Time] `query:"updatedFrom,omitzero" format:"date-time" json:"-"`
	// Will return stories with 'updatedAt' <= 'updatedTo'.
	UpdatedTo param.Opt[time.Time] `query:"updatedTo,omitzero" format:"date-time" json:"-"`
	// Filter local news by area. Applies only to local news, when this param is passed
	// non-local news will not be returned. If multiple parameters are passed, they
	// will be applied as OR operations.
	Area []string `query:"area,omitzero" json:"-"`
	// Filter by categories. Categories are general themes that the article is about.
	// Examples of categories: Tech, Politics, etc. If multiple parameters are passed,
	// they will be applied as OR operations. Use 'none' to search uncategorized
	// articles. More ➜
	Category []string `query:"category,omitzero" json:"-"`
	// Filter local news by city. Applies only to local news, when this param is passed
	// non-local news will not be returned. If multiple parameters are passed, they
	// will be applied as OR operations.
	City []string `query:"city,omitzero" json:"-"`
	// Filter to specific story. Passing a cluster ID will filter results to only the
	// content found within the cluster. Multiple params could be passed.
	ClusterID []string `query:"clusterId,omitzero" json:"-"`
	// List of company domains for filtering. Filtering is applied on topCompanies
	// field.
	CompanyDomain []string `query:"companyDomain,omitzero" json:"-"`
	// List of company IDs for filtering. Filtering is applied to topCompanies field.
	CompanyID []string `query:"companyId,omitzero" json:"-"`
	// List of company tickers for filtering. Filtering is applied on topCompanies
	// field.
	CompanySymbol []string `query:"companySymbol,omitzero" json:"-"`
	// Country code to filter by country. If multiple parameters are passed, they will
	// be applied as OR operations.
	Country []string `query:"country,omitzero" json:"-"`
	// Excludes specific stories from the results by their unique identifiers. Use this
	// parameter to filter out unwanted or previously seen stories.
	ExcludeClusterID []string `query:"excludeClusterId,omitzero" json:"-"`
	// List of person Wikidata IDs for filtering. Filter is applied on topPeople field.
	PersonWikidataID []string `query:"personWikidataId,omitzero" json:"-"`
	// Sort stories by count ('count'), total count ('totalCount'), creation date
	// ('createdAt'), last updated date ('updatedAt'), or relevance ('relevance'). By
	// default is sorted by 'createdAt'
	//
	// Any of "createdAt", "updatedAt", "relevance", "count", "totalCount".
	SortBy SortBy `query:"sortBy,omitzero" json:"-"`
	// Filter stories by sources that wrote articles belonging to this story. At least
	// 1 article is required for story to match. Multiple parameters could be passed.
	Source []string `query:"source,omitzero" json:"-"`
	// Filter stories by sources that wrote articles belonging to this story. Source
	// groups are expanded into a list of sources. At least 1 article by the source is
	// required for story to match. Multiple params could be passed.
	SourceGroup []string `query:"sourceGroup,omitzero" json:"-"`
	// Filter local news by state. Applies only to local news, when this param is
	// passed non-local news will not be returned. If multiple parameters are passed,
	// they will be applied as OR operations.
	State []string `query:"state,omitzero" json:"-"`
	// Filters by Google Content Categories. This field will accept 1 or more
	// categories, must pass the full name of the category. Example:
	// taxonomy=/Finance/Banking/Other, /Finance/Investing/Funds
	Taxonomy []string `query:"taxonomy,omitzero" json:"-"`
	// Filter by topics. Each topic is some kind of entity that the article is about.
	// Examples of topics: Markets, Joe Biden, Green Energy, Climate Change,
	// Cryptocurrency, etc. If multiple parameters are passed, they will be applied as
	// OR operations.
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
