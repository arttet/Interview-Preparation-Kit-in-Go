package main

import (
	"errors"
	"fmt"
	"testing"

	"github.com/arttet/problem-solving-with-algorithms-and-data-structures/utility"
	"github.com/stretchr/testify/assert"
)

func TestOK(t *testing.T) {
	tests := []int{1, 2, 3, 4}
	for _, i := range tests {
		test := utility.TestCase{
			In:  fmt.Sprintf("input/input%02d.txt", i),
			Out: fmt.Sprintf("output/output%02d.txt", i),
		}
		test.RunTest(t, main)
	}
}

func TestPanic(t *testing.T) {
	assert.Panics(t, func() { checkError(errors.New("")) }, "The code did not panic")
}
