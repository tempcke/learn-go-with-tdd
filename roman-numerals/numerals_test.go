package romanNumerals

import (
	"fmt"
	"gotest.tools/assert"
	"testing"
)

var IntRomMap = map[int]string {
	1: "I",
	2: "II",
	3: "III",
	4: "IV",
	5: "V",
	6: "VI",
	7: "VII",
	8: "VIII",
	9: "IX",
	10: "X",
	11: "XI",
	12: "XII",
	14: "XIV",
	15: "XV",
	18: "XVIII",
	19: "XIX",
	20: "XX",
	38: "XXXVIII",
	39: "XXXIX",
	40: "XL",
	47: "XLVII",
	49: "XLIX",
	50: "L",
	90: "XC",
	100: "C",
	400: "CD",
	500: "D",
	900: "CM",
	1000: "M",
	1984: "MCMLXXXIV",
	3999: "MMMCMXCIX",
	2014: "MMXIV",
	1006: "MVI",
	798: "DCCXCVIII",
}


func TestConvertToRoman(t *testing.T) {
	for n, r := range IntRomMap {
		testName := fmt.Sprintf("Test %d is %s", n, r)
		t.Run(testName, func(t *testing.T) {
			assertRoman(t, n, r)
		})
	}
}

func assertRoman(t *testing.T, n int, expected string) {
	r := ConvertToRoman(n)
	msg := "\nConvertToRoman(%d)\n Want: %s\n Got:  %s"
	assert.Equal(
		t, expected, r,
		msg, n, expected, r)
}

