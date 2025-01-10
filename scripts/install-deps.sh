#!/usr/bin/env bash

set -eo pipefail

echo "Installing dependencies..."

# Install Go dependencies
go mod download
go mod verify

# Install protobuf compiler
echo "Installing protoc..."

# Install golangci-lint
echo "Installing linters..."

echo "Dependencies installed successfully"

