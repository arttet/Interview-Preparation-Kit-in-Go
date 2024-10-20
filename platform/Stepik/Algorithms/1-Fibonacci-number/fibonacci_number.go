package main

import (
	"bufio"
	"fmt"
	"os"
)

// FibonacciNumber returns the nth value in the fibonacci series.
func FibonacciNumber(n, m uint) uint {
	F := [2][2]uint{{1, 1}, {1, 0}}
	power(&F, n-1, m)

	return F[0][0]
}

func power(fib *[2][2]uint, n, m uint) {
	if n == 0 || n == 1 {
		return
	}

	power(fib, n/2, m)
	multiply(fib, fib, m)

	if n%2 != 0 {
		M := [2][2]uint{{1, 1}, {1, 0}}
		multiply(fib, &M, m)
	}
}

func multiply(lhs, rhs *[2][2]uint, mod uint) {
	a := lhs[0][0]*rhs[0][0] + lhs[0][1]*rhs[1][0]
	b := lhs[0][0]*rhs[0][1] + lhs[0][1]*rhs[1][1]
	c := lhs[1][0]*rhs[0][0] + lhs[1][1]*rhs[1][0]
	d := lhs[1][0]*rhs[0][1] + lhs[1][1]*rhs[1][1]

	lhs[0][0] = a % mod
	lhs[0][1] = b % mod
	lhs[1][0] = c % mod
	lhs[1][1] = d % mod
}

func main() {
	stdin, err := os.Open(os.Getenv("INPUT_PATH"))
	if err != nil {
		stdin = os.Stdin
	} else {
		defer stdin.Close()
	}

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	if err != nil {
		stdout = os.Stdout
	} else {
		defer stdout.Close()
	}

	reader := bufio.NewReaderSize(stdin, 1024*1024)
	writer := bufio.NewWriterSize(stdout, 1024*1024)

	var n, m uint

	_, err = fmt.Fscan(reader, &n, &m)
	checkError(err)

	result := FibonacciNumber(n, m)
	fmt.Fprint(writer, result)

	err = writer.Flush()
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
