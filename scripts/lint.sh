#!/usr/bin/env bash

set -eo pipefail

echo "Running linters..."

# Format check
echo "Checking code formatting..."
gofmt -l .

# Run golangci-lint
echo "Running golangci-lint..."
golangci-lint run --timeout=10m

echo "Linting completed successfully"

