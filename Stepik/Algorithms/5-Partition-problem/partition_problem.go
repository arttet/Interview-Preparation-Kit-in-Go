package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func partition(n int) []int {
	k := int(math.Sqrt(1+4*2*float64(n)) / 2)

	result := make([]int, k)

	i, sum := 0, 0
	for ; i < k-1; i++ {
		result[i] = i + 1
		sum += result[i]
	}

	remainder := n - sum
	if remainder < k {
		result[i-1] += remainder
		result = result[:i]
	} else {
		result[i] = remainder
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
	_, err = fmt.Fscan(reader, &n)
	checkError(err)

	result := partition(n)
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
