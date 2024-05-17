# randomize

A Go package to generate random values from input types

### Getting Started

`go get github.com/papaya147/randomize@v0.1.2`

### Unsupported Types

- Func
- Uintptr
- Unsafe Pointers

Interfaces will be auto assigned as `nil` due to not having knowledge of implementation.

### Examples

```go
package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/google/uuid"
	"github.com/papaya147/randomize"
)

type MySubStruct struct {
	Field11 float32
	Field2  *float64
}

type MyStruct struct {
	Field1 int
	Field2 string
	Field3 MySubStruct
}

type StudentDegree string

const (
	StudentDegreeBachelors StudentDegree = "bachelors"
	StudentDegreeMasters   StudentDegree = "masters"
	StudentDegreeDoctorate StudentDegree = "doctorate"
)

func randomStudentDegree() StudentDegree {
	r := rand.New(rand.NewSource(time.Now().UnixMilli()))
	degrees := []StudentDegree{
		StudentDegreeBachelors,
		StudentDegreeMasters,
		StudentDegreeDoctorate,
	}
	return degrees[r.Intn(len(degrees))]
}

func main() {
	randomize.SetStringLength(10) // defaults to 10
	randomize.SetSliceLength(3)   // defaults to 5
	randomize.SetMapLength(3)     // defaults to 5

	a, _ := randomize.Do[int]()
	fmt.Println(a)
	// 5530747941541261317

	b, _ := randomize.Do[uuid.UUID]()
	fmt.Println(b)
	// 2e0b778b-b872-1827-584c-6b72e340089b

	c, _ := randomize.Do[MyStruct]()
	fmt.Println(c)
	// {6687195844413792927 czuuqhecce {0.78019506 0x1400012a168}}

	d, _ := randomize.Do[*string]()
	fmt.Println(d)
	// 0x14000112100
	fmt.Println(*d)
	// hsazbxhwmm

	e, _ := randomize.Do[[]int]()
	fmt.Println(e)
	// [8919747130222607929 7070119811876138339 2708903612550929244]

	f, _ := randomize.Do[map[string]int8]()
	fmt.Println(f)
	// map[tzjrnumodm:101 unjyttnfhm:-88 vlycruzfue:-118]

	g, _ := randomize.Do[[4]uint8]()
	fmt.Println(g)
	// [55 187 78 179]

	// register custom randomizer for user defined types
	randomize.RegisterCustomRandomizer[StudentDegree](randomStudentDegree)
	h, err := randomize.Do[StudentDegree]()
	if err != nil {
		panic(err)
	}
	fmt.Println(h)
	// bachelors
}
```
