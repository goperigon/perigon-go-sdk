// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

import (
	"github.com/goperigon/perigon-go-sdk/v2/internal/apijson"
	"github.com/goperigon/perigon-go-sdk/v2/packages/param"
	"github.com/goperigon/perigon-go-sdk/v2/packages/respjson"
)

// aliased to make [param.APIUnion] private when embedding
type paramUnion = param.APIUnion

// aliased to make [param.APIObject] private when embedding
type paramObj = param.APIObject

type CategoryHolder struct {
	Name string `json:"name,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CategoryHolder) RawJSON() string { return r.JSON.raw }
func (r *CategoryHolder) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Coordinate struct {
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
func (r Coordinate) RawJSON() string { return r.JSON.raw }
func (r *Coordinate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LocationHolder struct {
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
func (r LocationHolder) RawJSON() string { return r.JSON.raw }
func (r *LocationHolder) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SourceLocation struct {
	City        string     `json:"city,nullable"`
	Coordinates Coordinate `json:"coordinates"`
	Country     string     `json:"country,nullable"`
	County      string     `json:"county,nullable"`
	State       string     `json:"state,nullable"`
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
func (r SourceLocation) RawJSON() string { return r.JSON.raw }
func (r *SourceLocation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
