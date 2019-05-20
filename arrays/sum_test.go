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
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSumTail(t *testing.T) {
	numbers := []int{1, 2, 3, 5}
	sum := SumTail(numbers)
	expected := 10

	if expected != sum {
		t.Errorf("got %d want %d given, %v", sum, expected, numbers)
	}
}

func TestSumTails(t *testing.T) {
	sum := SumAllTails(
		[]int{1, 2},
		[]int{0, 9},
	)
	expected := []int{2, 9}

	if !reflect.DeepEqual(sum, expected) {
		t.Errorf("got %v want %v", sum, expected)
	}
}

func TestSumTails2(t *testing.T) {
	assertTailSums := func(expected []int, actual []int) {
		t.Helper()
		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("got %v want %v", actual, expected)
		}
	}

	t.Run("make the sums of some slices", func(t *testing.T) {
		actual := SumAllTails(
			[]int{1, 2, 3},
			[]int{5, 8, 13},
		)
		expected := []int{5, 21}
		
		assertTailSums(expected, actual)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		actual := SumAllTails(
			[]int{},
			[]int{3, 4, 5},
		)
		expected := []int{0, 9}

		assertTailSums(expected, actual)
	})
}
