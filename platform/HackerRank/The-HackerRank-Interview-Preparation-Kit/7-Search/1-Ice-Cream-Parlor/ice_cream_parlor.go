package main

import (
	"bufio"
	"fmt"
	"os"
)

func whatFlavors(costs []int, money int) (int, int) {
	var lhs, rhs int

	mapping := make(map[int]int)

	for i, v := range costs {
		cost, ok := mapping[money-v]
		if ok {
			lhs = cost
			rhs = i + 1
		} else {
			mapping[v] = i + 1
		}
	}

	return lhs, rhs
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

	var t int
	_, err = fmt.Fscan(reader, &t)
	checkError(err)

	var (
		money, n int
		lhs, rhs int
	)

	for range t {
		_, err = fmt.Fscan(reader, &money, &n)
		checkError(err)

		costs := make([]int, n)
		for j := range n {
			_, err = fmt.Fscan(reader, &costs[j])
			checkError(err)
		}

		lhs, rhs = whatFlavors(costs, money)
		fmt.Fprintln(writer, lhs, rhs)
	}

	err = writer.Flush()
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
