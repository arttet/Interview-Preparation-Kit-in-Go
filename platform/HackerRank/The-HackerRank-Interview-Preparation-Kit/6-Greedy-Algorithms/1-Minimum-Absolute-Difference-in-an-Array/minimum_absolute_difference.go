package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

func minimumAbsoluteDifference(arr []int) int {
	sort.Ints(arr)

	result := math.MaxInt32
	for i := range len(arr) - 1 {
		if diff := arr[i+1] - arr[i]; diff < result {
			result = diff
		}
	}

	return result
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
	_, err = fmt.Fscan(reader, &n)
	checkError(err)

	arr := make([]int, n)
	for i := range n {
		_, err = fmt.Fscan(reader, &arr[i])
		checkError(err)
	}

	result := minimumAbsoluteDifference(arr)
	fmt.Fprint(writer, result)

	err = writer.Flush()
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
