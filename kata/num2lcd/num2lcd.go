package num2lcd

import "strings"

func Num2Lcd(num int) string {
	toks := numberSegments(num)
	lcdNum := strings.Join(toks[:], "\n")
	return lcdNum + "\n"
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
	},{
		"   ",
		"|_|",
		"  |",
	},{
		" _ ",
		"|_ ",
		" _|",
	},{
		" _ ",
		"|_ ",
		"|_|",
	},{
		" _ ",
		"  |",
		"  |",
	},{
		" _ ",
		"|_|",
		"|_|",
	},{
		" _ ",
		"|_|",
		" _|",
	},
}

func numberSegments(n int) digit {
	return digits[n]
}
