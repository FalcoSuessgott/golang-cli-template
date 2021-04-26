package cmd

import (
	"errors"
	"strconv"

	"github.com/spf13/cobra"

	"github.com/FalcoSuessgott/golang-cli-template/internal/convert"
	"github.com/FalcoSuessgott/golang-cli-template/pkg/example"
)

type exampleOptions struct {
	multiply bool
	add      bool
}

func defaultExampleOptions() *exampleOptions {
	return &exampleOptions{}
}

func newExampleCmd() *cobra.Command {
	o := defaultExampleOptions()

	cmd := &cobra.Command{
		Use:          "example",
		Short:        "example subcommand which adds or multiplies two given integers",
		SilenceUsage: true,
		Args:         cobra.ExactArgs(2),
		RunE:         o.run,
	}

	cmd.Flags().BoolVarP(&o.multiply, "multiply", "m", o.multiply, "multiply")
	cmd.Flags().BoolVarP(&o.add, "add", "a", o.add, "add")

	return cmd
}

func (o *exampleOptions) parseArgs(args []string) ([]int, error) {
	values := make([]int, 2)

	for i, a := range args {
		v, err := convert.ConvertToInteger(a)
		if err != nil {
			return err
		}

		values[i] = v
	}

	return values, nil
}

func (o *exampleOptions) run(cmd *cobra.Command, args []string) error {
	values, err := o.parseArgs(args)
	if err != nil {
		return err
	}

	if o.multiply {
		example.Multiply(values[0], values[1])
	}

	if o.add {
		example.Add(values[0], values[1])
	}

	return nil
}
