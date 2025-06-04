// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package perigon_test

import (
	"context"
	"os"
	"testing"

	"github.com/goperigon/perigon-go-sdk/v2"
	"github.com/goperigon/perigon-go-sdk/v2/internal/testutil"
	"github.com/goperigon/perigon-go-sdk/v2/option"
)

func TestUsage(t *testing.T) {
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
	alls, err := client.All.List(context.TODO(), perigon.AllListParams{})
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	t.Logf("%+v\n", alls.Articles)
}
