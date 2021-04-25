.PHONY: fmtcheck fmt build lint cover test

build:
	go build -ldflags "-X codehub.sva.de/deploy4sva/d4sva/cmd.version=$(shell git describe --abbrev=0 --tags)" -o d4sva

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