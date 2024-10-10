package main

import (
	"bufio"
	"fmt"
	"os"
)

func maxSum(stairs []int) int {
	var current, last, previous int

	for i := range stairs {
		current = max(previous, last) + stairs[i]
		previous = last
		last = current
	}

	return last
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

	var n int
	_, err = fmt.Fscanln(reader, &n)
	checkError(err)

	stairs := make([]int, n)
	for i := range stairs {
		_, err = fmt.Fscan(reader, &stairs[i])
		checkError(err)
	}

	result := maxSum(stairs)
	fmt.Fprintln(writer, result)

	writer.Flush()
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
