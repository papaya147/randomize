package randomize

import "reflect"

var customRandomizers map[reflect.Type]func() reflect.Value

func init() {
	customRandomizers = make(map[reflect.Type]func() reflect.Value)
}

func RegisterCustomRandomizer[T any](f func() T) {
	var t T
	customRandomizers[reflect.TypeOf(t)] = func() reflect.Value {
		return reflect.ValueOf(f())
	}
}
