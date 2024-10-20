package main

import (
	"bufio"
	"fmt"
	"os"
)

func countTriplets(arr []int64, r int64) int64 {
	dictB := make(map[int64]int64)
	dictC := make(map[int64]int64)

	var numTriplets int64

	// Here v number can be an A or a B or a C from
	// Geometric progression triplet (A, B, C) with ratio r.
	for _, v := range arr {
		// number is a C the update the final count
		numTriplets += dictC[v]

		// number is a B
		// add number of Bs that exist into Cs
		dictC[v*r] += dictB[v]

		// number is an A
		// this means add the count of it being B
		dictB[v*r]++
	}

	return numTriplets
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

	var n, r int64
	_, err = fmt.Fscan(reader, &n, &r)
	checkError(err)

	arr := make([]int64, n)
	for i := range n {
		_, err = fmt.Fscan(reader, &arr[i])
		checkError(err)
	}

	result := countTriplets(arr, r)
	fmt.Fprintln(writer, result)

	err = writer.Flush()
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
