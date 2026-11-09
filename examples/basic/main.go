package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/goperigon/perigon-go-sdk/v3"
)

func main() {
	client, err := perigon.New(os.Getenv("PERIGON_API_KEY"))
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()

	fmt.Println("Perigon Go SDK - Basic Usage Examples")
	fmt.Println("======================================")

	fmt.Println("\n1. Basic Article Search")
	basicArticleSearch(client, ctx)

	fmt.Println("\n2. Filtering by Source and Date Range")
	filteredArticleSearch(client, ctx)

	fmt.Println("\n3. Story Search (Clustered Articles)")
	storySearch(client, ctx)

	fmt.Println("\n4. Company-focused Article Search")
	companySearch(client, ctx)

	fmt.Println("\n5. Error Handling Demonstration")
	errorHandlingDemo(client, ctx)
}

func basicArticleSearch(client *perigon.ClientWithResponses, ctx context.Context) {
	resp, err := client.SearchArticlesWithResponse(ctx, &perigon.SearchArticlesParams{
		Q:      perigon.Ptr("artificial intelligence"),
		Size:   perigon.Ptr(int32(5)),
		SortBy: perigon.Ptr(perigon.AllEndpointSortByDate),
	})
	if err != nil {
		log.Printf("Error in basic search: %v", err)
		return
	}
	body, err := perigon.Decoded(resp.StatusCode(), resp.Body, resp.JSON200)
	if err != nil {
		log.Printf("Error in basic search: %v", err)
		return
	}

	fmt.Printf("Found %d articles about AI:\n", len(body.Articles))
	for i, article := range body.Articles {
		if i >= 3 {
			break
		}
		fmt.Printf("  • %s\n", text(article.Title))
		fmt.Printf("    Published: %s\n", text(article.PubDate))
		fmt.Printf("    Source: %s\n", sourceDomain(article.Source))
	}
}

func filteredArticleSearch(client *perigon.ClientWithResponses, ctx context.Context) {
	lastWeek := time.Now().AddDate(0, 0, -7)
	resp, err := client.SearchArticlesWithResponse(ctx, &perigon.SearchArticlesParams{
		Q:        perigon.Ptr("technology OR startup"),
		Source:   perigon.Ptr([]string{"reuters.com", "bloomberg.com", "techcrunch.com"}),
		Category: perigon.Ptr([]string{"Tech", "Business"}),
		From:     perigon.Ptr(lastWeek),
		Size:     perigon.Ptr(int32(3)),
		SortBy:   perigon.Ptr(perigon.AllEndpointSortByDate),
	})
	if err != nil {
		log.Printf("Error in filtered search: %v", err)
		return
	}
	body, err := perigon.Decoded(resp.StatusCode(), resp.Body, resp.JSON200)
	if err != nil {
		log.Printf("Error in filtered search: %v", err)
		return
	}

	fmt.Printf("Found %d filtered tech articles:\n", len(body.Articles))
	for _, article := range body.Articles {
		fmt.Printf("  • %s\n", text(article.Title))
		fmt.Printf("    Source: %s | Category: %s\n", sourceDomain(article.Source), firstCategory(article.Categories))
	}
}

func storySearch(client *perigon.ClientWithResponses, ctx context.Context) {
	resp, err := client.SearchStoriesWithResponse(ctx, &perigon.SearchStoriesParams{
		Q:                perigon.Ptr("climate change"),
		MinUniqueSources: perigon.Ptr(int32(3)),
		SortBy:           perigon.Ptr(perigon.ClustersListSortByUpdatedAt),
		Size:             perigon.Ptr(int32(3)),
	})
	if err != nil {
		log.Printf("Error in story search: %v", err)
		return
	}
	body, err := perigon.Decoded(resp.StatusCode(), resp.Body, resp.JSON200)
	if err != nil {
		log.Printf("Error in story search: %v", err)
		return
	}

	fmt.Printf("Found %d stories about climate change:\n", len(body.Results))
	for _, story := range body.Results {
		fmt.Printf("  • %s\n", text(story.Name))
		fmt.Printf("    Articles: %d | Last Updated: %s\n", n32(story.UniqueCount), text(story.UpdatedAt))
		if summary := text(story.Summary); summary != "" {
			fmt.Printf("    Summary: %.100s...\n", summary)
		}
	}
}

func companySearch(client *perigon.ClientWithResponses, ctx context.Context) {
	companies, err := client.SearchCompaniesWithResponse(ctx, &perigon.SearchCompaniesParams{
		Name: perigon.Ptr("Tesla"),
		Size: perigon.Ptr(int32(2)),
	})
	if err != nil {
		log.Printf("Error in company search: %v", err)
		return
	}
	companyBody, err := perigon.Decoded(companies.StatusCode(), companies.Body, companies.JSON200)
	if err != nil {
		log.Printf("Error in company search: %v", err)
		return
	}

	fmt.Printf("Found %d companies matching 'Tesla':\n", len(companyBody.Results))
	for _, company := range companyBody.Results {
		domain := "N/A"
		if company.Domains != nil && len(*company.Domains) > 0 {
			domain = (*company.Domains)[0]
		}
		fmt.Printf("  • %s (%s)\n", text(company.Name), domain)
		fmt.Printf("    Country: %s | Industry: %s\n", text(company.Country), text(company.Industry))
	}

	articles, err := client.SearchArticlesWithResponse(ctx, &perigon.SearchArticlesParams{
		CompanyName: perigon.Ptr("Tesla"),
		Size:        perigon.Ptr(int32(3)),
		SortBy:      perigon.Ptr(perigon.AllEndpointSortByDate),
	})
	if err != nil {
		log.Printf("Error searching Tesla articles: %v", err)
		return
	}
	articleBody, err := perigon.Decoded(articles.StatusCode(), articles.Body, articles.JSON200)
	if err != nil {
		log.Printf("Error searching Tesla articles: %v", err)
		return
	}

	fmt.Printf("\nFound %d articles mentioning Tesla:\n", len(articleBody.Articles))
	for _, article := range articleBody.Articles {
		fmt.Printf("  • %s\n", text(article.Title))
		fmt.Printf("    Published: %s\n", text(article.PubDate))
	}
}

func errorHandlingDemo(client *perigon.ClientWithResponses, ctx context.Context) {
	resp, err := client.SearchArticlesWithResponse(ctx, &perigon.SearchArticlesParams{
		Size: perigon.Ptr(int32(1000)),
		Q:    perigon.Ptr(""),
	})
	if err != nil {
		fmt.Printf("Transport error: %v\n", err)
		return
	}
	if err := perigon.Failed(resp.StatusCode(), resp.Body); err != nil {
		if apiErr, ok := err.(*perigon.APIError); ok {
			fmt.Printf("API Error: HTTP %d\n", apiErr.StatusCode)
			fmt.Printf("Error details: %s\n", apiErr.Error())
		}
		fmt.Println("Error handling demonstrated successfully")
		return
	}
	fmt.Println("No error occurred")
}

func text(v *string) string {
	if v == nil {
		return ""
	}
	return *v
}

func n32(v *int32) int32 {
	if v == nil {
		return 0
	}
	return *v
}

func sourceDomain(source *perigon.SourceHolder) string {
	if source == nil || source.Domain == nil {
		return ""
	}
	return *source.Domain
}

func firstCategory(categories *[]perigon.CategoryHolder) string {
	if categories == nil || len(*categories) == 0 || (*categories)[0].Name == nil {
		return "N/A"
	}
	return *(*categories)[0].Name
}
