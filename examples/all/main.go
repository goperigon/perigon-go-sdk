package main

import (
	"context"
	"fmt"

	"github.com/goperigon/perigon-go-sdk/v2"
)

func main() {
	client := perigon.NewClient()
	ctx := context.Background()

	all, err := client.All.List(ctx, perigon.AllListParams{
		Q: perigon.String("Tariffs"),
	})
	if err != nil {
		panic(err)
	}

	for _, article := range all.Articles {
		fmt.Printf("Title: %s\n", article.Title)
		fmt.Printf("PubDate: %s\n", article.PubDate)

	}
}
