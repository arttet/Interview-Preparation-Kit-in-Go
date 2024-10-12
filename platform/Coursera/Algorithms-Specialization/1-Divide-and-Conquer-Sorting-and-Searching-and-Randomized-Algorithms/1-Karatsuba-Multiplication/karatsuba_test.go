package main

import (
	"errors"
	"fmt"
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/arttet/Interview-Preparation-Kit-in-Go/internal/utility"
)

const (
	lhs      = "3141592653589793238462643383279502884197169399375105820974944592"
	rhs      = "2718281828459045235360287471352662497757247093699959574966967627"
	expected = "8539734222673567065463550869546574495034888535765114961879601127067743044893204848617875072216249073013374895871952806582723184" //nolint: lll
)

var ErrTestPanicMock = errors.New("mock panic")

func TestOK(t *testing.T) {
	N := 35
	for i := 1; i <= N; i++ {
		test := utility.TestCase{
			In:  fmt.Sprintf("input/input%02d.txt", i),
			Out: fmt.Sprintf("output/output%02d.txt", i),
		}
		test.RunTest(t, main)
	}
}

func TestCoursera(t *testing.T) {
	t.Parallel()

	ast := assert.New(t)
	result := multiplication(lhs, rhs)
	ast.Equal(expected, result)
}

func BenchmarkBigMult(b *testing.B) {
	for i := 0; i < b.N; i++ {
		x, _ := big.NewInt(0).SetString(lhs, 10)
		y, _ := big.NewInt(0).SetString(rhs, 10)
		big.NewInt(0).Mul(x, y)
	}
}

func BenchmarkMult(b *testing.B) {
	for i := 0; i < b.N; i++ {
		multiplication(lhs, rhs)
	}
}

func TestPanic(t *testing.T) {
	t.Parallel()

	assert.Panics(t, func() { checkError(ErrTestPanicMock) }, "The code did not panic")
}

func TestDoCarryPanic(t *testing.T) {
	t.Parallel()

	x := newInteger("", 1)
	x[0] = -10
	assert.Panics(t, func() { x.doCarry() }, "The code did not panic")
}
