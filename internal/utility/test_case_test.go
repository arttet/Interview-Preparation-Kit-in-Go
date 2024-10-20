package utility_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/arttet/Interview-Preparation-Kit-in-Go/internal/utility"
)

func TestOK(t *testing.T) {
	test := utility.TestCase{
		In:  "input/input.txt",
		Out: "output/output.txt",
	}

	funcTest := func() {
		stdout, err := os.Create(os.Getenv(utility.OutputPathEnv))
		utility.CheckError(err)
		defer stdout.Close()

		_, err = fmt.Fprintln(stdout, "Test is done.")
		utility.CheckError(err)
	}

	assert.NotPanics(t, func() { test.RunTest(t, funcTest) }, "The function did panic")
}

func TestPanic(t *testing.T) {
	t.Parallel()

	emptyCase := utility.TestCase{}

	assert.Panics(t, func() { emptyCase.RunTest(t, func() {}) }, "The function did not panic")
	assert.Panics(t, func() { utility.CheckError(utility.ErrTestPanicMock) }, "The function did not panic")
}
