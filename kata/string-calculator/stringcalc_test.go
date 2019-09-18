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
}

func TestStringCalculator(t *testing.T) {
	for _, req := range requirements {
		t.Run(req.desc, func(t *testing.T) {
			checkRequirement(t, req)
		})
	}
}

func checkRequirement(t *testing.T, req requirement) {
	result := resolve(req.input)
	if req.expected != result {
		t.Errorf("\nWant  %d\nGot   %d\n", req.expected, result)
	}
}
