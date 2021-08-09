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
	test := utility.TestCase{
		In:       "jobs.txt",
		Expected: "69119377652\n67311454237",
	}
	test.RunTest(t, main)
}

func TestPanic(t *testing.T) {
	assert.Panics(t, func() { checkError(errors.New("")) }, "The code did not panic")
}
