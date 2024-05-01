package randomize

import (
	"math/rand"
	"reflect"
	"time"
)

var r *rand.Rand

func init() {
	r = rand.New(rand.NewSource(time.Now().UnixMilli()))
}

var sliceLength int = 5
var stringLength int = 10
var mapLength int = 5

func SetSliceLength(l int) {
	sliceLength = l
}

func SetStringLength(l int) {
	stringLength = l
}

func SetMapLength(l int) {
	mapLength = l
}

func Do[T any]() T {
	var t T
	if reflect.TypeOf(t) == nil {
		return t
	}

	switch reflect.TypeOf(t).Kind() {
	case reflect.Func, reflect.Uintptr, reflect.UnsafePointer:
		panic("types: func, uintptr, unsafe pointer are unsupported")
	case reflect.Bool,
		reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Float32, reflect.Float64,
		reflect.String:
		return randomize(reflect.TypeOf(t)).Interface().(T)
	case reflect.Slice:
		return randomSliceFromGeneric[T]()
	case reflect.Array:
		return randomArrayFromGeneric[T]()
	case reflect.Struct:
		return randomStructFromGeneric[T]()
	case reflect.Pointer:
		return randomPointerFromGeneric[T]()
	case reflect.Map:
		return randomMapFromGeneric[T]()
	default:
		return t
	}
}

func randomize(t reflect.Type) reflect.Value {
	switch t.Kind() {
	case reflect.Invalid:
		return reflect.ValueOf(t)
	case reflect.Bool:
		return reflect.ValueOf(randomBool())
	case reflect.Int:
		return reflect.ValueOf(randomInt())
	case reflect.Int8:
		return reflect.ValueOf(randomInt8())
	case reflect.Int16:
		return reflect.ValueOf(randomInt16())
	case reflect.Int32:
		return reflect.ValueOf(randomInt32())
	case reflect.Int64:
		return reflect.ValueOf(randomInt64())
	case reflect.Uint:
		return reflect.ValueOf(randomUint())
	case reflect.Uint8:
		return reflect.ValueOf(randomUint8())
	case reflect.Uint16:
		return reflect.ValueOf(randomUint16())
	case reflect.Uint32:
		return reflect.ValueOf(randomUint32())
	case reflect.Uint64:
		return reflect.ValueOf(randomUint64())
	case reflect.Float32:
		return reflect.ValueOf(randomFloat32())
	case reflect.Float64:
		return reflect.ValueOf(randomFloat64())
	case reflect.Array:
		return randomArrayFromReflectType(t)
	case reflect.Slice:
		return randomSliceFromReflectType(t)
	case reflect.String:
		return reflect.ValueOf(randomString(stringLength))
	case reflect.Struct:
		return randomStructFromReflectType(t)
	case reflect.Pointer:
		return randomPointerFromReflectType(t)
	case reflect.Map:
		return randomMapFromReflectType(t)
	}
	return reflect.ValueOf(t)
}