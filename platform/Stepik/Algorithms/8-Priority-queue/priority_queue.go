package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

// An MaxIntHeap is a min-heap of ints.
type MaxIntHeap []int

func (h MaxIntHeap) Len() int           { return len(h) }
func (h MaxIntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxIntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxIntHeap) Push(value any) {
	val, _ := value.(int)
	*h = append(*h, val)
}

func (h *MaxIntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]

	return x
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
	var value int
	var command string

	_, err = fmt.Fscan(reader, &n)
	checkError(err)

	maxHeap := &MaxIntHeap{}
	heap.Init(maxHeap)

	for ; n > 0; n-- {
		_, err = fmt.Fscan(reader, &command)
		checkError(err)

		if command == "Insert" {
			_, err = fmt.Fscan(reader, &value)
			checkError(err)

			heap.Push(maxHeap, value)
		} else if command == "ExtractMax" && maxHeap.Len() > 0 {
			fmt.Fprintf(writer, "%d\n", heap.Pop(maxHeap))
		}
	}

	err = writer.Flush()
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
