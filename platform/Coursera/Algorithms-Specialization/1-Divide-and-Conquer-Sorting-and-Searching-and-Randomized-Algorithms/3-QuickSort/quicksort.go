package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type pivotIDType int

const (
	firstPivot pivotIDType = iota
	lastPivot
	medianOfThreePivot
)

func countComparisons(arr []int, pivotID pivotIDType) int64 {
	return quickSort(arr, 0, len(arr)-1, pivotID)
}

func quickSort(arr []int, left, right int, pivotID pivotIDType) int64 {
	if left >= right {
		return 0
	}

	choosePivot(arr, left, right, pivotID)
	p := partition(arr, left, right)

	count := int64(right - left)
	count += quickSort(arr, left, p-1, pivotID)
	count += quickSort(arr, p+1, right, pivotID)

	return count
}

func choosePivot(arr []int, left, right int, pivotID pivotIDType) {
	switch pivotID {
	case firstPivot:
	case lastPivot:
		arr[left], arr[right] = arr[right], arr[left]
	case medianOfThreePivot:
		pivotIndex := arrayMax(arr, arrayMin(arr, left, right), arrayMin(arr, arrayMax(arr, left, right), (left+right)/2))
		arr[left], arr[pivotIndex] = arr[pivotIndex], arr[left]
	}
}

func partition(arr []int, left, right int) int {
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

func arrayMax(arr []int, left, right int) int {
	if arr[left] > arr[right] {
		return left
	}

	return right
}

func arrayMin(arr []int, left, right int) int {
	if arr[left] < arr[right] {
		return left
	}

	return right
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

	reader := bufio.NewScanner(stdin)
	writer := bufio.NewWriterSize(stdout, 1024*1024)

	const initialCapacity = 64
	arr := make([]int, 0, initialCapacity)
	for reader.Scan() {
		var value int
		value, err = strconv.Atoi(reader.Text())
		checkError(err)
		arr = append(arr, value)
	}

	pivots := []pivotIDType{firstPivot, lastPivot, medianOfThreePivot}
	for _, pivot := range pivots {
		array := make([]int, len(arr))
		copy(array, arr)

		result := countComparisons(array, pivot)
		fmt.Fprintln(writer, result)
	}

	err = writer.Flush()
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
