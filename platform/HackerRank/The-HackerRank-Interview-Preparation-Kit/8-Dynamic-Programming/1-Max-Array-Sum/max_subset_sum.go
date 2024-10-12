package main

import (
	"bufio"
	"fmt"
	"os"
)

func maxSubsetSum(arr []int) int {
	n := len(arr)

	previous, last := arr[0], max(arr[0], arr[1])
	for i := 2; i < n; i++ {
		current := max(previous+arr[i], last, previous, arr[i])
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

	arr := make([]int, n)
	for i := range n {
		_, err = fmt.Fscan(reader, &arr[i])
		checkError(err)
	}

	answer := maxSubsetSum(arr)
	fmt.Fprintln(writer, answer)

	err = writer.Flush()
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
