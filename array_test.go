package randomize

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRandomArrayFromGeneric(t *testing.T) {
	a, err := randomArrayFromGeneric[[5]string]()
	require.NoError(t, err)
	require.Len(t, a, 5)
	for _, e := range a {
		require.IsType(t, string(""), e)
	}

	b, err := randomArrayFromGeneric[[10]int]()
	require.NoError(t, err)
	require.Len(t, b, 10)
	for _, e := range b {
		require.IsType(t, int(0), e)
	}

	c, err := randomArrayFromGeneric[[15]float64]()
	require.NoError(t, err)
	require.Len(t, c, 15)
	for _, e := range c {
		require.IsType(t, float64(0), e)
	}
}

func TestRandomArrayFromReflectType(t *testing.T) {
	x := [2]string{"hi"}
	a, err := randomArrayFromReflectType(reflect.TypeOf(x))
	require.NoError(t, err)
	require.Len(t, a.Interface(), len(x))
	for _, e := range a.Interface().([2]string) {
		require.IsType(t, string(""), e)
	}
}
