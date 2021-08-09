package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Queue is a queue data structure.
type Queue struct {
	leftStack  []int
	rightStack []int
}

// Constructor initializes Queue.
func Constructor() Queue {
	return Queue{}
}

func (queue *Queue) balance() {
	for len(queue.leftStack) != 0 {
		element := queue.leftStack[len(queue.leftStack)-1]
		queue.leftStack = queue.leftStack[:len(queue.leftStack)-1]

		queue.rightStack = append(queue.rightStack, element)
	}
}

// Push element x to the end of the queue.
func (queue *Queue) Push(x int) {
	queue.leftStack = append(queue.leftStack, x)
}

// Pop removes the element on top of the queue. */
func (queue *Queue) Pop() {
	if len(queue.rightStack) == 0 {
		queue.balance()
	}

	queue.rightStack = queue.rightStack[:len(queue.rightStack)-1]
}

// Top gets the top element.
func (queue *Queue) Top() int {
	if len(queue.rightStack) == 0 {
		queue.balance()
	}

	return queue.rightStack[len(queue.rightStack)-1]
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

	queue := Constructor()
	for i := 0; i < int(q); i++ {
		query := strings.Fields(readLine(reader))
		queryType, err := strconv.ParseInt(query[0], 10, 64)
		checkError(err)

		switch queryType {
		case 1:
			value, err := strconv.ParseInt(query[1], 10, 64)
			checkError(err)
			queue.Push(int(value))
		case 2:
			queue.Pop()
		case 3:
			fmt.Fprintf(writer, "%d\n", queue.Top())
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
