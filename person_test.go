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

func TestPersonListWithOptionalParams(t *testing.T) {
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
	_, err := client.People.List(context.TODO(), perigonsdk.PersonListParams{
		Name:            perigonsdk.String("name"),
		OccupationID:    []string{"string"},
		OccupationLabel: perigonsdk.String("occupationLabel"),
		Page:            perigonsdk.Int(0),
		Size:            perigonsdk.Int(1),
		WikidataID:      []string{"string"},
	})
	if err != nil {
		var apierr *perigonsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
