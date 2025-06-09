package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/goperigon/perigon-go-sdk/v2"
)

func main() {
	// Initialize the SDK with environment variables
	// Make sure to set PERIGON_API_KEY environment variable
	client := perigon.NewClient()
	ctx := context.Background()

	fmt.Println("🚀 Perigon Go SDK - Basic Usage Examples")
	fmt.Println("======================================")

	// Example 1: Basic article search with queries
	fmt.Println("\n1️⃣ Basic Article Search")
	basicArticleSearch(&client, ctx)

	// Example 2: Filtering by source and date range
	fmt.Println("\n2️⃣ Filtering by Source and Date Range")
	filteredArticleSearch(&client, ctx)

	// Example 3: Story search (clustered articles)
	fmt.Println("\n3️⃣ Story Search (Clustered Articles)")
	storySearch(&client, ctx)

	// Example 4: Company-focused article search
	fmt.Println("\n4️⃣ Company-focused Article Search")
	companySearch(&client, ctx)

	// Example 5: Error handling demonstration
	fmt.Println("\n5️⃣ Error Handling Demonstration")
	errorHandlingDemo(&client, ctx)

	fmt.Println("\n✅ All basic examples completed!")
}

// basicArticleSearch demonstrates how to perform a simple article search
func basicArticleSearch(client *perigon.Client, ctx context.Context) {
	articles, err := client.All.List(ctx, perigon.AllListParams{
		Q:      perigon.String("artificial intelligence"),
		Size:   perigon.Int(5),
		SortBy: perigon.AllEndpointSortByDate,
	})

	if err != nil {
		log.Printf("Error in basic search: %v", err)
		return
	}

	fmt.Printf("Found %d articles about AI:\n", len(articles.Articles))
	for i, article := range articles.Articles {
		if i >= 3 { // Limit output
			break
		}
		fmt.Printf("  • %s\n", article.Title)
		fmt.Printf("    Published: %s\n", article.PubDate) // PubDate is already a string
		fmt.Printf("    Source: %s\n", article.Source.Domain)
	}
}

// filteredArticleSearch demonstrates advanced filtering options
func filteredArticleSearch(client *perigon.Client, ctx context.Context) {
	// Search for technology articles from specific sources in the last 7 days
	lastWeek := time.Now().AddDate(0, 0, -7)
	
	articles, err := client.All.List(ctx, perigon.AllListParams{
		Q:        perigon.String("technology OR startup"),
		Source:   []string{"reuters.com", "bloomberg.com", "techcrunch.com"},
		Category: []string{"Tech", "Business"},
		From:     perigon.Time(lastWeek),
		Size:     perigon.Int(3),
		SortBy:   perigon.AllEndpointSortByDate,
	})

	if err != nil {
		log.Printf("Error in filtered search: %v", err)
		return
	}

	fmt.Printf("Found %d filtered tech articles:\n", len(articles.Articles))
	for _, article := range articles.Articles {
		fmt.Printf("  • %s\n", article.Title)
		fmt.Printf("    Source: %s | Category: %s\n", 
			article.Source.Domain, 
			func() string {
				if len(article.Categories) > 0 {
					return article.Categories[0].Name
				}
				return "N/A"
			}())
	}
}

// storySearch demonstrates how to find trending stories
func storySearch(client *perigon.Client, ctx context.Context) {
	stories, err := client.Stories.List(ctx, perigon.StoryListParams{
		Q:                perigon.String("climate change"),
		MinUniqueSources: perigon.Int(3),
		SortBy:           perigon.SortByUpdatedAt,
		Size:             perigon.Int(3),
	})

	if err != nil {
		log.Printf("Error in story search: %v", err)
		return
	}

	fmt.Printf("Found %d stories about climate change:\n", len(stories.Results))
	for _, story := range stories.Results {
		fmt.Printf("  • %s\n", story.Name)
		fmt.Printf("    Articles: %d | Last Updated: %s\n", 
			story.UniqueCount, story.UpdatedAt) // Use UniqueCount instead of ArticleCount, UpdatedAt is already a string
		if story.Summary != "" {
			fmt.Printf("    Summary: %.100s...\n", story.Summary)
		}
	}
}

// companySearch demonstrates company-related searches
func companySearch(client *perigon.Client, ctx context.Context) {
	// Search for companies first
	companies, err := client.Companies.List(ctx, perigon.CompanyListParams{
		Name: perigon.String("Tesla"),
		Size: perigon.Int(2),
	})

	if err != nil {
		log.Printf("Error in company search: %v", err)
		return
	}

	fmt.Printf("Found %d companies matching 'Tesla':\n", len(companies.Results))
	for _, company := range companies.Results {
		domain := "N/A"
		if len(company.Domains) > 0 {
			domain = company.Domains[0]
		}
		fmt.Printf("  • %s (%s)\n", company.Name, domain)
		fmt.Printf("    Country: %s | Industry: %s\n", company.Country, company.Industry)
	}

	// Search for articles mentioning Tesla
	articles, err := client.All.List(ctx, perigon.AllListParams{
		CompanyName: perigon.String("Tesla"),
		Size:        perigon.Int(3),
		SortBy:      perigon.AllEndpointSortByDate,
	})

	if err != nil {
		log.Printf("Error searching Tesla articles: %v", err)
		return
	}

	fmt.Printf("\nFound %d articles mentioning Tesla:\n", len(articles.Articles))
	for _, article := range articles.Articles {
		fmt.Printf("  • %s\n", article.Title)
		fmt.Printf("    Published: %s\n", article.PubDate) // PubDate is already a string
	}
}

// errorHandlingDemo shows proper error handling techniques
func errorHandlingDemo(client *perigon.Client, ctx context.Context) {
	// Intentionally trigger an error with invalid parameters
	_, err := client.All.List(ctx, perigon.AllListParams{
		Size: perigon.Int(1000), // This might exceed limits
		Q:    perigon.String(""),  // Empty query might cause issues
	})

	if err != nil {
		// Check if it's a Perigon API error
		if apiErr, ok := err.(*perigon.Error); ok {
			fmt.Printf("API Error: HTTP %d\n", apiErr.StatusCode)
			fmt.Printf("Error details: %s\n", apiErr.Error())
			
			// You can also dump the full request/response for debugging
			// fmt.Println(string(apiErr.DumpRequest(true)))
			// fmt.Println(string(apiErr.DumpResponse(true)))
		} else {
			fmt.Printf("Other error: %v\n", err)
		}
		
		fmt.Println("✅ Error handling demonstrated successfully")
	} else {
		fmt.Println("No error occurred (unexpected)")
	}
} 