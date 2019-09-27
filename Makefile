HAVE_GOLINT:=$(shell which golint)

## Go
.PHONY: setup lint test build run
setup:
	@echo "Start setup"
	@env GO111MODULE=on go mod vendor

lint: setup golint
	@echo "Check lint"
	@golint $(shell go list ./...|grep -v vendor)
	@go vet ./...

test: setup
	@echo "go test"
	@go test

build: setup
	@echo "build"
	@GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ./bin/bot ./cmd/bot

run: setup
	@echo "go run"
	@go run ./cmd/bot/main.go

## Docker local
CONTAINER_PREFIX:=number-hit-bot

.PHONY: dstart dstop dstatus dlogin dclean dlog
dstart: setup
	@echo "docker start"
	@docker-compose up -d

dstop:
	@echo "docker stop"
	@docker-compose stop

# restart container
drestart:
	@make dstop
	@make dstart

dstatus:
	@echo "docker status"
	@docker ps --filter name=$(CONTAINER_PREFIX)

dlogin:
	@echo "docker login"
	@docker exec -it $(shell docker ps --all --format "{{.Names}}" | peco) /bin/sh

dclean:
	@echo "docker clean"
	@docker ps --all --filter name=$(CONTAINER_PREFIX) --quiet | xargs docker rm --force

dlog:
	@echo "docker log"
	@docker-compose logs -f $(shell docker ps --all --format "{{.Names}}" | peco | cut -d"_" -f2)


## Install package
.PHONY: golint
golint:
ifndef HAVE_GOLINT
	@echo "Installing linter"
	@go get -u github.com/golang/lint/golint
endif
