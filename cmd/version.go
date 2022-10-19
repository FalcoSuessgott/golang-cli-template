package cmd

import (
	"flag"
	"fmt"

	"github.com/google/subcommands"
)

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
					fmt.Printf("Printing version string: %s\n", rootOpts.version)
				}
				if rootOpts.version == "" {
					return fmt.Errorf("Version not set")
				}

				fmt.Printf("%s\n", rootOpts.version)
				return nil
			},
		},
	}
}
