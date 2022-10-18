package cmd

import (
	"context"
	"flag"
	"fmt"
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
		{[]string{"1", "2"}, subcommands.ExitUsageError},
		{[]string{"-x", "1", "2"}, subcommands.ExitUsageError},
		{[]string{"-m", "1", "2", "3"}, subcommands.ExitUsageError},
		{[]string{"-m", "foo", "bar"}, subcommands.ExitUsageError},
	}

	opts := rootOptions{
		version: "dev",
		verbose: false,
	}

	for _, tc := range testCases {

		command := exampleCommand(&opts)

		fs := flag.NewFlagSet("test", flag.ContinueOnError)

		command.SetFlags(fs)

		fs.Parse(tc.arguments)

		exitStatus := command.Execute(context.Background(), fs)

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
