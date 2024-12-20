package main

import (
	"bufio"
	"fmt"
	"os"
)

func countingSort(arr []int, maxValue int) []int {
	count := make([]int, maxValue)

	for _, n := range arr {
		count[n]++
	}

	for i := 1; i < len(count); i++ {
		count[i] += count[i-1]
	}

	sorted := make([]int, len(arr))
	for _, n := range arr {
		sorted[count[n]-1] = n
		count[n]--
	}

	return sorted
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

	var n int
	_, err = fmt.Fscanln(reader, &n)
	checkError(err)

	arr := make([]int, n)
	for i := range n {
		_, err = fmt.Fscan(reader, &arr[i])
		checkError(err)
	}

	const maxValue = 11
	result := countingSort(arr, maxValue)
	for _, counter := range result {
		fmt.Fprintf(writer, "%d ", counter)
	}

	err = writer.Flush()
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
