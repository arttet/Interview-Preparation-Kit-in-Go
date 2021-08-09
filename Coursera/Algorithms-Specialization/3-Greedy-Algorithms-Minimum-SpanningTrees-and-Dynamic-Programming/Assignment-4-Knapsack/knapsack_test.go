package main

import (
	"errors"
	"fmt"
	"testing"

	"github.com/arttet/problem-solving-with-algorithms-and-data-structures/utility"
	"github.com/stretchr/testify/assert"
)

func TestOK(t *testing.T) {
	N := 44
	for i := 1; i <= N; i++ {
		test := utility.TestCase{
			In:  fmt.Sprintf("input/input%02d.txt", i),
			Out: fmt.Sprintf("output/output%02d.txt", i),
		}
		test.RunTest(t, main)
	}
}

func TestCoursera(t *testing.T) {
	tests := []utility.TestCase{
		{
			In:       "knapsack1.txt",
			Expected: "2493893",
		},
		{
			In:       "knapsack_big.txt",
			Expected: "4243395",
		},
	}

	for i := range tests {
		tests[i].RunTest(t, main)
	}
}

func TestPanic(t *testing.T) {
	assert.Panics(t, func() { checkError(errors.New("")) }, "The code did not panic")
}
