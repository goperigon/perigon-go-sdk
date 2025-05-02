// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package perigon

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
	opts = append(r.Options[:], opts...)
	path := "v1/people/all"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type ImageHolder struct {
	URL string `json:"url,nullable"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		URL         resp.Field
		ExtraFields map[string]resp.Field
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
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Precision   resp.Field
		Time        resp.Field
		ExtraFields map[string]resp.Field
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
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Label       resp.Field
		WikidataID  resp.Field
		ExtraFields map[string]resp.Field
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
	NumResults int64                      `json:"numResults,nullable"`
	Results    []PersonListResponseResult `json:"results,nullable"`
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
func (r PersonListResponse) RawJSON() string { return r.JSON.raw }
func (r *PersonListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PersonListResponseResult struct {
	Abstract       string                                   `json:"abstract,nullable"`
	Aliases        []string                                 `json:"aliases,nullable"`
	CreatedAt      string                                   `json:"createdAt,nullable"`
	DateOfBirth    WikidataDateHolder                       `json:"dateOfBirth"`
	DateOfDeath    WikidataDateHolder                       `json:"dateOfDeath"`
	Description    string                                   `json:"description,nullable"`
	Gender         WikidataLabelHolder                      `json:"gender"`
	Image          ImageHolder                              `json:"image"`
	Name           string                                   `json:"name,nullable"`
	Occupation     []WikidataLabelHolder                    `json:"occupation,nullable"`
	PoliticalParty []PersonListResponseResultPoliticalParty `json:"politicalParty,nullable"`
	Position       []PersonListResponseResultPosition       `json:"position,nullable"`
	UpdatedAt      string                                   `json:"updatedAt,nullable"`
	WikidataID     string                                   `json:"wikidataId,nullable"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Abstract       resp.Field
		Aliases        resp.Field
		CreatedAt      resp.Field
		DateOfBirth    resp.Field
		DateOfDeath    resp.Field
		Description    resp.Field
		Gender         resp.Field
		Image          resp.Field
		Name           resp.Field
		Occupation     resp.Field
		PoliticalParty resp.Field
		Position       resp.Field
		UpdatedAt      resp.Field
		WikidataID     resp.Field
		ExtraFields    map[string]resp.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PersonListResponseResult) RawJSON() string { return r.JSON.raw }
func (r *PersonListResponseResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PersonListResponseResultPoliticalParty struct {
	EndTime    WikidataDateHolder `json:"endTime"`
	Label      string             `json:"label,nullable"`
	StartTime  WikidataDateHolder `json:"startTime"`
	WikidataID string             `json:"wikidataId,nullable"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		EndTime     resp.Field
		Label       resp.Field
		StartTime   resp.Field
		WikidataID  resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PersonListResponseResultPoliticalParty) RawJSON() string { return r.JSON.raw }
func (r *PersonListResponseResultPoliticalParty) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PersonListResponseResultPosition struct {
	Employer   WikidataLabelHolder `json:"employer"`
	EndTime    WikidataDateHolder  `json:"endTime"`
	Label      string              `json:"label,nullable"`
	StartTime  WikidataDateHolder  `json:"startTime"`
	WikidataID string              `json:"wikidataId,nullable"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Employer    resp.Field
		EndTime     resp.Field
		Label       resp.Field
		StartTime   resp.Field
		WikidataID  resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PersonListResponseResultPosition) RawJSON() string { return r.JSON.raw }
func (r *PersonListResponseResultPosition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PersonListParams struct {
	// Search by name of the person. Supports exact matching with quotes ("") and
	// Boolean operators (AND, OR, NOT).
	Name param.Opt[string] `query:"name,omitzero" json:"-"`
	// Search by occupation name. Supports exact matching with quotes ("") and Boolean
	// operators (AND, OR, NOT).
	OccupationLabel param.Opt[string] `query:"occupationLabel,omitzero" json:"-"`
	// The page number to retrieve.
	Page param.Opt[int64] `query:"page,omitzero" json:"-"`
	// The number of items per page.
	Size param.Opt[int64] `query:"size,omitzero" json:"-"`
	// Filter by Wikidata occupation ID(s). Use this to find people with specific
	// occupations.
	OccupationID []string `query:"occupationId,omitzero" json:"-"`
	// Filter by Wikidata entity ID(s). Use this to find specific people by their
	// Wikidata identifiers.
	WikidataID []string `query:"wikidataId,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f PersonListParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [PersonListParams]'s query parameters as `url.Values`.
func (r PersonListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
