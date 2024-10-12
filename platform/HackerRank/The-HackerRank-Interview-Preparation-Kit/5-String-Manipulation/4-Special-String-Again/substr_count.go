package main

import (
	"bufio"
	"fmt"
	"os"
)

func substrCount(n int, s string) int64 {
	var result int64
	for p := 0; p < n; p++ {
		var repeat int64 = 1
		for ; p < n-1 && s[p] == s[p+1]; p++ {
			repeat++
		}
		result += repeat * (repeat + 1) / 2

		for i := 1; p-i >= 0 && p+i < n && s[p+i] == s[p-1] && s[p-i] == s[p-1]; i++ {
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

	var n int
	var s string
	_, err = fmt.Fscan(reader, &n, &s)
	checkError(err)

	result := substrCount(n, s)
	fmt.Fprintln(writer, result)

	err = writer.Flush()
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
