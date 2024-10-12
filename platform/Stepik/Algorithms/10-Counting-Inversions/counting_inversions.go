package main

import (
	"bufio"
	"fmt"
	"os"
)

func countInversions(arr []int) int64 {
	return mergesort(arr, make([]int, len(arr)), 0, len(arr))
}

func mergesort(arr, temp []int, left, right int) int64 {
	if left == right-1 {
		return 0
	}

	middle := (left + right) / 2

	var count int64
	count += mergesort(arr, temp, left, middle)
	count += mergesort(arr, temp, middle, right)
	count += merge(arr, temp, left, middle, right)

	return count
}

func merge(arr, temp []int, left, middle, right int) int64 {
	var inversions int64

	i, j := left, middle
	k := left
	for ; i < middle && j < right; k++ {
		if arr[i] <= arr[j] {
			temp[k] = arr[i]
			i++
		} else {
			temp[k] = arr[j]
			j++
			inversions += int64(middle - i)
		}
	}

	k += copy(temp[k:], arr[i:middle])
	copy(temp[k:], arr[j:right])
	copy(arr[left:right], temp[left:right])

	return inversions
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

	result := countInversions(arr)
	fmt.Fprint(writer, result)

	err = writer.Flush()
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
