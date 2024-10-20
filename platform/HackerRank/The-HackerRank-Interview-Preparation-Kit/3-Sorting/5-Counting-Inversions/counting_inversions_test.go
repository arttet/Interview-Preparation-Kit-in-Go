package main

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/arttet/Interview-Preparation-Kit-in-Go/internal/utility"
)

func TestOK(t *testing.T) {
	tests := []int{0, 14, 15}
	for _, i := range tests {
		test := utility.TestCase{
			In:  fmt.Sprintf("input/input%02d.txt", i),
			Out: fmt.Sprintf("output/output%02d.txt", i),
		}
		test.RunTest(t, main)
	}
}

func TestSort(t *testing.T) {
	t.Parallel()

	ast := assert.New(t)

	input := []int{2, 1, 3, 1, 2}
	ast.Equal(int64(4), countInversions(input))
	ast.True(sort.IntsAreSorted(input))
}

func TestPanic(t *testing.T) {
	t.Parallel()

	assert.Panics(t, func() { main() }, "The function did not panic")
	assert.Panics(t, func() { checkError(utility.ErrTestPanicMock) }, "The function did not panic")
}
