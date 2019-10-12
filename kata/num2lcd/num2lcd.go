package num2lcd

import "strings"

func Num2Lcd(num int) string {
	toks := numberSegments(num)
	return strings.Join(toks[:], "\n") + "\n"
}

func numberSegments(n int) [3]string {
	switch n {
	case 0: return [3]string{
		" _ ",
		"| |",
		"|_|",
	}
	case 1: return [3]string{
		"   ",
		"  |",
		"  |",
	}
	case 2: return [3]string{
		" _ ",
		" _|",
		"|_ ",
	}
	case 3: return [3]string{
		" _ ",
		" _|",
		" _|",
	}
	case 4: return [3]string{
		"   ",
		"|_|",
		"  |",
	}
	case 5: return [3]string{
		" _ ",
		"|_ ",
		" _|",
	}
	case 6: return [3]string{
		" _ ",
		"|_ ",
		"|_|",
	}
	case 7: return [3]string{
		" _ ",
		"  |",
		"  |",
	}
	case 8: return [3]string{
		" _ ",
		"|_|",
		"|_|",
	}
	case 9: return [3]string{
		" _ ",
		"|_|",
		" _|",
	}
	default: return [3]string{}
	}
}

