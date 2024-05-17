package randomize

import "reflect"

func randomMapFromGeneric[T any]() (T, error) {
	var t T
	randomMap, err := randomMapFromReflectType(reflect.TypeOf(t))
	if err != nil {
		return t, err
	}
	return randomMap.Interface().(T), nil
}

func randomMapFromReflectType(t reflect.Type) (reflect.Value, error) {
	keyType := t.Key()
	valueType := t.Elem()
	newMap := reflect.MakeMap(t)
	for range mapLength {
		k, err := randomize(keyType)
		if err != nil {
			return newMap, err
		}
		v, err := randomize(valueType)
		if err != nil {
			return newMap, err
		}
		newMap.SetMapIndex(k, v)
	}
	return newMap, nil
}
