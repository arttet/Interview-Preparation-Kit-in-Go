package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// ValueType is a type of data.
type ValueType uint

// MinStack is a variant of a stack data structure that can return the largest value.
type MinStack struct {
	stack []item
}

type item struct {
	value   ValueType
	minimum ValueType
}

// MinQueue is a variant of a queue data structure that can return the largest value.
type MinQueue struct {
	leftStack  MinStack
	rightStack MinStack
}

// Constructor initializes MaxQueue.
func Constructor(k int) MinQueue {
	return MinQueue{
		leftStack: MinStack{
			stack: make([]item, 0, k),
		},
		rightStack: MinStack{
			stack: make([]item, 0, k),
		},
	}
}

// Push element x onto stack.
func (stack *MinStack) Push(x ValueType) {
	minimum := x
	if !stack.Empty() && stack.GetMin() < x {
		minimum = stack.GetMin()
	}
	stack.stack = append(stack.stack, item{value: x, minimum: minimum})
}

// Pop removes the element on top of the stack. */
func (stack *MinStack) Pop() {
	stack.stack = stack.stack[:stack.Size()-1]
}

// Top gets the top element.
func (stack *MinStack) Top() ValueType {
	return stack.stack[stack.Size()-1].value
}

// GetMin returns the largest value.
func (stack *MinStack) GetMin() ValueType {
	return stack.stack[stack.Size()-1].minimum
}

// Size returns the number of elements in the stack.
func (stack *MinStack) Size() int {
	return len(stack.stack)
}

// Empty returns whether the stack is empty.
func (stack *MinStack) Empty() bool {
	return stack.Size() == 0
}

// Clear removes all elements from the stack (
func (stack *MinStack) Clear() {
	stack.stack = stack.stack[:0]
}

// Push element x onto queue.
func (queue *MinQueue) Push(x ValueType) {
	queue.leftStack.Push(x)
}

func (queue *MinQueue) balance() {
	for !queue.leftStack.Empty() {
		element := queue.leftStack.Top()
		queue.leftStack.Pop()

		queue.rightStack.Push(element)
	}
}

// Pop removes the element on top of the queue. */
func (queue *MinQueue) Pop() {
	if queue.rightStack.Empty() {
		queue.balance()
	}

	queue.rightStack.Pop()
}

// GetMin returns the largest value.
func (queue *MinQueue) GetMin() ValueType {
	if queue.rightStack.Empty() {
		queue.balance()
	}

	if !queue.leftStack.Empty() {
		return min(queue.leftStack.GetMin(), queue.rightStack.GetMin())
	}

	return queue.rightStack.GetMin()
}

// Clear removes all elements from the queue (
func (queue *MinQueue) Clear() {
	queue.leftStack.Clear()
	queue.rightStack.Clear()
}

func min(a, b ValueType) ValueType {
	if a < b {
		return a
	}
	return b
}

func max(nums []ValueType) ValueType {
	maximum := nums[0]
	for _, num := range nums {
		if num > maximum {
			maximum = num
		}
	}
	return maximum
}

func (queue *MinQueue) maximumOfTheMinimumSlidingWindow(arr []ValueType, n int, k int) ValueType {
	var i, j int
	for ; i < k; i++ {
		queue.Push(arr[i])
	}

	var maximumOfTheMinimum, min ValueType
	for ; i < n; i, j = i+1, j+1 {
		min = queue.GetMin()
		if maximumOfTheMinimum < min {
			maximumOfTheMinimum = min
		}

		queue.Pop()
		queue.Push(arr[i])
	}

	min = queue.GetMin()
	if maximumOfTheMinimum < min {
		maximumOfTheMinimum = min
	}

	return maximumOfTheMinimum
}

func riddle(arr []ValueType) []ValueType {
	n := len(arr)

	result := make([]ValueType, n)
	result[0] = max(arr)

	queue := Constructor(n)
	for i := 1; i < n; i++ {
		result[i] = queue.maximumOfTheMinimumSlidingWindow(arr, n, i+1)
		queue.Clear()
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

	n, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)

	arr := make([]ValueType, int(n))

	nums := strings.Fields(readLine(reader))
	for i := range nums {
		value, err := strconv.Atoi(nums[i])
		checkError(err)
		arr[i] = ValueType(value)
	}

	result := riddle(arr)
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
