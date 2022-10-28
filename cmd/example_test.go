package cmd

import (
	"context"
	"flag"
	"github.com/google/subcommands"
	"testing"
)

func TestExampleCommand(t *testing.T) {

	testCases := []struct {
		arguments          []string
		expectedExitStatus subcommands.ExitStatus
	}{
		{[]string{"-m", "1", "2"}, subcommands.ExitSuccess},
		{[]string{"-a", "1", "2"}, subcommands.ExitSuccess},
		{[]string{}, subcommands.ExitUsageError},
		{[]string{"-m", "-a", "1", "2"}, subcommands.ExitUsageError},
		{[]string{"1", "2"}, subcommands.ExitUsageError},
		{[]string{"-x", "1", "2"}, subcommands.ExitUsageError},
		{[]string{"-m", "1", "2", "3"}, subcommands.ExitUsageError},
		{[]string{"-m", "foo", "bar"}, subcommands.ExitUsageError},
	}

	opts := rootOptions{
		verbose: false,
	}

	for i, tc := range testCases {

		command := exampleCommand()

		fs := flag.NewFlagSet("test", flag.ContinueOnError)

		command.SetFlags(fs)

		err := fs.Parse(tc.arguments)
		if err != nil {
			t.Logf("[%d] Ignoring argument parsing error: %v", i, err)
		}

		exitStatus := command.Execute(context.Background(), fs, &opts)

		if exitStatus != tc.expectedExitStatus {
			t.Errorf(
				"[%d] Args: %v, expected: %v, got: %v\n",
				i,
				tc.arguments,
				tc.expectedExitStatus,
				exitStatus,
			)
		}
	}
}
