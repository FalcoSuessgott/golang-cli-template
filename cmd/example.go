package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

type exampleOptions struct {
	multiply bool
	sum bool
}

func defaultExampleOptions() *exampleOptions {
	return &exampleOptions{}
}

func newExampleCmd() *cobra.Command {

	o := defaultExampleOptions()

	cmd := &cobra.Command{
		Use:          "example",
		Short:        "example subcommand",
		Args:         cobra.NoArgs,
		SilenceUsage: true,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Fprintf(cmd.OutOrStdout(), "%s\n", version)
		},
	}

	cmd.Flags().BoolVarP(&o.multiply, "multiply", "m", o.multiply, "multiply")
	cmd.Flags().BoolVarP(&o.sum, "sum", "s", o.sum, "sum")

	return cmd
}

