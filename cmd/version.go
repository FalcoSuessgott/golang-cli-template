package cmd

import (
	"flag"
	"fmt"

	"github.com/google/subcommands"
)

// Substituted at build time using -X linker flag
// https://pkg.go.dev/cmd/link
// See Makefile and Goreleaser configuration
var Name = "go-bench-away"
var Version = "dev"
var SHA = "?"
var BuildDate = "?"

type versionCmd struct {
	metaCommand
}

func versionCommand() subcommands.Command {
	return &versionCmd{
		metaCommand: metaCommand{
			name:     "version",
			synopsis: "print version",
			usage:    "version",
			setFlags: func(_ *flag.FlagSet) {},
			execute: func(rootOpts *rootOptions, f *flag.FlagSet) error {
				if rootOpts.verbose {
					fmt.Printf("Printing version string\n")
				}

				fmt.Printf(
					"%s version:%s (%s) (built: %s)\n",
					Name,
					Version,
					SHA,
					BuildDate,
				)
				return nil
			},
		},
	}
}
