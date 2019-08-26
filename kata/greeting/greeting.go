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

func Greet(names... string) string {
	if len(names) == 1 {
		name := names[0]
		n := Name{name}
		if n.isUpper() {
			return fmt.Sprintf("HELLO %s!", n.String())
		}
		return fmt.Sprintf("Hello, %s.", n.String())
	}
	return fmt.Sprintf("Hello, %s.", strings.Join(names, " and "))
}