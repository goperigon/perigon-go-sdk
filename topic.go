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
)

// TopicService contains methods and other services that help with interacting with
// the perigon API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTopicService] method instead.
type TopicService struct {
	Options []option.RequestOption
}

// NewTopicService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewTopicService(opts ...option.RequestOption) (r TopicService) {
	r = TopicService{}
	r.Options = opts
	return
}

// Search through all available Topics that exist within the Perigon Database.
func (r *TopicService) List(ctx context.Context, query TopicListParams, opts ...option.RequestOption) (res *TopicListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/topics/all"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Topic search result
type TopicListResponse struct {
	Data  []TopicListResponseData `json:"data,required"`
	Total int64                   `json:"total,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Total       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TopicListResponse) RawJSON() string { return r.JSON.raw }
func (r *TopicListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TopicListResponseData struct {
	ID        int64                       `json:"id,nullable"`
	CreatedAt time.Time                   `json:"createdAt,nullable" format:"date-time"`
	Labels    TopicListResponseDataLabels `json:"labels"`
	Name      string                      `json:"name,nullable"`
	UpdatedAt time.Time                   `json:"updatedAt,nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Labels      respjson.Field
		Name        respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TopicListResponseData) RawJSON() string { return r.JSON.raw }
func (r *TopicListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TopicListResponseDataLabels struct {
	Category    string `json:"category,nullable"`
	Subcategory string `json:"subcategory,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Category    respjson.Field
		Subcategory respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TopicListResponseDataLabels) RawJSON() string { return r.JSON.raw }
func (r *TopicListResponseDataLabels) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TopicListParams struct {
	// Filter topics by broad article categories such as Politics, Tech, Sports,
	// Business, Finance, Entertainment, etc.
	Category param.Opt[string] `query:"category,omitzero" json:"-"`
	// Search for topics by exact name or partial text match. Does not support
	// wildcards. Examples include Markets, Cryptocurrency, Climate Change, etc.
	Name param.Opt[string] `query:"name,omitzero" json:"-"`
	// The specific page of results to retrieve in the paginated response. Starts at 0.
	Page param.Opt[int64] `query:"page,omitzero" json:"-"`
	// The number of topics to return per page in the paginated response.
	Size param.Opt[int64] `query:"size,omitzero" json:"-"`
	// Filter topics by their specific subcategory. Subcategories provide more granular
	// classification beyond the main category, such as TV or Event.
	Subcategory param.Opt[string] `query:"subcategory,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [TopicListParams]'s query parameters as `url.Values`.
func (r TopicListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
