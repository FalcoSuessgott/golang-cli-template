package cmd

import (
	"bytes"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExampleCommandMultiply(t *testing.T) {
	cmd := newExampleCmd()
	b := bytes.NewBufferString("")

	cmd.SetArgs([]string{"2", "3", "--multiply"})
	cmd.SetOut(b)

	err := cmd.Execute()
	require.NoError(t, err)

	out, err := ioutil.ReadAll(b)
	require.NoError(t, err)

	assert.Equal(t, "6", string(out))
}

func TestExampleCommandMultiplyInvalidArgs(t *testing.T) {
	cmd := newExampleCmd()
	b := bytes.NewBufferString("")

	cmd.SetArgs([]string{"2", "s", "--multiply"})
	cmd.SetOut(b)

	err := cmd.Execute()
	assert.Error(t, err)
}

func TestExampleAdd(t *testing.T) {
	cmd := newExampleCmd()
	b := bytes.NewBufferString("")

	cmd.SetArgs([]string{"2", "3", "-a"})
	cmd.SetOut(b)

	err := cmd.Execute()
	require.NoError(t, err)

	out, err := ioutil.ReadAll(b)
	require.NoError(t, err)

	assert.Equal(t, "5", string(out))
}
