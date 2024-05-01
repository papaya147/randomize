package randomize

import (
	"reflect"
)

func randomSliceFromGeneric[T any]() T {
	var t T
	return randomSliceFromReflectType(reflect.TypeOf(t)).Interface().(T)
}

func randomSliceFromReflectType(t reflect.Type) reflect.Value {
	elementType := t.Elem()
	newSlice := reflect.MakeSlice(t, sliceLength, sliceLength)
	for i := 0; i < newSlice.Len(); i++ {
		value := randomize(elementType)
		newSlice.Index(i).Set(value)
	}
	return newSlice
}
