package randomize

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

type TestStruct struct {
	Field1 int
	Field2 string
	Field3 []float32
}

type TestStructWrapper struct {
	S1     TestStruct
	Field1 uint
}

func TestRandomStructFromGeneric(t *testing.T) {
	a := randomStructFromGeneric[TestStruct]()
	require.IsType(t, TestStruct{}, a)

	b := randomStructFromGeneric[TestStructWrapper]()
	require.IsType(t, TestStructWrapper{}, b)
}

func TestRandomStructFromReflectType(t *testing.T) {
	x := TestStruct{}
	a := randomStructFromReflectType(reflect.TypeOf(x))
	require.IsType(t, x, a.Interface())
}
