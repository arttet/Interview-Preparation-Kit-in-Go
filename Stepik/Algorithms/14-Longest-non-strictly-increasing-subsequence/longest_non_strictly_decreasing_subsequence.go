package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
	"sort"
)

const (
	maxInt int = (1<<bits.UintSize)/2 - 1
	minInt int = (1 << bits.UintSize) / -2
)

func longestNonStrictlyDecreasingSubsequence(arr []int) []int {
	n := len(arr)
	if n == 0 {
		return nil
	}

	dp := make([]int, n+1)
	position := make([]int, n+1)
	previous := make([]int, n)

	dp[0] = maxInt
	for i := 1; i <= n; i++ {
		dp[i] = minInt
	}

	position[0] = -1

	length := 0

	for i, num := range arr {
		j := sort.Search(n, func(i int) bool { return num > dp[i] })
		if num >= dp[j] {
			dp[j] = num
			position[j] = i
			previous[i] = position[j-1]
			if j > length {
				length = j
			}
		}
	}

	result := make([]int, length)
	for i, j := position[length], length-1; i != -1; i = previous[i] {
		result[j] = i + 1
		j--
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

	var n int
	_, err = fmt.Fscanln(reader, &n)
	checkError(err)

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		_, err = fmt.Fscan(reader, &arr[i])
		checkError(err)
	}

	result := longestNonStrictlyDecreasingSubsequence(arr)
	fmt.Fprintln(writer, len(result))
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
