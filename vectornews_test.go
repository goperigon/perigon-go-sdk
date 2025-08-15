// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package perigon_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/goperigon/perigon-go-sdk/v2"
	"github.com/goperigon/perigon-go-sdk/v2/internal/testutil"
	"github.com/goperigon/perigon-go-sdk/v2/option"
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
		Prompt: "Enter your prompt",
		Filter: perigon.ArticleSearchFilterParam{
			And:           []perigon.ArticleSearchFilterParam{},
			ArticleID:     []string{""},
			Category:      []string{""},
			City:          []string{""},
			ClusterID:     []string{""},
			CompanyDomain: []string{""},
			CompanyID:     []string{""},
			CompanyName:   []string{""},
			CompanySymbol: []string{""},
			Coordinates: perigon.CoordinateFilterParam{
				Lat:    perigon.Float(0),
				Lon:    perigon.Float(0),
				Radius: perigon.Float(0),
			},
			Country:                 []string{""},
			County:                  []string{""},
			ExcludeCity:             []string{""},
			ExcludeCompanyDomain:    []string{""},
			ExcludeCompanyID:        []string{""},
			ExcludeCompanySymbol:    []string{""},
			ExcludeCountry:          []string{""},
			ExcludeCounty:           []string{""},
			ExcludeLabel:            []string{""},
			ExcludeLanguage:         []string{""},
			ExcludeLocationsCountry: []string{""},
			ExcludePersonName:       []string{""},
			ExcludePersonWikidataID: []string{""},
			ExcludeSource:           []string{""},
			ExcludeState:            []string{""},
			ExcludeTopic:            []string{""},
			Label:                   []string{""},
			Language:                []string{""},
			LocationsCountry:        []string{""},
			Not:                     []perigon.ArticleSearchFilterParam{},
			Or:                      []perigon.ArticleSearchFilterParam{},
			PersonName:              []string{""},
			PersonWikidataID:        []string{""},
			Source:                  []string{""},
			SourceCity:              []string{""},
			SourceCoordinates: perigon.CoordinateFilterParam{
				Lat:    perigon.Float(0),
				Lon:    perigon.Float(0),
				Radius: perigon.Float(0),
			},
			SourceCountry: []string{""},
			SourceCounty:  []string{""},
			SourceGroup:   []string{""},
			SourceState:   []string{""},
			State:         []string{""},
			Taxonomy:      []string{""},
			Topic:         []string{""},
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
