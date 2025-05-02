// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package perigon_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/goperigon/perigon-go-sdk"
	"github.com/goperigon/perigon-go-sdk/internal/testutil"
	"github.com/goperigon/perigon-go-sdk/option"
)

func TestSourceListWithOptionalParams(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
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
	_, err := client.Sources.List(context.TODO(), perigon.SourceListParams{
		Category:          []string{"string"},
		Country:           []string{"string"},
		Domain:            []string{"string"},
		Label:             []string{"string"},
		MaxMonthlyPosts:   perigon.Int(0),
		MaxMonthlyVisits:  perigon.Int(0),
		MinMonthlyPosts:   perigon.Int(0),
		MinMonthlyVisits:  perigon.Int(0),
		Name:              perigon.String("name"),
		Page:              perigon.Int(0),
		Paywall:           perigon.Bool(true),
		ShowNumResults:    perigon.Bool(true),
		ShowSubdomains:    perigon.Bool(true),
		Size:              perigon.Int(1),
		SortBy:            perigon.SortByCreatedAt,
		SourceCity:        []string{"string"},
		SourceCountry:     []string{"string"},
		SourceCounty:      []string{"string"},
		SourceGroup:       perigon.String("sourceGroup"),
		SourceLat:         perigon.Float(-90),
		SourceLon:         perigon.Float(-180),
		SourceMaxDistance: perigon.Float(1),
		SourceState:       []string{"string"},
		Topic:             []string{"string"},
	})
	if err != nil {
		var apierr *perigon.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
