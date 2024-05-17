package randomize

import (
	"reflect"
)

func randomSliceFromGeneric[T any]() (T, error) {
	var t T
	randomSlice, err := randomSliceFromReflectType(reflect.TypeOf(t))
	if err != nil {
		return t, err
	}
	return randomSlice.Interface().(T), nil
}

func randomSliceFromReflectType(t reflect.Type) (reflect.Value, error) {
	elementType := t.Elem()
	newSlice := reflect.MakeSlice(t, sliceLength, sliceLength)
	for i := 0; i < newSlice.Len(); i++ {
		value, err := randomize(elementType)
		if err != nil {
			return reflect.Value{}, err
		}
		newSlice.Index(i).Set(value)
	}
	return newSlice, nil
}
