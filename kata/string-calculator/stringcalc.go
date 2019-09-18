package string_calculator

import "strconv"

func resolve(input string) int {
	if n, err := str2int(input); err == nil {
		return n
	}
	return 0
}

func str2int(input string) (i int, err error) {
	n, err := strconv.ParseInt(input, 10, 64)
	if err != nil {
		return 0, err
	}
	return int(n), nil
}