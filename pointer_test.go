package randomize

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRandomPointerFromGeneric(t *testing.T) {
	a := randomPointerFromGeneric[*int]()
	require.IsType(t, int(0), *a)

	b := randomPointerFromGeneric[*TestStruct1]()
	require.IsType(t, TestStruct1{}, *b)

	c := randomPointerFromGeneric[*[]string]()
	require.IsType(t, []string{}, *c)
}

func TestRandomPointerFromReflectType(t *testing.T) {
	var x *int
	a := randomPointerFromReflectType(reflect.TypeOf(x))
	require.IsType(t, x, a.Interface())
}
