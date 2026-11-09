# Perigon Go API Library

The Perigon Go library provides access to the [Perigon REST API](https://docs.perigon.io/discuss) from applications
written in Go.

It is generated with [oapi-codegen](https://github.com/oapi-codegen/oapi-codegen) from the public OpenAPI specification.

## Installation

```sh
go get github.com/goperigon/perigon-go-sdk/v3
```

This library requires Go 1.22.5+. Regenerating it requires Go 1.25+.

## Usage

```go
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
		Q:    perigon.Ptr("artificial intelligence"),
		Size: perigon.Ptr(int32(5)),
	})
	if err != nil {
		panic(err)
	}
	body, err := perigon.Decoded(resp.StatusCode(), resp.Body, resp.JSON200)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", body.Articles)
}
```

`New` sends the API key as `Authorization: Bearer` and returns an error when the key is empty. It does not read
`PERIGON_API_KEY`, `PERIGON_BASE_URL`, or `PERIGON_CUSTOM_HEADERS`. `WithBaseURL` changes the host, which defaults to
`https://api.perigon.io`. `WithHTTPClient` replaces the HTTP client. `WithRequestEditorFn` edits the request before it
is sent.

Method names come from the OpenAPI `operationId`. Optional fields are pointers; use `perigon.Ptr`. `perigon.Decoded`
returns the decoded 2xx body. It returns an `*APIError` for a non-2xx status and an error when that body was not
decoded. `perigon.Failed` is the status check `Decoded` uses.

More samples are in [examples](examples).

## Tests

Unit tests check the client against a fake HTTP transport. They do not call the Perigon API and do not need an API key.

```sh
go test ./...
```

Integration tests call `https://api.perigon.io` and use your quota. They need `PERIGON_API_KEY`.

```sh
PERIGON_API_KEY="your-api-key" go test -tags integration -run TestIntegration ./...
```

## Publish

Set the version in `VERSION`. It must be `3.x.y` so the tag matches this module path. Merge to `main` and run the
Publish workflow. It tests the module, tags `v{version}`, and asks the Go module proxy to fetch that tag. It does not
change the version.

## Regenerating

```sh
./scripts/generate.sh
```

The script downloads the public spec, applies [scripts/spec-filters.json](scripts/spec-filters.json), regenerates
`perigon.gen.go`, and runs the tests. Run it with Go 1.25 or newer. See [CONTRIBUTING.md](CONTRIBUTING.md).

## License

MIT
