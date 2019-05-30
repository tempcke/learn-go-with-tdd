package main

import (
	"testing"
	"reflect"
)

func TestPerimeter(t *testing.T) {
	assertPerimeter := func(t *testing.T, expected float64, actual float64) {
		t.Helper()

		if actual != expected {
			t.Errorf("got %.2f want %.2f", actual, expected)
		}
	}

	t.Run("rectangles", func(t *testing.T) {

		r1 := Rectangle{Width: 10.0, Height: 10.0}
		r2 := Rectangle{Width: 10.0, Height: 20.0}

		assertPerimeter(t, 40.0, r1.Perimeter())
		assertPerimeter(t, 60.0, r2.Perimeter())
	})
}

func TestArea(t *testing.T) {
	for _, test := range shapesForAreaTests() {
		testName := reflect.TypeOf(test.shape).String()
		t.Run(testName, func(t *testing.T) {
			assertShapeArea(t, test.expect, test.shape)
		})
	}
}

type areaTest struct {
	shape Shape
	expect float64
}

func shapesForAreaTests() []areaTest {
	return []areaTest {
		{shape: Rectangle{Width: 2.0, Height: 3.0}, expect: 6.0},
		{shape: Rectangle{Width: 3.0, Height: 4.0}, expect: 12.0},
		{shape: Rectangle{Width: 12, Height: 6}, expect: 72.0},
		{shape: Circle{Radius: 10}, expect: 314.1592653589793},
		{shape: Triangle{Width: 12, Height: 6}, expect: 36.0},
	}
}

func assertShapeArea(t *testing.T, expected float64, shape Shape) {
	actual := shape.Area()
	if actual != expected {
		t.Errorf("got %.2f want %.2f", actual, expected)
	}
}
