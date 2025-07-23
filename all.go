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

// AllService contains methods and other services that help with interacting with
// the perigon API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAllService] method instead.
type AllService struct {
	Options []option.RequestOption
}

// NewAllService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAllService(opts ...option.RequestOption) (r AllService) {
	r = AllService{}
	r.Options = opts
	return
}

// Search and filter all news articles available via the Perigon API. The result
// includes a list of individual articles that were matched to your specific
// criteria.
func (r *AllService) List(ctx context.Context, query AllListParams, opts ...option.RequestOption) (res *AllListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/articles/all"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type AllEndpointSortBy string

const (
	AllEndpointSortByRelevance      AllEndpointSortBy = "relevance"
	AllEndpointSortByDate           AllEndpointSortBy = "date"
	AllEndpointSortByReverseDate    AllEndpointSortBy = "reverseDate"
	AllEndpointSortByReverseAddDate AllEndpointSortBy = "reverseAddDate"
	AllEndpointSortByAddDate        AllEndpointSortBy = "addDate"
	AllEndpointSortByPubDate        AllEndpointSortBy = "pubDate"
	AllEndpointSortByRefreshDate    AllEndpointSortBy = "refreshDate"
)

type Article struct {
	AddDate               string                  `json:"addDate,nullable"`
	ArticleID             string                  `json:"articleId,nullable"`
	AuthorsByline         string                  `json:"authorsByline,nullable"`
	Categories            []shared.CategoryHolder `json:"categories,nullable"`
	Claim                 string                  `json:"claim,nullable"`
	Cluster               NewsCluster             `json:"cluster"`
	ClusterID             string                  `json:"clusterId,nullable"`
	Companies             []ArticleCompany        `json:"companies,nullable"`
	Content               string                  `json:"content,nullable"`
	Country               string                  `json:"country,nullable"`
	Description           string                  `json:"description,nullable"`
	Entities              []ArticleEntity         `json:"entities,nullable"`
	EventTypes            []ArticleEventType      `json:"eventTypes,nullable"`
	Highlights            map[string][]string     `json:"highlights,nullable"`
	ImageURL              string                  `json:"imageUrl,nullable"`
	Journalists           []Journalist            `json:"journalists,nullable"`
	Keywords              []ArticleKeyword        `json:"keywords,nullable"`
	Labels                []ArticleLabel          `json:"labels,nullable"`
	Language              string                  `json:"language,nullable"`
	Links                 []string                `json:"links,nullable"`
	Locations             []shared.LocationHolder `json:"locations,nullable"`
	MatchedAuthors        []ArticleMatchedAuthor  `json:"matchedAuthors,nullable"`
	Medium                string                  `json:"medium,nullable"`
	People                []ArticlePerson         `json:"people,nullable"`
	Places                []ArticlePlace          `json:"places,nullable"`
	PubDate               string                  `json:"pubDate,nullable"`
	RefreshDate           string                  `json:"refreshDate,nullable"`
	Reprint               bool                    `json:"reprint,nullable"`
	ReprintGroupID        string                  `json:"reprintGroupId,nullable"`
	Score                 float64                 `json:"score,nullable"`
	Sentiment             ArticleSentiment        `json:"sentiment"`
	ShortSummary          string                  `json:"shortSummary,nullable"`
	Source                ArticleSource           `json:"source"`
	Summary               string                  `json:"summary,nullable"`
	Taxonomies            []ArticleTaxonomy       `json:"taxonomies,nullable"`
	Title                 string                  `json:"title,nullable"`
	Topics                []ArticleTopic          `json:"topics,nullable"`
	TranslatedDescription string                  `json:"translatedDescription,nullable"`
	TranslatedSummary     string                  `json:"translatedSummary,nullable"`
	TranslatedTitle       string                  `json:"translatedTitle,nullable"`
	Translation           string                  `json:"translation,nullable"`
	URL                   string                  `json:"url,nullable"`
	Verdict               string                  `json:"verdict,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AddDate               respjson.Field
		ArticleID             respjson.Field
		AuthorsByline         respjson.Field
		Categories            respjson.Field
		Claim                 respjson.Field
		Cluster               respjson.Field
		ClusterID             respjson.Field
		Companies             respjson.Field
		Content               respjson.Field
		Country               respjson.Field
		Description           respjson.Field
		Entities              respjson.Field
		EventTypes            respjson.Field
		Highlights            respjson.Field
		ImageURL              respjson.Field
		Journalists           respjson.Field
		Keywords              respjson.Field
		Labels                respjson.Field
		Language              respjson.Field
		Links                 respjson.Field
		Locations             respjson.Field
		MatchedAuthors        respjson.Field
		Medium                respjson.Field
		People                respjson.Field
		Places                respjson.Field
		PubDate               respjson.Field
		RefreshDate           respjson.Field
		Reprint               respjson.Field
		ReprintGroupID        respjson.Field
		Score                 respjson.Field
		Sentiment             respjson.Field
		ShortSummary          respjson.Field
		Source                respjson.Field
		Summary               respjson.Field
		Taxonomies            respjson.Field
		Title                 respjson.Field
		Topics                respjson.Field
		TranslatedDescription respjson.Field
		TranslatedSummary     respjson.Field
		TranslatedTitle       respjson.Field
		Translation           respjson.Field
		URL                   respjson.Field
		Verdict               respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Article) RawJSON() string { return r.JSON.raw }
func (r *Article) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ArticleCompany struct {
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
func (r ArticleCompany) RawJSON() string { return r.JSON.raw }
func (r *ArticleCompany) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ArticleEntity struct {
	Data     string `json:"data,nullable"`
	Mentions int64  `json:"mentions,nullable"`
	Type     string `json:"type,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Mentions    respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ArticleEntity) RawJSON() string { return r.JSON.raw }
func (r *ArticleEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ArticleEventType struct {
	Name string `json:"name,nullable"`
	Type string `json:"type,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ArticleEventType) RawJSON() string { return r.JSON.raw }
func (r *ArticleEventType) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ArticleKeyword struct {
	Name   string  `json:"name,nullable"`
	Weight float64 `json:"weight,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Weight      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ArticleKeyword) RawJSON() string { return r.JSON.raw }
func (r *ArticleKeyword) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ArticleLabel struct {
	Name string `json:"name,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ArticleLabel) RawJSON() string { return r.JSON.raw }
func (r *ArticleLabel) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ArticleMatchedAuthor struct {
	ID   string `json:"id,nullable"`
	Name string `json:"name,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ArticleMatchedAuthor) RawJSON() string { return r.JSON.raw }
func (r *ArticleMatchedAuthor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ArticlePerson struct {
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
func (r ArticlePerson) RawJSON() string { return r.JSON.raw }
func (r *ArticlePerson) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ArticlePlace struct {
	Amenity       string            `json:"amenity,nullable"`
	City          string            `json:"city,nullable"`
	Coordinates   shared.Coordinate `json:"coordinates"`
	Country       string            `json:"country,nullable"`
	CountryCode   string            `json:"countryCode,nullable"`
	County        string            `json:"county,nullable"`
	Neighbourhood string            `json:"neighbourhood,nullable"`
	OsmID         string            `json:"osmId,nullable"`
	Postcode      string            `json:"postcode,nullable"`
	Quarter       string            `json:"quarter,nullable"`
	Road          string            `json:"road,nullable"`
	State         string            `json:"state,nullable"`
	StateDistrict string            `json:"stateDistrict,nullable"`
	Suburb        string            `json:"suburb,nullable"`
	Town          string            `json:"town,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amenity       respjson.Field
		City          respjson.Field
		Coordinates   respjson.Field
		Country       respjson.Field
		CountryCode   respjson.Field
		County        respjson.Field
		Neighbourhood respjson.Field
		OsmID         respjson.Field
		Postcode      respjson.Field
		Quarter       respjson.Field
		Road          respjson.Field
		State         respjson.Field
		StateDistrict respjson.Field
		Suburb        respjson.Field
		Town          respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ArticlePlace) RawJSON() string { return r.JSON.raw }
func (r *ArticlePlace) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ArticleSentiment struct {
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
func (r ArticleSentiment) RawJSON() string { return r.JSON.raw }
func (r *ArticleSentiment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ArticleSource struct {
	Domain   string                `json:"domain,nullable"`
	Location shared.SourceLocation `json:"location"`
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
func (r ArticleSource) RawJSON() string { return r.JSON.raw }
func (r *ArticleSource) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ArticleTaxonomy struct {
	Name  string  `json:"name,nullable"`
	Score float64 `json:"score,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Score       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ArticleTaxonomy) RawJSON() string { return r.JSON.raw }
func (r *ArticleTaxonomy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ArticleTopic struct {
	Name string `json:"name,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ArticleTopic) RawJSON() string { return r.JSON.raw }
func (r *ArticleTopic) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AllListResponse struct {
	Articles   []Article `json:"articles,required"`
	NumResults int64     `json:"numResults,required"`
	Status     int64     `json:"status,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Articles    respjson.Field
		NumResults  respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AllListResponse) RawJSON() string { return r.JSON.raw }
func (r *AllListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AllListParams struct {
	// Filter for articles added to Perigon's system after this date. Accepts ISO 8601
	// format (e.g., 2022-02-01T00:00:00) or yyyy-mm-dd format.
	AddDateFrom param.Opt[time.Time] `query:"addDateFrom,omitzero" format:"date-time" json:"-"`
	// Filter for articles added to Perigon's system before this date. Accepts ISO 8601
	// format (e.g., 2022-02-01T23:59:59) or yyyy-mm-dd format.
	AddDateTo param.Opt[time.Time] `query:"addDateTo,omitzero" format:"date-time" json:"-"`
	// Filter articles by company name mentions. Performs an exact match on company
	// names.
	CompanyName param.Opt[string] `query:"companyName,omitzero" json:"-"`
	// Search within the full article body content. Supports Boolean logic, exact
	// phrase matching with quotes, and wildcards for comprehensive content searching.
	Content param.Opt[string] `query:"content,omitzero" json:"-"`
	// Search within article description fields. Supports Boolean expressions, exact
	// phrase matching with quotes, and wildcards for flexible pattern matching.
	Desc param.Opt[string] `query:"desc,omitzero" json:"-"`
	// Filter for articles published after this date. Accepts ISO 8601 format (e.g.,
	// 2023-03-01T00:00:00) or yyyy-mm-dd format.
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
	// Latitude of the center point to search places
	Lat param.Opt[float64] `query:"lat,omitzero" json:"-"`
	// Returns only articles that contain links to the specified URL pattern. Matches
	// against the 'links' field in article responses.
	LinkTo param.Opt[string] `query:"linkTo,omitzero" json:"-"`
	// Longitude of the center point to search places
	Lon param.Opt[float64] `query:"lon,omitzero" json:"-"`
	// Maximum distance (in km) from starting point to search articles by tagged places
	MaxDistance param.Opt[float64] `query:"maxDistance,omitzero" json:"-"`
	// Filter articles with a negative sentiment score greater than or equal to the
	// specified value. Scores range from 0 to 1, with higher values indicating
	// stronger negative tone.
	NegativeSentimentFrom param.Opt[float64] `query:"negativeSentimentFrom,omitzero" json:"-"`
	// Filter articles with a negative sentiment score less than or equal to the
	// specified value. Scores range from 0 to 1, with higher values indicating
	// stronger negative tone.
	NegativeSentimentTo param.Opt[float64] `query:"negativeSentimentTo,omitzero" json:"-"`
	// Filter articles with a neutral sentiment score greater than or equal to the
	// specified value. Scores range from 0 to 1, with higher values indicating
	// stronger neutral tone.
	NeutralSentimentFrom param.Opt[float64] `query:"neutralSentimentFrom,omitzero" json:"-"`
	// Filter articles with a neutral sentiment score less than or equal to the
	// specified value. Scores range from 0 to 1, with higher values indicating
	// stronger neutral tone.
	NeutralSentimentTo param.Opt[float64] `query:"neutralSentimentTo,omitzero" json:"-"`
	// The specific page of results to retrieve in the paginated response. Starts at 0.
	Page param.Opt[int64] `query:"page,omitzero" json:"-"`
	// Filter to show only results where the source has a paywall (true) or does not
	// have a paywall (false).
	Paywall param.Opt[bool] `query:"paywall,omitzero" json:"-"`
	// Filter articles with a positive sentiment score greater than or equal to the
	// specified value. Scores range from 0 to 1, with higher values indicating
	// stronger positive tone.
	PositiveSentimentFrom param.Opt[float64] `query:"positiveSentimentFrom,omitzero" json:"-"`
	// Filter articles with a positive sentiment score less than or equal to the
	// specified value. Scores range from 0 to 1, with higher values indicating
	// stronger positive tone.
	PositiveSentimentTo param.Opt[float64] `query:"positiveSentimentTo,omitzero" json:"-"`
	// Filters by Google Content Categories. This field will filter by the category
	// prefix only. Example: prefixTaxonomy=/Finance
	PrefixTaxonomy param.Opt[string] `query:"prefixTaxonomy,omitzero" json:"-"`
	// Primary search query for filtering articles based on their title, description,
	// and content. Supports Boolean operators (AND, OR, NOT), exact phrases with
	// quotes, and wildcards (\* and ?) for flexible searching.
	Q param.Opt[string] `query:"q,omitzero" json:"-"`
	// Filter for articles refreshed/updated in Perigon's system after this date. In
	// most cases yields similar results to addDateFrom but can differ for updated
	// content. Accepts ISO 8601 format (e.g., 2022-02-01T00:00:00) or yyyy-mm-dd
	// format.
	RefreshDateFrom param.Opt[time.Time] `query:"refreshDateFrom,omitzero" format:"date-time" json:"-"`
	// Filter for articles refreshed/updated in Perigon's system before this date. In
	// most cases yields similar results to addDateTo but can differ for updated
	// content. Accepts ISO 8601 format (e.g., 2022-02-01T23:59:59) or yyyy-mm-dd
	// format.
	RefreshDateTo param.Opt[time.Time] `query:"refreshDateTo,omitzero" format:"date-time" json:"-"`
	// Returns all articles in a specific reprint group, including the original article
	// and all its known reprints. Use when you want to see all versions of the same
	// content.
	ReprintGroupID param.Opt[string] `query:"reprintGroupId,omitzero" json:"-"`
	// Expand search to include translated content fields for non-English articles.
	// When true, searches translated title, description, and content fields.
	SearchTranslation param.Opt[bool] `query:"searchTranslation,omitzero" json:"-"`
	// When set to true, enables text highlighting in search results.
	ShowHighlighting param.Opt[bool] `query:"showHighlighting,omitzero" json:"-"`
	// Whether to show the total number of all matched articles. Default value is false
	// which makes queries a bit more efficient but also counts up to 10000 articles.
	ShowNumResults param.Opt[bool] `query:"showNumResults,omitzero" json:"-"`
	// Controls whether to include reprinted content in results. When true (default),
	// shows syndicated articles from wire services like AP or Reuters that appear on
	// multiple sites.
	ShowReprints param.Opt[bool] `query:"showReprints,omitzero" json:"-"`
	// The number of articles to return per page in the paginated response.
	Size param.Opt[int64] `query:"size,omitzero" json:"-"`
	// Latitude of the center point to search articles created by local publications.
	SourceLat param.Opt[float64] `query:"sourceLat,omitzero" json:"-"`
	// Latitude of the center point to search articles created by local publications.
	SourceLon param.Opt[float64] `query:"sourceLon,omitzero" json:"-"`
	// Maximum distance from starting point to search articles created by local
	// publications.
	SourceMaxDistance param.Opt[float64] `query:"sourceMaxDistance,omitzero" json:"-"`
	// Search specifically within article headlines/titles. Supports Boolean operators,
	// exact phrases with quotes, and wildcards for matching title variations.
	Title param.Opt[string] `query:"title,omitzero" json:"-"`
	// Filter for articles published before this date. Accepts ISO 8601 format (e.g.,
	// 2022-02-01T23:59:59) or yyyy-mm-dd format.
	To param.Opt[time.Time] `query:"to,omitzero" format:"date-time" json:"-"`
	// Search within article URLs to find content from specific website sections or
	// domains. Supports wildcards (\* and ?) for partial URL matching.
	URL param.Opt[string] `query:"url,omitzero" json:"-"`
	// Filters articles where a specified area, such as a neighborhood, borough, or
	// district, plays a central role in the content, beyond mere mentions, to ensure
	// the results are deeply relevant to the area in question. If multiple parameters
	// are passed, they will be applied as OR operations.
	Area []string `query:"area,omitzero" json:"-"`
	// Retrieve specific news articles by their unique article identifiers. Multiple
	// IDs can be provided to return a collection of specific articles.
	ArticleID []string `query:"articleId,omitzero" json:"-"`
	// Filter articles by specific author names. Works as an exact match for each name.
	// Multiple values create an OR filter to find articles by any of the specified
	// authors.
	Author []string `query:"author,omitzero" json:"-"`
	// Filter articles by author bylines. Works as an exact match for each author name
	// provided. Multiple values create an OR filter to find articles by any of the
	// specified authors.
	Byline []string `query:"byline,omitzero" json:"-"`
	// Filter by broad content categories such as Politics, Tech, Sports, Business, or
	// Finance. Use 'none' to find uncategorized articles. Multiple values create an OR
	// filter.
	Category []string `query:"category,omitzero" json:"-"`
	// Filters articles where a specified city plays a central role in the content,
	// beyond mere mentions, to ensure the results are deeply relevant to the urban
	// area in question. If multiple parameters are passed, they will be applied as OR
	// operations.
	City []string `query:"city,omitzero" json:"-"`
	// Filter results to only show content within a specific related content cluster.
	// Returns articles grouped together as part of Perigon Stories based on topic
	// relevance.
	ClusterID []string `query:"clusterId,omitzero" json:"-"`
	// Filter articles by company domains (e.g., apple.com). For available company
	// entities, consult the /companies endpoint.
	CompanyDomain []string `query:"companyDomain,omitzero" json:"-"`
	// Filter articles by company identifiers. For a complete list of tracked
	// companies, refer to the /companies endpoint.
	CompanyID []string `query:"companyId,omitzero" json:"-"`
	// Filter articles by company stock symbols. For available company entities and
	// their symbols, consult the /companies endpoint.
	CompanySymbol []string `query:"companySymbol,omitzero" json:"-"`
	// Country code to filter by country. If multiple parameters are passed, they will
	// be applied as OR operations.
	Country []string `query:"country,omitzero" json:"-"`
	// A list of counties to include (or specify) in the search results. This field
	// filters the returned articles based on the county associated with the event or
	// news. Only articles tagged with one of these counties will be included.
	County []string `query:"county,omitzero" json:"-"`
	// Exclude articles written by specific authors. Any article with an author name
	// matching an entry in this list will be omitted from results. Multiple values
	// create an AND-exclude filter.
	ExcludeAuthor []string `query:"excludeAuthor,omitzero" json:"-"`
	// Exclude articles with specific categories. Multiple values create an AND-exclude
	// filter, removing all content with any of these categories.
	ExcludeCategory []string `query:"excludeCategory,omitzero" json:"-"`
	// A list of cities to exclude from the results. Articles that are associated with
	// any of the specified cities will be filtered out.
	ExcludeCity []string `query:"excludeCity,omitzero" json:"-"`
	// Exclude articles related to companies with specific domains. Creates an
	// AND-exclude filter to remove content about these companies.
	ExcludeCompanyDomain []string `query:"excludeCompanyDomain,omitzero" json:"-"`
	// Exclude articles mentioning companies with specific identifiers. Creates an
	// AND-exclude filter to remove content about these corporate entities.
	ExcludeCompanyID []string `query:"excludeCompanyId,omitzero" json:"-"`
	// A list of stock symbols (ticker symbols) that identify companies to be excluded.
	// Articles related to companies using any of these symbols will be omitted, which
	// is useful for targeting or avoiding specific public companies.
	ExcludeCompanySymbol []string `query:"excludeCompanySymbol,omitzero" json:"-"`
	// Excludes articles from specific counties or administrative divisions in the
	// search results. Accepts either a single county name or a list of county names.
	// County names should match the format used in article metadata (e.g., 'Los
	// Angeles County', 'Cook County'). This parameter allows for more granular
	// geographic filter
	ExcludeCounty []string `query:"excludeCounty,omitzero" json:"-"`
	// Exclude articles written by specific journalists identified by their unique IDs.
	// Multiple values create an AND-exclude filter.
	ExcludeJournalistID []string `query:"excludeJournalistId,omitzero" json:"-"`
	// Exclude articles with specific editorial labels. Multiple values create an
	// AND-exclude filter, removing all content with any of these labels.
	ExcludeLabel []string `query:"excludeLabel,omitzero" json:"-"`
	// Exclude articles in specific languages using ISO-639 two-letter codes. Multiple
	// values create an AND-exclude filter.
	ExcludeLanguage []string `query:"excludeLanguage,omitzero" json:"-"`
	// Excludes articles where a specified country plays a central role in the content,
	// ensuring results are not deeply relevant to the country in question. If multiple
	// parameters are passed, they will be applied as AND operations, excluding
	// articles relevant to any of the specified countries.
	ExcludeLocationsCountry []string `query:"excludeLocationsCountry,omitzero" json:"-"`
	// Exclude articles mentioning specific people by name. Creates an AND-exclude
	// filter to remove content about these individuals.
	ExcludePersonName []string `query:"excludePersonName,omitzero" json:"-"`
	// Exclude articles mentioning people with specific Wikidata IDs. Creates an
	// AND-exclude filter to remove content about these individuals. Uses precise
	// identifiers to avoid name ambiguity.
	ExcludePersonWikidataID []string `query:"excludePersonWikidataId,omitzero" json:"-"`
	// Exclude articles from specific publisher domains or subdomains. Supports
	// wildcards (_ and ?) for pattern matching (e.g., _.cnn.com). Multiple values
	// create an AND-exclude filter.
	ExcludeSource []string `query:"excludeSource,omitzero" json:"-"`
	// Exclude articles from specified Perigon source groups. Multiple values create an
	// AND-exclude filter, removing content from publishers in any of the specified
	// bundles (e.g., top10, top100).
	ExcludeSourceGroup []string `query:"excludeSourceGroup,omitzero" json:"-"`
	// A list of states to exclude. Articles that include, or are associated with, any
	// of the states provided here will be filtered out. This is especially useful if
	// you want to ignore news tied to certain geographical areas (e.g., US states).
	ExcludeState []string `query:"excludeState,omitzero" json:"-"`
	// Exclude articles with specific topics. Multiple values create an AND-exclude
	// filter, removing all content with any of these topics.
	ExcludeTopic []string `query:"excludeTopic,omitzero" json:"-"`
	// Filter by unique journalist identifiers which can be found through the
	// Journalist API or in the matchedAuthors field. Multiple values create an OR
	// filter.
	JournalistID []string `query:"journalistId,omitzero" json:"-"`
	// Filter articles by editorial labels such as Opinion, Paid-news, Non-news, Fact
	// Check, or Press Release. Multiple values create an OR filter.
	Label []string `query:"label,omitzero" json:"-"`
	// Filter articles by their language using ISO-639 two-letter codes (e.g., en, es,
	// fr). Multiple values create an OR filter.
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
	// Filter articles by their primary medium type. Accepts Article for written
	// content or Video for video-based stories. Multiple values create an OR filter.
	Medium []string `query:"medium,omitzero" json:"-"`
	// Filter articles by exact person name matches. Does not support Boolean or
	// complex logic. For available person entities, consult the /people endpoint.
	PersonName []string `query:"personName,omitzero" json:"-"`
	// Filter articles by Wikidata IDs of mentioned people. Refer to the /people
	// endpoint for a complete list of tracked individuals.
	PersonWikidataID []string `query:"personWikidataId,omitzero" json:"-"`
	// Determines the article sorting order. Options include relevance (default),
	// date/pubDate (newest publication date first), reverseDate (oldest publication
	// date first), addDate (newest ingestion date first), reverseAddDate (oldest
	// ingestion date first), and refreshDate (most recently updated in system first,
	// often identical to addDate).
	//
	// Any of "relevance", "date", "reverseDate", "reverseAddDate", "addDate",
	// "pubDate", "refreshDate".
	SortBy AllEndpointSortBy `query:"sortBy,omitzero" json:"-"`
	// Filter articles by specific publisher domains or subdomains. Supports wildcards
	// (_ and ?) for pattern matching (e.g., _.cnn.com). Multiple values create an OR
	// filter.
	Source []string `query:"source,omitzero" json:"-"`
	// Find articles published by sources that are located within a given city.
	SourceCity []string `query:"sourceCity,omitzero" json:"-"`
	// Find articles published by sources that are located within a given country. Must
	// be 2 character country code (i.e. us, gb, etc).
	SourceCountry []string `query:"sourceCountry,omitzero" json:"-"`
	// Find articles published by sources that are located within a given county.
	SourceCounty []string `query:"sourceCounty,omitzero" json:"-"`
	// Filter articles using Perigon's curated publisher bundles (e.g., top100,
	// top25crypto). Multiple values create an OR filter to include articles from any
	// of the specified bundles.
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
	// taxonomy=/Finance/Banking/Other, /Finance/Investing/Funds.
	// [Full list](https://cloud.google.com/natural-language/docs/categories)
	Taxonomy []string `query:"taxonomy,omitzero" json:"-"`
	// Filter by specific topics such as Markets, Crime, Cryptocurrency, or College
	// Sports. Topics are more granular than categories, and articles can have multiple
	// topics. Use the /topics endpoint for a complete list of available topics.
	// Multiple values create an OR filter.
	Topic []string `query:"topic,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AllListParams]'s query parameters as `url.Values`.
func (r AllListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
