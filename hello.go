package main

import "fmt"

const englishPrefix = "Hello, "
const defaultSubject = "World"

func subject(name string) string {
	if name == "" {
		return defaultSubject
	}
	return name
}

func Hello(name string) string {
	return englishPrefix + subject(name)
}

func main() {
	fmt.Println(Hello(""))
}

