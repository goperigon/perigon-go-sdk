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

func TestSummarizeNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Summarize.New(context.TODO(), perigon.SummarizeNewParams{
		AddDateFrom:             perigon.Time(time.Now()),
		AddDateTo:               perigon.Time(time.Now()),
		Area:                    []string{"string"},
		ArticleID:               []string{"string"},
		Author:                  []string{"string"},
		Byline:                  []string{"string"},
		Category:                []string{"string"},
		City:                    []string{"string"},
		ClusterID:               []string{"string"},
		CompanyDomain:           []string{"string"},
		CompanyID:               []string{"string"},
		CompanyName:             perigon.String("companyName"),
		CompanySymbol:           []string{"string"},
		Content:                 perigon.String("content"),
		Country:                 []string{"string"},
		County:                  []string{"string"},
		Desc:                    perigon.String("desc"),
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
		From:                    perigon.Time(time.Now()),
		JournalistID:            []string{"string"},
		Label:                   []string{"string"},
		Language:                []string{"string"},
		Lat:                     perigon.Float(-90),
		LinkTo:                  perigon.String("linkTo"),
		Location:                []string{"string"},
		LocationsCountry:        []string{"string"},
		Lon:                     perigon.Float(-180),
		MaxDistance:             perigon.Float(1),
		Medium:                  []string{"string"},
		NegativeSentimentFrom:   perigon.Float(0),
		NegativeSentimentTo:     perigon.Float(0),
		NeutralSentimentFrom:    perigon.Float(0),
		NeutralSentimentTo:      perigon.Float(0),
		Page:                    perigon.Int(0),
		Paywall:                 perigon.Bool(true),
		PersonName:              []string{"string"},
		PersonWikidataID:        []string{"string"},
		PositiveSentimentFrom:   perigon.Float(0),
		PositiveSentimentTo:     perigon.Float(0),
		PrefixTaxonomy:          perigon.String("prefixTaxonomy"),
		Q:                       perigon.String("q"),
		RefreshDateFrom:         perigon.Time(time.Now()),
		RefreshDateTo:           perigon.Time(time.Now()),
		ReprintGroupID:          perigon.String("reprintGroupId"),
		SearchTranslation:       perigon.Bool(true),
		ShowNumResults:          perigon.Bool(true),
		ShowReprints:            perigon.Bool(true),
		Size:                    perigon.Int(0),
		SortBy:                  perigon.AllEndpointSortByRelevance,
		Source:                  []string{"string"},
		SourceCity:              []string{"string"},
		SourceCountry:           []string{"string"},
		SourceCounty:            []string{"string"},
		SourceGroup:             []string{"string"},
		SourceLat:               perigon.Float(-90),
		SourceLon:               perigon.Float(-180),
		SourceMaxDistance:       perigon.Float(1),
		SourceState:             []string{"string"},
		State:                   []string{"string"},
		Taxonomy:                []string{"string"},
		Title:                   perigon.String("title"),
		To:                      perigon.Time(time.Now()),
		Topic:                   []string{"string"},
		URL:                     perigon.String("url"),
		MaxArticleCount:         perigon.Int(1),
		MaxTokens:               perigon.Int(0),
		Method:                  perigon.SummarizeNewParamsMethodArticles,
		Model:                   perigon.SummarizeNewParamsModelGpt4o,
		Prompt:                  perigon.String("prompt"),
		ReturnedArticleCount:    perigon.Int(1),
		SummarizeFields:         perigon.SummarizeNewParamsSummarizeFieldsTitle,
		Temperature:             perigon.Float(0),
		TopP:                    perigon.Float(0),
	})
	if err != nil {
		var apierr *perigon.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
