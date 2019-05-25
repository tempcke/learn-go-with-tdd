package main

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	assertPerimeter := func(t *testing.T, expected float64, actual float64) {
		t.Helper()

		if actual != expected {
			t.Errorf("got %.2f want %.2f", actual, expected)
		}
	}

	t.Run("rectangles", func(t *testing.T) {

		r1 := Rectangle{10.0, 10.0}
		r2 := Rectangle{10.0, 20.0}

		assertPerimeter(t, 40.0, r1.Perimeter())
		assertPerimeter(t, 60.0, r2.Perimeter())
	})
}

func TestArea(t *testing.T) {
	for _, test := range shapesForAreaTests() {
		assertShapeArea(t, test.expect, test.shape)
	}
}

type areaTest struct {
	shape Shape
	expect float64
}

func shapesForAreaTests() []areaTest {
	return []areaTest {
		{Rectangle{2.0, 3.0}, 6.0},
		{Rectangle{3.0, 4.0}, 12.0},
		{Rectangle{12, 6}, 72.0},
		{Circle{10}, 314.1592653589793},
		{Triangle{12, 6}, 36.0},
	}
}

func assertShapeArea(t *testing.T, expected float64, shape Shape) {
	actual := shape.Area()
	if actual != expected {
		t.Errorf("got %.2f want %.2f", actual, expected)
	}
}
