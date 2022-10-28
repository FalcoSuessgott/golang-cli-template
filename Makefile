project?=github.com/mprimi/golang-cli-template
projectname?=golang-cli-template
version?=dev
sha?=$(shell git rev-parse --short HEAD)
date?=$(shell date "+%Y-%m-%d_%H:%M:%S")
linkflags="-X $(project)/cmd.Version=$(version) -X $(project)/cmd.SHA=$(sha) -X $(project)/cmd.BuildDate=$(date)"

default: build

.PHONY: build install run test clean cover vet fmt lint mod check

build:
	@go build -ldflags $(linkflags) -o $(projectname)

install:
	@go install -ldflags $(linkflags)

run:
	@go run -ldflags $(linkflags) main.go

test:
	@go test -v -count=1 ./...

clean:
	@rm -rf coverage.out dist/ $(projectname)

cover:
	@go test -race $(shell go list ./... | grep -v /vendor/) -v -coverprofile=coverage.out
	@go tool cover -func=coverage.out
	@go tool cover -html coverage.out -o coverage.html

vet:
	@go vet ./...

fmt:
	@gofmt -w -s  .

lint: # Depends on https://github.com/golangci/golangci-lint
	@golangci-lint run -c .golangci-lint.yml --fix

mod:
	@go mod tidy

check: mod fmt test lint cover vet
