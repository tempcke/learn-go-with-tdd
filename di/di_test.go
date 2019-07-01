package main

import (
	"bytes"
	"gotest.tools/assert"
	"testing"
)

//func TestDi(t *testing.T) {
//	t.Error("https://github.com/quii/learn-go-with-tests/blob/master/dependency-injection.md")
//}

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Chris")

	assert.Equal(t, "hello, Chris", buffer.String())
}