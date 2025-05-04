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
