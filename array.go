package randomize

import "reflect"

func randomArrayFromGeneric[T any]() T {
	var t T
	return randomArrayFromReflectType(reflect.TypeOf(t)).Interface().(T)
}

func randomArrayFromReflectType(t reflect.Type) reflect.Value {
	newArray := reflect.New(t).Elem()
	for i := 0; i < newArray.Len(); i++ {
		value := randomize(t.Elem())
		newArray.Index(i).Set(value)
	}
	return newArray
}
