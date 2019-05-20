package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	assertSum := func(expected int, numbers []int) {
		t.Helper() // required so that fail output does not point inside this function but to where this function was called

		sum := Sum(numbers)

		if expected != sum {
			t.Errorf("sum %d want %d given, %v", sum, expected, numbers)
		}
	}

	t.Run("collection of 5 numbers", func(t *testing.T) {
		assertSum(15, []int{1, 2, 3, 4, 5})
	})

	t.Run("collection of any size", func(t *testing.T) {
		assertSum(6, []int{1, 2, 3})
	})
}

func TestSumAll(t *testing.T) {

	got := SumAll([]int{1,2}, []int{0,9})
	want := []int{3, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
