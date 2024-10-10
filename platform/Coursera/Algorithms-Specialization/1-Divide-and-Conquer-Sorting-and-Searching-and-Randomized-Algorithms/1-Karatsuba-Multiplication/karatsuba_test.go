package main

import (
	"errors"
	"fmt"
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/arttet/Interview-Preparation-Kit-in-Go/internal/utility"
)

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

const x = "3141592653589793238462643383279502884197169399375105820974944592"
const y = "2718281828459045235360287471352662497757247093699959574966967627"
const z = "8539734222673567065463550869546574495034888535765114961879601127067743044893204848617875072216249073013374895871952806582723184"

func TestCoursera(t *testing.T) {
	ast := assert.New(t)
	result := multiplication(x, y)
	ast.Equal(result, z)
}

func BenchmarkBigMult(b *testing.B) {
	for i := 0; i < b.N; i++ {
		x, _ := big.NewInt(0).SetString(x, 10)
		y, _ := big.NewInt(0).SetString(y, 10)
		big.NewInt(0).Mul(x, y)
	}
}

func BenchmarkMult(b *testing.B) {
	for i := 0; i < b.N; i++ {
		multiplication(x, y)
	}
}

func TestPanic(t *testing.T) {
	assert.Panics(t, func() { checkError(errors.New("")) }, "The code did not panic")
}

func TestDoCarryPanic(t *testing.T) {
	x := newInteger("", 1)
	x[0] = -10
	assert.Panics(t, func() { x.doCarry() }, "The code did not panic")
}
