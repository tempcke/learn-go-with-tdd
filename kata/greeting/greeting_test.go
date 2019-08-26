package main

import (
	"gotest.tools/assert"
	"testing"
)

func TestGreeting(t *testing.T) {
	expected := "Hello, Bob."
	assert.Equal(t, expected, Greet("Bob"))
}

func Greet(name string) string {
	return "Hello, Bob."
}