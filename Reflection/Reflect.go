package Reflection

import "reflect"

func Walk(x interface{}, fn func(input string)) {
	val := getValue(x)

	walkValue := func(value reflect.Value) {
		Walk(value.Interface(), fn)
	}

	switch val.Kind() {
	case reflect.String:
		fn(val.String())
	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			walkValue(val.Field(i))
		}
	case reflect.Slice, reflect.Array:
		for i := 0; i < val.Len(); i++ {
			walkValue(val.Index(i))
		}
	case reflect.Map:
		for _, key := range val.MapKeys() {
			walkValue(val.MapIndex(key))
		}
	case reflect.Chan:
		for v, ok := val.Recv(); ok; v, ok = val.Recv() {
			walkValue(v)
		}

	case reflect.Func:
		valFnResult := val.Call(nil)
		for _, res := range valFnResult {
			walkValue(res)
		}

	}

}

func getValue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)
	// if they pass us a pointer we need to get the underlying element of that as a pointer will just be a memory address
	// we do this by changing val to get the underlying element of it
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	return val
}
