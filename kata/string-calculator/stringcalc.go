package string_calculator

import (
	"strconv"
	"strings"
)

func Resolve(input string) int {
	if n, err := str2int(input); err == nil {
		return n
	}
	d := detectDelim(input)
	nums := str2slice(input, d)
	return sum(nums...)
}

func str2int(input string) (i int, err error) {
	if input == "" {
		return 0, nil
	}
	n, err := strconv.ParseInt(input, 10, 64)
	if err != nil {
		return 0, err
	}
	return int(n), nil
}

func str2slice(input string, delim string) []int {
	var result []int
	toks := strings.Split(input, delim)
	for _, tok := range toks {
		i, _ := str2int(tok)
		result = append(result, i)
	}
	return result
}

func sum(ints ...int) int {
	sum := 0
	for i := range ints {
		sum += ints[i]
	}
	return sum
}

func detectDelim(input string) string {
	if strings.Contains(input, "\n") {
		return "\n"
	}
	return ","
}
