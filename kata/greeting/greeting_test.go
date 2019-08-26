package main

import (
	"gotest.tools/assert"
	"testing"
)

func TestGreeting(t *testing.T) {
	t.Run("Requirement 1, simple greeting", func(t *testing.T) {
	  expected := "Hello, Bob."
	  assert.Equal(t, expected, Greet("Bob"))
  })

	t.Run("Requirment 2, handle null name", func(t *testing.T) {
		var name string
		expected := "Hello, my friend."
		assert.Equal(t, expected, Greet(name))
	})
}

