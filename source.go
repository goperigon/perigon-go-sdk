// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package perigonsdk

import (
	"context"
	"net/http"
	"net/url"

	"github.com/goperigon/perigon-go-sdk/internal/apijson"
	"github.com/goperigon/perigon-go-sdk/internal/apiquery"
	"github.com/goperigon/perigon-go-sdk/internal/requestconfig"
	"github.com/goperigon/perigon-go-sdk/option"
	"github.com/goperigon/perigon-go-sdk/packages/param"
	"github.com/goperigon/perigon-go-sdk/packages/resp"
)

// SourceService contains methods and other services that help with interacting
// with the perigon-sdk API.
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
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Count       resp.Field
		Name        resp.Field
		ExtraFields map[string]resp.Field
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
	NumResults int64                      `json:"numResults,nullable"`
	Results    []SourceListResponseResult `json:"results,nullable"`
	Status     int64                      `json:"status,nullable"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		NumResults  resp.Field
		Results     resp.Field
		Status      resp.Field
		ExtraFields map[string]resp.Field
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
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ID                 resp.Field
		AdFontesBiasRating resp.Field
		AllSidesBiasRating resp.Field
		AltNames           resp.Field
		AvgBiasRating      resp.Field
		AvgMonthlyPosts    resp.Field
		Description        resp.Field
		Domain             resp.Field
		GlobalRank         resp.Field
		Location           resp.Field
		LogoFavIcon        resp.Field
		LogoLarge          resp.Field
		LogoSquare         resp.Field
		MbfcBiasRating     resp.Field
		MonthlyVisits      resp.Field
		Name               resp.Field
		Paywall            resp.Field
		PrimaryRecordID    resp.Field
		TopCategories      resp.Field
		TopCountries       resp.Field
		TopLabels          resp.Field
		TopTopics          resp.Field
		UpdatedAt          resp.Field
		ExtraFields        map[string]resp.Field
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
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		City        resp.Field
		Coordinates resp.Field
		Country     resp.Field
		County      resp.Field
		State       resp.Field
		ExtraFields map[string]resp.Field
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
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Lat         resp.Field
		Lon         resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SourceListResponseResultLocationCoordinates) RawJSON() string { return r.JSON.raw }
func (r *SourceListResponseResultLocationCoordinates) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SourceListParams struct {
	// Returns the sources that have at most this number of average monthly posts.
	MaxMonthlyPosts param.Opt[int64] `query:"maxMonthlyPosts,omitzero" json:"-"`
	// Enter a maximum number of monthly visits that the source must have in order to
	// match your query.
	MaxMonthlyVisits param.Opt[int64] `query:"maxMonthlyVisits,omitzero" json:"-"`
	// Returns the sources that have at least this number of average monthly posts.
	MinMonthlyPosts param.Opt[int64] `query:"minMonthlyPosts,omitzero" json:"-"`
	// Filter by popularity. Enter a minimum number of monthly visits that the source
	// must have in order to match your query.
	MinMonthlyVisits param.Opt[int64] `query:"minMonthlyVisits,omitzero" json:"-"`
	// Search by name of source. This parameter supports complex boolean search
	// operators, and also searches the altNames field for alternative names of the
	// source.
	Name param.Opt[string] `query:"name,omitzero" json:"-"`
	// The page number to retrieve.
	Page param.Opt[int64] `query:"page,omitzero" json:"-"`
	// Use 'true' to find only sources known to have a paywall, or use 'false' to
	// filter for only sources that do not have a paywall.
	Paywall param.Opt[bool] `query:"paywall,omitzero" json:"-"`
	// If 'true', shows accurate number of results matched by the query. By default,
	// the counter is accurate only up to 10,000 results due performance reasons.
	ShowNumResults param.Opt[bool] `query:"showNumResults,omitzero" json:"-"`
	// Controls whether subdomains are included in the response. When set to true
	// (default), all relevant subdomains of media sources will be returned as separate
	// results. Set to false to consolidate results to parent domains only.
	ShowSubdomains param.Opt[bool] `query:"showSubdomains,omitzero" json:"-"`
	// The number of items per page.
	Size param.Opt[int64] `query:"size,omitzero" json:"-"`
	// Find all sources within a sourceGroup. Find Source Groups in the guided part of
	// our documentation...
	SourceGroup param.Opt[string] `query:"sourceGroup,omitzero" json:"-"`
	// Latitude of the center point to search local publications.
	SourceLat param.Opt[float64] `query:"sourceLat,omitzero" json:"-"`
	// Longitude of the center point to search local publications.
	SourceLon param.Opt[float64] `query:"sourceLon,omitzero" json:"-"`
	// Maximum distance from starting point to search local publications.
	SourceMaxDistance param.Opt[float64] `query:"sourceMaxDistance,omitzero" json:"-"`
	// Filter by categories. Categories are general themes that the article is about.
	// Examples of categories: Tech, Politics, etc. If multiple parameters are passed,
	// they will be applied as OR operations.
	Category []string `query:"category,omitzero" json:"-"`
	// Country code to filter sources by the countries in which they most commonly
	// cover. If multiple parameters are passed, they will be applied as OR operations.
	Country []string `query:"country,omitzero" json:"-"`
	// Domain name for the media source to lookup. This parameter supports wildcard
	// queries, ie. "\*.cnn.com" will match all subdomains for cnn.com.
	Domain []string `query:"domain,omitzero" json:"-"`
	// Filter sources by label. For example, searching 'Opinion' will return the
	// sources where 'Opinion'-type articles is one of the top labels for the articles
	// they publish.
	Label []string `query:"label,omitzero" json:"-"`
	// Use 'relevance' to sort by relevance to the query, 'globalRank' for top ranked
	// sources based on popularity, 'monthlyVisits' for sources with the largest
	// audience, 'avgMonthlyPosts' for sources with the highest publishing frequency.
	// Defaults to 'relevance'.
	//
	// Any of "createdAt", "updatedAt", "relevance", "count", "totalCount".
	SortBy SortBy `query:"sortBy,omitzero" json:"-"`
	// Find all local publications that are located within a given city.
	SourceCity []string `query:"sourceCity,omitzero" json:"-"`
	// Find all local publications that are located within a given country.
	SourceCountry []string `query:"sourceCountry,omitzero" json:"-"`
	// Find all local publications that are located within a given county.
	SourceCounty []string `query:"sourceCounty,omitzero" json:"-"`
	// Find all local publications that are located within a given state.
	SourceState []string `query:"sourceState,omitzero" json:"-"`
	// Find sources by topic. For example, searching 'Markets' will return the sources
	// where 'Markets' is one of the top 10 topics that they cover.
	Topic []string `query:"topic,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f SourceListParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [SourceListParams]'s query parameters as `url.Values`.
func (r SourceListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
