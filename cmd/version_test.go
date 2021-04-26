package cmd

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVersionCommand(t *testing.T) {
	version = "v1.0.0"
	cmd := newVersionCmd()
	b := bytes.NewBufferString("")
	cmd.SetOut(b)
	cmd.Execute()

	out, err := ioutil.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}

	output := strings.Split(string(out), "\n")

	assert.Equal(t, fmt.Sprintf("%s", version), output[0])
}
