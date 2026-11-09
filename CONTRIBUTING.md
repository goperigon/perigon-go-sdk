## Setting up the environment

To set up the repository, run:

```sh
$ ./scripts/bootstrap
$ ./scripts/lint
```

This installs the Go module dependencies. Using the module needs [Go 1.22.5+](https://go.dev/doc/install). Regenerating
it needs Go 1.25+.

## Regenerating the SDK

Most of the client is generated from the public OpenAPI spec
by [oapi-codegen](https://github.com/oapi-codegen/oapi-codegen). `./scripts/generate.sh` downloads
`https://api.perigon.io/v1/openapi/public-sdk`, applies `scripts/spec-filters.json`, and writes `perigon.gen.go`.

```sh
$ ./scripts/generate.sh
```

`./scripts/generate.sh` needs Go 1.25 or newer because the generator module in `tools/go.mod` requires it.

Do not edit `perigon.gen.go` by hand. The hand-written files are `client.go`, `errors.go`, `ptr.go`, `generate.go`,
`client_test.go`, and `integration_test.go`. `scripts/spec-filters.json` lists the hidden fields and excluded paths.
Remove an entry there to expose a field or endpoint again.

The preprocessed snapshot is written to `.openapi-spec/spec.local.json` and is not committed.

The generator is pinned in `tools/go.mod`. `./scripts/generate.sh` is the regeneration entry point, and `go generate`
runs that script from the repository root. `New` does not read `PERIGON_API_KEY`, `PERIGON_BASE_URL`, or
`PERIGON_CUSTOM_HEADERS`.

## Adding and running examples

```sh
$ go run ./examples/<your-example>
```

Set `PERIGON_API_KEY` before running an example against the live API.

## Using the repository from source

```sh
$ go mod edit -replace github.com/goperigon/perigon-go-sdk/v3=/path/to/perigon-go-sdk
```

## Running tests

```sh
$ ./scripts/test
```

`./scripts/test` runs the unit tests. They use a fake HTTP transport and do not call the Perigon API. Integration tests
call `https://api.perigon.io` and need `PERIGON_API_KEY`:

```sh
$ PERIGON_API_KEY="your-api-key" go test -tags integration -run TestIntegration ./...
```

## Formatting

```sh
$ ./scripts/format
```
