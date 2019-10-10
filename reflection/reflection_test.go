package main

import (
	"reflect"
	"testing"
)

func TestWalk(t *testing.T) {

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
		Name string
		Profile Profile
	}

	cases := []requirement{
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
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.Expected) {
				t.Errorf("\ngot  %v \nwant %v", got, test.Expected)
			}
		})
	}
}
