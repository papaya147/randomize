package randomize

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRandomPointerFromGeneric(t *testing.T) {
	a, err := randomPointerFromGeneric[*int]()
	require.NoError(t, err)
	require.IsType(t, int(0), *a)

	b, err := randomPointerFromGeneric[*TestStruct1]()
	require.NoError(t, err)
	require.IsType(t, TestStruct1{}, *b)

	c, err := randomPointerFromGeneric[*[]string]()
	require.NoError(t, err)
	require.IsType(t, []string{}, *c)
}

func TestRandomPointerFromReflectType(t *testing.T) {
	var x *int
	a, err := randomPointerFromReflectType(reflect.TypeOf(x))
	require.NoError(t, err)
	require.IsType(t, x, a.Interface())
}
