package main

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/goperigon/perigon-go-sdk/v2"
)

func main() {
	// Initialize the SDK with environment variables
	// Make sure to set PERIGON_API_KEY environment variable
	client := perigon.NewClient()
	ctx := context.Background()

	fmt.Println("🤖 Perigon Go SDK - Advanced Usage Examples")
	fmt.Println("=========================================")

	// Example 1: AI-powered article summarization
	fmt.Println("\n1️⃣ AI-Powered Article Summarization")
	aiSummarization(&client, ctx)

	// Example 2: Sentiment analysis filtering
	fmt.Println("\n2️⃣ Sentiment Analysis Filtering")
	sentimentAnalysis(&client, ctx)

	// Example 3: Geographic news filtering
	fmt.Println("\n3️⃣ Geographic News Filtering")
	geographicFiltering(&client, ctx)

	// Example 4: Complex multi-parameter searches
	fmt.Println("\n4️⃣ Complex Multi-Parameter Search")
	complexSearch(&client, ctx)

	// Example 5: Vector search (semantic search)
	fmt.Println("\n5️⃣ Vector Search (Semantic Search)")
	vectorSearch(&client, ctx)

	// Example 6: Entity and people search
	fmt.Println("\n6️⃣ Entity and People Search")
	entitySearch(&client, ctx)

	// Example 7: Advanced source and journalist analysis
	fmt.Println("\n7️⃣ Advanced Source and Journalist Analysis")
	sourceAnalysis(&client, ctx)

	// Example 8: Topic-based story discovery
	fmt.Println("\n8️⃣ Topic-based Story Discovery")
	topicDiscovery(&client, ctx)

	fmt.Println("\n✅ All advanced examples completed!")
}

// aiSummarization demonstrates AI-powered summarization
func aiSummarization(client *perigon.Client, ctx context.Context) {
	summary, err := client.Summarize.New(ctx, perigon.SummarizeNewParams{
		Q:                    perigon.String("artificial intelligence AND machine learning"),
		MaxArticleCount:      perigon.Int(10),
		ReturnedArticleCount: perigon.Int(5),
		Prompt:               perigon.String("Summarize the key developments in AI and ML, focusing on breakthrough technologies and their potential impact on society."),
		MaxTokens:            perigon.Int(300),
		Temperature:          perigon.Float(0.7),
		From:                 perigon.Time(time.Now().AddDate(0, 0, -7)), // Last 7 days
		Category:             []string{"Tech", "Business"},
	})

	if err != nil {
		log.Printf("Error in AI summarization: %v", err)
		return
	}

	fmt.Printf("Generated Summary (%d articles analyzed):\n", summary.NumResults)
	fmt.Printf("%s\n", summary.Summary)
	
	fmt.Printf("\nSample articles used for summary:\n")
	for i, article := range summary.Results {
		if i >= 3 { // Limit output
			break
		}
		fmt.Printf("  • %s\n", article.Title)
		fmt.Printf("    Score: %.2f | Source: %s\n", article.Score, article.Source.Domain)
	}
}

// sentimentAnalysis demonstrates sentiment-based filtering
func sentimentAnalysis(client *perigon.Client, ctx context.Context) {
	// Search for positive sentiment articles about renewable energy
	positiveArticles, err := client.All.List(ctx, perigon.AllListParams{
		Q:                     perigon.String("renewable energy OR solar OR wind power"),
		PositiveSentimentFrom: perigon.Float(0.7), // High positive sentiment
		Size:                  perigon.Int(5),
		SortBy:                perigon.AllEndpointSortByDate,
		From:                  perigon.Time(time.Now().AddDate(0, 0, -14)), // Last 2 weeks
	})

	if err != nil {
		log.Printf("Error in positive sentiment search: %v", err)
		return
	}

	fmt.Printf("Found %d positive articles about renewable energy:\n", len(positiveArticles.Articles))
	for _, article := range positiveArticles.Articles {
		fmt.Printf("  • %s\n", article.Title)
		fmt.Printf("    Sentiment: Positive=%.2f, Neutral=%.2f, Negative=%.2f\n", 
			article.Sentiment.Positive, article.Sentiment.Neutral, article.Sentiment.Negative)
		fmt.Printf("    Source: %s\n", article.Source.Domain)
	}

	// Also search for negative sentiment articles for comparison
	negativeArticles, err := client.All.List(ctx, perigon.AllListParams{
		Q:                     perigon.String("renewable energy OR solar OR wind power"),
		NegativeSentimentFrom: perigon.Float(0.7), // High negative sentiment
		Size:                  perigon.Int(2),
		SortBy:                perigon.AllEndpointSortByDate,
	})

	if err == nil && len(negativeArticles.Articles) > 0 {
		fmt.Printf("\nNegative sentiment articles for comparison:\n")
		for _, article := range negativeArticles.Articles {
			fmt.Printf("  • %s\n", article.Title)
			fmt.Printf("    Sentiment: Negative=%.2f\n", article.Sentiment.Negative)
		}
	}
}

// geographicFiltering demonstrates location-based filtering
func geographicFiltering(client *perigon.Client, ctx context.Context) {
	// Search for news from specific geographic regions
	articles, err := client.All.List(ctx, perigon.AllListParams{
		Q:               perigon.String("technology startup"),
		SourceCountry:   []string{"us", "gb", "ca"}, // US, UK, Canada
		LocationsCountry: []string{"us"}, // Stories focused on US
		City:            []string{"San Francisco", "New York", "Austin"},
		Size:            perigon.Int(5),
		SortBy:          perigon.AllEndpointSortByRelevance,
	})

	if err != nil {
		log.Printf("Error in geographic search: %v", err)
		return
	}

	fmt.Printf("Found %d tech startup articles from English-speaking countries:\n", len(articles.Articles))
	for _, article := range articles.Articles {
		fmt.Printf("  • %s\n", article.Title)
		
		// Display location information
		locations := []string{}
		for _, loc := range article.Locations {
			if loc.City != "" && loc.Country != "" {
				locations = append(locations, fmt.Sprintf("%s, %s", loc.City, loc.Country))
			}
		}
		if len(locations) > 0 {
			fmt.Printf("    Locations: %s\n", strings.Join(locations, "; "))
		}
		
		fmt.Printf("    Source Location: %s\n", article.Source.Location.Country)
	}
}

// complexSearch demonstrates advanced filtering with multiple parameters
func complexSearch(client *perigon.Client, ctx context.Context) {
	// Complex search with multiple filters
	articles, err := client.All.List(ctx, perigon.AllListParams{
		Q:           perigon.String("(artificial intelligence OR machine learning) AND (healthcare OR medical)"),
		Category:    []string{"Tech", "Health"},
		Language:    []string{"en"},
		Source:      []string{"reuters.com", "bloomberg.com", "nature.com", "sciencedaily.com"},
		ExcludeLabel: []string{"Opinion", "Press Release"},
		From:        perigon.Time(time.Now().AddDate(0, -1, 0)), // Last month
		Size:        perigon.Int(10),
		SortBy:      perigon.AllEndpointSortByRelevance,
		ShowHighlighting: perigon.Bool(true),
	})

	if err != nil {
		log.Printf("Error in complex search: %v", err)
		return
	}

	fmt.Printf("Found %d AI healthcare articles with complex filtering:\n", len(articles.Articles))
	for i, article := range articles.Articles {
		if i >= 5 { // Limit output
			break
		}
		fmt.Printf("  • %s\n", article.Title)
		
		// Show categories
		categories := []string{}
		for _, cat := range article.Categories {
			categories = append(categories, cat.Name)
		}
		if len(categories) > 0 {
			fmt.Printf("    Categories: %s\n", strings.Join(categories, ", "))
		}

		// Show companies mentioned
		companies := []string{}
		for _, comp := range article.Companies {
			companies = append(companies, comp.Name)
		}
		if len(companies) > 0 && len(companies) <= 3 {
			fmt.Printf("    Companies: %s\n", strings.Join(companies, ", "))
		}
		
		fmt.Printf("    Relevance Score: %.2f\n", article.Score)
	}
}

// vectorSearch demonstrates semantic search capabilities
func vectorSearch(client *perigon.Client, ctx context.Context) {
	response, err := client.Vector.News.Search(ctx, perigon.VectorNewsSearchParams{
		Prompt:      "Find articles about the impact of remote work on corporate culture and employee productivity",
		Size:        perigon.Int(5),
		PubDateFrom: perigon.Time(time.Now().AddDate(0, 0, -30)), // Last 30 days
		Filter: perigon.ArticleSearchFilterParam{
			Category: []string{"Business", "Tech"},
			Language: []string{"en"},
		},
	})

	if err != nil {
		log.Printf("Error in vector search: %v", err)
		return
	}

	fmt.Printf("Found %d semantically relevant articles about remote work:\n", len(response.Results))
	for _, result := range response.Results {
		article := result.Data
		fmt.Printf("  • %s\n", article.Title)
		fmt.Printf("    Semantic Score: %.3f\n", result.Score)
		fmt.Printf("    Published: %s | Source: %s\n", article.PubDate, article.Source.Domain)
		
		// Show a snippet of the description
		if article.Description != "" && len(article.Description) > 100 {
			fmt.Printf("    Description: %s...\n", article.Description[:97])
		}
	}
}

// entitySearch demonstrates searching for people and entities
func entitySearch(client *perigon.Client, ctx context.Context) {
	// Search for people
	people, err := client.People.List(ctx, perigon.PersonListParams{
		Name: perigon.String("Elon Musk"),
		Size: perigon.Int(1),
	})

	if err != nil {
		log.Printf("Error searching people: %v", err)
		return
	}

	if len(people.Results) > 0 {
		person := people.Results[0]
		fmt.Printf("Found person: %s\n", person.Name)
		fmt.Printf("Description: %s\n", person.Description)
		fmt.Printf("Wikidata ID: %s\n", person.WikidataID)
	}

	// Search for articles mentioning this person
	articles, err := client.All.List(ctx, perigon.AllListParams{
		PersonName: []string{"Elon Musk"},
		Size:       perigon.Int(3),
		SortBy:     perigon.AllEndpointSortByDate,
		From:       perigon.Time(time.Now().AddDate(0, 0, -7)), // Last week
	})

	if err != nil {
		log.Printf("Error searching articles by person: %v", err)
		return
	}

	fmt.Printf("\nRecent articles mentioning Elon Musk:\n")
	for _, article := range articles.Articles {
		fmt.Printf("  • %s\n", article.Title)
		fmt.Printf("    Published: %s\n", article.PubDate)
		
		// Show other people mentioned in the same article
		if len(article.People) > 1 {
			otherPeople := []string{}
			for _, p := range article.People {
				if p.Name != "Elon Musk" && p.Name != "" {
					otherPeople = append(otherPeople, p.Name)
				}
			}
			if len(otherPeople) > 0 && len(otherPeople) <= 3 {
				fmt.Printf("    Other people mentioned: %s\n", strings.Join(otherPeople, ", "))
			}
		}
	}
}

// sourceAnalysis demonstrates source and journalist analysis
func sourceAnalysis(client *perigon.Client, ctx context.Context) {
	// Analyze tech-focused sources
	sources, err := client.Sources.List(ctx, perigon.SourceListParams{
		Category:       []string{"Tech"},
		MinMonthlyPosts: perigon.Int(100), // Active sources
		Size:           perigon.Int(5),
		SortBy:         perigon.SortByRelevance,
	})

	if err != nil {
		log.Printf("Error in source analysis: %v", err)
		return
	}

	fmt.Printf("Found %d active tech sources:\n", len(sources.Results))
	for _, source := range sources.Results {
		fmt.Printf("  • %s (%s)\n", source.Name, source.Domain)
		fmt.Printf("    Monthly Posts: %d | Monthly Visits: %d\n", 
			source.AvgMonthlyPosts, source.MonthlyVisits)
		fmt.Printf("    Country: %s | Paywall: %t\n", 
			source.Location.Country, source.Paywall)
	}

	// Search for journalists in tech
	journalists, err := client.Journalists.List(ctx, perigon.JournalistListParams{
		Category:        []string{"Tech"},
		MinMonthlyPosts: perigon.Int(10),
		Size:            perigon.Int(3),
	})

	if err != nil {
		log.Printf("Error in journalist search: %v", err)
		return
	}

	fmt.Printf("\nActive tech journalists:\n")
	for _, journalist := range journalists.Results {
		fmt.Printf("  • %s\n", journalist.FullName)
		fmt.Printf("    Title: %s\n", journalist.Title)
		fmt.Printf("    Monthly Posts: %d\n", journalist.AvgMonthlyPosts)
		if journalist.TwitterHandle != "" {
			fmt.Printf("    Twitter: @%s\n", journalist.TwitterHandle)
		}
	}
}

// topicDiscovery demonstrates topic-based content discovery
func topicDiscovery(client *perigon.Client, ctx context.Context) {
	// Discover trending topics
	topics, err := client.Topics.List(ctx, perigon.TopicListParams{
		Category: perigon.String("Tech"),
		Size:     perigon.Int(10),
	})

	if err != nil {
		log.Printf("Error in topic discovery: %v", err)
		return
	}

	fmt.Printf("Available tech topics:\n")
	for i, topic := range topics.Data {
		if i >= 5 { // Limit output
			break
		}
		fmt.Printf("  • %s\n", topic.Name)
	}

	// Search stories for a specific topic
	if len(topics.Data) > 0 {
		selectedTopic := topics.Data[0].Name
		stories, err := client.Stories.List(ctx, perigon.StoryListParams{
			Topic:            []string{selectedTopic},
			MinUniqueSources: perigon.Int(3),
			Size:             perigon.Int(3),
			SortBy:           perigon.SortByUpdatedAt,
		})

		if err != nil {
			log.Printf("Error searching stories by topic: %v", err)
			return
		}

		fmt.Printf("\nStories about '%s':\n", selectedTopic)
		for _, story := range stories.Results {
			fmt.Printf("  • %s\n", story.Name)
			fmt.Printf("    Articles: %d | Sources: %d\n", 
				story.UniqueCount, len(story.UniqueSources))
			fmt.Printf("    Last Updated: %s\n", story.UpdatedAt)
			
			// Show top companies mentioned in the story
			if len(story.TopCompanies) > 0 {
				companies := []string{}
				for i, comp := range story.TopCompanies {
					if i >= 3 { // Limit to top 3
						break
					}
					companies = append(companies, comp.Name)
				}
				fmt.Printf("    Top Companies: %s\n", strings.Join(companies, ", "))
			}
		}
	}
} 