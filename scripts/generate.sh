#!/usr/bin/env bash
set -euo pipefail

cd "$(dirname "$0")/.."

echo "==> Preprocessing OpenAPI spec"
go run ./scripts/preprocess

echo "==> Generating client"
go run -modfile=tools/go.mod github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen \
  -config oapi-codegen.yaml \
  .openapi-spec/spec.local.json

echo "==> Formatting"
gofmt -s -w perigon.gen.go

echo "==> Testing"
go test ./...
