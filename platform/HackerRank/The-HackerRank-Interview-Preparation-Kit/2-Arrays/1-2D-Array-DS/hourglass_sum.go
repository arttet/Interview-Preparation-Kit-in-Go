package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func hourglassSum(arr [][]int) int {
	maximumSum := math.MinInt32

	m, n := len(arr), len(arr[0])
	for i := range m - 2 {
		for j := range n - 2 {
			sum := arr[i][j] + arr[i][j+1] + arr[i][j+2]
			sum += arr[i+1][j+1]
			sum += arr[i+2][j] + arr[i+2][j+1] + arr[i+2][j+2]

			if sum > maximumSum {
				maximumSum = sum
			}
		}
	}

	return maximumSum
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

	const size = 6

	arr := make([][]int, size)
	for i := range size {
		arr[i] = make([]int, size)
		for j := range size {
			_, err = fmt.Fscan(reader, &arr[i][j])
			checkError(err)
		}
	}

	result := hourglassSum(arr)
	fmt.Fprint(writer, result)

	err = writer.Flush()
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
