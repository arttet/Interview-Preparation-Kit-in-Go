package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type pair struct {
	char  rune
	index int
}

func isBalanced(str string) string {
	stack := make([]pair, 0, len(str))

	parentheses := map[rune]byte{
		'{': 1,
		'[': 1,
		'(': 1,
		'}': 2,
		']': 2,
		')': 2,
	}

	for i, char := range str {
		if value := parentheses[char]; value == 1 {
			stack = append(stack, pair{
				char:  char,
				index: i,
			})
		} else if value == 2 {
			n := len(stack)
			if n == 0 {
				return strconv.Itoa(i + 1)
			}

			top := stack[n-1]
			stack = stack[:n-1]

			if top.char != '{' && char == '}' || top.char != '[' && char == ']' || top.char != '(' && char == ')' {
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

	err = writer.Flush()
	checkError(err)
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	checkError(err)

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
