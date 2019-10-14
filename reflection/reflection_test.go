package main

import (
	"reflect"
	"testing"
)

type requirement struct {
	Name     string
	Input    interface{}
	Expected []string
}

type Profile struct {
	Age  int
	City string
}

type Person struct {
	Name    string
	Profile Profile
}

var requirements = []requirement{
	{
		"Struct with one string field",
		struct {
			Name string
		}{"Chris"},
		[]string{"Chris"},
	}, {
		"Struct with two string fields",
		struct {
			Name string
			City string
		}{"Chris", "London"},
		[]string{"Chris", "London"},
	}, {
		"Struct with non string field",
		struct {
			Name string
			Age  int
		}{"Chris", 33},
		[]string{"Chris"},
	}, {
		"Nested fields",
		Person{
			"Chris",
			Profile{33, "London"},
		},
		[]string{"Chris", "London"},
	}, {
		"Pointers to things",
		&Person{
			"Chris",
			Profile{33, "London"},
		},
		[]string{"Chris", "London"},
	}, {
		"Slices",
		[]Profile{
			{33, "London"},
			{34, "Chicago"},
		},
		[]string{"London", "Chicago"},
	}, {
		"Arrays",
		[2]Profile{
			{33, "London"},
			{34, "Chicago"},
		},
		[]string{"London", "Chicago"},
	},
}

var mapReq = requirement{
	"Maps",
	map[string]string{
		"Foo": "Bar",
		"Baz": "Boz",
	},
	[]string{"Bar", "Boz"},
}

func TestWalk(t *testing.T) {
	for _, test := range requirements {
		t.Run(test.Name, func(t *testing.T) {
			got := walkResults(test)

			if !reflect.DeepEqual(got, test.Expected) {
				t.Errorf("\ngot  %v \nwant %v", got, test.Expected)
			}
		})
	}

	// https://github.com/quii/learn-go-with-tests/blob/master/reflection.md#one-final-problem
	// Remember that maps in Go do not guarantee order. So your tests will sometimes
	// fail because we assert that the calls to fn are done in a particular order.
	//
	// To fix this, we'll need to move our assertion with the maps to a new test
	// where we do not care about the order.

	t.Run(mapReq.Name, func(t *testing.T) {
		got := walkResults(mapReq)
		for _, expected := range mapReq.Expected {
			assertContains(t, got, expected)
		}
	})
}

func walkResults(test requirement) []string {
	var got []string
	walk(test.Input, func(input string) {
		got = append(got, input)
	})
	return got
}

func assertContains(t *testing.T, haystack []string, needle string)  {
	contains := false
	for _, x := range haystack {
		if x == needle {
			contains = true
		}
	}
	if !contains {
		t.Errorf("expected %+v to contain %q but it didn't", haystack, needle)
	}
}