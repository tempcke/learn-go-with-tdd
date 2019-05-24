package main

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	assertPerimeter := func(expected float64, actual float64) {
		t.Helper()

		if actual != expected {
			t.Errorf("got %.2f want %.2f", actual, expected)
		}
	}

	t.Run("rectangles", func(t *testing.T) {

		r1 := Rectangle{10.0, 10.0}
		r2 := Rectangle{10.0, 20.0}

		assertPerimeter(40.0, r1.Perimeter())
		assertPerimeter(60.0, r2.Perimeter())
	})
}

func TestArea(t *testing.T) {
	assertArea := func(expected float64, actual float64) {
		t.Helper()

		if actual != expected {
			t.Errorf("got %.2f want %.2f", actual, expected)
		}
	}

	t.Run("rectangles", func(t *testing.T) {

		r1 := Rectangle{2.0,3.0}
		r2 := Rectangle{3.0,4.0}

		assertArea(6.0, r1.Area())
		assertArea(12.0, r2.Area())
	})

	t.Run("circles", func(t *testing.T) {
		assertArea(314.1592653589793, Circle{10}.Area())
	})
}