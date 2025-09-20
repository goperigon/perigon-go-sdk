// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package perigon

import (
	"context"
	"net/http"
	"net/url"
	"slices"

	"github.com/goperigon/perigon-go-sdk/v2/internal/apijson"
	"github.com/goperigon/perigon-go-sdk/v2/internal/apiquery"
	"github.com/goperigon/perigon-go-sdk/v2/internal/requestconfig"
	"github.com/goperigon/perigon-go-sdk/v2/option"
	"github.com/goperigon/perigon-go-sdk/v2/packages/param"
	"github.com/goperigon/perigon-go-sdk/v2/packages/respjson"
)

// PersonService contains methods and other services that help with interacting
// with the perigon API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPersonService] method instead.
type PersonService struct {
	Options []option.RequestOption
}

// NewPersonService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewPersonService(opts ...option.RequestOption) (r PersonService) {
	r = PersonService{}
	r.Options = opts
	return
}

// Search and retrieve additional information on known persons that exist within
// Perigon's entity database and as referenced in any article response object. Our
// database contains over 650,000 people from around the world and is refreshed
// frequently. People data is derived from Wikidata and includes a wikidataId field
// that can be used to lookup even more information on Wikidata's website.
func (r *PersonService) List(ctx context.Context, query PersonListParams, opts ...option.RequestOption) (res *PersonListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/people/all"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type ImageHolder struct {
	URL string `json:"url,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ImageHolder) RawJSON() string { return r.JSON.raw }
func (r *ImageHolder) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WikidataDateHolder struct {
	Precision string `json:"precision,nullable"`
	Time      string `json:"time,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Precision   respjson.Field
		Time        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WikidataDateHolder) RawJSON() string { return r.JSON.raw }
func (r *WikidataDateHolder) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WikidataLabelHolder struct {
	Label      string `json:"label,nullable"`
	WikidataID string `json:"wikidataId,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Label       respjson.Field
		WikidataID  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WikidataLabelHolder) RawJSON() string { return r.JSON.raw }
func (r *WikidataLabelHolder) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Person search result
type PersonListResponse struct {
	NumResults int64                      `json:"numResults,required"`
	Results    []PersonListResponseResult `json:"results,required"`
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
func (r PersonListResponse) RawJSON() string { return r.JSON.raw }
func (r *PersonListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PersonListResponseResult struct {
	Abstract       string                                   `json:"abstract,nullable"`
	Aliases        []string                                 `json:"aliases,nullable"`
	CreatedAt      string                                   `json:"createdAt,nullable"`
	DateOfBirth    WikidataDateHolder                       `json:"dateOfBirth,nullable"`
	DateOfDeath    WikidataDateHolder                       `json:"dateOfDeath,nullable"`
	Description    string                                   `json:"description,nullable"`
	Gender         WikidataLabelHolder                      `json:"gender,nullable"`
	Image          ImageHolder                              `json:"image,nullable"`
	Name           string                                   `json:"name,nullable"`
	Occupation     []WikidataLabelHolder                    `json:"occupation,nullable"`
	PoliticalParty []PersonListResponseResultPoliticalParty `json:"politicalParty,nullable"`
	Position       []PersonListResponseResultPosition       `json:"position,nullable"`
	UpdatedAt      string                                   `json:"updatedAt,nullable"`
	WikidataID     string                                   `json:"wikidataId,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Abstract       respjson.Field
		Aliases        respjson.Field
		CreatedAt      respjson.Field
		DateOfBirth    respjson.Field
		DateOfDeath    respjson.Field
		Description    respjson.Field
		Gender         respjson.Field
		Image          respjson.Field
		Name           respjson.Field
		Occupation     respjson.Field
		PoliticalParty respjson.Field
		Position       respjson.Field
		UpdatedAt      respjson.Field
		WikidataID     respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PersonListResponseResult) RawJSON() string { return r.JSON.raw }
func (r *PersonListResponseResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PersonListResponseResultPoliticalParty struct {
	EndTime    WikidataDateHolder `json:"endTime,nullable"`
	Label      string             `json:"label,nullable"`
	StartTime  WikidataDateHolder `json:"startTime,nullable"`
	WikidataID string             `json:"wikidataId,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EndTime     respjson.Field
		Label       respjson.Field
		StartTime   respjson.Field
		WikidataID  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PersonListResponseResultPoliticalParty) RawJSON() string { return r.JSON.raw }
func (r *PersonListResponseResultPoliticalParty) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PersonListResponseResultPosition struct {
	Employer   WikidataLabelHolder `json:"employer,nullable"`
	EndTime    WikidataDateHolder  `json:"endTime,nullable"`
	Label      string              `json:"label,nullable"`
	StartTime  WikidataDateHolder  `json:"startTime,nullable"`
	WikidataID string              `json:"wikidataId,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Employer    respjson.Field
		EndTime     respjson.Field
		Label       respjson.Field
		StartTime   respjson.Field
		WikidataID  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PersonListResponseResultPosition) RawJSON() string { return r.JSON.raw }
func (r *PersonListResponseResultPosition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PersonListParams struct {
	// String. Search by person's name. Supports Boolean operators (AND, OR, NOT),
	// exact phrases with quotes, and wildcards (\* and ?) for flexible searching.
	Name param.Opt[string] `query:"name,omitzero" json:"-"`
	// String. Search by occupation name (e.g., politician, actor, CEO, athlete).
	// Supports Boolean operators (AND, OR, NOT), exact phrases with quotes, and
	// wildcards (\* and ?) for flexible searching.
	OccupationLabel param.Opt[string] `query:"occupationLabel,omitzero" json:"-"`
	// Integer. The specific page of results to retrieve in the paginated response.
	// Starts at 0.
	Page param.Opt[int64] `query:"page,omitzero" json:"-"`
	// Integer. The number of people to return per page in the paginated response.
	Size param.Opt[int64] `query:"size,omitzero" json:"-"`
	// String Array. Filter by Wikidata occupation IDs (e.g., Q82955 for politician,
	// Q33999 for actor, Q19546 for businessman). Finds people with specific
	// professions. Multiple values create an OR filter.
	OccupationID []string `query:"occupationId,omitzero" json:"-"`
	// String Array. Filter by Wikidata entity IDs (e.g., Q7747, Q937). These are
	// unique identifiers from Wikidata.org that precisely identify public figures and
	// eliminate name ambiguity. Multiple values create an OR filter.
	WikidataID []string `query:"wikidataId,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PersonListParams]'s query parameters as `url.Values`.
func (r PersonListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
