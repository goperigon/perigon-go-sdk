// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package perigonsdk_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/perigon-sdk-go"
	"github.com/stainless-sdks/perigon-sdk-go/internal/testutil"
	"github.com/stainless-sdks/perigon-sdk-go/option"
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
	client := perigonsdk.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Sources.List(context.TODO(), perigonsdk.SourceListParams{
		Category:          []string{"string"},
		Country:           []string{"string"},
		Domain:            []string{"string"},
		Label:             []string{"string"},
		MaxMonthlyPosts:   perigonsdk.Int(0),
		MaxMonthlyVisits:  perigonsdk.Int(0),
		MinMonthlyPosts:   perigonsdk.Int(0),
		MinMonthlyVisits:  perigonsdk.Int(0),
		Name:              perigonsdk.String("name"),
		Page:              perigonsdk.Int(0),
		Paywall:           perigonsdk.Bool(true),
		ShowNumResults:    perigonsdk.Bool(true),
		ShowSubdomains:    perigonsdk.Bool(true),
		Size:              perigonsdk.Int(1),
		SortBy:            perigonsdk.SortByCreatedAt,
		SourceCity:        []string{"string"},
		SourceCountry:     []string{"string"},
		SourceCounty:      []string{"string"},
		SourceGroup:       perigonsdk.String("sourceGroup"),
		SourceLat:         perigonsdk.Float(-90),
		SourceLon:         perigonsdk.Float(-180),
		SourceMaxDistance: perigonsdk.Float(1),
		SourceState:       []string{"string"},
		Topic:             []string{"string"},
	})
	if err != nil {
		var apierr *perigonsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
