package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// MaxStack is a variant of a stack data structure that can return the largest value
type MaxStack struct {
	stack []item
}

type item struct {
	value   int
	maximum int
}

// Constructor initializes MaxStack.
func Constructor() MaxStack {
	return MaxStack{}
}

// Push element x onto stack.
func (stack *MaxStack) Push(x int) {
	maximum := x
	if len(stack.stack) > 0 && stack.GetMax() > x {
		maximum = stack.GetMax()
	}
	stack.stack = append(stack.stack, item{value: x, maximum: maximum})
}

// Pop removes the element on top of the stack. */
func (stack *MaxStack) Pop() {
	stack.stack = stack.stack[:stack.Size()-1]
}

// Top gets the top element.
func (stack *MaxStack) Top() int {
	return stack.stack[stack.Size()-1].value
}

// GetMax returns the largest value.
func (stack *MaxStack) GetMax() int {
	return stack.stack[stack.Size()-1].maximum
}

// Size returns the number of elements in the stack.
func (stack *MaxStack) Size() int {
	return len(stack.stack)
}

// Empty returns whether the stack is empty.
func (stack *MaxStack) Empty() bool {
	return stack.Size() == 0
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

	q, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)

	stack := Constructor()

	for i := 0; i < int(q); i++ {
		query := strings.Fields(readLine(reader))
		queryType := query[0]

		switch queryType {
		case "push":
			value, err := strconv.ParseInt(query[1], 10, 64)
			checkError(err)
			stack.Push(int(value))
		case "pop":
			stack.Pop()
		case "max":
			fmt.Fprintf(writer, "%d\n", stack.GetMax())
		default:
		}
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
