package mumbleing

import "testing"

type requirment struct {
	input    string
	expected string
}

var requirments = []requirment{
	{"", ""},
	{"a", "A"},
	{"ab", "A-Bb"},
	{"abC", "A-Bb-Ccc"},
	{"aBCd", "A-Bb-Ccc-Dddd"},
	{"QWERTY", "Q-Ww-Eee-Rrrr-Ttttt-Yyyyyy"},
}

func TestMumble(t *testing.T) {
	for _, req := range requirments {
		t.Run(req.input+" to "+req.expected, func(t *testing.T) {
			checkRequirement(t, req)
		})
	}
}

func checkRequirement(t *testing.T, req requirment) {
	result := mumbleLetters(req.input)
	if req.expected != result {
		t.Errorf(
			"\nInput %s\nWant  %s\nGot   %s\n",
			req.input,
			req.expected,
			result)
	}
}
