package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func maximumToys(prices []int, k int) int {
	sort.Ints(prices)

	var result int

	for _, price := range prices {
		if k < price {
			break
		}

		k -= price
		result++
	}

	return result
}

func main() {
	stdin, err := os.Open(os.Getenv("INPUT_PATH"))
	if err != nil {
		stdin = os.Stdin
	} else {
		defer func() { err = stdin.Close(); checkError(err) }()
	}

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	if err != nil {
		stdout = os.Stdout
	} else {
		defer func() { err = stdout.Close(); checkError(err) }()
	}

	reader := bufio.NewReaderSize(stdin, 1024*1024)
	writer := bufio.NewWriterSize(stdout, 1024*1024)

	var n, k int

	_, err = fmt.Fscan(reader, &n, &k)
	checkError(err)

	arr := make([]int, n)
	for i := range n {
		_, err = fmt.Fscan(reader, &arr[i])
		checkError(err)
	}

	result := maximumToys(arr, k)
	_, err = fmt.Fprint(writer, result)
	checkError(err)

	err = writer.Flush()
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
