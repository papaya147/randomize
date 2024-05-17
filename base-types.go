package randomize

import (
	"reflect"
	"strings"
)

func randomUint64() uint64 {
	return r.Uint64()
}

func randomUint32() uint32 {
	return r.Uint32()
}

func randomUint16() uint16 {
	return uint16(r.Intn(1 << 16))
}

func randomUint8() uint8 {
	return uint8(r.Intn(1 << 8))
}

func randomUint() uint {
	return uint(r.Uint32())
}

func randomInt64() int64 {
	return r.Int63() - 1<<62
}

func randomInt32() int32 {
	return r.Int31() - 1<<30
}

func randomInt16() int16 {
	return int16(r.Intn(1<<16)) - 1<<14
}

func randomInt8() int8 {
	return int8(r.Intn(1<<8)) - 1<<6
}

func randomInt() int {
	return int(r.Int())
}

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func randomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[r.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func randomFloat64() float64 {
	return r.Float64()
}

func randomFloat32() float32 {
	return r.Float32()
}

func randomBool() bool {
	return r.Intn(2) == 0
}

func isBaseType(kind reflect.Kind) bool {
	switch kind {
	case reflect.Bool,
		reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Float32, reflect.Float64,
		reflect.String:
		return true
	default:
		return false
	}
}
