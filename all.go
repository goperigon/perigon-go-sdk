// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package perigon

import (
	"github.com/goperigon/perigon-go-sdk/v2/internal/apijson"
	"github.com/goperigon/perigon-go-sdk/v2/option"
	"github.com/goperigon/perigon-go-sdk/v2/packages/respjson"
)

// AllService contains methods and other services that help with interacting with
// the perigon API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAllService] method instead.
type AllService struct {
	Options []option.RequestOption
}

// NewAllService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAllService(opts ...option.RequestOption) (r AllService) {
	r = AllService{}
	r.Options = opts
	return
}

type AllEndpointSortBy string

const (
	AllEndpointSortByRelevance      AllEndpointSortBy = "relevance"
	AllEndpointSortByDate           AllEndpointSortBy = "date"
	AllEndpointSortByReverseDate    AllEndpointSortBy = "reverseDate"
	AllEndpointSortByReverseAddDate AllEndpointSortBy = "reverseAddDate"
	AllEndpointSortByAddDate        AllEndpointSortBy = "addDate"
	AllEndpointSortByPubDate        AllEndpointSortBy = "pubDate"
	AllEndpointSortByRefreshDate    AllEndpointSortBy = "refreshDate"
)

type Article struct {
	AddDate               string                 `json:"addDate,nullable"`
	ArticleID             string                 `json:"articleId,nullable"`
	AuthorsByline         string                 `json:"authorsByline,nullable"`
	Categories            []ArticleCategory      `json:"categories,nullable"`
	Claim                 string                 `json:"claim,nullable"`
	Cluster               NewsCluster            `json:"cluster"`
	ClusterID             string                 `json:"clusterId,nullable"`
	Companies             []ArticleCompany       `json:"companies,nullable"`
	Content               string                 `json:"content,nullable"`
	Country               string                 `json:"country,nullable"`
	Description           string                 `json:"description,nullable"`
	Entities              []ArticleEntity        `json:"entities,nullable"`
	EventTypes            []ArticleEventType     `json:"eventTypes,nullable"`
	Highlights            map[string][]string    `json:"highlights,nullable"`
	ImageURL              string                 `json:"imageUrl,nullable"`
	Journalists           []Journalist           `json:"journalists,nullable"`
	Keywords              []ArticleKeyword       `json:"keywords,nullable"`
	Labels                []ArticleLabel         `json:"labels,nullable"`
	Language              string                 `json:"language,nullable"`
	Links                 []string               `json:"links,nullable"`
	Locations             []ArticleLocation      `json:"locations,nullable"`
	MatchedAuthors        []ArticleMatchedAuthor `json:"matchedAuthors,nullable"`
	Medium                string                 `json:"medium,nullable"`
	People                []ArticlePerson        `json:"people,nullable"`
	Places                []ArticlePlace         `json:"places,nullable"`
	PubDate               string                 `json:"pubDate,nullable"`
	RefreshDate           string                 `json:"refreshDate,nullable"`
	Reprint               bool                   `json:"reprint,nullable"`
	ReprintGroupID        string                 `json:"reprintGroupId,nullable"`
	Score                 float64                `json:"score,nullable"`
	Sentiment             ArticleSentiment       `json:"sentiment"`
	ShortSummary          string                 `json:"shortSummary,nullable"`
	Source                ArticleSource          `json:"source"`
	Summary               string                 `json:"summary,nullable"`
	Taxonomies            []ArticleTaxonomy      `json:"taxonomies,nullable"`
	Title                 string                 `json:"title,nullable"`
	Topics                []ArticleTopic         `json:"topics,nullable"`
	TranslatedDescription string                 `json:"translatedDescription,nullable"`
	TranslatedSummary     string                 `json:"translatedSummary,nullable"`
	TranslatedTitle       string                 `json:"translatedTitle,nullable"`
	Translation           string                 `json:"translation,nullable"`
	URL                   string                 `json:"url,nullable"`
	Verdict               string                 `json:"verdict,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AddDate               respjson.Field
		ArticleID             respjson.Field
		AuthorsByline         respjson.Field
		Categories            respjson.Field
		Claim                 respjson.Field
		Cluster               respjson.Field
		ClusterID             respjson.Field
		Companies             respjson.Field
		Content               respjson.Field
		Country               respjson.Field
		Description           respjson.Field
		Entities              respjson.Field
		EventTypes            respjson.Field
		Highlights            respjson.Field
		ImageURL              respjson.Field
		Journalists           respjson.Field
		Keywords              respjson.Field
		Labels                respjson.Field
		Language              respjson.Field
		Links                 respjson.Field
		Locations             respjson.Field
		MatchedAuthors        respjson.Field
		Medium                respjson.Field
		People                respjson.Field
		Places                respjson.Field
		PubDate               respjson.Field
		RefreshDate           respjson.Field
		Reprint               respjson.Field
		ReprintGroupID        respjson.Field
		Score                 respjson.Field
		Sentiment             respjson.Field
		ShortSummary          respjson.Field
		Source                respjson.Field
		Summary               respjson.Field
		Taxonomies            respjson.Field
		Title                 respjson.Field
		Topics                respjson.Field
		TranslatedDescription respjson.Field
		TranslatedSummary     respjson.Field
		TranslatedTitle       respjson.Field
		Translation           respjson.Field
		URL                   respjson.Field
		Verdict               respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Article) RawJSON() string { return r.JSON.raw }
func (r *Article) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ArticleCategory struct {
	Name string `json:"name,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ArticleCategory) RawJSON() string { return r.JSON.raw }
func (r *ArticleCategory) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ArticleCompany struct {
	ID      string   `json:"id,nullable"`
	Domains []string `json:"domains,nullable"`
	Name    string   `json:"name,nullable"`
	Symbols []string `json:"symbols,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Domains     respjson.Field
		Name        respjson.Field
		Symbols     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ArticleCompany) RawJSON() string { return r.JSON.raw }
func (r *ArticleCompany) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ArticleEntity struct {
	Data     string `json:"data,nullable"`
	Mentions int64  `json:"mentions,nullable"`
	Type     string `json:"type,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Mentions    respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ArticleEntity) RawJSON() string { return r.JSON.raw }
func (r *ArticleEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ArticleEventType struct {
	Name string `json:"name,nullable"`
	Type string `json:"type,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ArticleEventType) RawJSON() string { return r.JSON.raw }
func (r *ArticleEventType) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ArticleKeyword struct {
	Name   string  `json:"name,nullable"`
	Weight float64 `json:"weight,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Weight      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ArticleKeyword) RawJSON() string { return r.JSON.raw }
func (r *ArticleKeyword) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ArticleLabel struct {
	Name string `json:"name,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ArticleLabel) RawJSON() string { return r.JSON.raw }
func (r *ArticleLabel) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ArticleLocation struct {
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
func (r ArticleLocation) RawJSON() string { return r.JSON.raw }
func (r *ArticleLocation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ArticleMatchedAuthor struct {
	ID   string `json:"id,nullable"`
	Name string `json:"name,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ArticleMatchedAuthor) RawJSON() string { return r.JSON.raw }
func (r *ArticleMatchedAuthor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ArticlePerson struct {
	Name       string `json:"name,nullable"`
	WikidataID string `json:"wikidataId,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		WikidataID  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ArticlePerson) RawJSON() string { return r.JSON.raw }
func (r *ArticlePerson) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ArticlePlace struct {
	Amenity       string                  `json:"amenity,nullable"`
	City          string                  `json:"city,nullable"`
	Coordinates   ArticlePlaceCoordinates `json:"coordinates"`
	Country       string                  `json:"country,nullable"`
	CountryCode   string                  `json:"countryCode,nullable"`
	County        string                  `json:"county,nullable"`
	Neighbourhood string                  `json:"neighbourhood,nullable"`
	OsmID         string                  `json:"osmId,nullable"`
	Postcode      string                  `json:"postcode,nullable"`
	Quarter       string                  `json:"quarter,nullable"`
	Road          string                  `json:"road,nullable"`
	State         string                  `json:"state,nullable"`
	StateDistrict string                  `json:"stateDistrict,nullable"`
	Suburb        string                  `json:"suburb,nullable"`
	Town          string                  `json:"town,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amenity       respjson.Field
		City          respjson.Field
		Coordinates   respjson.Field
		Country       respjson.Field
		CountryCode   respjson.Field
		County        respjson.Field
		Neighbourhood respjson.Field
		OsmID         respjson.Field
		Postcode      respjson.Field
		Quarter       respjson.Field
		Road          respjson.Field
		State         respjson.Field
		StateDistrict respjson.Field
		Suburb        respjson.Field
		Town          respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ArticlePlace) RawJSON() string { return r.JSON.raw }
func (r *ArticlePlace) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ArticlePlaceCoordinates struct {
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
func (r ArticlePlaceCoordinates) RawJSON() string { return r.JSON.raw }
func (r *ArticlePlaceCoordinates) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ArticleSentiment struct {
	Negative float64 `json:"negative,nullable"`
	Neutral  float64 `json:"neutral,nullable"`
	Positive float64 `json:"positive,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Negative    respjson.Field
		Neutral     respjson.Field
		Positive    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ArticleSentiment) RawJSON() string { return r.JSON.raw }
func (r *ArticleSentiment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ArticleSource struct {
	Domain   string                `json:"domain,nullable"`
	Location ArticleSourceLocation `json:"location"`
	Paywall  bool                  `json:"paywall,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Domain      respjson.Field
		Location    respjson.Field
		Paywall     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ArticleSource) RawJSON() string { return r.JSON.raw }
func (r *ArticleSource) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ArticleSourceLocation struct {
	City        string                           `json:"city,nullable"`
	Coordinates ArticleSourceLocationCoordinates `json:"coordinates"`
	Country     string                           `json:"country,nullable"`
	County      string                           `json:"county,nullable"`
	State       string                           `json:"state,nullable"`
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
func (r ArticleSourceLocation) RawJSON() string { return r.JSON.raw }
func (r *ArticleSourceLocation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ArticleSourceLocationCoordinates struct {
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
func (r ArticleSourceLocationCoordinates) RawJSON() string { return r.JSON.raw }
func (r *ArticleSourceLocationCoordinates) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ArticleTaxonomy struct {
	Name  string  `json:"name,nullable"`
	Score float64 `json:"score,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Score       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ArticleTaxonomy) RawJSON() string { return r.JSON.raw }
func (r *ArticleTaxonomy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ArticleTopic struct {
	Name string `json:"name,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ArticleTopic) RawJSON() string { return r.JSON.raw }
func (r *ArticleTopic) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
