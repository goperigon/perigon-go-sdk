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
	NumResults int64        `json:"numResults,nullable"`
	Results    []Journalist `json:"results,nullable"`
	Status     int64        `json:"status,nullable"`
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
	// Returns the journalist with the maximum indicated number of average monthly
	// posts.
	MaxMonthlyPosts param.Opt[int64] `query:"maxMonthlyPosts,omitzero" json:"-"`
	// Returns the journalists with the minimum indicated number of average monthly
	// posts.
	MinMonthlyPosts param.Opt[int64] `query:"minMonthlyPosts,omitzero" json:"-"`
	// Searches through journalist names, scores and ranks them, returns results sorted
	// by relevance.
	Name param.Opt[string] `query:"name,omitzero" json:"-"`
	// The page number to retrieve.
	Page param.Opt[int64] `query:"page,omitzero" json:"-"`
	// Searches through name, title, twitterBio fields with priority given to the name,
	// then to the title, then to the twitter bio. Returns results sorted by relevance.
	Q param.Opt[string] `query:"q,omitzero" json:"-"`
	// If 'true', shows accurate number of results matched by the query. By default,
	// the counter is accurate only up to 10,000 results due performance reasons.
	ShowNumResults param.Opt[bool] `query:"showNumResults,omitzero" json:"-"`
	// The number of items per page.
	Size param.Opt[int64] `query:"size,omitzero" json:"-"`
	// Searches for journalists by (exact match) twitter handle.
	Twitter param.Opt[string] `query:"twitter,omitzero" json:"-"`
	// Starting date when the record was last updated.
	UpdatedAtFrom param.Opt[time.Time] `query:"updatedAtFrom,omitzero" format:"date-time" json:"-"`
	// Ending date when the record was last updated.
	UpdatedAtTo param.Opt[time.Time] `query:"updatedAtTo,omitzero" format:"date-time" json:"-"`
	// Filter by journalist ID. Journalist IDs are unique journalist identifiers which
	// can be found through the Journalist API, or in the matchedAuthors field.
	ID []string `query:"id,omitzero" json:"-"`
	// Filter by categories. Categories are general themes that the article is about.
	// Examples of categories: Tech, Politics, etc. If multiple parameters are passed,
	// they will be applied as OR operations.
	Category []string `query:"category,omitzero" json:"-"`
	// Country code to filter by country. If multiple parameters are passed, they will
	// be applied as OR operations.
	Country []string `query:"country,omitzero" json:"-"`
	// Filter journalists by label. For example, searching 'Opinion' will return the
	// journalists where 'Opinion'-type articles is one of the top labels for the
	// articles they publish.
	Label []string `query:"label,omitzero" json:"-"`
	// Search for journalist by the publisher's domain can include a subdomain. If
	// multiple parameters are passed, they will be applied as OR operations. Wildcards
	// (_ and ?) are suported (e.g. _.cnn.com).
	Source []string `query:"source,omitzero" json:"-"`
	// Searches for journalists by topic.
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
