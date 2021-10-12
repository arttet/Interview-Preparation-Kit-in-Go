package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func maxMin(k int, arr []int) int {
	sort.SliceStable(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})

	result := arr[len(arr)-1]
	for i := 0; i < len(arr)-k+1; i++ {
		if value := arr[i+k-1] - arr[i]; value < result {
			result = value
		}
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

	var n, k int
	_, err = fmt.Fscan(reader, &n, &k)
	checkError(err)

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		_, err = fmt.Fscan(reader, &arr[i])
		checkError(err)
	}

	result := maxMin(k, arr)
	fmt.Fprint(writer, result)
	writer.Flush()
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
