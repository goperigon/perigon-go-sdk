package perigon

import (
	"context"
	"io"
	"net/http"
	"reflect"
	"strings"
	"testing"
	"time"
)

type roundTripFunc func(*http.Request) (*http.Response, error)

func (f roundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req)
}

func TestNewSetsBearerAndBaseURL(t *testing.T) {
	var got *http.Request
	doer := &http.Client{Transport: roundTripFunc(func(req *http.Request) (*http.Response, error) {
		got = req
		return &http.Response{
			StatusCode: http.StatusOK,
			Header:     http.Header{"Content-Type": []string{"application/json"}},
			Body:       io.NopCloser(strings.NewReader(`{"articles":[],"numResults":0,"status":200}`)),
			Request:    req,
		}, nil
	})}

	client, err := New("secret", WithHTTPClient(doer))
	if err != nil {
		t.Fatal(err)
	}
	from := time.Date(2022, 2, 1, 0, 0, 0, 0, time.UTC)
	resp, err := client.SearchArticlesWithResponse(context.Background(), &SearchArticlesParams{
		Q:      Ptr("ai"),
		Source: Ptr([]string{"reuters.com", "apnews.com"}),
		From:   Ptr(from),
	})
	if err != nil {
		t.Fatal(err)
	}
	if err := Failed(resp.StatusCode(), resp.Body); err != nil {
		t.Fatal(err)
	}
	if got == nil {
		t.Fatal("request was not sent")
	}
	if got.Header.Get("Authorization") != "Bearer secret" {
		t.Fatalf("authorization = %q", got.Header.Get("Authorization"))
	}
	if got.URL.Scheme != "https" || got.URL.Host != "api.perigon.io" {
		t.Fatalf("url = %s", got.URL.String())
	}
	if !strings.HasSuffix(got.URL.Path, "/v1/articles/all") {
		t.Fatalf("path = %s", got.URL.Path)
	}
	query := got.URL.Query()
	if query.Get("q") != "ai" {
		t.Fatalf("query = %s", got.URL.RawQuery)
	}
	if !reflect.DeepEqual(query["source"], []string{"reuters.com", "apnews.com"}) {
		t.Fatalf("source = %#v", query["source"])
	}
	if query.Get("from") != from.Format(time.RFC3339Nano) {
		t.Fatalf("from = %s", query.Get("from"))
	}
	if resp.JSON200 == nil {
		t.Fatal("JSON200 is nil")
	}
}

func TestNewWithBaseURL(t *testing.T) {
	var gotHost string
	doer := &http.Client{Transport: roundTripFunc(func(req *http.Request) (*http.Response, error) {
		gotHost = req.URL.Host
		return &http.Response{
			StatusCode: http.StatusOK,
			Header:     http.Header{"Content-Type": []string{"application/json"}},
			Body:       io.NopCloser(strings.NewReader(`{}`)),
			Request:    req,
		}, nil
	})}

	client, err := New("secret", WithBaseURL("http://127.0.0.1:4010"), WithHTTPClient(doer))
	if err != nil {
		t.Fatal(err)
	}
	if _, err := client.SearchArticlesWithResponse(context.Background(), &SearchArticlesParams{}); err != nil {
		t.Fatal(err)
	}
	if gotHost != "127.0.0.1:4010" {
		t.Fatalf("host = %s", gotHost)
	}
}

func TestNewRequiresAPIKey(t *testing.T) {
	for _, key := range []string{"", " ", "\t"} {
		if _, err := New(key); err == nil {
			t.Fatalf("New(%q) succeeded", key)
		}
	}
}

func TestNewRejectsRelativeBaseURL(t *testing.T) {
	doer := &http.Client{Transport: roundTripFunc(func(req *http.Request) (*http.Response, error) {
		t.Fatal("request was sent")
		return nil, nil
	})}
	for _, baseURL := range []string{"api.perigon.io", "https://api.perigon.io?x=1", "https://api.perigon.io#frag"} {
		if _, err := New("secret", WithBaseURL(baseURL), WithHTTPClient(doer)); err == nil {
			t.Fatalf("New base URL %q succeeded", baseURL)
		}
	}
}

func TestNewDefaultTransportTimeout(t *testing.T) {
	client, err := New("secret")
	if err != nil {
		t.Fatal(err)
	}
	raw, ok := client.ClientInterface.(*Client)
	if !ok {
		t.Fatalf("client type = %T", client.ClientInterface)
	}
	httpClient, ok := raw.Client.(*http.Client)
	if !ok {
		t.Fatalf("doer type = %T", raw.Client)
	}
	transport, ok := httpClient.Transport.(*http.Transport)
	if !ok {
		t.Fatalf("transport type = %T", httpClient.Transport)
	}
	if transport.ResponseHeaderTimeout != defaultResponseHeaderTimeout {
		t.Fatalf("timeout = %s", transport.ResponseHeaderTimeout)
	}
}

func TestFailed(t *testing.T) {
	if err := Failed(200, []byte(`{"ok":true}`)); err != nil {
		t.Fatalf("2xx returned %v", err)
	}
	if err := Failed(204, nil); err != nil {
		t.Fatalf("204 returned %v", err)
	}
	err := Failed(401, []byte(`{"message":"no"}`))
	apiErr, ok := err.(*APIError)
	if !ok {
		t.Fatalf("error type = %T", err)
	}
	if apiErr.StatusCode != 401 {
		t.Fatalf("status = %d", apiErr.StatusCode)
	}
	if string(apiErr.Body) != `{"message":"no"}` {
		t.Fatalf("body = %s", apiErr.Body)
	}
}

func TestDecoded(t *testing.T) {
	value := &QuerySearchResult{}
	got, err := Decoded(200, []byte(`{}`), value)
	if err != nil || got != value {
		t.Fatalf("got %v err %v", got, err)
	}
	if _, err := Decoded[QuerySearchResult](200, nil, nil); err == nil {
		t.Fatal("nil body succeeded")
	}
	err = nil
	if _, err = Decoded(401, []byte(`{"message":"no"}`), value); err == nil {
		t.Fatal("401 succeeded")
	}
	if _, ok := err.(*APIError); !ok {
		t.Fatalf("error type = %T", err)
	}
}
