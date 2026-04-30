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
	Name string `json:"name" api:"nullable"`
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
	Lat float64 `json:"lat" api:"nullable"`
	Lon float64 `json:"lon" api:"nullable"`
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
	Area    string `json:"area" api:"nullable"`
	City    string `json:"city" api:"nullable"`
	Country string `json:"country" api:"nullable"`
	County  string `json:"county" api:"nullable"`
	State   string `json:"state" api:"nullable"`
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
	City        string     `json:"city" api:"nullable"`
	Coordinates Coordinate `json:"coordinates" api:"nullable"`
	Country     string     `json:"country" api:"nullable"`
	County      string     `json:"county" api:"nullable"`
	State       string     `json:"state" api:"nullable"`
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
