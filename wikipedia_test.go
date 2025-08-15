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

func TestWikipediaSearchWithOptionalParams(t *testing.T) {
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
	_, err := client.Wikipedia.Search(context.TODO(), perigon.WikipediaSearchParams{
		ID:                      []string{"string"},
		Category:                []string{"string"},
		Page:                    perigon.Int(0),
		PageviewsFrom:           perigon.Int(0),
		PageviewsTo:             perigon.Int(0),
		Q:                       perigon.String("q"),
		Reference:               perigon.String("reference"),
		ScrapedAtFrom:           perigon.Time(time.Now()),
		ScrapedAtTo:             perigon.Time(time.Now()),
		SectionID:               []string{"string"},
		ShowNumResults:          perigon.Bool(true),
		Size:                    perigon.Int(0),
		SortBy:                  perigon.WikipediaSearchParamsSortByRelevance,
		Summary:                 perigon.String("summary"),
		Text:                    perigon.String("text"),
		Title:                   perigon.String("title"),
		WikiCode:                []string{"string"},
		WikidataID:              []string{"string"},
		WikidataInstanceOfID:    []string{"string"},
		WikidataInstanceOfLabel: []string{"string"},
		WikiNamespace:           []int64{0},
		WikiPageID:              []int64{0},
		WikiRevisionFrom:        perigon.Time(time.Now()),
		WikiRevisionID:          []int64{0},
		WikiRevisionTo:          perigon.Time(time.Now()),
		WithPageviews:           perigon.Bool(true),
	})
	if err != nil {
		var apierr *perigon.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWikipediaVectorSearchWithOptionalParams(t *testing.T) {
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
	_, err := client.Wikipedia.VectorSearch(context.TODO(), perigon.WikipediaVectorSearchParams{
		Prompt: "Enter your prompt",
		Filter: perigon.WikipediaSearchFilterParam{
			And:                     []perigon.WikipediaSearchFilterParam{},
			Not:                     []perigon.WikipediaSearchFilterParam{},
			Or:                      []perigon.WikipediaSearchFilterParam{},
			PageID:                  []string{""},
			SectionID:               []string{""},
			WikiCode:                []string{""},
			WikidataID:              []string{""},
			WikidataInstanceOfID:    []string{""},
			WikidataInstanceOfLabel: []string{""},
			WikiNamespace:           []int64{0},
			WikiPageID:              []int64{0},
			WikiRevisionID:          []int64{0},
		},
		Page:             perigon.Int(0),
		PageviewsFrom:    perigon.Int(0),
		PageviewsTo:      perigon.Int(0),
		Size:             perigon.Int(1),
		WikiRevisionFrom: perigon.Time(time.Now()),
		WikiRevisionTo:   perigon.Time(time.Now()),
	})
	if err != nil {
		var apierr *perigon.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
