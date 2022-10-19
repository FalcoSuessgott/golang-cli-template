package cmd

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/google/subcommands"
)

type rootOptions struct {
	// Build-time properties
	name    string
	version string
	// Top-level options
	verbose bool
}

type metaCommand struct {
	name     string
	synopsis string
	usage    string
	setFlags func(*flag.FlagSet)
	execute  func(*rootOptions, *flag.FlagSet) error
}

func (mc *metaCommand) Name() string { return mc.name }

func (mc *metaCommand) Synopsis() string { return mc.synopsis }

func (mc *metaCommand) Usage() string { return mc.usage }

func (mc *metaCommand) SetFlags(f *flag.FlagSet) {
	if mc.setFlags != nil {
		mc.setFlags(f)
	}
}

func (mc *metaCommand) Execute(_ context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	if mc.execute == nil {
		fmt.Fprintf(os.Stderr, "Not implemented")
		return subcommands.ExitFailure
	}

	var rootOpts *rootOptions = args[0].(*rootOptions)

	err := mc.execute(rootOpts, f)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Command %s failed: %v", mc.name, err)
		return subcommands.ExitFailure
	}

	return subcommands.ExitSuccess
}

// Entry point
func Run(name, version string, args []string) int {
	var rootOps = rootOptions{
		name:    name,
		version: version,
	}

	rootFs := flag.NewFlagSet("", flag.ExitOnError)
	rootFs.BoolVar(&rootOps.verbose, "v", false, "verbose")

	cmdr := subcommands.NewCommander(rootFs, name)

	commandsMap := map[string][]subcommands.Command{
		"example": {
			exampleCommand(),
		},
		"help": {
			versionCommand(),
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
	return int(cmdr.Execute(context.Background(), &rootOps))
}
