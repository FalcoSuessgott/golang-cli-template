.PHONY: fmtcheck fmt build lint cover test

build:
	go build -ldflags "-X github.com/FalcoSuessgott/golang-cli-template/cmd.version=$(shell git describe --abbrev=0 --tags)" -o golang-cli-template

fmtcheck:
	@sh -c "'$(CURDIR)/scripts/gofmtcheck.sh'"

test:
	go test -v ./...

cover:
	go test -v -race $(shell go list ./... | grep -v /vendor/) -v -coverprofile=coverage.out
	go tool cover -func=coverage.out

fmt:
	gofmt -w -s  .

lint:
	golangci-lint run -c .golang-ci.yml