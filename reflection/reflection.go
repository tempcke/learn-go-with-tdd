package main

import "reflect"

func walk(x interface{}, fn func(input string)) {
	w := fnWalk{fn}
	w.walk(x)
}

type fnWalk struct {
	fn func(input string)
}

func (w *fnWalk) walk(x interface{}) {
	val := w.getValue(x)

	switch val.Kind() {

	case reflect.String:
		w.fn(val.String())

	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			w.walkValue(val.Field(i))
		}

	case reflect.Slice, reflect.Array:
		for i := 0; i < val.Len(); i++ {
			w.walkValue(val.Index(i))
		}

	case reflect.Map:
		for _, key := range val.MapKeys() {
			w.walkValue(val.MapIndex(key))
		}
	}
}

func (w *fnWalk) getValue(x interface{}) reflect.Value {
	v := reflect.ValueOf(x)

	if v.Kind() == reflect.Ptr {
		return v.Elem()
	}
	return v
}

func (w *fnWalk) walkValue(value reflect.Value) {
	w.walk(value.Interface())
}
