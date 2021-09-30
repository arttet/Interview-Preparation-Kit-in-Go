package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type pair struct {
	index  int
	prefix int64
}

func maximumSum(arr []int64, m int64) int64 {
	n := len(arr)
	if n == 0 {
		return 0
	}

	prefix := make([]pair, n)

	var current, maxSum int64
	for i := 0; i < n; i++ {
		current = (arr[i]%m + current) % m
		prefix[i].prefix = current
		prefix[i].index = i
		if current > maxSum {
			maxSum = current
		}
	}

	if maxSum == m-1 {
		return maxSum
	}

	sort.SliceStable(prefix, func(i, j int) bool {
		return prefix[i].prefix < prefix[j].prefix
	})

	for i := 1; i < n; i++ {
		if prefix[i].index < prefix[i-1].index {
			diff := prefix[i].prefix - prefix[i-1].prefix
			if m-diff > maxSum && diff != 0 {
				maxSum = m - diff
				if maxSum == m-1 {
					break
				}
			}
		}
	}

	return maxSum
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

	var q int
	fmt.Fscan(reader, &q)

	for i := 0; i < q; i++ {
		var n int
		var m int64

		fmt.Fscan(reader, &n, &m)

		arr := make([]int64, n)
		for j := 0; j < n; j++ {
			fmt.Fscan(reader, &arr[j])
		}

		answer := maximumSum(arr, m)
		fmt.Fprintln(writer, answer)
	}

	writer.Flush()
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
