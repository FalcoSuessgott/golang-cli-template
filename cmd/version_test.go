package cmd

import (
	"bytes"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVersionCommand(t *testing.T) {
	version = "v1.0.0"
	cmd := newVersionCmd()
	b := bytes.NewBufferString("")
	cmd.SetOut(b)

	err := cmd.Execute()
	require.NoError(t, err)

	out, err := ioutil.ReadAll(b)
	require.NoError(t, err)

	assert.Equal(t, version, string(out))
}
