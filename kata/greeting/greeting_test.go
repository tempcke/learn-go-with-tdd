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

	t.Run("Requirement 2, handle null name", func(t *testing.T) {
		var name string
		expected := "Hello, my friend."
		assert.Equal(t, expected, Greet(name))
	})

	t.Run("Requirement 3, shout for uppercase names", func(t *testing.T) {
		name := "JERRY"
		expected := "HELLO JERRY!"
		assert.Equal(t, expected, Greet(name))
	})

	t.Run("Requirement 4, multiple names", func(t *testing.T) {
		expected := "Hello, Jill and Jane."
		actual := Greet("Jill", "Jane")
		assert.Equal(t, expected, actual)
	})
}

