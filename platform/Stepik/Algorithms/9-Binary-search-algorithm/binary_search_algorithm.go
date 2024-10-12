package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func binarySearchAlgorithm(arr []int, x int) int {
	i := sort.Search(len(arr), func(i int) bool { return arr[i] >= x })
	if i < len(arr) && arr[i] == x {
		return i + 1
	}

	return -1
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
	for i := range n {
		_, err = fmt.Fscan(reader, &arr[i])
		checkError(err)
	}

	var k, value int
	_, err = fmt.Fscan(reader, &k)
	checkError(err)

	for range k {
		_, err = fmt.Fscan(reader, &value)
		checkError(err)

		result := binarySearchAlgorithm(arr, value)
		fmt.Fprintf(writer, "%d ", result)
	}

	err = writer.Flush()
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
