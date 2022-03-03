package example

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	testCases := []struct {
		name     string
		a        int
		b        int
		expected int
	}{
		{
			name:     "",
			a:        1,
			b:        2,
			expected: 3,
		},
		{
			name:     "",
			a:        2,
			b:        2,
			expected: 4,
		},
	}

	for _, tc := range testCases {
		assert.Equal(t, tc.expected, Add(tc.a, tc.b))
	}
}

func TestMultiply(t *testing.T) {
	testCases := []struct {
		name     string
		a        int
		b        int
		expected int
	}{
		{
			name:     "",
			a:        1,
			b:        2,
			expected: 2,
		},
		{
			name:     "",
			a:        2,
			b:        2,
			expected: 4,
		},
	}

	for _, tc := range testCases {
		assert.Equal(t, tc.expected, Multiply(tc.a, tc.b))
	}
}
