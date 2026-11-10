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

	resp, err := client.VectorSearchArticlesWithResponse(context.Background(), perigon.VectorSearchArticlesJSONRequestBody{
		Prompt: "Elon Musk impact on Tesla stock price",
	})
	if err != nil {
		panic(err)
	}
	body, err := perigon.Decoded(resp.StatusCode(), resp.Body, resp.JSON200)
	if err != nil {
		panic(err)
	}

	for _, result := range body.Results {
		if result.Data == nil {
			continue
		}
		fmt.Printf("Title: %s\n", text(result.Data.Title))
		fmt.Printf("PubDate: %s\n", text(result.Data.PubDate))
	}
}

func text(v *string) string {
	if v == nil {
		return ""
	}
	return *v
}
