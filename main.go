package main

import (
	"os"

	"github.com/mprimi/golang-cli-template/cmd"
)

func main() {
	os.Exit(cmd.Run(os.Args[1:]))
}
