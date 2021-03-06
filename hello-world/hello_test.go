package main

import (
	"testing"
)

func TestHello(t *testing.T) {

	assertEquals := func(expected string, actual string) {
		t.Helper() // required so that fail output does not point inside this function but to where this function was called

		if actual != expected {
			t.Errorf("got '%s' want '%s'", actual, expected)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		assertEquals("Hello, Chris", Hello("Chris", ""))
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		assertEquals("Hello, World", Hello("", ""))
	})

	t.Run("can get greeting in spanish", func(t *testing.T) {
		greeting := Hello("Chris", "spanish")
		assertEquals("Hola, Chris", greeting)
	})

	t.Run("can get greeting in french", func(t *testing.T) {
		greeting := Hello("Chris", "french")
		assertEquals("Bonjour, Chris", greeting)
	})
}