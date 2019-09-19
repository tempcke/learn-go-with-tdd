package string_calculator

import (
	"testing"
)

type requirement struct {
	desc     string
	input    string
	expected int
}

var requirements = []requirement{
	{
		"An empty string returns zero",
		"",
		0,
	},
	{
		"A single number returns the value",
		"42",
		42,
	},
	{
		"Two numbers, comma delimited, returns the sum",
		"2,3",
		5,
	},
	{
		"Two numbers, newline delimited, returns the sum",
		"3\n4",
		7,
	},
	{
		"Three numbers, delimited either way, returns the sum",
		"3\n4\n5",
		12,
	},
	{
		"Three numbers, delimited either way, returns the sum",
		"3,4,5",
		12,
	},
}

func TestStringCalculator(t *testing.T) {
	for _, req := range requirements {
		t.Run(req.desc, func(t *testing.T) {
			checkRequirement(t, req)
		})
	}
}

func TestErrorHandling(t *testing.T) {
	t.Run("Negative numbers result in an error", func(t *testing.T) {
		input := "-42"
		n, e := Resolve(input)
		if e == nil {
			t.Errorf("\nExpected error due to negative number\n\tInput: %s\n\tResult: %d", input, n)
		}
	})
}

func checkRequirement(t *testing.T, req requirement) {
	result, e := Resolve(req.input)
	if e != nil {
		t.Errorf("\nUnexpected error\n\tInput: %s\n\tError: %s\n", req.input, e.Error())
	}
	if req.expected != result {
		t.Errorf("\nWant  %d\nGot   %d\n", req.expected, result)
	}
}
