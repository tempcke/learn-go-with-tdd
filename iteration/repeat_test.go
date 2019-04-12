package iteration

import "testing"

func TestRepeat(t *testing.T) {

	assertEquals := func(expected string, actual string) {
		t.Helper() // required so that fail output does not point inside this function but to where this function was called

		if actual != expected {
			t.Errorf("got '%s' want '%s'", actual, expected)
		}
	}

	t.Run("should repeat string five times", func(t *testing.T) {
		assertEquals("aaaaa", Repeat("a"))
	})
}