package cmd

import (
	"context"
	"flag"
	"fmt"
	"github.com/google/subcommands"
	"testing"
)

func TestVersionCommand(t *testing.T) {

	testCases := []struct {
		verbose            bool
		version            string
		arguments          []string
		expectedExitStatus subcommands.ExitStatus
	}{
		{false, "test", []string{}, subcommands.ExitSuccess},
		{true, "test", []string{}, subcommands.ExitSuccess},
		{true, "test", []string{"blah"}, subcommands.ExitSuccess},
		{true, "", []string{}, subcommands.ExitFailure},
	}

	for _, tc := range testCases {

		opts := rootOptions{
			version: tc.version,
			verbose: tc.verbose,
		}
		fs := flag.NewFlagSet("test", flag.ContinueOnError)
		fs.Parse(tc.arguments)

		exitStatus := versionCommand(&opts).Execute(context.Background(), fs)

		if exitStatus != tc.expectedExitStatus {
			t.Error(
				fmt.Sprintf(
					"Args: %v, expected: %v, got: %v\n",
					tc.arguments,
					tc.expectedExitStatus,
					exitStatus,
				),
			)
		}
	}
}
