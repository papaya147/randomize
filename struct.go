package randomize

import (
	"reflect"
)

func randomStructFromGeneric[T any]() (T, error) {
	var t T
	randomStruct, err := randomize(reflect.TypeOf(t))
	if err != nil {
		return t, err
	}
	return randomStruct.Interface().(T), nil
}

func randomStructFromReflectType(t reflect.Type) (reflect.Value, error) {
	newStruct := reflect.New(t).Elem()
	for i := range newStruct.NumField() {
		field := t.Field(i)
		fieldValue := newStruct.Field(i)
		if fieldValue.CanSet() {
			value, err := randomize(field.Type)
			if err != nil {
				return reflect.Value{}, err
			}
			fieldValue.Set(value)
		}
	}
	return newStruct, nil
}
