package main

import (
	"context"
	"fmt"
	"os"

	"github.com/goperigon/perigon-go-sdk/v3"
)

func main() {
	client, err := perigon.New(os.Getenv("PERIGON_API_KEY"))
	if err != nil {
		panic(err)
	}

	resp, err := client.SearchStoriesWithResponse(context.Background(), &perigon.SearchStoriesParams{
		Q:    perigon.Ptr(`"Ukraine War"`),
		Name: perigon.Ptr(`"Ukraine" AND "War"`),
	})
	if err != nil {
		panic(err)
	}
	body, err := perigon.Decoded(resp.StatusCode(), resp.Body, resp.JSON200)
	if err != nil {
		panic(err)
	}

	for _, story := range body.Results {
		fmt.Printf("Story Name: %s\n", text(story.Name))
		fmt.Printf("Story Last Updated: %s\n", text(story.UpdatedAt))
	}
}

func text(v *string) string {
	if v == nil {
		return ""
	}
	return *v
}
