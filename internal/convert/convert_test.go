package convert

import (
	"fmt"
	"testing"
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
			input:    123,
			expected: 123,
			err:      false,
		},
		{
			name:     "valid string",
			input:    "123",
			expected: 123,
			err:      false,
		},
		{
			name:  "invalid string",
			input: "foo",
			err:   true,
		},
		{
			name:  "invalid bool",
			input: false,
			err:   true,
		},
	}

	for i, tc := range testCases {
		res, err := ToInteger(tc.input)

		if tc.err != (err != nil) {
			t.Error(
				fmt.Sprintf(
					"[%d] Convert '%v', expected error: %v, got error: %v",
					i,
					tc.input,
					tc.err,
					(err != nil),
				),
			)
		} else if err == nil && tc.expected != res {
			fmt.Sprintf(
				"[%d] Convert '%v', expected: %d, got: %d",
				i,
				tc.input,
				tc.expected,
				res,
			)
		}
	}
}
