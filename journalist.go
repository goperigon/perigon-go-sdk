// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package perigonsdk

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/stainless-sdks/perigon-sdk-go/internal/apijson"
	"github.com/stainless-sdks/perigon-sdk-go/internal/apiquery"
	"github.com/stainless-sdks/perigon-sdk-go/internal/requestconfig"
	"github.com/stainless-sdks/perigon-sdk-go/option"
	"github.com/stainless-sdks/perigon-sdk-go/packages/param"
	"github.com/stainless-sdks/perigon-sdk-go/packages/resp"
)

// JournalistService contains methods and other services that help with interacting
// with the perigon-sdk API.
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
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ID                  resp.Field
		AvgMonthlyPosts     resp.Field
		BlogURL             resp.Field
		Description         resp.Field
		FacebookURL         resp.Field
		FullName            resp.Field
		Headline            resp.Field
		ImageURL            resp.Field
		InstagramURL        resp.Field
		LinkedinConnections resp.Field
		LinkedinFollowers   resp.Field
		LinkedinURL         resp.Field
		Locations           resp.Field
		Name                resp.Field
		Title               resp.Field
		TopCategories       resp.Field
		TopCountries        resp.Field
		TopLabels           resp.Field
		TopSources          resp.Field
		TopTopics           resp.Field
		TumblrURL           resp.Field
		TwitterBio          resp.Field
		TwitterHandle       resp.Field
		UpdatedAt           resp.Field
		WebsiteURL          resp.Field
		YoutubeURL          resp.Field
		ExtraFields         map[string]resp.Field
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
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Area        resp.Field
		City        resp.Field
		Country     resp.Field
		County      resp.Field
		State       resp.Field
		ExtraFields map[string]resp.Field
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
func (r NameCount) RawJSON() string { return r.JSON.raw }
func (r *NameCount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Journalist search result
type JournalistListResponse struct {
	NumResults int64        `json:"numResults,nullable"`
	Results    []Journalist `json:"results,nullable"`
	Status     int64        `json:"status,nullable"`
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

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f JournalistListParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [JournalistListParams]'s query parameters as `url.Values`.
func (r JournalistListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
