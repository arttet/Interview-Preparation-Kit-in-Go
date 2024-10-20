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
	prefix := make([]pair, n)

	var current, maxSum int64
	for i := range n {
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
	} else {
		defer stdin.Close()
	}

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	if err != nil {
		stdout = os.Stdout
	} else {
		defer stdout.Close()
	}

	reader := bufio.NewReaderSize(stdin, 1024*1024)
	writer := bufio.NewWriterSize(stdout, 1024*1024)

	var q int
	_, err = fmt.Fscan(reader, &q)
	checkError(err)

	for range q {
		var n int
		var m int64

		_, err = fmt.Fscan(reader, &n, &m)
		checkError(err)

		arr := make([]int64, n)
		for j := range n {
			_, err = fmt.Fscan(reader, &arr[j])
			checkError(err)
		}

		answer := maximumSum(arr, m)
		fmt.Fprintln(writer, answer)
	}

	err = writer.Flush()
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
