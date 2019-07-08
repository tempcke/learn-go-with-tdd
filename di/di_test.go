package main

import (
	"bytes"
	"gotest.tools/assert"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Chris")

	assert.Equal(t, "hello, Chris", buffer.String())
}