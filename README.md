 > This is Marco's fork of:

# [golang-cli-template](https://github.com/mprimi/golang-cli-template)
A general purpose project template for golang CLI applications

The upstream version has a lot of nice features. This is a stripped down version for personal use.

Removed in this fork:

 * A well-written README, including live badges
 * GH Pages with Jekyll theme and assets
 * GitLab CI
 * Install script
 * GH Semantic Releaser
 * Docker
 * `deb`, `rpm`, `apk`
 * Pre-commit checks
 * Codecov
 * Dropped `testify` dependency
 * Replace `cobra` (~300 indirect dependencies!), with [`subcommands`](https://pkg.go.dev/github.com/google/subcommands)

Modified:

 * Makefile
 * Subcommands and tests

Loose ends:
 * Dig into linting and coverage settings
