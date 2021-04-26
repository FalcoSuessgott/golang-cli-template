# golang-cli-template
A general purpose project template for golang CLI applications

[![Test](https://github.com/FalcoSuessgott/golang-cli-template/actions/workflows/test.yml/badge.svg)](https://github.com/FalcoSuessgott/golang-cli-template/actions/workflows/test.yml) [![golangci-lint](https://github.com/FalcoSuessgott/golang-cli-template/actions/workflows/lint.yml/badge.svg)](https://github.com/FalcoSuessgott/golang-cli-template/actions/workflows/lint.yml) [![Go Report Card](https://goreportcard.com/badge/github.com/FalcoSuessgott/golang-cli-template)](https://goreportcard.com/report/github.com/FalcoSuessgott/golang-cli-template) [![Go Reference](https://pkg.go.dev/badge/github.com/FalcoSuessgott/golang-cli-template.svg)](https://pkg.go.dev/github.com/FalcoSuessgott/golang-cli-template) [![codecov](https://codecov.io/gh/FalcoSuessgott/golang-cli-template/branch/main/graph/badge.svg?token=Y5K4SID71F)](https://codecov.io/gh/FalcoSuessgott/golang-cli-template)

This template servers as a starting point for golang commandline applications based on [project-layout](https://github.com/golang-standards/project-layout).

# Features
- [goreleaser](https://goreleaser.com/)
- [golangci-lint](https://golangci-lint.run/)
- Golang Github Actions Stages (Linting, Testing, Releasing)
- Fully tested [cobra](https://cobra.dev/) setup
- Makefile

# Project Layout
* [assets/](https://pkg.go.dev/github.com/FalcoSuessgott/golang-cli-template/assets) => docs, images, etc
* [cmd/](https://pkg.go.dev/github.com/FalcoSuessgott/golang-cli-template/cmd)  => commandline configurartions (flags, subcommands)
* [pkg/](https://pkg.go.dev/github.com/FalcoSuessgott/golang-cli-template/pkg)  => packages that are okay to import for other projects
* [internal/](https://pkg.go.dev/github.com/FalcoSuessgott/golang-cli-template/pkg)  => packages that are only for project internal purposes

# Makefile Targes

## format
```sh
make fmt
```

## formatcheck
```sh
make fmtcheck
```

## build
```sh
make build
```

## test
```sh
make test
```

## cover
```sh
make cover
```