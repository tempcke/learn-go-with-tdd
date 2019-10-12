package num2lcd

import "strings"

func Num2Lcd(num int) string {
	toks := numberSegments(num)
	return strings.Join(toks[:], "\n") + "\n"
}

type digit [3]string

func numberSegments(n int) digit {
	switch n {
	case 0: return digit{
		" _ ",
		"| |",
		"|_|",
	}
	case 1: return digit{
		"   ",
		"  |",
		"  |",
	}
	case 2: return digit{
		" _ ",
		" _|",
		"|_ ",
	}
	case 3: return digit{
		" _ ",
		" _|",
		" _|",
	}
	case 4: return digit{
		"   ",
		"|_|",
		"  |",
	}
	case 5: return digit{
		" _ ",
		"|_ ",
		" _|",
	}
	case 6: return digit{
		" _ ",
		"|_ ",
		"|_|",
	}
	case 7: return digit{
		" _ ",
		"  |",
		"  |",
	}
	case 8: return digit{
		" _ ",
		"|_|",
		"|_|",
	}
	case 9: return digit{
		" _ ",
		"|_|",
		" _|",
	}
	default: return digit{}
	}
}

