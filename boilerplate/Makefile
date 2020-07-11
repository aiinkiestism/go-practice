# Meta info
NAME := boilerplate
VERSION := $(gobump show -r)
REEVISION := $(shell git rev-parse --short HEAD)
LDFLAGS := "-X main.revision=$(REVISION)"

export GO111MODULE=on

## Install dependencies
.PHONY: deps
test: deps
	go test ./...

## Lint
.PHONY: lint
lint: devel-deps
	go vet ./...
	golint -set_exit_status ./...

## Build buinaries ex. make bin/boilerplate
bin/%: cmd/%/main.go deps
	go build -ldflags "$(LDFLAGS)" -o $@ $<

## Build binary
.PHONY: build
build: bin/boilerplate

## Show help
.PHONY: help
help:
	@make2help $(MAKEFILE_LIST)
