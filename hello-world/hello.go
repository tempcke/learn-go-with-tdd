package main

import "fmt"

const englishPrefix = "Hello"
const spanishPrefix = "Hola"
const frenchPrefix = "Bonjour"
const defaultSubject = "World"

func greetingSubject(name string) string {
	if name == "" {
		return defaultSubject
	}
	return name
}

func greetingPrefix(language string) string {
	switch language {
	case "spanish":
		return spanishPrefix
	case "french":
		return frenchPrefix
	case "english":
		return englishPrefix
	default:
		return englishPrefix
	}
}

func Hello(name string, language string) string {
	return fmt.Sprintf("%s, %s", greetingPrefix(language), greetingSubject(name));
}

func main() {
	fmt.Println(Hello("", "english"))
}

