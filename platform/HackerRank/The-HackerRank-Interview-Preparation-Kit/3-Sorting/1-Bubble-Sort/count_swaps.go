package main

import (
	"bufio"
	"fmt"
	"os"
)

func countSwaps(arr []int) ([]int, int) {
	var i, j int
	n := len(arr)

	var numSwaps int
	for i = 0; i < n; i++ {
		for j = 0; j < n-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				numSwaps++
			}
		}
	}

	return arr, numSwaps
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
	_, err = fmt.Fscan(reader, &n)
	checkError(err)

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		_, err = fmt.Fscan(reader, &arr[i])
		checkError(err)
	}

	arr, numSwaps := countSwaps(arr)

	_, _ = fmt.Fprintf(writer, "Array is sorted in %d swaps.\n", numSwaps)
	_, _ = fmt.Fprintf(writer, "First Element: %d\n", arr[0])
	_, _ = fmt.Fprintf(writer, "Last Element: %d\n", arr[n-1])

	writer.Flush()
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
