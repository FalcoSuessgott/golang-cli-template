package cmd

import (
	"github.com/spf13/cobra"
)

var (
	version = ""
)

func NewRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "golang-cli-template",
		Short: "Develope and deploy infrastructure configurations on multiple provider",
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}

	return cmd
}

// Execute invokes the command.
func Execute() error {
	return NewRootCmd().Execute()
}
