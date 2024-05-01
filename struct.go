package randomize

import (
	"reflect"
)

func randomStructFromGeneric[T any]() T {
	var t T
	return randomStructFromReflectType(reflect.TypeOf(t)).Interface().(T)
}

func randomStructFromReflectType(t reflect.Type) reflect.Value {
	newStruct := reflect.New(t).Elem()
	for i := range newStruct.NumField() {
		field := t.Field(i)
		fieldValue := newStruct.Field(i)
		if fieldValue.CanSet() {
			fieldValue.Set(randomize(field.Type))
		}
	}
	return newStruct
}
