package convert

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidInteger(t *testing.T) {
	i, ok := ConvertToInteger(2)
	assert.Equal(t, i, 2)
	assert.Nil(t, ok, "expect error to be nil")
}

func TestInValidIntegers(t *testing.T) {
	_, ok := ConvertToInteger("string")
	assert.Error(t, ok)
}
