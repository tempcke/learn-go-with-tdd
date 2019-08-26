package main

import (
	"fmt"
	"strings"
)

const defaultName = "my friend"

type Name struct {
	name string
}

func (n Name) String() string {
	if n.name == "" {
		return defaultName
	}
	return n.name
}

func (n Name) isUpper() bool {
	return n.String() == strings.ToUpper(n.name)
}

func NameList(names []string) string {
	if len(names) == 2 {
		return strings.Join(names, " and ")
	}
	var result string
	for i, name := range names {
		switch i {
		case 0: result = name
		case len(names) - 1: result = fmt.Sprintf("%s, and %s", result, name)
		default: result = fmt.Sprintf("%s, %s", result, name)
		}
	}
	return result
}

func Greet(names... string) string {
	if len(names) == 1 {
		name := names[0]
		n := Name{name}
		if n.isUpper() {
			return fmt.Sprintf("HELLO %s!", n.String())
		}
		return fmt.Sprintf("Hello, %s.", n.String())
	}
	return fmt.Sprintf("Hello, %s.", NameList(names))
}