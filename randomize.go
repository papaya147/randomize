package randomize

import (
	"errors"
	"fmt"
	"github.com/shopspring/decimal"
	"math/rand"
	"reflect"
	"time"
)

var r *rand.Rand

func init() {
	r = rand.New(rand.NewSource(time.Now().UnixMilli()))
	RegisterCustomRandomizer(func() time.Month {
		args := []time.Month{
			time.January,
			time.February,
			time.March,
			time.April,
			time.May,
			time.June,
			time.July,
			time.August,
			time.September,
			time.October,
			time.November,
			time.December,
		}
		return args[r.Intn(len(args))]
	})
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

// Must calls Do and panics if an error is returned
func Must[T any]() T {
	out, err := Do[T]()
	if err != nil {
		panic(err)
	}
	return out
}

// Do returns a random variable of the type T
func Do[T any]() (T, error) {
	var t T
	typ := reflect.TypeOf(t)
	if typ == nil {
		return t, nil
	}

	if randomizer, ok := customRandomizers[typ]; ok {
		return randomizer().Interface().(T), nil
	}

	if typ.String() != typ.Kind().String() && isBaseType(typ.Kind()) {
		return t, fmt.Errorf("%s does not have a custom mapping, but may have enumerated types", typ.String())
	}

	switch typ.Kind() {
	case reflect.Func, reflect.Uintptr, reflect.UnsafePointer:
		return t, errors.New("types: func, uintptr, unsafe pointer are unsupported")
	case reflect.Bool,
		reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Float32, reflect.Float64,
		reflect.String:
		value, err := randomize(reflect.TypeOf(t))
		if err != nil {
			return t, err
		}
		return value.Interface().(T), nil
	case reflect.Slice:
		return randomSliceFromGeneric[T]()
	case reflect.Array:
		return randomArrayFromGeneric[T]()
	case reflect.Struct:
		if typ == reflect.TypeOf(time.Time{}) {
			return randomTime().Interface().(T), nil
		}
		if typ == reflect.TypeOf(decimal.Decimal{}) {
			return randomDecimal().Interface().(T), nil
		}
		return randomStructFromGeneric[T]()
	case reflect.Pointer:
		return randomPointerFromGeneric[T]()
	case reflect.Map:
		return randomMapFromGeneric[T]()
	default:
		return t, nil
	}
}

func randomize(t reflect.Type) (reflect.Value, error) {
	if randomizer, ok := customRandomizers[t]; ok {
		return randomizer(), nil
	}

	if t.String() != t.Kind().String() && isBaseType(t.Kind()) {
		return reflect.Value{}, fmt.Errorf("%s does not have a custom mapping, but may have enumerated types", t.String())
	}

	switch t.Kind() {
	case reflect.Invalid:
		return reflect.ValueOf(t), nil
	case reflect.Bool:
		return reflect.ValueOf(randomBool()), nil
	case reflect.Int:
		return reflect.ValueOf(randomInt()), nil
	case reflect.Int8:
		return reflect.ValueOf(randomInt8()), nil
	case reflect.Int16:
		return reflect.ValueOf(randomInt16()), nil
	case reflect.Int32:
		return reflect.ValueOf(randomInt32()), nil
	case reflect.Int64:
		return reflect.ValueOf(randomInt64()), nil
	case reflect.Uint:
		return reflect.ValueOf(randomUint()), nil
	case reflect.Uint8:
		return reflect.ValueOf(randomUint8()), nil
	case reflect.Uint16:
		return reflect.ValueOf(randomUint16()), nil
	case reflect.Uint32:
		return reflect.ValueOf(randomUint32()), nil
	case reflect.Uint64:
		return reflect.ValueOf(randomUint64()), nil
	case reflect.Float32:
		return reflect.ValueOf(randomFloat32()), nil
	case reflect.Float64:
		return reflect.ValueOf(randomFloat64()), nil
	case reflect.Array:
		return randomArrayFromReflectType(t)
	case reflect.Slice:
		return randomSliceFromReflectType(t)
	case reflect.String:
		return reflect.ValueOf(randomString(stringLength)), nil
	case reflect.Struct:
		if t == reflect.TypeOf(time.Time{}) {
			return randomTime(), nil
		}
		if t == reflect.TypeOf(decimal.Decimal{}) {
			return randomDecimal(), nil
		}
		return randomStructFromReflectType(t)
	case reflect.Pointer:
		return randomPointerFromReflectType(t)
	case reflect.Map:
		return randomMapFromReflectType(t)
	default:
		return reflect.ValueOf(t), nil
	}
}

func randomTime() reflect.Value {
	t := time.Date(r.Intn(9000)+1000, time.Month(r.Intn(12)), r.Intn(28)+1, r.Intn(24), r.Intn(60), r.Intn(60), 0, time.UTC)
	return reflect.ValueOf(t)
}

func randomDecimal() reflect.Value {
	d := decimal.NewFromFloat(Must[float64]())
	return reflect.ValueOf(d)
}
