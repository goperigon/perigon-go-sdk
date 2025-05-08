package main

import (
	"context"
	"fmt"

	"github.com/goperigon/perigon-go-sdk/v2"
)

func main() {
	client := perigon.NewClient()
	ctx := context.Background()

	resp, err := client.Vector.News.Search(ctx, perigon.VectorNewsSearchParams{
		Prompt: "Elon Musk impact on Tesla stock price",
	})

	if err != nil {
		panic(err)
	}

	for _, res := range resp.Results {
		article := res.Data
		fmt.Printf("Title: %s\n", article.Title)
		fmt.Printf("PubDate: %s\n", article.PubDate)

	}
}
