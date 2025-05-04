NAME := siggy
VERSION := 0.1.0
REVISION := $(shell git rev-parse --short HEAD)
LDFLAGS := "-X main.revision=$(REVISION)"

export GO111MODULE=on

## Start containers for services
.PHONY: run
run:
	./docker-compose-up.sh

## Execute the app
.PHONY: exec
exec:
	docker compose exec app bash

## Install dependencies
.PHONY: deps
deps:
	go mod download

## Run tests
.PHONY: test
test: deps
	ENV=test go test ./... -p=1

## Lint
.PHONY: lint
lint:
	golangci-lint run -c .golangci-lint.yml ./...
