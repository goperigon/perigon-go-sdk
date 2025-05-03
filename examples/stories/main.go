package main

import (
	"context"
	"fmt"

	"github.com/goperigon/perigon-go-sdk"
)

func main() {
	client := perigon.NewClient()
	ctx := context.Background()

	stories, err := client.Stories.List(ctx, perigon.StoryListParams{
		Q:    perigon.String("\"Ukraine War\""),
		Name: perigon.String("\"Ukraine\" AND \"War\""),
	})
	if err != nil {
		panic(err)
	}

	for _, story := range stories.Results {
		fmt.Printf("Story Name: %s\n", story.Name)
		fmt.Printf("Story Last Updated: %s\n", story.UpdatedAt)
	}
}
