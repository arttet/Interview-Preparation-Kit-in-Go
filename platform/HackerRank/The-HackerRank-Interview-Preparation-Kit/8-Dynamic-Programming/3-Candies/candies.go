package main

import (
	"bufio"
	"fmt"
	"os"
)

func candies(n int, rating []int) int64 {
	candies := make([]int, n)
	candies[0] = 1

	for i := 1; i < n; i++ {
		candies[i] = 1
		if rating[i] > rating[i-1] {
			candies[i] += candies[i-1]
		}
	}

	for i := n - 2; i >= 0; i-- {
		if rating[i] > rating[i+1] && candies[i] < candies[i+1]+1 {
			candies[i] = candies[i+1] + 1
		}
	}

	var result int64
	for i := range candies {
		result += int64(candies[i])
	}

	return result
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

	var n int
	_, err = fmt.Fscanln(reader, &n)
	checkError(err)

	arr := make([]int, n)
	for i := range n {
		_, err = fmt.Fscanln(reader, &arr[i])
		checkError(err)
	}

	result := candies(n, arr)
	fmt.Fprintf(writer, "%d\n", result)

	err = writer.Flush()
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
