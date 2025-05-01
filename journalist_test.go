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

func TestJournalistGet(t *testing.T) {
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
	_, err := client.Journalists.Get(context.TODO(), "id")
	if err != nil {
		var apierr *perigonsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestJournalistListWithOptionalParams(t *testing.T) {
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
	_, err := client.Journalists.List(context.TODO(), perigonsdk.JournalistListParams{
		ID:              []string{"string"},
		Category:        []string{"string"},
		Country:         []string{"string"},
		Label:           []string{"string"},
		MaxMonthlyPosts: perigonsdk.Int(0),
		MinMonthlyPosts: perigonsdk.Int(0),
		Name:            perigonsdk.String("name"),
		Page:            perigonsdk.Int(0),
		Q:               perigonsdk.String("q"),
		ShowNumResults:  perigonsdk.Bool(true),
		Size:            perigonsdk.Int(0),
		Source:          []string{"string"},
		Topic:           []string{"string"},
		Twitter:         perigonsdk.String("twitter"),
		UpdatedAtFrom:   perigonsdk.Time(time.Now()),
		UpdatedAtTo:     perigonsdk.Time(time.Now()),
	})
	if err != nil {
		var apierr *perigonsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
