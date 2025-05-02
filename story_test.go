// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package perigonsdk_test

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

func TestStoryListWithOptionalParams(t *testing.T) {
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
	_, err := client.Stories.List(context.TODO(), perigonsdk.StoryListParams{
		Area:                  []string{"string"},
		Category:              []string{"string"},
		City:                  []string{"string"},
		ClusterID:             []string{"string"},
		CompanyDomain:         []string{"string"},
		CompanyID:             []string{"string"},
		CompanyName:           perigonsdk.String("companyName"),
		CompanySymbol:         []string{"string"},
		Country:               []string{"string"},
		ExcludeClusterID:      []string{"string"},
		From:                  perigonsdk.Time(time.Now()),
		InitializedFrom:       perigonsdk.Time(time.Now()),
		InitializedTo:         perigonsdk.Time(time.Now()),
		MaxClusterSize:        perigonsdk.Int(0),
		MinClusterSize:        perigonsdk.Int(1),
		MinUniqueSources:      perigonsdk.Int(1),
		Name:                  perigonsdk.String("name"),
		NameExists:            perigonsdk.Bool(true),
		NegativeSentimentFrom: perigonsdk.Float(0),
		NegativeSentimentTo:   perigonsdk.Float(0),
		NeutralSentimentFrom:  perigonsdk.Float(0),
		NeutralSentimentTo:    perigonsdk.Float(0),
		Page:                  perigonsdk.Int(0),
		PersonName:            perigonsdk.String("personName"),
		PersonWikidataID:      []string{"string"},
		PositiveSentimentFrom: perigonsdk.Float(0),
		PositiveSentimentTo:   perigonsdk.Float(0),
		Q:                     perigonsdk.String("q"),
		ShowDuplicates:        perigonsdk.Bool(true),
		ShowNumResults:        perigonsdk.Bool(true),
		ShowStoryPageInfo:     perigonsdk.Bool(true),
		Size:                  perigonsdk.Int(0),
		SortBy:                perigonsdk.SortByCreatedAt,
		Source:                []string{"string"},
		SourceGroup:           []string{"string"},
		State:                 []string{"string"},
		Taxonomy:              []string{"string"},
		To:                    perigonsdk.Time(time.Now()),
		Topic:                 []string{"string"},
		UpdatedFrom:           perigonsdk.Time(time.Now()),
		UpdatedTo:             perigonsdk.Time(time.Now()),
	})
	if err != nil {
		var apierr *perigonsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
