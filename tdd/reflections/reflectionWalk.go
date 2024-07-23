package reflections

import "reflect"

func walk(x interface{}, fn func(input string)) {
	// val := reflect.ValueOf(x)
	val := getValue(x)

	// numberOfValues := 0
	// var getField func(int) reflect.Value

	walkValue := func(value reflect.Value) {
		walk(value.Interface(), fn)
	}

	switch val.Kind() {
	case reflect.String:
		fn(val.String())
	case reflect.Struct:
		// numberOfValues = val.NumField()
		// getField = val.Field
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
		// case reflect.Slice:
		// 	numberOfValues = val.Len()
		// 	getField = val.Index
	}

	// for i := 0; i < numberOfValues; i++ {
	// 	walk(getField(i).Interface(), fn)
	// }

	// for i := 0; i < val.NumField(); i++ {
	// field := val.Field(i)
	// fn(field.String())

	// if field.Kind() == reflect.String {
	// 	fn(field.String())
	// }

	// switch field.Kind() {
	// case reflect.String:
	// 	fn(field.String())
	// case reflect.Struct:
	// 	walk(field.Interface(), fn)
	// }
	// field := val.Field(0)
	// fn(field.String())

}

func getValue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	return val
}
