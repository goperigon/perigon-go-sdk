package perigon

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// DefaultBaseURL is the Perigon API host used when WithBaseURL is not set.
const DefaultBaseURL = "https://api.perigon.io"

// defaultResponseHeaderTimeout bounds the time between a fully written request
// and the server's response headers. It does not apply to the response body.
const defaultResponseHeaderTimeout = 10 * time.Minute

// Option configures the client returned by New.
// WithBaseURL, WithHTTPClient, and WithRequestEditorFn are accepted options.
type Option = ClientOption

// New returns a client for the Perigon API.
// apiKey is sent as Authorization: Bearer. An empty key is an error.
// Pass WithBaseURL or WithHTTPClient to override the host or the HTTP client.
// New does not read PERIGON_API_KEY, PERIGON_BASE_URL, or PERIGON_CUSTOM_HEADERS.
func New(apiKey string, opts ...Option) (*ClientWithResponses, error) {
	apiKey = strings.TrimSpace(apiKey)
	if apiKey == "" {
		return nil, fmt.Errorf("perigon: api key is required")
	}

	// A sentinel server shows whether WithBaseURL ran. Applying options here
	// does not register request editors on the client that is returned.
	const untouched = "\x00"
	probe := &Client{Server: untouched}
	for _, opt := range opts {
		if err := opt(probe); err != nil {
			return nil, err
		}
	}
	if probe.Server != untouched {
		if err := validateBaseURL(probe.Server); err != nil {
			return nil, err
		}
	}

	clientOpts := make([]ClientOption, 0, len(opts)+2)
	clientOpts = append(clientOpts, WithRequestEditorFn(func(_ context.Context, req *http.Request) error {
		req.Header.Set("Authorization", "Bearer "+apiKey)
		return nil
	}))
	if probe.Client == nil {
		clientOpts = append(clientOpts, WithHTTPClient(defaultHTTPClient()))
	}
	clientOpts = append(clientOpts, opts...)
	return NewClientWithResponses(DefaultBaseURL, clientOpts...)
}

func validateBaseURL(raw string) error {
	parsed, err := url.Parse(raw)
	if err != nil || parsed.Scheme == "" || parsed.Host == "" || parsed.RawQuery != "" || parsed.Fragment != "" {
		return fmt.Errorf("perigon: base URL must be absolute and include a host")
	}
	return nil
}

func defaultHTTPClient() *http.Client {
	transport := http.DefaultTransport.(*http.Transport).Clone()
	transport.ResponseHeaderTimeout = defaultResponseHeaderTimeout
	return &http.Client{Transport: transport}
}
