package num2lcd

import "strings"

func Num2Lcd(num int) string {
	if num <= 9 {
		toks := numberSegments(num)
		lcdNum := strings.Join(toks[:], "\n")
		return lcdNum + "\n"
	}

	digits := IntToSlice(num)
	var result [3]string
	for i := len(digits)-1; i >= 0; i-- {
		lcdNum := numberSegments(digits[i])
		for j:=0; j<3; j++ {
			result[j] += lcdNum[j]
		}
	}
	return strings.Join(result[:], "\n") + "\n"
}



func IntToSlice(n int) []int {
	return intToSliceIterator(n, []int{})
}

func intToSliceIterator(i int, s []int) []int {
	if i == 0 {
		return s
	}
	return intToSliceIterator(i/10, append(s, i% 10))
}

type digit [3]string

var digits = [...]digit{
	{
		" _ ",
		"| |",
		"|_|",
	}, {
		"   ",
		"  |",
		"  |",
	}, {
		" _ ",
		" _|",
		"|_ ",
	}, {
		" _ ",
		" _|",
		" _|",
	}, {
		"   ",
		"|_|",
		"  |",
	}, {
		" _ ",
		"|_ ",
		" _|",
	}, {
		" _ ",
		"|_ ",
		"|_|",
	}, {
		" _ ",
		"  |",
		"  |",
	}, {
		" _ ",
		"|_|",
		"|_|",
	}, {
		" _ ",
		"|_|",
		" _|",
	},
}

func numberSegments(n int) digit {
	return digits[n]
}
