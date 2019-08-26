package main

import "fmt"

const DEFAULT_NAME = "my friend"

func Greet(name string) string {
	return fmt.Sprintf("Hello, %s.", Name(name))
}

func Name(name string) string {
	if name == "" {
		return DEFAULT_NAME
	}
	return name
}