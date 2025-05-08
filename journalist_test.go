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

func TestJournalistGet(t *testing.T) {
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
	_, err := client.Journalists.Get(context.TODO(), "id")
	if err != nil {
		var apierr *perigon.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestJournalistListWithOptionalParams(t *testing.T) {
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
	_, err := client.Journalists.List(context.TODO(), perigon.JournalistListParams{
		ID:              []string{"string"},
		Category:        []string{"string"},
		Country:         []string{"string"},
		Label:           []string{"string"},
		MaxMonthlyPosts: perigon.Int(0),
		MinMonthlyPosts: perigon.Int(0),
		Name:            perigon.String("name"),
		Page:            perigon.Int(0),
		Q:               perigon.String("q"),
		ShowNumResults:  perigon.Bool(true),
		Size:            perigon.Int(0),
		Source:          []string{"string"},
		Topic:           []string{"string"},
		Twitter:         perigon.String("twitter"),
		UpdatedAtFrom:   perigon.Time(time.Now()),
		UpdatedAtTo:     perigon.Time(time.Now()),
	})
	if err != nil {
		var apierr *perigon.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
