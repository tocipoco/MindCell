#!/usr/bin/make -f

VERSION := $(shell echo $(shell git describe --tags) | sed 's/^v//')
COMMIT := $(shell git log -1 --format='%H')
LEDGER_ENABLED ?= true
BINDIR ?= $(GOPATH)/bin
BUILDDIR ?= $(CURDIR)/build
DOCKER := $(shell which docker)

export GO111MODULE = on

###############################################################################
###                                  Build                                  ###
###############################################################################

all: install

build: go.sum
	@echo "Building MindCell..."
	@go build -mod=readonly -o build/mindcelld ./cmd/mindcelld

install: go.sum
	@echo "Installing MindCell..."
	@go install -mod=readonly ./cmd/mindcelld

go.sum: go.mod
	@echo "Ensuring dependencies..."
	@go mod tidy
	@go mod download

###############################################################################
###                                  Proto                                  ###
###############################################################################

proto-all: proto-gen proto-swagger-gen

proto-gen:
	@echo "Generating protobuf files..."
	@./scripts/protocgen.sh

proto-swagger-gen:
	@echo "Generating Swagger documentation..."
	@./scripts/protoc-swagger-gen.sh

###############################################################################
###                                 Testing                                 ###
###############################################################################

test:
	@go test -mod=readonly ./...

test-coverage:
	@go test -mod=readonly -coverprofile=coverage.txt -covermode=atomic ./...

test-verbose:
	@go test -mod=readonly -v ./...

###############################################################################
###                                Linting                                  ###
###############################################################################

format:
	@find . -name '*.go' -type f -not -path "./vendor*" -not -path "*.git*" | xargs gofmt -w -s
	@find . -name '*.go' -type f -not -path "./vendor*" -not -path "*.git*" | xargs goimports -w

lint:
	@golangci-lint run --out-format=tab

###############################################################################
###                                Cleaning                                 ###
###############################################################################

clean:
	@rm -rf $(BUILDDIR)/
	@rm -rf vendor/

.PHONY: all build install go.sum proto-all proto-gen proto-swagger-gen test test-coverage test-verbose format lint clean

