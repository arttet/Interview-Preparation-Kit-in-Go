package main

import (
	"errors"
	"fmt"
	"sort"
	"testing"

	"github.com/arttet/problem-solving-with-algorithms-and-data-structures/utility"
	"github.com/stretchr/testify/assert"
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
	ast := assert.New(t)

	input := []int{2, 1, 3, 1, 2}
	ast.Equal(int64(4), countInversions(input))
	ast.Equal(true, sort.IntsAreSorted(input))
}

func TestPanic(t *testing.T) {
	assert.Panics(t, func() { checkError(errors.New("")) }, "The code did not panic")
}
