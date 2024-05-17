package randomize

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRandomMapFromGeneric(t *testing.T) {
	a, err := randomMapFromGeneric[map[int]string]()
	require.NoError(t, err)
	require.IsType(t, map[int]string{}, a)
	require.Len(t, a, mapLength)
}
