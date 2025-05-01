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

func TestAllListWithOptionalParams(t *testing.T) {
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
	_, err := client.All.List(context.TODO(), perigonsdk.AllListParams{
		AddDateFrom:             perigonsdk.Time(time.Now()),
		AddDateTo:               perigonsdk.Time(time.Now()),
		Area:                    []string{"string"},
		ArticleID:               []string{"string"},
		Author:                  []string{"string"},
		Byline:                  []string{"string"},
		Category:                []string{"string"},
		City:                    []string{"string"},
		ClusterID:               []string{"string"},
		CompanyDomain:           []string{"string"},
		CompanyID:               []string{"string"},
		CompanyName:             perigonsdk.String("companyName"),
		CompanySymbol:           []string{"string"},
		Content:                 perigonsdk.String("content"),
		Country:                 []string{"string"},
		County:                  []string{"string"},
		Desc:                    perigonsdk.String("desc"),
		ExcludeAuthor:           []string{"string"},
		ExcludeCategory:         []string{"string"},
		ExcludeCity:             []string{"string"},
		ExcludeCompanyDomain:    []string{"string"},
		ExcludeCompanyID:        []string{"string"},
		ExcludeCompanySymbol:    []string{"string"},
		ExcludeCounty:           []string{"string"},
		ExcludeJournalistID:     []string{"string"},
		ExcludeLabel:            []string{"string"},
		ExcludeLanguage:         []string{"string"},
		ExcludeLocationsCountry: []string{"string"},
		ExcludePersonName:       []string{"string"},
		ExcludePersonWikidataID: []string{"string"},
		ExcludeSource:           []string{"string"},
		ExcludeSourceGroup:      []string{"string"},
		ExcludeState:            []string{"string"},
		ExcludeTopic:            []string{"string"},
		From:                    perigonsdk.Time(time.Now()),
		JournalistID:            []string{"string"},
		Label:                   []string{"string"},
		Language:                []string{"string"},
		Lat:                     perigonsdk.Float(-90),
		LinkTo:                  perigonsdk.String("linkTo"),
		Location:                []string{"string"},
		LocationsCountry:        []string{"string"},
		Lon:                     perigonsdk.Float(-180),
		MaxDistance:             perigonsdk.Float(1),
		Medium:                  []string{"string"},
		NegativeSentimentFrom:   perigonsdk.Float(0),
		NegativeSentimentTo:     perigonsdk.Float(0),
		NeutralSentimentFrom:    perigonsdk.Float(0),
		NeutralSentimentTo:      perigonsdk.Float(0),
		Page:                    perigonsdk.Int(0),
		Paywall:                 perigonsdk.Bool(true),
		PersonName:              []string{"string"},
		PersonWikidataID:        []string{"string"},
		PositiveSentimentFrom:   perigonsdk.Float(0),
		PositiveSentimentTo:     perigonsdk.Float(0),
		PrefixTaxonomy:          perigonsdk.String("prefixTaxonomy"),
		Q:                       perigonsdk.String("q"),
		RefreshDateFrom:         perigonsdk.Time(time.Now()),
		RefreshDateTo:           perigonsdk.Time(time.Now()),
		ReprintGroupID:          perigonsdk.String("reprintGroupId"),
		SearchTranslation:       perigonsdk.Bool(true),
		ShowNumResults:          perigonsdk.Bool(true),
		ShowReprints:            perigonsdk.Bool(true),
		Size:                    perigonsdk.Int(0),
		SortBy:                  perigonsdk.AllEndpointSortByRelevance,
		Source:                  []string{"string"},
		SourceCity:              []string{"string"},
		SourceCountry:           []string{"string"},
		SourceCounty:            []string{"string"},
		SourceGroup:             []string{"string"},
		SourceLat:               perigonsdk.Float(-90),
		SourceLon:               perigonsdk.Float(-180),
		SourceMaxDistance:       perigonsdk.Float(1),
		SourceState:             []string{"string"},
		State:                   []string{"string"},
		Taxonomy:                []string{"string"},
		Title:                   perigonsdk.String("title"),
		To:                      perigonsdk.Time(time.Now()),
		Topic:                   []string{"string"},
		URL:                     perigonsdk.String("url"),
	})
	if err != nil {
		var apierr *perigonsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
