package main

import (
	"testing"
)

func TestGreeting(t *testing.T) {
	t.Run("Requirement 1, simple greeting", func(t *testing.T) {
	  expected := "Hello, Bob."
	  assertEqual(t, expected, Greet("Bob"))
  })

	t.Run("Requirement 2, handle null name", func(t *testing.T) {
		var name string
		expected := "Hello, my friend."
		assertEqual(t, expected, Greet(name))
	})

	t.Run("Requirement 3, shout for uppercase names", func(t *testing.T) {
		name := "JERRY"
		expected := "HELLO JERRY!"
		assertEqual(t, expected, Greet(name))
	})

	t.Run("Requirement 4, two names", func(t *testing.T) {
		expected := "Hello, Jill and Jane."
		actual := Greet("Jill", "Jane")
		assertEqual(t, expected, actual)
	})

	t.Run("Requirement 5, three or more names", func(t *testing.T) {
		expected := "Hello, Amy, Brian, and Charlotte."
		actual := Greet("Amy", "Brian", "Charlotte")
		assertEqual(t, expected, actual)
	})

	t.Run("Req 6, separate normal and shouting names", func(t *testing.T) {
		expected := "Hello, Amy and Charlotte. AND HELLO BRIAN!"
		actual := Greet("Amy", "BRIAN", "Charlotte")
		assertEqual(t, expected, actual)
	})

	t.Run("Req 7, detect multiple names within each string arg", func(t *testing.T) {
		expected := "Hello, Bob, Charlie, and Dianne."
		actual := Greet("Bob", "Charlie, Dianne")
		assertEqual(t, expected, actual)
	})
}

func assertEqual(t *testing.T, expected, actual string) {
	if expected != actual {
		t.Errorf("\nWant  %s\nGot   %s\n", expected, actual)
	}
}
