package string_calculator

import (
	"strconv"
	"strings"
)

type InputError string

func (e InputError) Error() string {
	return string(e)
}

const (
	ErrNegNums = InputError("negative numbers are not permitted")
)

func Resolve(input string) (int, error) {
	if n, err := str2int(input); err == nil {
		return validate(n)
	}
	s, d := inputAndDelim(input)
	nums := str2slice(s, d)
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

func sum(ints ...int) (int, error) {
	sum := 0
	for i := range ints {
		n, e := validate(ints[i])
		if e != nil {
			return n, e
		}
		sum += n
	}
	return sum, nil
}

func validate(n int) (int, error) {
	if n < 0 {
		return n, ErrNegNums
	}
	if n > 1000 {
		return 0, nil
	}
	return n, nil
}

func inputAndDelim(input string) (string, string) {
	delim := detectDelim(input)

	if strings.HasPrefix(input, "//") {
		input = input[strings.Index(input,"\n")+1:]
	}

	return input, delim
}

func detectDelim(input string) string {
	if strings.HasPrefix(input, "//") {
		nlPos := strings.Index(input,"\n")
		d := input[2:nlPos]
		if strings.HasPrefix(d, "[") {
			d = strings.Trim(d, "[]")
		}
		return d
	}
	if strings.Contains(input, "\n") {
		return "\n"
	}
	return ","
}
