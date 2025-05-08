// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package perigon

import (
	"context"
	"net/http"
	"net/url"

	"github.com/goperigon/perigon-go-sdk/v2/internal/apijson"
	"github.com/goperigon/perigon-go-sdk/v2/internal/apiquery"
	"github.com/goperigon/perigon-go-sdk/v2/internal/requestconfig"
	"github.com/goperigon/perigon-go-sdk/v2/option"
	"github.com/goperigon/perigon-go-sdk/v2/packages/param"
	"github.com/goperigon/perigon-go-sdk/v2/packages/respjson"
)

// SourceService contains methods and other services that help with interacting
// with the perigon API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSourceService] method instead.
type SourceService struct {
	Options []option.RequestOption
}

// NewSourceService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewSourceService(opts ...option.RequestOption) (r SourceService) {
	r = SourceService{}
	r.Options = opts
	return
}

// Search and filter the 142,000+ media sources available via the Perigon API. The
// result includes a list of individual media sources that were matched to your
// specific criteria.
func (r *SourceService) List(ctx context.Context, query SourceListParams, opts ...option.RequestOption) (res *SourceListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/sources/all"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type SortBy string

const (
	SortByCreatedAt  SortBy = "createdAt"
	SortByUpdatedAt  SortBy = "updatedAt"
	SortByRelevance  SortBy = "relevance"
	SortByCount      SortBy = "count"
	SortByTotalCount SortBy = "totalCount"
)

type SourceTopStatHolder struct {
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
func (r SourceTopStatHolder) RawJSON() string { return r.JSON.raw }
func (r *SourceTopStatHolder) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Source search result
type SourceListResponse struct {
	NumResults int64                      `json:"numResults,required"`
	Results    []SourceListResponseResult `json:"results,required"`
	Status     int64                      `json:"status,required"`
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
func (r SourceListResponse) RawJSON() string { return r.JSON.raw }
func (r *SourceListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SourceListResponseResult struct {
	ID                 string                           `json:"id,nullable"`
	AdFontesBiasRating string                           `json:"adFontesBiasRating,nullable"`
	AllSidesBiasRating string                           `json:"allSidesBiasRating,nullable"`
	AltNames           []string                         `json:"altNames,nullable"`
	AvgBiasRating      string                           `json:"avgBiasRating,nullable"`
	AvgMonthlyPosts    int64                            `json:"avgMonthlyPosts,nullable"`
	Description        string                           `json:"description,nullable"`
	Domain             string                           `json:"domain,nullable"`
	GlobalRank         int64                            `json:"globalRank,nullable"`
	Location           SourceListResponseResultLocation `json:"location"`
	LogoFavIcon        ImageHolder                      `json:"logoFavIcon"`
	LogoLarge          ImageHolder                      `json:"logoLarge"`
	LogoSquare         ImageHolder                      `json:"logoSquare"`
	MbfcBiasRating     string                           `json:"mbfcBiasRating,nullable"`
	MonthlyVisits      int64                            `json:"monthlyVisits,nullable"`
	Name               string                           `json:"name,nullable"`
	Paywall            bool                             `json:"paywall,nullable"`
	PrimaryRecordID    string                           `json:"primaryRecordId,nullable"`
	TopCategories      []SourceTopStatHolder            `json:"topCategories,nullable"`
	TopCountries       []SourceTopStatHolder            `json:"topCountries,nullable"`
	TopLabels          []SourceTopStatHolder            `json:"topLabels,nullable"`
	TopTopics          []SourceTopStatHolder            `json:"topTopics,nullable"`
	UpdatedAt          string                           `json:"updatedAt,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		AdFontesBiasRating respjson.Field
		AllSidesBiasRating respjson.Field
		AltNames           respjson.Field
		AvgBiasRating      respjson.Field
		AvgMonthlyPosts    respjson.Field
		Description        respjson.Field
		Domain             respjson.Field
		GlobalRank         respjson.Field
		Location           respjson.Field
		LogoFavIcon        respjson.Field
		LogoLarge          respjson.Field
		LogoSquare         respjson.Field
		MbfcBiasRating     respjson.Field
		MonthlyVisits      respjson.Field
		Name               respjson.Field
		Paywall            respjson.Field
		PrimaryRecordID    respjson.Field
		TopCategories      respjson.Field
		TopCountries       respjson.Field
		TopLabels          respjson.Field
		TopTopics          respjson.Field
		UpdatedAt          respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SourceListResponseResult) RawJSON() string { return r.JSON.raw }
func (r *SourceListResponseResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SourceListResponseResultLocation struct {
	City        string                                      `json:"city,nullable"`
	Coordinates SourceListResponseResultLocationCoordinates `json:"coordinates"`
	Country     string                                      `json:"country,nullable"`
	County      string                                      `json:"county,nullable"`
	State       string                                      `json:"state,nullable"`
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
func (r SourceListResponseResultLocation) RawJSON() string { return r.JSON.raw }
func (r *SourceListResponseResultLocation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SourceListResponseResultLocationCoordinates struct {
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
func (r SourceListResponseResultLocationCoordinates) RawJSON() string { return r.JSON.raw }
func (r *SourceListResponseResultLocationCoordinates) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SourceListParams struct {
	// Filter for sources that publish no more than this many articles per month. Used
	// to target publishers by content volume.
	MaxMonthlyPosts param.Opt[int64] `query:"maxMonthlyPosts,omitzero" json:"-"`
	// Filter for sources with no more than this many monthly visitors. Used to target
	// publishers by audience size.
	MaxMonthlyVisits param.Opt[int64] `query:"maxMonthlyVisits,omitzero" json:"-"`
	// Filter for sources that publish at least this many articles per month. Used to
	// target publishers by content volume.
	MinMonthlyPosts param.Opt[int64] `query:"minMonthlyPosts,omitzero" json:"-"`
	// Filter for sources with at least this many monthly visitors. Used to target
	// publishers by audience size.
	MinMonthlyVisits param.Opt[int64] `query:"minMonthlyVisits,omitzero" json:"-"`
	// Search by source name or alternative names. Supports Boolean operators (AND, OR,
	// NOT), exact phrases with quotes, and wildcards (\* and ?) for flexible
	// searching.
	Name param.Opt[string] `query:"name,omitzero" json:"-"`
	// The specific page of results to retrieve in the paginated response. Starts at 0.
	Page param.Opt[int64] `query:"page,omitzero" json:"-"`
	// Filter by paywall status. Set to true to find sources with paywalls, or false to
	// find sources without paywalls.
	Paywall param.Opt[bool] `query:"paywall,omitzero" json:"-"`
	// Controls whether to return the exact result count. When false (default), counts
	// are capped at 10,000 for performance reasons. Set to true for precise counts in
	// smaller result sets.
	ShowNumResults param.Opt[bool] `query:"showNumResults,omitzero" json:"-"`
	// Controls whether subdomains are included as separate results. When true
	// (default), subdomains appear as distinct sources. When false, results are
	// consolidated to parent domains only.
	ShowSubdomains param.Opt[bool] `query:"showSubdomains,omitzero" json:"-"`
	// The number of sources to return per page in the paginated response.
	Size param.Opt[int64] `query:"size,omitzero" json:"-"`
	// Filter by predefined publisher bundles (e.g., top100, top50tech). Returns all
	// sources within the specified group. See documentation for available source
	// groups.
	SourceGroup param.Opt[string] `query:"sourceGroup,omitzero" json:"-"`
	// Latitude coordinate for filtering local publications by geographic proximity.
	// Used with sourceLon and sourceMaxDistance for radius search.
	SourceLat param.Opt[float64] `query:"sourceLat,omitzero" json:"-"`
	// Longitude coordinate for filtering local publications by geographic proximity.
	// Used with sourceLat and sourceMaxDistance for radius search.
	SourceLon param.Opt[float64] `query:"sourceLon,omitzero" json:"-"`
	// Maximum distance in kilometers from the coordinates defined by sourceLat and
	// sourceLon. Defines the radius for local publication searches.
	SourceMaxDistance param.Opt[float64] `query:"sourceMaxDistance,omitzero" json:"-"`
	// Filter sources by their primary content categories such as Politics, Tech,
	// Sports, Business, or Finance. Returns sources that frequently cover these
	// topics. Multiple values create an OR filter.
	Category []string `query:"category,omitzero" json:"-"`
	// Filter sources by countries they commonly cover in their reporting. Uses ISO
	// 3166-1 alpha-2 two-letter country codes in lowercase (e.g., us, gb, jp). See
	// documentation for supported country codes. Multiple values create an OR filter.
	Country []string `query:"country,omitzero" json:"-"`
	// Filter by specific publisher domains or subdomains. Supports wildcards (_ and ?)
	// for pattern matching (e.g., _.cnn.com, us?.nytimes.com). Multiple values create
	// an OR filter.
	Domain []string `query:"domain,omitzero" json:"-"`
	// Filter sources by their content label patterns (e.g., Opinion, Paid-news,
	// Non-news). Returns sources where the specified label is common in their
	// published content. See documentation for all available labels. Multiple values
	// create an OR filter.
	Label []string `query:"label,omitzero" json:"-"`
	// Determines the source sorting order. Options include relevance (default, best
	// match to query), globalRank (by overall traffic and popularity), monthlyVisits
	// (by total monthly visitor count), and avgMonthlyPosts (by number of articles
	// published monthly).
	//
	// Any of "createdAt", "updatedAt", "relevance", "count", "totalCount".
	SortBy SortBy `query:"sortBy,omitzero" json:"-"`
	// Filter for local publications based in specific cities. Multiple values create
	// an OR filter.
	SourceCity []string `query:"sourceCity,omitzero" json:"-"`
	// Filter for local publications based in specific countries. Uses ISO 3166-1
	// alpha-2 two-letter country codes in lowercase (e.g., us, gb, jp). See
	// documentation for supported country codes. Multiple values create an OR filter.
	SourceCountry []string `query:"sourceCountry,omitzero" json:"-"`
	// Filter for local publications based in specific counties. Multiple values create
	// an OR filter.
	SourceCounty []string `query:"sourceCounty,omitzero" json:"-"`
	// Filter for local publications based in specific states or regions. Uses standard
	// two-letter state codes in lowercase (e.g., ca, ny, tx). See documentation for
	// supported state codes. Multiple values create an OR filter.
	SourceState []string `query:"sourceState,omitzero" json:"-"`
	// Filter sources by their frequently covered topics (e.g., Markets,
	// Cryptocurrency, Climate Change). Returns sources where the specified topic is
	// among their top 10 covered areas. Multiple values create an OR filter.
	Topic []string `query:"topic,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SourceListParams]'s query parameters as `url.Values`.
func (r SourceListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
