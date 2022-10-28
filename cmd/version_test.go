package cmd

import (
	"context"
	"flag"
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
	}

	for i, tc := range testCases {

		opts := rootOptions{
			verbose: tc.verbose,
		}
		fs := flag.NewFlagSet("test", flag.ContinueOnError)

		command := versionCommand()

		command.SetFlags(fs)

		err := fs.Parse(tc.arguments)
		if err != nil {
			t.Logf("[%d] Ignoring argument parsing error: %v", i, err)
		}

		exitStatus := command.Execute(context.Background(), fs, &opts)

		if exitStatus != tc.expectedExitStatus {
			t.Errorf(
				"[%d], Args: %v, expected: %v, got: %v\n",
				i,
				tc.arguments,
				tc.expectedExitStatus,
				exitStatus,
			)
		}
	}
}
