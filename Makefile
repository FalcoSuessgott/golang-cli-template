projectname?=golang-cli-template

default: build

.PHONY: build install run test clean cover vet fmt lint mod check

build:
	@go build -ldflags "-X main.version=dev" -o $(projectname)

install:
	@go install -ldflags "-X main.version=dev"

run:
	@go run -ldflags "-X main.version=dev"  main.go

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
