package romanNumerals

import (
	"strings"
)

type RomanNumeral struct {
	Value int
	Symbol string
}

var RomanNumerals = []RomanNumeral{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func ConvertToRoman(n int) string {
	var result strings.Builder

	for _, rn := range RomanNumerals {
		for n >= rn.Value {
			result.WriteString(rn.Symbol)
			n -= rn.Value
		}
	}

	return result.String()
}

var RNs = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

func ConvertRomanToInt(roman string) int {
	total := 0
	prev := 0
	for i:=len(roman)-1; i>=0; i-- {
		char := string(roman[i])
		v := RNs[char]
		if v < prev {
			total -= v
		}	else {
			total += v
		}
		prev = v
	}
	return total
}