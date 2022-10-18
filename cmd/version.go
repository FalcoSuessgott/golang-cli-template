package cmd

import (
	"context"
	"flag"
	"fmt"

	"github.com/google/subcommands"
)

type versionCmd struct {
	rootOpts *rootOptions
}

func versionCommand(rootOpts *rootOptions) subcommands.Command {
	return &versionCmd{
		rootOpts: rootOpts,
	}
}

func (_ *versionCmd) Name() string { return "version" }

func (_ *versionCmd) Synopsis() string { return "Print version" }

func (_ *versionCmd) Usage() string { return "version\n" }

func (_ *versionCmd) SetFlags(f *flag.FlagSet) {}

func (cmd *versionCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	if cmd.rootOpts.verbose {
		fmt.Printf("Printing version string: %s\n", cmd.rootOpts.version)
	}
	if cmd.rootOpts.version == "" {
		return subcommands.ExitFailure
	}
	fmt.Printf("%s\n", cmd.rootOpts.version)
	return subcommands.ExitSuccess
}
