package main

import "reflect"


func walk(x interface{}, fn func(input string)) {
	val := getValue(x)

	switch val.Kind() {
	case reflect.String:
		fn(val.String())
	case reflect.Struct:
		walkEach(val.Field, val.NumField(), fn)
	case reflect.Slice:
		walkEach(val.Index, val.Len(), fn)
	}
}

func getValue(x interface{}) reflect.Value {
	v := reflect.ValueOf(x)

	if v.Kind() == reflect.Ptr {
		return v.Elem()
	}
	return v
}

func walkEach(field func(int) reflect.Value, len int, fn func(input string)) {
	for i:=0; i<len; i++ {
		walk(field(i).Interface(), fn)
	}
}