package main

import (
	"errors"
	"fmt"
	"testing"

	utility "github.com/arttet/Interview-Preparation-Kit-in-Go/internal/test"
	"github.com/stretchr/testify/assert"
)

func TestOK(t *testing.T) {
	tests := []int{0, 1, 2, 12, 14}
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
