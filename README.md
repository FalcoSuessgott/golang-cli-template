# golang-cli-template
A general purpose project template for golang CLI applications

[![Test](https://github.com/FalcoSuessgott/golang-cli-template/actions/workflows/test.yml/badge.svg)](https://github.com/FalcoSuessgott/golang-cli-template/actions/workflows/test.yml) [![golangci-lint](https://github.com/FalcoSuessgott/golang-cli-template/actions/workflows/lint.yml/badge.svg)](https://github.com/FalcoSuessgott/golang-cli-template/actions/workflows/lint.yml) [![Go Report Card](https://goreportcard.com/badge/github.com/FalcoSuessgott/golang-cli-template)](https://goreportcard.com/report/github.com/FalcoSuessgott/golang-cli-template) [![Go Reference](https://pkg.go.dev/badge/github.com/FalcoSuessgott/golang-cli-template.svg)](https://pkg.go.dev/github.com/FalcoSuessgott/golang-cli-template) [![codecov](https://codecov.io/gh/FalcoSuessgott/golang-cli-template/branch/main/graph/badge.svg?token=Y5K4SID71F)](https://codecov.io/gh/FalcoSuessgott/golang-cli-template)

This template serves as a starting point for golang commandline applications it is based on golang projects that I consider high quality and various other useful blog posts that helped me understanding golang better.

# Features
- [goreleaser](https://goreleaser.com/) with `deb.` and `.rpm` package releasing
- [golangci-lint](https://golangci-lint.run/) for linting and formatting
- [Github Actions](.github/worflows) Stages (Linting, Testing, Releasing)
- [Gitlab CI](.gitlab-ci.yml) Configuration (Linting, Testing, Releasing)
- [cobra](https://cobra.dev/) example setup including tests
- [Makefile](Makefile) - with various useful targets (see Makefile Targets)
- [Github Pages](_config.yml) using [jekyll-theme-minimal](https://github.com/pages-themes/minimal) (checkout [https://falcosuessgott.github.io/golang-cli-template/](https://falcosuessgott.github.io/)golang-cli-template/)

# Project Layout
* [assets/](https://pkg.go.dev/github.com/FalcoSuessgott/golang-cli-template/assets) => docs, images, etc
* [cmd/](https://pkg.go.dev/github.com/FalcoSuessgott/golang-cli-template/cmd)  => commandline configurartions (flags, subcommands)
* [pkg/](https://pkg.go.dev/github.com/FalcoSuessgott/golang-cli-template/pkg)  => packages that are okay to import for other projects
* [internal/](https://pkg.go.dev/github.com/FalcoSuessgott/golang-cli-template/pkg)  => packages that are only for project internal purposes

# How to use this template
```sh
GITHUB_USER="my-github-user"
PROJECT="new-golang-project"
git clone git@github.com:FalcoSuessgott/golang-cli-template.git "$PROJECT"
cd "$PROJECT"
rm -rf .git
find . -type f -exec sed -i "s/FalcoSuessgott\/golang-cli-template/$GITHUB_USER\/$PROJECT/g" {} +
make fmt
git init
git add .
git commit -m "initial commit"
git remote add origin "git@github.com:$GITHUB_USER/$PROJECT.git"
```

# Demo Application

```sh
$> golang-cli-template                                                  
golang-cli project template demo application

Usage:
  golang-cli-template [flags]
  golang-cli-template [command]

Available Commands:
  example     example subcommand which adds or multiplies two given integers
  help        Help about any command
  version     Displays d4sva binary version

Flags:
  -h, --help   help for golang-cli-template

Use "golang-cli-template [command] --help" for more information about a command.
```

```sh
$> golang-cli-template 2 5 --add                                                
7

$> golang-cli-template 2 5 --multiply
10
```

# Makefile Targets

## lint
```sh
make lint
```

## format
```sh
make fmt
```

## formatcheck
```sh
make fmtcheck
```

## deps
```sh
make deps
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

# Contribute
If you find issues in that setup or have some nice features / improvements, I would welcome an issue or a PR :)


# Ideas
- [ ] implement a `create` subcommand that preconfigures this project setup and all its dependencies
- [ ] introduce viper config examples
