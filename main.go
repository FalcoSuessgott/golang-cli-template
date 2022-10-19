package main

import (
	"os"

	"github.com/mprimi/golang-cli-template/cmd"
)

var version = "dev"

const (
	name    = "golang-cli-template"
)

func main() {
	os.Exit(cmd.Run(name, version, os.Args[1:]))
}
