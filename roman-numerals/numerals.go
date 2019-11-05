package romanNumerals

import "strings"

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
	//for n > 0 {
	//	switch {
	//	case n > 9:
	//		result.WriteString("X")
	//		n -= 10
	//	case n > 8:
	//		result.WriteString("IX")
	//		n -= 9
	//	case n > 4:
	//		result.WriteString("V")
	//		n -= 5
	//	case n > 3:
	//		result.WriteString("IV")
	//		n -= 4
	//	default:
	//		result.WriteString("I")
	//		n -= 1
	//	}
	//}
	return result.String()
}