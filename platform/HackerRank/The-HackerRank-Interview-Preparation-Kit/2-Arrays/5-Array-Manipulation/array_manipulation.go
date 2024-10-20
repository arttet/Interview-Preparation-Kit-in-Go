package main

import (
	"bufio"
	"fmt"
	"os"
)

func arrayManipulation(n int, queries [][]int) int64 {
	arr := make([]int64, n)

	for _, query := range queries {
		a, b, k := query[0]-1, query[1], int64(query[2])
		arr[a] += k
		if b < n {
			arr[b] -= k
		}
	}

	for i := 1; i < len(arr); i++ {
		arr[i] += arr[i-1]
	}

	maximum := arr[0]
	for _, value := range arr {
		if maximum < value {
			maximum = value
		}
	}

	return maximum
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

	var n, m int
	_, err = fmt.Fscan(reader, &n, &m)
	checkError(err)

	queries := make([][]int, m)
	for i := range m {
		const size = 3
		queries[i] = make([]int, size)
		_, err = fmt.Fscan(reader, &queries[i][0], &queries[i][1], &queries[i][2])
		checkError(err)
	}

	result := arrayManipulation(n, queries)
	fmt.Fprint(writer, result)

	err = writer.Flush()
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
