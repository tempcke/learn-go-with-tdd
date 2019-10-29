package romanNumerals

import (
	"gotest.tools/assert"
	"testing"
)

func TestRomanNumerals(t *testing.T) {
	n := 1
	expect := "I"
	assertRoman(t, n, expect)
}

func assertRoman(t *testing.T, n int, expected string) {
	r := ConvertToRoman(1)
	msg := "\nConvertToRoman(%d)\n Want: %s\n Got:  %s"
	assert.Equal(
		t, expected, r,
		msg, n, expected, r)
}