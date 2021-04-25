package main

import (
	"os"

	"github.com/FalcoSuessgott/golang-cli-template/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
