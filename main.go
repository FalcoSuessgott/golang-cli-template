package main

import (
	"os"

	"github.com/mprimi/golang-cli-template/cmd"
)

const (
	name    = "golang-cli-template"
	version = "dev" // Replaced at build time in release
)

func main() {
	os.Exit(cmd.Run(name, version, os.Args[1:]))
}
