package main

import "fmt"

const defaultName = "my friend"

type Name string

func (n Name) String() string {
	if n == "" {
		return defaultName
	}
	return string(n)
}

func Greet(name string) string {
	n := Name(name)
	return fmt.Sprintf("Hello, %s.", n.String())
}