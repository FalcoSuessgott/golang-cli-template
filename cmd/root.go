package cmd

import (
	"github.com/spf13/cobra"
)

var (
	version = ""
)

func newRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "golang-cli-template",
		Short: "Develop and deploy infrastructure configurations on multiple provider",
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}

	cmd.AddCommand(newVersionCmd()) // version subcommand
	cmd.AddCommand(newExampleCmd()) // examble subcommand

	return cmd
}

// Execute invokes the command.
func Execute() error {
	return newRootCmd().Execute()
}
