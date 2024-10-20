package main

import (
	"bufio"
	"fmt"
	"os"
)

func repeatedString(s string, n int64) int64 {
	q := n / int64(len(s))
	r := n % int64(len(s))

	var result int64
	for i, ch := range s {
		if ch == 'a' {
			result += q
			if int64(i) < r {
				result++
			}
		}
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

	var str string
	_, err = fmt.Fscan(reader, &str)
	checkError(err)

	var n int64
	_, err = fmt.Fscan(reader, &n)
	checkError(err)

	result := repeatedString(str, n)
	fmt.Fprintln(writer, result)

	err = writer.Flush()
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
