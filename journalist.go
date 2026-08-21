// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package perigon

import (
	"context"
	"errors"
	"fmt"
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
// JournalistService contains methods and other services that help with interacting
// with the perigon API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewJournalistService] method instead.
type JournalistService struct {
	Options []option.RequestOption
}

// NewJournalistService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewJournalistService(opts ...option.RequestOption) (r JournalistService) {
	r = JournalistService{}
	r.Options = opts
	return
}

// Find additional details on a journalist by using the journalist ID found in an
// article response object.
func (r *JournalistService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *Journalist, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/journalists/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Search journalists using broad search attributes. Our database contains over
// 230,000 journalists from around the world and is refreshed frequently.
func (r *JournalistService) List(ctx context.Context, query JournalistListParams, opts ...option.RequestOption) (res *JournalistListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/journalists/all"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Expanded journalist details, when requested and available.
type Journalist struct {
	// Unique Perigon identifier for the journalist.
	ID string `json:"id" api:"nullable"`
	// Average number of articles published by the journalist per month.
	AvgMonthlyPosts int64 `json:"avgMonthlyPosts" api:"nullable"`
	// Blog URL for the journalist.
	BlogURL string `json:"blogUrl" api:"nullable"`
	// Biographical description of the journalist.
	Description string `json:"description" api:"nullable"`
	// Journalist's email.
	Email string `json:"email" api:"nullable"`
	// Facebook profile URL for the journalist.
	FacebookURL string `json:"facebookUrl" api:"nullable"`
	// Full name of the journalist.
	FullName string `json:"fullName" api:"nullable"`
	// Professional headline associated with the journalist.
	Headline string `json:"headline" api:"nullable"`
	// Profile image URL for the journalist.
	ImageURL string `json:"imageUrl" api:"nullable"`
	// Instagram profile URL for the journalist.
	InstagramURL string `json:"instagramUrl" api:"nullable"`
	// LinkedIn connection count for the journalist, when available.
	LinkedinConnections int64 `json:"linkedinConnections" api:"nullable"`
	// LinkedIn follower count for the journalist, when available.
	LinkedinFollowers int64 `json:"linkedinFollowers" api:"nullable"`
	// LinkedIn profile URL for the journalist.
	LinkedinURL string `json:"linkedinUrl" api:"nullable"`
	// Locations associated with the journalist.
	Locations []shared.LocationHolder `json:"locations" api:"nullable"`
	// Common display name of the journalist.
	Name string `json:"name" api:"nullable"`
	// Professional title of the journalist.
	Title string `json:"title" api:"nullable"`
	// Categories most frequently covered by the journalist.
	TopCategories []NameCount `json:"topCategories" api:"nullable"`
	// Countries most frequently associated with the journalist's articles.
	TopCountries []NameCount `json:"topCountries" api:"nullable"`
	// Editorial labels most frequently associated with the journalist's articles.
	TopLabels []NameCount `json:"topLabels" api:"nullable"`
	// Publishers most frequently associated with the journalist.
	TopSources []NameCount `json:"topSources" api:"nullable"`
	// Topics most frequently covered by the journalist.
	TopTopics []NameCount `json:"topTopics" api:"nullable"`
	// Tumblr profile URL for the journalist.
	TumblrURL string `json:"tumblrUrl" api:"nullable"`
	// Biography from the journalist's X or Twitter profile.
	TwitterBio string `json:"twitterBio" api:"nullable"`
	// Journalist's X or Twitter handle.
	TwitterHandle string `json:"twitterHandle" api:"nullable"`
	// Date and time the journalist record was last refreshed, in ISO 8601 format.
	UpdatedAt string `json:"updatedAt" api:"nullable"`
	// Personal or professional website URL for the journalist.
	WebsiteURL string `json:"websiteUrl" api:"nullable"`
	// YouTube channel URL for the journalist.
	YoutubeURL string `json:"youtubeUrl" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                  respjson.Field
		AvgMonthlyPosts     respjson.Field
		BlogURL             respjson.Field
		Description         respjson.Field
		Email               respjson.Field
		FacebookURL         respjson.Field
		FullName            respjson.Field
		Headline            respjson.Field
		ImageURL            respjson.Field
		InstagramURL        respjson.Field
		LinkedinConnections respjson.Field
		LinkedinFollowers   respjson.Field
		LinkedinURL         respjson.Field
		Locations           respjson.Field
		Name                respjson.Field
		Title               respjson.Field
		TopCategories       respjson.Field
		TopCountries        respjson.Field
		TopLabels           respjson.Field
		TopSources          respjson.Field
		TopTopics           respjson.Field
		TumblrURL           respjson.Field
		TwitterBio          respjson.Field
		TwitterHandle       respjson.Field
		UpdatedAt           respjson.Field
		WebsiteURL          respjson.Field
		YoutubeURL          respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Journalist) RawJSON() string { return r.JSON.raw }
func (r *Journalist) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Countries most frequently associated with the journalist's articles.
type NameCount struct {
	// Number of matching records for the value.
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
func (r NameCount) RawJSON() string { return r.JSON.raw }
func (r *NameCount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Journalist search result
type JournalistListResponse struct {
	NumResults int64        `json:"numResults" api:"required"`
	Results    []Journalist `json:"results" api:"required"`
	Status     int64        `json:"status" api:"required"`
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
func (r JournalistListResponse) RawJSON() string { return r.JSON.raw }
func (r *JournalistListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type JournalistListParams struct {
	// Integer. Filter for journalists who publish no more than this many articles per
	// month.
	MaxMonthlyPosts param.Opt[int64] `query:"maxMonthlyPosts,omitzero" json:"-"`
	// Integer. Filter for journalists who publish at least this many articles per
	// month. Used to identify more active journalists.
	MinMonthlyPosts param.Opt[int64] `query:"minMonthlyPosts,omitzero" json:"-"`
	// String. Search specifically within journalist names. Supports Boolean operators
	// (AND, OR, NOT), exact phrases with quotes, and wildcards (\* and ?) for flexible
	// searching.
	Name param.Opt[string] `query:"name,omitzero" json:"-"`
	// Integer. The specific page of results to retrieve in the paginated response.
	// Starts at 0.
	Page param.Opt[int64] `query:"page,omitzero" json:"-"`
	// String. Primary search query for filtering journalists based on their name,
	// title, and Twitter bio. Supports Boolean operators (AND, OR, NOT), exact phrases
	// with quotes, and wildcards (\* and ?) for flexible searching.
	Q param.Opt[string] `query:"q,omitzero" json:"-"`
	// Boolean. Controls whether to return the exact result count. When false
	// (default), counts are capped at 10,000 for performance reasons. Set to true for
	// precise counts in smaller result sets.
	ShowNumResults param.Opt[bool] `query:"showNumResults,omitzero" json:"-"`
	// Integer. The number of journalists to return per page in the paginated response.
	Size param.Opt[int64] `query:"size,omitzero" json:"-"`
	// String. Filter journalists by their exact Twitter handle, without the @ symbol.
	Twitter param.Opt[string] `query:"twitter,omitzero" json:"-"`
	// Date. Filter for journalist profiles updated on or after this date. Accepts ISO
	// 8601 format (e.g., 2023-03-01T00:00:00) or yyyy-mm-dd format.
	UpdatedAtFrom param.Opt[time.Time] `query:"updatedAtFrom,omitzero" format:"date-time" json:"-"`
	// Date. Filter for journalist profiles updated on or before this date. Accepts ISO
	// 8601 format (e.g., 2023-03-01T23:59:59) or yyyy-mm-dd format.
	UpdatedAtTo param.Opt[time.Time] `query:"updatedAtTo,omitzero" format:"date-time" json:"-"`
	// String Array. Filter by unique journalist identifiers. Multiple values create an
	// OR filter to find journalists matching any of the specified IDs.
	ID []string `query:"id,omitzero" json:"-"`
	// String Array. Filter journalists by the top categories they cover. Categories
	// are general themes that the article is about. Examples of categories: Tech,
	// Politics, etc. If multiple parameters are passed, they will be applied as OR
	// operations. (Searches inside the topCategories data field.)
	Category []string `query:"category,omitzero" json:"-"`
	// String Array. Filter journalists by countries they commonly cover in their
	// reporting. Uses ISO 3166-1 alpha-2 two-letter country codes in lowercase (e.g.,
	// us, gb, jp). Multiple values create an OR filter.
	Country []string `query:"country,omitzero" json:"-"`
	// String Array. Filter journalists by the most common label tagged to the articles
	// they publish. This accepts labels like 'Opinion' or 'Pop Culture'. (Searches the
	// topLabels data field.)
	Label []string `query:"label,omitzero" json:"-"`
	// String Array. Filter journalists by their profile location area — neighborhood,
	// borough, or district (`locations.area`). Multiple values create an OR filter.
	LocationArea []string `query:"locationArea,omitzero" json:"-"`
	// String Array. Filter journalists by their profile location city
	// (`locations.city`). Multiple values create an OR filter.
	LocationCity []string `query:"locationCity,omitzero" json:"-"`
	// String Array. Filter journalists by their profile location country
	// (`locations.country`). Uses ISO 3166-1 alpha-2 two-letter country codes in
	// lowercase (e.g., us, gb, jp). `uk` is accepted as an alias for `gb`. Distinct
	// from `country`, which filters reporting focus (`topCountries`). Multiple values
	// create an OR filter.
	LocationCountry []string `query:"locationCountry,omitzero" json:"-"`
	// String Array. Filter journalists by their profile location county
	// (`locations.county`). County names should match stored metadata (e.g., 'Los
	// Angeles County', 'Cook County'). Multiple values create an OR filter.
	LocationCounty []string `query:"locationCounty,omitzero" json:"-"`
	// String Array. Filter journalists by their profile location state or region
	// (`locations.state`). For US locations this is the two-letter state code (e.g.,
	// NY, CA). Multiple values create an OR filter.
	LocationState []string `query:"locationState,omitzero" json:"-"`
	// String. Sort journalists from highest to lowest by avgMonthlyPosts,
	// linkedinConnections, or linkedinFollowers. When omitted, results are sorted by
	// relevance.
	//
	// Any of "avgMonthlyPosts", "linkedinConnections", "linkedinFollowers".
	SortBy JournalistListParamsSortBy `query:"sortBy,omitzero" json:"-"`
	// String Array. Filter journalists by the publisher domains they write for.
	// Supports wildcards (_ and ?) for pattern matching (e.g., _.cnn.com). Multiple
	// values create an OR filter.
	Source []string `query:"source,omitzero" json:"-"`
	// String Array. Filter journalists by the top topics they cover. Topics are more
	// specific themes that the article is about. Examples of topics: 'Economy', 'Real
	// Estate', 'Cryptocurrency'. If multiple parameters are passed, they will be
	// applied as OR operations. (Searches inside the topCategories data field.)
	Topic []string `query:"topic,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [JournalistListParams]'s query parameters as `url.Values`.
func (r JournalistListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// String. Sort journalists from highest to lowest by avgMonthlyPosts,
// linkedinConnections, or linkedinFollowers. When omitted, results are sorted by
// relevance.
type JournalistListParamsSortBy string

const (
	JournalistListParamsSortByAvgMonthlyPosts     JournalistListParamsSortBy = "avgMonthlyPosts"
	JournalistListParamsSortByLinkedinConnections JournalistListParamsSortBy = "linkedinConnections"
	JournalistListParamsSortByLinkedinFollowers   JournalistListParamsSortBy = "linkedinFollowers"
)
