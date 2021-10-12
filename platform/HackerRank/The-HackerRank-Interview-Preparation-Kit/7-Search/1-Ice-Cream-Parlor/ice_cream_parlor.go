package main

import (
	"bufio"
	"fmt"
	"os"
)

func whatFlavors(costs []int, money int) (lhs int, rhs int) {
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
	}
	defer stdin.Close()

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	if err != nil {
		stdout = os.Stdout
	}
	defer stdout.Close()

	reader := bufio.NewReaderSize(stdin, 1024*1024)
	writer := bufio.NewWriterSize(stdout, 1024*1024)

	var t int
	_, err = fmt.Fscan(reader, &t)
	checkError(err)

	var money, n int
	var lhs, rhs int

	for i := 0; i < t; i++ {
		_, err = fmt.Fscan(reader, &money, &n)
		checkError(err)

		costs := make([]int, n)
		for j := 0; j < n; j++ {
			_, err = fmt.Fscan(reader, &costs[j])
			checkError(err)
		}

		lhs, rhs = whatFlavors(costs, money)
		fmt.Fprintln(writer, lhs, rhs)
	}

	writer.Flush()
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
