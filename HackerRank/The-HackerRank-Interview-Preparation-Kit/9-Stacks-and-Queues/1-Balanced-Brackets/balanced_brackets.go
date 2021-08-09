package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func isBalanced(s string) string {
	stack := make([]rune, 0, len(s))

	for _, ch := range s {
		if ch == '{' || ch == '[' || ch == '(' {
			stack = append(stack, ch)
		} else {
			n := len(stack)
			if n == 0 {
				return "NO"
			}

			top := stack[n-1]
			stack = stack[:n-1]

			if top != '{' && ch == '}' || top != '[' && ch == ']' || top != '(' && ch == ')' {
				return "NO"
			}
		}
	}

	if len(stack) != 0 {
		return "NO"
	}

	return "YES"
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

	n, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)

	for i := 0; i < int(n); i++ {
		s := readLine(reader)
		result := isBalanced(s)

		fmt.Fprintf(writer, "%s\n", result)
	}

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}
	checkError(err)
	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
