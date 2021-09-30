package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type pair struct {
	ch    rune
	index int
}

func isBalanced(s string) string {
	stack := make([]pair, 0, len(s))

	parentheses := map[rune]byte{
		'{': 1,
		'[': 1,
		'(': 1,
		'}': 2,
		']': 2,
		')': 2,
	}

	for i, ch := range s {
		if value := parentheses[ch]; value == 1 {
			stack = append(stack, pair{
				ch:    ch,
				index: i,
			})
		} else if value == 2 {
			n := len(stack)
			if n == 0 {
				return strconv.Itoa(i + 1)
			}

			top := stack[n-1]
			stack = stack[:n-1]

			if top.ch != '{' && ch == '}' || top.ch != '[' && ch == ']' || top.ch != '(' && ch == ')' {
				return strconv.Itoa(i + 1)
			}
		}
	}

	if n := len(stack); n != 0 {
		return strconv.Itoa(stack[n-1].index + 1)
	}

	return "Success"
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

	s := readLine(reader)
	result := isBalanced(s)

	fmt.Fprintf(writer, "%s\n", result)

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
