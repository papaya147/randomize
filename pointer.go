package randomize

import (
	"reflect"
)

func randomPointerFromGeneric[T any]() (T, error) {
	var t T
	randomPointer, err := randomPointerFromReflectType(reflect.TypeOf(t))
	if err != nil {
		return t, err
	}
	return randomPointer.Interface().(T), nil
}

func randomPointerFromReflectType(t reflect.Type) (reflect.Value, error) {
	newPointer := reflect.New(t.Elem())
	elem := newPointer.Elem()
	value, err := randomize(elem.Type())
	if err != nil {
		return reflect.Value{}, err
	}
	elem.Set(value)
	return newPointer, nil
}
