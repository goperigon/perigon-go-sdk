// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package perigon_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/goperigon/perigon-go-sdk"
	"github.com/goperigon/perigon-go-sdk/internal/testutil"
	"github.com/goperigon/perigon-go-sdk/option"
)

func TestVectorNewsSearchWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := perigon.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Vector.News.Search(context.TODO(), perigon.VectorNewsSearchParams{
		Prompt: "prompt",
		Filter: perigon.ArticleSearchFilterParam{
			And:           []perigon.ArticleSearchFilterParam{},
			ArticleID:     []string{"string"},
			Category:      []string{"string"},
			City:          []string{"string"},
			ClusterID:     []string{"string"},
			CompanyDomain: []string{"string"},
			CompanyID:     []string{"string"},
			CompanyName:   []string{"string"},
			CompanySymbol: []string{"string"},
			Coordinates: perigon.ArticleSearchFilterCoordinatesParam{
				Lat:    perigon.Float(0),
				Lon:    perigon.Float(0),
				Radius: perigon.Float(0),
			},
			Country:                 []string{"string"},
			County:                  []string{"string"},
			ExcludeCity:             []string{"string"},
			ExcludeCompanyDomain:    []string{"string"},
			ExcludeCompanyID:        []string{"string"},
			ExcludeCompanySymbol:    []string{"string"},
			ExcludeCountry:          []string{"string"},
			ExcludeCounty:           []string{"string"},
			ExcludeLabel:            []string{"string"},
			ExcludeLanguage:         []string{"string"},
			ExcludeLocationsCountry: []string{"string"},
			ExcludePersonName:       []string{"string"},
			ExcludePersonWikidataID: []string{"string"},
			ExcludeSource:           []string{"string"},
			ExcludeState:            []string{"string"},
			ExcludeTopic:            []string{"string"},
			Label:                   []string{"string"},
			Language:                []string{"string"},
			LocationsCountry:        []string{"string"},
			Not:                     []perigon.ArticleSearchFilterParam{},
			Or:                      []perigon.ArticleSearchFilterParam{},
			PersonName:              []string{"string"},
			PersonWikidataID:        []string{"string"},
			Source:                  []string{"string"},
			SourceCity:              []string{"string"},
			SourceCoordinates: perigon.ArticleSearchFilterSourceCoordinatesParam{
				Lat:    perigon.Float(0),
				Lon:    perigon.Float(0),
				Radius: perigon.Float(0),
			},
			SourceCountry: []string{"string"},
			SourceCounty:  []string{"string"},
			SourceGroup:   []string{"string"},
			SourceState:   []string{"string"},
			State:         []string{"string"},
			Taxonomy:      []string{"string"},
			Topic:         []string{"string"},
		},
		Page:         perigon.Int(0),
		PubDateFrom:  perigon.Time(time.Now()),
		PubDateTo:    perigon.Time(time.Now()),
		ShowReprints: perigon.Bool(true),
		Size:         perigon.Int(1),
	})
	if err != nil {
		var apierr *perigon.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
