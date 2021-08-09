package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// MaxStack is a variant of a stack data structure that can return the largest value.
type MaxStack struct {
	stack []item
}

type item struct {
	value   int
	maximum int
}

// MaxQueue is a variant of a queue data structure that can return the largest value.
type MaxQueue struct {
	leftStack  MaxStack
	rightStack MaxStack
}

// Constructor initializes MaxQueue.
func Constructor(k int) MaxQueue {
	return MaxQueue{
		leftStack: MaxStack{
			stack: make([]item, 0, k),
		},
		rightStack: MaxStack{
			stack: make([]item, 0, k),
		},
	}
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

// Push element x onto queue.
func (queue *MaxQueue) Push(x int) {
	queue.leftStack.Push(x)
}

func (queue *MaxQueue) balance() {
	for !queue.leftStack.Empty() {
		element := queue.leftStack.Top()
		queue.leftStack.Pop()

		queue.rightStack.Push(element)
	}
}

// Pop removes the element on top of the queue. */
func (queue *MaxQueue) Pop() {
	if queue.rightStack.Empty() {
		queue.balance()
	}

	queue.rightStack.Pop()
}

// GetMax returns the largest value.
func (queue *MaxQueue) GetMax() int {
	if queue.rightStack.Empty() {
		queue.balance()
	}

	if !queue.leftStack.Empty() {
		return max(queue.leftStack.GetMax(), queue.rightStack.GetMax())
	}

	return queue.rightStack.GetMax()
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxSlidingWindow(arr []int, k int) []int {
	n := len(arr)
	queue := Constructor(k)

	var i, j int
	for ; i < k; i++ {
		queue.Push(arr[i])
	}

	result := make([]int, n-k+1)
	for ; i < n; i, j = i+1, j+1 {
		result[j] = queue.GetMax()

		queue.Pop()
		queue.Push(arr[i])
	}
	result[j] = queue.GetMax()

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

	n, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)

	arr := make([]int, int(n))

	nums := strings.Fields(readLine(reader))
	for i := range nums {
		arr[i], err = strconv.Atoi(nums[i])
		checkError(err)
	}

	m, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)

	result := maxSlidingWindow(arr, int(m))
	for i := range result {
		fmt.Fprintf(writer, "%d ", result[i])
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
