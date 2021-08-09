package main

import (
	"bufio"
	"fmt"
	"os"
)

// FibonacciNumber returns the nth value in the fibonacci series.
func FibonacciNumber(n int, m int) int {
	if n == 0 {
		return 0
	}

	F := [2][2]int{{1, 1}, {1, 0}}
	power(&F, n-1, m)

	return F[0][0]
}

func power(F *[2][2]int, n int, m int) {
	if n == 0 || n == 1 {
		return
	}

	power(F, n/2, m)
	multiply(F, F, m)

	if n%2 != 0 {
		M := [2][2]int{{1, 1}, {1, 0}}
		multiply(F, &M, m)
	}
}

func multiply(F *[2][2]int, M *[2][2]int, m int) {
	a := F[0][0]*M[0][0] + F[0][1]*M[1][0]
	b := F[0][0]*M[0][1] + F[0][1]*M[1][1]
	c := F[1][0]*M[0][0] + F[1][1]*M[1][0]
	d := F[1][0]*M[0][1] + F[1][1]*M[1][1]

	F[0][0] = a % m
	F[0][1] = b % m
	F[1][0] = c % m
	F[1][1] = d % m
}

func main() {
	stdin, err := os.Open(os.Getenv("INPUT_PATH"))
	if err != nil {
		stdin = os.Stdin
	}
	defer stdin.Close()

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	if err != nil {
		stdout = os.Stdout
	}
	defer stdout.Close()

	reader := bufio.NewReaderSize(stdin, 1024*1024)
	writer := bufio.NewWriterSize(stdout, 1024*1024)

	var n, m int

	_, err = fmt.Fscan(reader, &n, &m)
	checkError(err)

	result := FibonacciNumber(n, m)
	fmt.Fprint(writer, result)

	writer.Flush()
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
