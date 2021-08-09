package main

import (
	"bufio"
	"fmt"
	"os"
)

func leftRotation(arr []int, d int) []int {
	n := len(arr)
	result := make([]int, n)

	for i := 0; i < n; i++ {
		shift := (n - d + i) % n
		result[shift] = arr[i]
	}

	return result
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

	var n, d int
	_, err = fmt.Fscan(reader, &n, &d)
	checkError(err)

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		_, err = fmt.Fscan(reader, &arr[i])
		checkError(err)
	}

	result := leftRotation(arr, d)
	for i := range result {
		fmt.Fprintf(writer, "%d ", result[i])
	}
	writer.Flush()
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
