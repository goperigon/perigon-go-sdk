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

// Google Content Categories most strongly associated with the story.
type CategoryHolder struct {
	// Content category assigned to the article.
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

// Geographic coordinates of the place.
type Coordinate struct {
	// Latitude in decimal degrees.
	Lat float64 `json:"lat" api:"nullable"`
	// Longitude in decimal degrees.
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

// Locations most strongly associated with the story.
type LocationHolder struct {
	// Neighborhood, borough, or district central to the article.
	Area string `json:"area" api:"nullable"`
	// City central to the article.
	City string `json:"city" api:"nullable"`
	// Country central to the article.
	Country string `json:"country" api:"nullable"`
	// County central to the article.
	County string `json:"county" api:"nullable"`
	// State or region central to the article.
	State string `json:"state" api:"nullable"`
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

// Geographic location of the publisher, when available.
type SourceLocation struct {
	// City where the publisher is located.
	City string `json:"city" api:"nullable"`
	// Geographic coordinates of the place.
	Coordinates Coordinate `json:"coordinates" api:"nullable"`
	// Country where the publisher is located.
	Country string `json:"country" api:"nullable"`
	// County where the publisher is located.
	County string `json:"county" api:"nullable"`
	// State or region where the publisher is located.
	State string `json:"state" api:"nullable"`
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
