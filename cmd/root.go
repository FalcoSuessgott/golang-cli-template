package cmd

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/google/subcommands"
)

type rootOptions struct {
	version string
	verbose bool
}

// Entry point
func Run(name, version string, args []string) int {
	var rootOps rootOptions
	rootOps.version = version

	rootFs := flag.NewFlagSet("", flag.ExitOnError)
	rootFs.BoolVar(&rootOps.verbose, "v", false, "verbose")

	cmdr := subcommands.NewCommander(rootFs, name)

	commandsMap := map[string][]subcommands.Command{
		"example": {
			exampleCommand(&rootOps),
		},
		"help": {
			versionCommand(&rootOps),
			cmdr.HelpCommand(),
			cmdr.FlagsCommand(),
			cmdr.CommandsCommand(),
		},
	}

	for groupName, commands := range commandsMap {
		for _, command := range commands {
			cmdr.Register(command, groupName)
		}
	}

	err := rootFs.Parse(args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to parse arguments: %v\n", err)
		return 1
	}
	return int(cmdr.Execute(context.Background()))
}
