package example

import (
	"testing"
)

func TestSum(t *testing.T) {
	testCases := []struct {
		a                int
		b                int
		expectedAdd      int
		expectedMultiply int
	}{
		{
			a:                1,
			b:                2,
			expectedAdd:      3,
			expectedMultiply: 2,
		},
		{
			a:                2,
			b:                2,
			expectedAdd:      4,
			expectedMultiply: 4,
		},
	}

	for i, tc := range testCases {
		add := Add(tc.a, tc.b)
		multiply := Multiply(tc.a, tc.b)

		if tc.expectedAdd != add {
			t.Errorf(
				"[%d] Expected %d+%d=%d, got %d",
				i,
				tc.a,
				tc.b,
				tc.expectedAdd,
				add,
			)
		} else if tc.expectedMultiply != multiply {
			t.Errorf(
				"[%d], Expected %d*%d=%d, got %d",
				i,
				tc.a,
				tc.b,
				tc.expectedMultiply,
				multiply,
			)
		}
	}
}
