package main

import "reflect"

func walk(x interface{}, fn func(input string)) {
	v := getValue(x)

	for i := 0; i < getLen(v); i++ {
		field := getField(v, i)

		switch field.Kind() {
		case reflect.String:
			fn(field.String())
		case reflect.Struct:
			walk(field.Interface(), fn)
		}
	}
}

func getValue(x interface{}) reflect.Value {
	v := reflect.ValueOf(x)

	if v.Kind() == reflect.Ptr {
		return v.Elem()
	}
	return v
}

func getLen(value reflect.Value) int {
	if value.Kind() == reflect.Slice {
		return value.Len()
	}
	return value.NumField()
}

func getField(value reflect.Value, i int) reflect.Value {
	if value.Kind() == reflect.Slice {
		return value.Index(i)
	}
	return value.Field(i)
}