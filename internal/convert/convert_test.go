package convert

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestToInteger(t *testing.T) {
	testCases := []struct {
		name     string
		input    interface{}
		expected int
		err      bool
	}{
		{
			name:     "valid int",
			input:    1,
			expected: 1,
			err:      false,
		},
		{
			name:  "invalid bool",
			input: false,
			err:   true,
		},
	}

	for _, tc := range testCases {
		res, err := ToInteger(tc.input)

		if tc.err {
			require.Error(t, err, tc.name)
		} else {
			require.Equal(t, tc.expected, res)
		}
	}
}
