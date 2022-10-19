package main

import (
	"os"

	"github.com/mprimi/golang-cli-template/cmd"
)

const (
	name = "golang-cli-template"
)

// Substituted at build time by goreleaser
var version = "dev"
var sha = "?"
var buildDate = "?"

func main() {
	os.Exit(cmd.Run(name, version, sha, buildDate, os.Args[1:]))
}
