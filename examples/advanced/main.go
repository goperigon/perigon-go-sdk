package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/goperigon/perigon-go-sdk/v3"
)

func main() {
	client, err := perigon.New(os.Getenv("PERIGON_API_KEY"))
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()

	fmt.Println("Perigon Go SDK - Advanced Usage Examples")
	aiSummarization(client, ctx)
	sentimentAnalysis(client, ctx)
	geographicFiltering(client, ctx)
	complexSearch(client, ctx)
	vectorSearch(client, ctx)
	entitySearch(client, ctx)
	sourceAnalysis(client, ctx)
	topicDiscovery(client, ctx)
}

func aiSummarization(client *perigon.ClientWithResponses, ctx context.Context) {
	from := time.Now().AddDate(0, 0, -7)
	resp, err := client.SearchSummarizerWithResponse(ctx, &perigon.SearchSummarizerParams{
		Q:        perigon.Ptr("artificial intelligence AND machine learning"),
		From:     perigon.Ptr(from),
		Category: perigon.Ptr([]string{"Tech", "Business"}),
	}, perigon.SearchSummarizerJSONRequestBody{
		MaxArticleCount:      perigon.Ptr(int32(10)),
		ReturnedArticleCount: perigon.Ptr(int32(5)),
		Prompt:               perigon.Ptr("Summarize the key developments in AI and ML."),
		MaxTokens:            perigon.Ptr(int32(300)),
		Temperature:          perigon.Ptr(0.7),
	})
	if err != nil {
		log.Printf("request: %v", err)
		return
	}
	body, err := perigon.Decoded(resp.StatusCode(), resp.Body, resp.JSON200)
	if err != nil {
		log.Printf("API error: %v", err)
		return
	}
	fmt.Printf("Generated summary (%d articles):\n%s\n", body.NumResults, body.Summary)
}

func sentimentAnalysis(client *perigon.ClientWithResponses, ctx context.Context) {
	from := time.Now().AddDate(0, 0, -14)
	resp, err := client.SearchArticlesWithResponse(ctx, &perigon.SearchArticlesParams{
		Q:                     perigon.Ptr("renewable energy OR solar OR wind power"),
		PositiveSentimentFrom: perigon.Ptr(float32(0.7)),
		Size:                  perigon.Ptr(int32(5)),
		SortBy:                perigon.Ptr(perigon.AllEndpointSortByDate),
		From:                  perigon.Ptr(from),
	})
	if err != nil {
		log.Printf("request: %v", err)
		return
	}
	body, err := perigon.Decoded(resp.StatusCode(), resp.Body, resp.JSON200)
	if err != nil {
		log.Printf("API error: %v", err)
		return
	}
	fmt.Printf("Found %d positive renewable-energy articles\n", len(body.Articles))
}

func geographicFiltering(client *perigon.ClientWithResponses, ctx context.Context) {
	resp, err := client.SearchArticlesWithResponse(ctx, &perigon.SearchArticlesParams{
		Q:                perigon.Ptr("technology startup"),
		SourceCountry:    perigon.Ptr([]string{"us", "gb", "ca"}),
		LocationsCountry: perigon.Ptr([]string{"us"}),
		City:             perigon.Ptr([]string{"San Francisco", "New York", "Austin"}),
		Size:             perigon.Ptr(int32(5)),
		SortBy:           perigon.Ptr(perigon.AllEndpointSortByRelevance),
	})
	if err != nil {
		log.Printf("request: %v", err)
		return
	}
	body, err := perigon.Decoded(resp.StatusCode(), resp.Body, resp.JSON200)
	if err != nil {
		log.Printf("API error: %v", err)
		return
	}
	fmt.Printf("Found %d geographic matches\n", len(body.Articles))
}

func complexSearch(client *perigon.ClientWithResponses, ctx context.Context) {
	from := time.Now().AddDate(0, -1, 0)
	resp, err := client.SearchArticlesWithResponse(ctx, &perigon.SearchArticlesParams{
		Q:                perigon.Ptr("(artificial intelligence OR machine learning) AND (healthcare OR medical)"),
		Category:         perigon.Ptr([]string{"Tech", "Health"}),
		Language:         perigon.Ptr([]string{"en"}),
		Source:           perigon.Ptr([]string{"reuters.com", "bloomberg.com", "nature.com"}),
		ExcludeLabel:     perigon.Ptr([]string{"Opinion", "Press Release"}),
		From:             perigon.Ptr(from),
		Size:             perigon.Ptr(int32(10)),
		SortBy:           perigon.Ptr(perigon.AllEndpointSortByRelevance),
		ShowHighlighting: perigon.Ptr(true),
	})
	if err != nil {
		log.Printf("request: %v", err)
		return
	}
	body, err := perigon.Decoded(resp.StatusCode(), resp.Body, resp.JSON200)
	if err != nil {
		log.Printf("API error: %v", err)
		return
	}
	fmt.Printf("Found %d complex-search articles\n", len(body.Articles))
}

func vectorSearch(client *perigon.ClientWithResponses, ctx context.Context) {
	from := time.Now().AddDate(0, 0, -30)
	resp, err := client.VectorSearchArticlesWithResponse(ctx, perigon.VectorSearchArticlesJSONRequestBody{
		Prompt:      "Find articles about the impact of remote work on corporate culture",
		Size:        perigon.Ptr(int32(5)),
		PubDateFrom: perigon.Ptr(from),
		Filter: &perigon.ArticleSearchFilter{
			Category: perigon.Ptr([]string{"Business", "Tech"}),
			Language: perigon.Ptr([]string{"en"}),
		},
	})
	if err != nil {
		log.Printf("request: %v", err)
		return
	}
	body, err := perigon.Decoded(resp.StatusCode(), resp.Body, resp.JSON200)
	if err != nil {
		log.Printf("API error: %v", err)
		return
	}
	fmt.Printf("Found %d vector matches\n", len(body.Results))
}

func entitySearch(client *perigon.ClientWithResponses, ctx context.Context) {
	resp, err := client.SearchPeopleWithResponse(ctx, &perigon.SearchPeopleParams{
		Name: perigon.Ptr("Ada Lovelace"),
		Size: perigon.Ptr(int32(3)),
	})
	if err != nil {
		log.Printf("request: %v", err)
		return
	}
	body, err := perigon.Decoded(resp.StatusCode(), resp.Body, resp.JSON200)
	if err != nil {
		log.Printf("API error: %v", err)
		return
	}
	names := make([]string, 0, len(body.Results))
	for _, person := range body.Results {
		if person.Name != nil {
			names = append(names, *person.Name)
		}
	}
	fmt.Printf("People: %s\n", strings.Join(names, ", "))
}

func sourceAnalysis(client *perigon.ClientWithResponses, ctx context.Context) {
	resp, err := client.SearchSourcesWithResponse(ctx, &perigon.SearchSourcesParams{
		Name:   perigon.Ptr("Reuters"),
		SortBy: perigon.Ptr(perigon.SourcesSortByGlobalRank),
		Size:   perigon.Ptr(int32(3)),
	})
	if err != nil {
		log.Printf("request: %v", err)
		return
	}
	body, err := perigon.Decoded(resp.StatusCode(), resp.Body, resp.JSON200)
	if err != nil {
		log.Printf("API error: %v", err)
		return
	}
	fmt.Printf("Found %d sources\n", len(body.Results))
}

func topicDiscovery(client *perigon.ClientWithResponses, ctx context.Context) {
	resp, err := client.SearchTopicsWithResponse(ctx, &perigon.SearchTopicsParams{
		Name: perigon.Ptr("Climate"),
		Size: perigon.Ptr(int32(5)),
	})
	if err != nil {
		log.Printf("request: %v", err)
		return
	}
	body, err := perigon.Decoded(resp.StatusCode(), resp.Body, resp.JSON200)
	if err != nil {
		log.Printf("API error: %v", err)
		return
	}
	fmt.Printf("Found %d topics\n", len(body.Data))
}
