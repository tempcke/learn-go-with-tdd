package integers

import (
	"fmt"
	"testing"
)

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}

func TestAdder(t *testing.T) {
	assertSumIs := func(x, y, expected int) {
		t.Helper()
		sum := Add(x, y)

		if sum != expected {
			t.Errorf("expected '%d' but got '%d'", expected, sum)
		}
	}

	t.Run("2 + 2 = 4", func(t *testing.T) {
		assertSumIs(2, 2, 4)
	})

	t.Run("3 + 0 = 3", func(t *testing.T) {
		assertSumIs(3, 0, 3)
	})
}