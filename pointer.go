package randomize

import (
	"reflect"
)

func randomPointerFromGeneric[T any]() T {
	var t T
	return randomPointerFromReflectType(reflect.TypeOf(t)).Interface().(T)
}

func randomPointerFromReflectType(t reflect.Type) reflect.Value {
	newPointer := reflect.New(t.Elem())
	elem := newPointer.Elem()
	elem.Set(randomize(elem.Type()))
	return newPointer
}
