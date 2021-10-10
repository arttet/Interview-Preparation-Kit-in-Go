package main

import (
	"bufio"
	"fmt"
	"os"
)

func lengthOfLDS(arr []int) int {
	n := len(arr)
	if n == 0 {
		return 0
	}

	dp := make([]int, n)
	dp[0] = 1

	result := 1
	for i := 1; i < n; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if arr[j] <= arr[i] && arr[i]%arr[j] == 0 {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}

		result = max(result, dp[i])
	}

	return result
}

func max(lhs, rhs int) int {
	if lhs > rhs {
		return lhs
	}

	return rhs
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
	_, err = fmt.Fscanln(reader, &n)
	checkError(err)

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		_, err = fmt.Fscan(reader, &arr[i])
		checkError(err)
	}

	result := lengthOfLDS(arr)
	fmt.Fprint(writer, result)

	writer.Flush()
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}