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

func TestStoryListWithOptionalParams(t *testing.T) {
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
	_, err := client.Stories.List(context.TODO(), perigon.StoryListParams{
		Area:                  []string{"string"},
		Category:              []string{"string"},
		City:                  []string{"string"},
		ClusterID:             []string{"string"},
		CompanyDomain:         []string{"string"},
		CompanyID:             []string{"string"},
		CompanyName:           perigon.String("companyName"),
		CompanySymbol:         []string{"string"},
		Country:               []string{"string"},
		ExcludeClusterID:      []string{"string"},
		From:                  perigon.Time(time.Now()),
		InitializedFrom:       perigon.Time(time.Now()),
		InitializedTo:         perigon.Time(time.Now()),
		MaxClusterSize:        perigon.Int(0),
		MinClusterSize:        perigon.Int(1),
		MinUniqueSources:      perigon.Int(1),
		Name:                  perigon.String("name"),
		NameExists:            perigon.Bool(true),
		NegativeSentimentFrom: perigon.Float(0),
		NegativeSentimentTo:   perigon.Float(0),
		NeutralSentimentFrom:  perigon.Float(0),
		NeutralSentimentTo:    perigon.Float(0),
		Page:                  perigon.Int(0),
		PersonName:            perigon.String("personName"),
		PersonWikidataID:      []string{"string"},
		PositiveSentimentFrom: perigon.Float(0),
		PositiveSentimentTo:   perigon.Float(0),
		Q:                     perigon.String("q"),
		ShowDuplicates:        perigon.Bool(true),
		ShowNumResults:        perigon.Bool(true),
		ShowStoryPageInfo:     perigon.Bool(true),
		Size:                  perigon.Int(0),
		SortBy:                perigon.SortByCreatedAt,
		Source:                []string{"string"},
		SourceGroup:           []string{"string"},
		State:                 []string{"string"},
		Taxonomy:              []string{"string"},
		To:                    perigon.Time(time.Now()),
		Topic:                 []string{"string"},
		UpdatedFrom:           perigon.Time(time.Now()),
		UpdatedTo:             perigon.Time(time.Now()),
	})
	if err != nil {
		var apierr *perigon.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
