package cmd

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/google/subcommands"
	"github.com/mprimi/golang-cli-template/internal/convert"
	"github.com/mprimi/golang-cli-template/pkg/example"
)

type exampleCmd struct {
	metaCommand
	multiply bool
	add      bool
}

func exampleCommand() subcommands.Command {
	return &exampleCmd{
		metaCommand: metaCommand{
			name:     "example",
			synopsis: "perform numeric operations",
			usage:    "example (-m|-a) <number> <number>",
		},
	}
}

func (ec *exampleCmd) SetFlags(f *flag.FlagSet) {
	f.BoolVar(&ec.multiply, "m", false, "multiply")
	f.BoolVar(&ec.add, "a", false, "add")
}

func (ec *exampleCmd) Execute(_ context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	const (
		numberOfArgs = 2
	)

	var rootOpts *rootOptions = args[0].(*rootOptions)

	if rootOpts.verbose {
		fmt.Printf("Args: %v\n", f.Args())
	}

	if len(f.Args()) != numberOfArgs {
		fmt.Fprintf(os.Stderr, "Must pass two arguments")
		return subcommands.ExitUsageError
	} else if ec.add == ec.multiply { // XOR
		fmt.Fprintf(os.Stderr, "Must chose add or multiply")
		return subcommands.ExitUsageError
	}
	values := make([]int, numberOfArgs)
	for i, a := range f.Args() {
		v, err := convert.ToInteger(a)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Invalid number argument: %s", a)
			return subcommands.ExitUsageError
		}
		values[i] = v
	}

	if ec.multiply {
		fmt.Printf("%d\n", example.Multiply(values[0], values[1]))
	} else if ec.add {
		fmt.Printf("%d\n", example.Add(values[0], values[1]))
	}
	return subcommands.ExitSuccess
}
