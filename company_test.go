// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package perigon_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/goperigon/perigon-go-sdk"
	"github.com/goperigon/perigon-go-sdk/internal/testutil"
	"github.com/goperigon/perigon-go-sdk/option"
)

func TestCompanyListWithOptionalParams(t *testing.T) {
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
	_, err := client.Companies.List(context.TODO(), perigon.CompanyListParams{
		ID:               []string{"string"},
		Country:          []string{"string"},
		Domain:           []string{"string"},
		Exchange:         []string{"string"},
		Industry:         perigon.String("industry"),
		IpoFrom:          perigon.Time(time.Now()),
		IpoTo:            perigon.Time(time.Now()),
		Name:             perigon.String("name"),
		NumEmployeesFrom: perigon.Int(0),
		NumEmployeesTo:   perigon.Int(0),
		Page:             perigon.Int(0),
		Q:                perigon.String("q"),
		Sector:           perigon.String("sector"),
		Size:             perigon.Int(1),
		Symbol:           []string{"string"},
	})
	if err != nil {
		var apierr *perigon.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
