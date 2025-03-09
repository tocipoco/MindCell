#!/usr/bin/env bash

set -eo pipefail

echo "Running tests..."

# Run unit tests
echo "Running unit tests..."
go test ./... -v -race -coverprofile=coverage.txt -covermode=atomic

# Display coverage
echo "Test coverage:"
go tool cover -func=coverage.txt

echo "Tests completed successfully"

