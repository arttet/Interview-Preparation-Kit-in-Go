package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type pivotID int

const (
	firstPivot pivotID = iota
	lastPivot
	medianOfThreePivot
)

func countComparisons(arr []int, id pivotID) int64 {
	return quickSort(arr, 0, len(arr)-1, id)
}

func quickSort(arr []int, left int, right int, id pivotID) int64 {
	if left >= right {
		return 0
	}

	choosePivot(arr, left, right, id)
	p := partition(arr, left, right)

	count := int64(right - left)
	count += quickSort(arr, left, p-1, id)
	count += quickSort(arr, p+1, right, id)

	return count
}

func choosePivot(arr []int, left int, right int, id pivotID) {
	switch id {
	case firstPivot:
	case lastPivot:
		arr[left], arr[right] = arr[right], arr[left]
	case medianOfThreePivot:
		pivotIndex := max(arr, min(arr, left, right), min(arr, max(arr, left, right), (left+right)/2))
		arr[left], arr[pivotIndex] = arr[pivotIndex], arr[left]
	}
}

func partition(arr []int, left int, right int) int {
	pivot := arr[left]

	i := left
	for j := left + 1; j <= right; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	arr[i], arr[left] = arr[left], arr[i]
	return i
}

func max(arr []int, left int, right int) int {
	if arr[left] > arr[right] {
		return left
	}
	return right
}

func min(arr []int, left int, right int) int {
	if arr[left] < arr[right] {
		return left
	}
	return right
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

	reader := bufio.NewScanner(stdin)
	writer := bufio.NewWriterSize(stdout, 1024*1024)

	var arr []int
	for value := 0; reader.Scan(); arr = append(arr, value) {
		value, err = strconv.Atoi(reader.Text())
		checkError(err)
	}

	pivots := []pivotID{firstPivot, lastPivot, medianOfThreePivot}
	for _, pivot := range pivots {
		array := make([]int, len(arr))
		copy(array, arr)

		result := countComparisons(array, pivot)
		fmt.Fprintln(writer, result)
	}

	writer.Flush()
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
