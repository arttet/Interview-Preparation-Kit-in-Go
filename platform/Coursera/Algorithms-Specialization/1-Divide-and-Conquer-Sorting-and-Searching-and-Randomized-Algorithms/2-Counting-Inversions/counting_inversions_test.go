package main

import (
	"errors"
	"fmt"
	"sort"
	"testing"

	utility "github.com/arttet/Interview-Preparation-Kit-in-Go/internal/test"
	"github.com/stretchr/testify/assert"
)

func TestOK(t *testing.T) {
	N := 68
	for i := 1; i <= N; i++ {
		test := utility.TestCase{
			In:  fmt.Sprintf("input/input%02d.txt", i),
			Out: fmt.Sprintf("output/output%02d.txt", i),
		}
		test.RunTest(t, main)
	}
}

func TestCoursera(t *testing.T) {
	test := utility.TestCase{
		In:       "CourseraIntegerArray.txt",
		Expected: "2407905288",
	}
	test.RunTest(t, main)
}

func TestSort(t *testing.T) {
	ast := assert.New(t)

	input := []int{2, 1, 3, 1, 2}
	ast.Equal(int64(4), countInversions(input))
	ast.Equal(true, sort.IntsAreSorted(input))
}

func TestPanic(t *testing.T) {
	assert.Panics(t, func() { checkError(errors.New("")) }, "The code did not panic")
}