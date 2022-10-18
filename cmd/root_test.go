package cmd

import (
	"testing"
)

func TestRun(t *testing.T) {

	testCases := []struct {
		arguments        []string
		expectedExitCode int
	}{
		// Usage errors
		{[]string{}, 2},
		{[]string{"blah"}, 2},
		// Valid
		{[]string{"commands"}, 0},
		{[]string{"flags"}, 0},
		{[]string{"help"}, 0},
		{[]string{"help", "version"}, 0},
		{[]string{"help", "example"}, 0},
		{[]string{"-v", "version"}, 0},
		{[]string{"example", "-m", "0", "0"}, 0},
	}

	for i, tc := range testCases {
		exitCode := Run("test", "v", tc.arguments)
		if exitCode != tc.expectedExitCode {
			t.Errorf(
				"[%d] Args: %v expected: %d, got: %d",
				i,
				tc.arguments,
				tc.expectedExitCode,
				exitCode,
			)
		}
	}
}
