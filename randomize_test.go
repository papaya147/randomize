package randomize

import (
	"testing"
	"time"

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

type StudentDegree string

const (
	StudentDegreeBachelors StudentDegree = "bachelors"
	StudentDegreeMasters   StudentDegree = "masters"
	StudentDegreeDoctorate StudentDegree = "doctorate"
)

func randomStudentDegree() StudentDegree {
	degrees := []StudentDegree{
		StudentDegreeBachelors,
		StudentDegreeMasters,
		StudentDegreeDoctorate,
	}
	return degrees[r.Intn(len(degrees))]
}

func TestDo(t *testing.T) {
	a, err := Do[string]()
	require.NoError(t, err)
	require.NotZero(t, a)
	require.IsType(t, string(""), a)

	b, err := Do[bool]()
	require.NoError(t, err)
	require.IsType(t, bool(true), b)

	c, err := Do[[][4]string]()
	require.NoError(t, err)
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

	d, err := Do[TestStruct1]()
	require.NoError(t, err)
	require.NotZero(t, d)
	require.IsType(t, TestStruct1{}, d)

	e, err := Do[[]*string]()
	require.NoError(t, err)
	require.NotZero(t, e)
	require.Len(t, e, sliceLength)
	for _, f := range e {
		require.NotZero(t, f)
		require.NotZero(t, *f)
		require.IsType(t, string(""), *f)
	}

	f, err := Do[map[string]string]()
	require.NoError(t, err)
	require.NotZero(t, f)
	require.Len(t, f, mapLength)
	for k, v := range f {
		require.NotZero(t, k)
		require.NotZero(t, v)
		require.IsType(t, string(""), k)
		require.IsType(t, string(""), v)
	}

	g, err := Do[uuid.UUID]()
	require.NoError(t, err)
	require.NotZero(t, g)
	require.IsType(t, uuid.UUID{}, g)

	h, err := Do[interface{}]()
	require.NoError(t, err)
	require.Nil(t, h)

	RegisterCustomRandomizer[StudentDegree](randomStudentDegree)
	i, err := Do[StudentDegree]()
	require.NoError(t, err)
	require.NotZero(t, i)
	require.IsType(t, StudentDegree(""), i)
	require.Contains(t, []StudentDegree{StudentDegreeBachelors, StudentDegreeMasters, StudentDegreeDoctorate}, i)

	j, err := Do[time.Time]()
	require.NoError(t, err)
	require.NotZero(t, j)
}
