package main

import (
	"bufio"
	"fmt"
	"os"
)

func pairs(k int, arr []int) int {
	dict := make(map[int]bool)

	var item int
	for _, item = range arr {
		dict[item] = true
	}

	var result int
	for item = range dict {
		if dict[item+k] {
			result++
		}
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

	var n, k int
	_, err = fmt.Fscan(reader, &n, &k)
	checkError(err)

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		_, err = fmt.Fscan(reader, &arr[i])
		checkError(err)
	}

	result := pairs(k, arr)
	fmt.Fprint(writer, result)

	writer.Flush()
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
