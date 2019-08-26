package main

import (
	"fmt"
	"regexp"
	"strings"
)

const defaultName = "my friend"

type Names struct {
	normal []string
	shout  []string
}

func (names Names) list(list []string) string {
	if len(list) == 2 {
		return strings.Join(list, " and ")
	}
	var result string
	for i, name := range list {
		switch i {
		case 0: result = name
		case len(list) - 1: result = fmt.Sprintf("%s, and %s", result, name)
		default: result = fmt.Sprintf("%s, %s", result, name)
		}
	}
	return result
}

func (names Names) hasNormal() bool {
	return len(names.normal) >= 1
}

func (names Names) hasShout() bool {
	return len(names.shout) >= 1
}

func (names Names) normalGreet() string {
	if !names.hasNormal() {
		return ""
	}
	return fmt.Sprintf("Hello, %s.", names.list(names.normal))
}

func (names Names) shoutGreet() string {
	if !names.hasShout() {
		return ""
	}
	return fmt.Sprintf("HELLO %s!", names.list(names.shout))
}

func (names Names) Greet() string {
	normalGreet := names.normalGreet()
	shoutGreet  := names.shoutGreet()

	if names.hasShout() {
		if names.hasNormal() {
			return fmt.Sprintf("%s AND %s", normalGreet, shoutGreet)
		}
		return shoutGreet
	}
	return normalGreet
}

func Greet(nameList ... string) string {
	names := buildNames(nameList)
	return names.Greet()
}


func buildNames(names []string) Names {
	names = parseNameList(names)
	if len(names) == 1 && names[0] == "" {
		names[0] = defaultName
	}
	var normal []string
	var shout  []string
	for _, name := range names {
		if name == strings.ToUpper(name) {
			shout = append(shout, name)
			continue
		}
		normal = append(normal, name)
	}
	return Names{ normal, shout	}
}

func parseNameList(input []string) []string {
	var names []string
	for _, val := range input {
		if !isQuoted(val) && strings.Contains(val, ",") {
			for _, name := range strings.Split(val, ",") {
				names = append(names, strings.Trim(name, " "))
			}
			continue
		}
		names = append(names, strings.Trim(val,"\""))
	}
	return names
}

func isQuoted(name string) bool {
	var quoted = regexp.MustCompile(`^\".+\"$`)
	return quoted.MatchString(name)
}