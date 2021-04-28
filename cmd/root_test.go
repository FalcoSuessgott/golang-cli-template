package cmd

import (
	"bytes"
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
	if cmdErr != nil {
		t.Fail()
	}

	out, err := ioutil.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "golang-cli project template demo application\n\n"+cmd.UsageString(), string(out))
	assert.Nil(t, cmdErr)
}
