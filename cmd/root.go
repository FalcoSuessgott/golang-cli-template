package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var version = ""

func newRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "golang-cli-template",
		Short: "golang-cli project template demo application",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println(cmd.UsageString())
			return nil
		},
	}

	cmd.AddCommand(newVersionCmd()) // version subcommand
	cmd.AddCommand(newExampleCmd()) // example subcommand

	return cmd
}

// Execute invokes the command.
func Execute() error {
	return newRootCmd().Execute()
}
