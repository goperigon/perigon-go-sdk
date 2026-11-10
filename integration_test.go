//go:build integration

package perigon

import (
	"context"
	"os"
	"testing"
	"time"
)

func TestIntegrationSearchArticles(t *testing.T) {
	apiKey := os.Getenv("PERIGON_API_KEY")
	if apiKey == "" {
		t.Fatal("PERIGON_API_KEY must be set to run integration tests")
	}

	client, err := New(apiKey)
	if err != nil {
		t.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	resp, err := client.SearchArticlesWithResponse(ctx, &SearchArticlesParams{
		Q:    Ptr("perigon"),
		Size: Ptr(int32(1)),
	})
	if err != nil {
		t.Fatal(err)
	}
	if _, err := Decoded(resp.StatusCode(), resp.Body, resp.JSON200); err != nil {
		t.Fatal(err)
	}
}
