package randomize

import "reflect"

func randomMapFromGeneric[T any]() T {
	var t T
	return randomMapFromReflectType(reflect.TypeOf(t)).Interface().(T)
}

func randomMapFromReflectType(t reflect.Type) reflect.Value {
	keyType := t.Key()
	valueType := t.Elem()
	newMap := reflect.MakeMap(t)
	for range mapLength {
		k := randomize(keyType)
		v := randomize(valueType)
		newMap.SetMapIndex(k, v)
	}
	return newMap
}
