package randomize

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

type TestStruct1 struct {
	Field1 int
	Field2 []string
	Field3 struct {
		A float32
		B bool
		C []uint64
	}
	Field4 [5]bool
}

func TestRandomize(t *testing.T) {
	a := Do[string]()
	require.NotZero(t, a)
	require.IsType(t, string(""), a)

	b := Do[bool]()
	require.IsType(t, bool(true), b)

	c := Do[[][4]string]()
	require.NotZero(t, c)
	require.Len(t, c, sliceLength)
	for _, e := range c {
		require.NotZero(t, e)
		require.Len(t, e, 4)
		require.IsType(t, [4]string{}, e)
		for _, f := range e {
			require.NotZero(t, f)
			require.IsType(t, string(""), f)
		}
	}

	d := Do[TestStruct1]()
	require.NotZero(t, d)
	require.IsType(t, TestStruct1{}, d)

	e := Do[[]*string]()
	require.NotZero(t, e)
	require.Len(t, e, sliceLength)
	for _, f := range e {
		require.NotZero(t, f)
		require.NotZero(t, *f)
		require.IsType(t, string(""), *f)
	}

	f := Do[map[string]string]()
	require.NotZero(t, f)
	require.Len(t, f, mapLength)
	for k, v := range f {
		require.NotZero(t, k)
		require.NotZero(t, v)
		require.IsType(t, string(""), k)
		require.IsType(t, string(""), v)
	}

	g := Do[uuid.UUID]()
	require.NotZero(t, g)
	require.IsType(t, uuid.UUID{}, g)

	h := Do[interface{}]()
	require.Nil(t, h)
}
