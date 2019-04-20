package iteration

import (
	"fmt"
	"testing"
)

func ExampleRepeat() {
	str := Repeat("Foo", 5)
	fmt.Println(str)
	// Output: FooFooFooFooFoo
}

func TestRepeat(t *testing.T) {

	assertEquals := func(expected string, actual string) {
		t.Helper() // required so that fail output does not point inside this function but to where this function was called

		if actual != expected {
			t.Errorf("got '%s' want '%s'", actual, expected)
		}
	}

	t.Run("should repeat string five times", func(t *testing.T) {
		assertEquals("", Repeat("a", 0))
		assertEquals("a", Repeat("a", 1))
		assertEquals("aa", Repeat("a", 2))
		assertEquals("aaaaa", Repeat("a", 5))
		assertEquals("aaaaaaaaaa", Repeat("a", 10))
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}