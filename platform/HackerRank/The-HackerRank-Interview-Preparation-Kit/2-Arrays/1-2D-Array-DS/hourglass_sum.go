package main

import (
	"bufio"
	"fmt"
	"os"
)

func hourglassSum(arr [][]int) int {
	maximumSum := -2147483648

	m, n := len(arr), len(arr[0])
	for i := 0; i < m-2; i++ {
		for j := 0; j < n-2; j++ {
			var sum int
			sum += arr[i][j] + arr[i][j+1] + arr[i][j+2]
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

	arr := make([][]int, 6)
	for i := 0; i < 6; i++ {
		arr[i] = make([]int, 6)
		for j := 0; j < 6; j++ {
			_, err = fmt.Fscan(reader, &arr[i][j])
			checkError(err)
		}
	}

	result := hourglassSum(arr)
	fmt.Fprint(writer, result)
	writer.Flush()
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
