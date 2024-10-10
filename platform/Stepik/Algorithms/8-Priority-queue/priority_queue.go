package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
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

	h := &IntHeap{}
	heap.Init(h)

	for ; n > 0; n-- {
		_, _ = fmt.Fscan(reader, &command)
		if command == "Insert" {
			_, err = fmt.Fscan(reader, &value)
			checkError(err)

			heap.Push(h, value)
		} else if command == "ExtractMax" && h.Len() > 0 {
			fmt.Fprintf(writer, "%d\n", heap.Pop(h))
		}
	}

	writer.Flush()
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
