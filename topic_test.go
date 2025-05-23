// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package perigon_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/goperigon/perigon-go-sdk/v2"
	"github.com/goperigon/perigon-go-sdk/v2/internal/testutil"
	"github.com/goperigon/perigon-go-sdk/v2/option"
)

func TestTopicListWithOptionalParams(t *testing.T) {
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
	_, err := client.Topics.List(context.TODO(), perigon.TopicListParams{
		Category:    perigon.String("category"),
		Name:        perigon.String("name"),
		Page:        perigon.Int(0),
		Size:        perigon.Int(0),
		Subcategory: perigon.String("subcategory"),
	})
	if err != nil {
		var apierr *perigon.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
