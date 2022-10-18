projectname?=golang-cli-template

default: build

.PHONY: build
build:
	@go build -ldflags "-X main.version=dev" -o $(projectname)

.PHONY: install
install:
	@go install -ldflags "-X main.version=dev"

.PHONY: run
run:
	@go run -ldflags "-X main.version=dev"  main.go

.PHONY: fmtcheck
fmtcheck:
	@sh -c "find . -name '*.go' | grep -v vendor | xargs gofumpt -l -s"

PHONY: test
test:
	go test -v ./...

PHONY: clean
clean:
	@rm -rf coverage.out dist/ $(projectname)

PHONY: cover
cover:
	go test -v -race $(shell go list ./... | grep -v /vendor/) -v -coverprofile=coverage.out
	go tool cover -func=coverage.out

PHONY: fmt
fmt:
	gofumpt -w -s  .

PHONY: lint
lint:
	golangci-lint run -c .golang-ci.yml

PHONY: lint-fix
lint-fix:
	golangci-lint run -c .golang-ci.yml --fix

PHONY: install-tools
install-tools:
	@go install fmt
