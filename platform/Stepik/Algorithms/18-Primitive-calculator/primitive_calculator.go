package main

import (
	"bufio"
	"fmt"
	"os"
)

func minOperations(n int) []int {
	dp := make([]int, n+1)
	dp[0] = 1

	for i := 1; i <= n; i++ {
		dp[i] = dp[i-1] + 1

		if i%2 == 0 {
			dp[i] = min(dp[i/2]+1, dp[i])
		}

		if i%3 == 0 {
			dp[i] = min(dp[i/3]+1, dp[i])
		}
	}

	result := make([]int, 0, n)
	for i := n; i >= 1; {
		result = append(result, i)
		if dp[i-1] == dp[i]-1 { // nolint: gocritic
			i--
		} else if i%2 == 0 && dp[i/2] == dp[i]-1 {
			i /= 2
		} else if i%3 == 0 && dp[i/3] == dp[i]-1 {
			i /= 3
		}
	}

	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	return result
}

func min(lhs, rhs int) int {
	if lhs < rhs {
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

	result := minOperations(n)
	fmt.Fprintln(writer, len(result)-1)
	for i := range result {
		fmt.Fprintf(writer, "%d ", result[i])
	}

	writer.Flush()
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
