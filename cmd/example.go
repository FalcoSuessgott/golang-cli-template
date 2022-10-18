package cmd

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/google/subcommands"
	"github.com/mprimi/golang-cli-template/internal/convert"
	"github.com/mprimi/golang-cli-template/pkg/example"
)

type exampleCmd struct {
	rootOpts *rootOptions
	multiply bool
	add      bool
}

func exampleCommand(rootOpts *rootOptions) subcommands.Command {
	return &exampleCmd{
		rootOpts: rootOpts,
	}
}

func (_ *exampleCmd) Name() string { return "example" }

func (_ *exampleCmd) Synopsis() string { return "Example add/subtract" }

func (_ *exampleCmd) Usage() string { return "example (-a|-m) number number\n" }

func (cmd *exampleCmd) SetFlags(f *flag.FlagSet) {
	f.BoolVar(&cmd.multiply, "m", false, "multiply")
	f.BoolVar(&cmd.add, "a", false, "add")
}

func (cmd *exampleCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {

	const (
		numberOfArgs = 2
	)

	if cmd.rootOpts.verbose {
		fmt.Printf("Args: %v\n", f.Args())
	}

	if len(f.Args()) != numberOfArgs {
		log.Printf("Takes two arguments\n")
		return subcommands.ExitUsageError
	}

	if cmd.add == cmd.multiply { // XOR
		log.Printf("Must select add or multiply\n")
		return subcommands.ExitUsageError
	}

	values := make([]int, numberOfArgs)
	for i, a := range f.Args() {
		v, err := convert.ToInteger(a)
		if err != nil {
			log.Printf("Invalid number: %s\n", a)
			return subcommands.ExitUsageError
		}
		values[i] = v
	}

	if cmd.multiply {
		fmt.Printf("%d\n", example.Multiply(values[0], values[1]))
	} else if cmd.add {
		fmt.Printf("%d\n", example.Add(values[0], values[1]))
	}

	return subcommands.ExitSuccess
}
