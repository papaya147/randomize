package randomize

import "reflect"

func randomArrayFromGeneric[T any]() (T, error) {
	var t T
	randomArray, err := randomArrayFromReflectType(reflect.TypeOf(t))
	if err != nil {
		return t, err
	}
	return randomArray.Interface().(T), nil
}

func randomArrayFromReflectType(t reflect.Type) (reflect.Value, error) {
	newArray := reflect.New(t).Elem()
	for i := 0; i < newArray.Len(); i++ {
		value, err := randomize(t.Elem())
		if err != nil {
			return reflect.Value{}, nil
		}
		newArray.Index(i).Set(value)
	}
	return newArray, nil
}
