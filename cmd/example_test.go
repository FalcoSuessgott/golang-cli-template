package cmd

import (
	"bytes"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExampleCommand(t *testing.T) {
	testCases := []struct {
		name     string
		args     []string
		expected string
		err      bool
	}{
		{
			name:     "normal multiply",
			args:     []string{"2", "3", "--multiply"},
			expected: "6\n",
			err:      false,
		},
		{
			name:     "invalid multiply",
			args:     []string{"2", "s", "--multiply"},
			expected: "",
			err:      true,
		},
		{
			name:     "valid add",
			args:     []string{"2", "3", "-a"},
			expected: "5\n",
			err:      false,
		},
		{
			name:     "invalid add",
			args:     []string{"s", "3", "-a"},
			expected: "",
			err:      true,
		},
	}

	for _, tc := range testCases {
		cmd := newExampleCmd()
		b := bytes.NewBufferString("")

		cmd.SetArgs(tc.args)
		cmd.SetOut(b)

		err := cmd.Execute()
		out, _ := io.ReadAll(b)

		if tc.err {
			assert.Error(t, err, tc.name)
		} else {
			assert.Equal(t, tc.expected, string(out), tc.name)
		}
	}
}
