package cmd

import (
	"bytes"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRootCommandOutput(t *testing.T) {
	cmd := newRootCmd()
	b := bytes.NewBufferString("")

	cmd.SetArgs([]string{"-h"})
	cmd.SetOut(b)

	cmdErr := cmd.Execute()
	require.NoError(t, cmdErr)

	out, err := ioutil.ReadAll(b)
	require.NoError(t, err)

	assert.Equal(t, "golang-cli project template demo application\n\n"+cmd.UsageString(), string(out))
	assert.Nil(t, cmdErr)
}
