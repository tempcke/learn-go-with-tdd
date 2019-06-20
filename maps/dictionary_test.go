package main

import "testing"

func TestSearch(t *testing.T) {
	dictionary := map[string]string{"test": "this is just a test"}

	assertEquals(t, "this is just a test", Search(dictionary, "test"))
}

func assertEquals(t *testing.T, expected string, actual string) {
	t.Helper()

	if actual != expected {
		t.Errorf(
			"Strings are not the same\nexpected\t'%s'\nreceived\t'%s'",
			expected,
			actual)
	}
}