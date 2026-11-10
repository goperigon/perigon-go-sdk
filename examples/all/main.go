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

	resp, err := client.SearchArticlesWithResponse(context.Background(), &perigon.SearchArticlesParams{
		Q: perigon.Ptr("Tariffs"),
	})
	if err != nil {
		panic(err)
	}
	body, err := perigon.Decoded(resp.StatusCode(), resp.Body, resp.JSON200)
	if err != nil {
		panic(err)
	}

	for _, article := range body.Articles {
		fmt.Printf("Title: %s\n", text(article.Title))
		fmt.Printf("PubDate: %s\n", text(article.PubDate))
	}
}

func text(v *string) string {
	if v == nil {
		return ""
	}
	return *v
}
