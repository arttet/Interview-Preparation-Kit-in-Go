package main

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/arttet/Interview-Preparation-Kit-in-Go/internal/utility"
)

var ErrTestPanicMock = errors.New("mock panic")

func TestOK(t *testing.T) {
	N := 20
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
		In:       "CourseraQuickSort.txt",
		Expected: "162085\n164123\n138382",
	}
	test.RunTest(t, main)
}

func TestPanic(t *testing.T) {
	t.Parallel()

	assert.Panics(t, func() { checkError(ErrTestPanicMock) }, "The code did not panic")
}
