// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package perigonsdk_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/stainless-sdks/perigon-sdk-go"
	"github.com/stainless-sdks/perigon-sdk-go/internal/testutil"
	"github.com/stainless-sdks/perigon-sdk-go/option"
)

func TestVectorNewsSearchWithOptionalParams(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := perigonsdk.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Vector.News.Search(context.TODO(), perigonsdk.VectorNewsSearchParams{
		Prompt:      "prompt",
		AddDateFrom: perigonsdk.Time(time.Now()),
		AddDateTo:   perigonsdk.Time(time.Now()),
		Filter: perigonsdk.ArticleSearchFilterParam{
			And:           []perigonsdk.ArticleSearchFilterParam{},
			ArticleID:     []string{"string"},
			Category:      []string{"string"},
			City:          []string{"string"},
			ClusterID:     []string{"string"},
			CompanyDomain: []string{"string"},
			CompanyID:     []string{"string"},
			CompanyName:   []string{"string"},
			CompanySymbol: []string{"string"},
			Coordinates: perigonsdk.ArticleSearchFilterCoordinatesParam{
				Lat:    perigonsdk.Float(0),
				Lon:    perigonsdk.Float(0),
				Radius: perigonsdk.Float(0),
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
			Not:                     []perigonsdk.ArticleSearchFilterParam{},
			Or:                      []perigonsdk.ArticleSearchFilterParam{},
			PersonName:              []string{"string"},
			PersonWikidataID:        []string{"string"},
			Source:                  []string{"string"},
			SourceCity:              []string{"string"},
			SourceCoordinates: perigonsdk.ArticleSearchFilterSourceCoordinatesParam{
				Lat:    perigonsdk.Float(0),
				Lon:    perigonsdk.Float(0),
				Radius: perigonsdk.Float(0),
			},
			SourceCountry: []string{"string"},
			SourceCounty:  []string{"string"},
			SourceGroup:   []string{"string"},
			SourceState:   []string{"string"},
			State:         []string{"string"},
			Taxonomy:      []string{"string"},
			Topic:         []string{"string"},
		},
		Page:           perigonsdk.Int(0),
		PubDateFrom:    perigonsdk.Time(time.Now()),
		PubDateTo:      perigonsdk.Time(time.Now()),
		ScoreThreshold: perigonsdk.Float(0),
		ShowReprints:   perigonsdk.Bool(true),
		Size:           perigonsdk.Int(1),
	})
	if err != nil {
		var apierr *perigonsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
