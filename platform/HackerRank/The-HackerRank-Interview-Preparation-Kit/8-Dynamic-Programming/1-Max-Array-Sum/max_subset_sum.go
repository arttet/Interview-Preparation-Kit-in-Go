package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
)

const minInt int = (1 << bits.UintSize) / -2

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

func max(nums ...int) int {
	maximum := minInt
	for _, num := range nums {
		if maximum < num {
			maximum = num
		}
	}
	return maximum
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
	fmt.Fscanln(reader, &n)

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &arr[i])
	}

	answer := maxSubsetSum(arr)
	fmt.Fprintln(writer, answer)

	writer.Flush()
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
