package randomize

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRandomSliceFromGeneric(t *testing.T) {
	a := randomSliceFromGeneric[[]string]()
	require.Len(t, a, sliceLength)
	for _, e := range a {
		require.IsType(t, string(""), e)
	}

	SetSliceLength(10)
	b := randomSliceFromGeneric[[]int]()
	require.Len(t, b, sliceLength)
	for _, e := range b {
		require.IsType(t, int(0), e)
	}
}
