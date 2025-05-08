// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package perigon

import (
	"context"
	"errors"
	"fmt"
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
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v1/journalists/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Search journalists using broad search attributes. Our database contains over
// 230,000 journalists from around the world and is refreshed frequently.
func (r *JournalistService) List(ctx context.Context, query JournalistListParams, opts ...option.RequestOption) (res *JournalistListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/journalists/all"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type Journalist struct {
	ID                  string               `json:"id,nullable"`
	AvgMonthlyPosts     int64                `json:"avgMonthlyPosts,nullable"`
	BlogURL             string               `json:"blogUrl,nullable"`
	Description         string               `json:"description,nullable"`
	FacebookURL         string               `json:"facebookUrl,nullable"`
	FullName            string               `json:"fullName,nullable"`
	Headline            string               `json:"headline,nullable"`
	ImageURL            string               `json:"imageUrl,nullable"`
	InstagramURL        string               `json:"instagramUrl,nullable"`
	LinkedinConnections int64                `json:"linkedinConnections,nullable"`
	LinkedinFollowers   int64                `json:"linkedinFollowers,nullable"`
	LinkedinURL         string               `json:"linkedinUrl,nullable"`
	Locations           []JournalistLocation `json:"locations,nullable"`
	Name                string               `json:"name,nullable"`
	Title               string               `json:"title,nullable"`
	TopCategories       []NameCount          `json:"topCategories,nullable"`
	TopCountries        []NameCount          `json:"topCountries,nullable"`
	TopLabels           []NameCount          `json:"topLabels,nullable"`
	TopSources          []NameCount          `json:"topSources,nullable"`
	TopTopics           []NameCount          `json:"topTopics,nullable"`
	TumblrURL           string               `json:"tumblrUrl,nullable"`
	TwitterBio          string               `json:"twitterBio,nullable"`
	TwitterHandle       string               `json:"twitterHandle,nullable"`
	UpdatedAt           string               `json:"updatedAt,nullable"`
	WebsiteURL          string               `json:"websiteUrl,nullable"`
	YoutubeURL          string               `json:"youtubeUrl,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                  respjson.Field
		AvgMonthlyPosts     respjson.Field
		BlogURL             respjson.Field
		Description         respjson.Field
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

type JournalistLocation struct {
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
func (r JournalistLocation) RawJSON() string { return r.JSON.raw }
func (r *JournalistLocation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NameCount struct {
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
func (r NameCount) RawJSON() string { return r.JSON.raw }
func (r *NameCount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Journalist search result
type JournalistListResponse struct {
	NumResults int64        `json:"numResults,required"`
	Results    []Journalist `json:"results,required"`
	Status     int64        `json:"status,required"`
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
	// Filter for journalists who publish no more than this many articles per month.
	MaxMonthlyPosts param.Opt[int64] `query:"maxMonthlyPosts,omitzero" json:"-"`
	// Filter for journalists who publish at least this many articles per month. Used
	// to identify more active journalists.
	MinMonthlyPosts param.Opt[int64] `query:"minMonthlyPosts,omitzero" json:"-"`
	// Search specifically within journalist names. Supports Boolean operators (AND,
	// OR, NOT), exact phrases with quotes, and wildcards (\* and ?) for flexible
	// searching.
	Name param.Opt[string] `query:"name,omitzero" json:"-"`
	// The specific page of results to retrieve in the paginated response. Starts at 0.
	Page param.Opt[int64] `query:"page,omitzero" json:"-"`
	// Primary search query for filtering journalists based on their name, title, and
	// Twitter bio. Supports Boolean operators (AND, OR, NOT), exact phrases with
	// quotes, and wildcards (\* and ?) for flexible searching.
	Q param.Opt[string] `query:"q,omitzero" json:"-"`
	// Controls whether to return the exact result count. When false (default), counts
	// are capped at 10,000 for performance reasons. Set to true for precise counts in
	// smaller result sets.
	ShowNumResults param.Opt[bool] `query:"showNumResults,omitzero" json:"-"`
	// The number of journalists to return per page in the paginated response.
	Size param.Opt[int64] `query:"size,omitzero" json:"-"`
	// Filter journalists by their exact Twitter handle, without the @ symbol.
	Twitter param.Opt[string] `query:"twitter,omitzero" json:"-"`
	// Filter for journalist profiles updated on or after this date. Accepts ISO 8601
	// format (e.g., 2023-03-01T00:00:00) or yyyy-mm-dd format.
	UpdatedAtFrom param.Opt[time.Time] `query:"updatedAtFrom,omitzero" format:"date-time" json:"-"`
	// Filter for journalist profiles updated on or before this date. Accepts ISO 8601
	// format (e.g., 2023-03-01T23:59:59) or yyyy-mm-dd format.
	UpdatedAtTo param.Opt[time.Time] `query:"updatedAtTo,omitzero" format:"date-time" json:"-"`
	// Filter by unique journalist identifiers. Multiple values create an OR filter to
	// find journalists matching any of the specified IDs.
	ID []string `query:"id,omitzero" json:"-"`
	// Filter journalists by the content categories they typically write about (e.g.,
	// Politics, Tech, Sports, Business). Multiple values create an OR filter.
	Category []string `query:"category,omitzero" json:"-"`
	// Filter journalists by countries they commonly cover in their reporting. Uses ISO
	// 3166-1 alpha-2 two-letter country codes in lowercase (e.g., us, gb, jp).
	// Multiple values create an OR filter.
	Country []string `query:"country,omitzero" json:"-"`
	// Filter journalists by the type of content they typically produce (e.g., Opinion,
	// Paid-news, Non-news). Multiple values create an OR filter.
	Label []string `query:"label,omitzero" json:"-"`
	// Filter journalists by the publisher domains they write for. Supports wildcards
	// (_ and ?) for pattern matching (e.g., _.cnn.com). Multiple values create an OR
	// filter.
	Source []string `query:"source,omitzero" json:"-"`
	// Filter journalists by the topics they frequently cover. Multiple values create
	// an OR filter to find journalists covering any of the specified topics.
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
