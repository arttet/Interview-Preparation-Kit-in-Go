package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"math"
	"os"
	"sort"
)

type HuffmanTree interface {
	Freq() int
}

type HuffmanLeaf struct {
	freq  int
	value rune
}

type HuffmanNode struct {
	freq  int
	left  HuffmanTree
	right HuffmanTree
}

func (self HuffmanLeaf) Freq() int {
	return self.freq
}

func (self HuffmanNode) Freq() int {
	return self.freq
}

type treeHeap []HuffmanTree

func (th *treeHeap) Len() int {
	return len(*th)
}

func (th *treeHeap) Less(i, j int) bool {
	return (*th)[i].Freq() < (*th)[j].Freq()
}

func (th *treeHeap) Push(element interface{}) {
	*th = append(*th, element.(HuffmanTree))
}

func (th *treeHeap) Pop() (popped interface{}) {
	popped = (*th)[len(*th)-1]
	*th = (*th)[:len(*th)-1]
	return
}

func (th treeHeap) Swap(i, j int) {
	th[i], th[j] = th[j], th[i]
}

func buildTree(frequencies map[rune]int) HuffmanTree {
	var trees treeHeap

	keys := make([]rune, 0, len(frequencies))
	for k := range frequencies {
		keys = append(keys, k)
	}
	sort.SliceStable(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})

	for _, ch := range keys {
		trees = append(trees, HuffmanLeaf{
			freq:  frequencies[ch],
			value: ch,
		})
	}

	heap.Init(&trees)
	for trees.Len() > 1 {
		a := heap.Pop(&trees).(HuffmanTree)
		b := heap.Pop(&trees).(HuffmanTree)
		heap.Push(&trees, HuffmanNode{a.Freq() + b.Freq(), a, b})
	}

	return heap.Pop(&trees).(HuffmanTree)
}

func buildCodes(tree *HuffmanTree, codes *map[rune]string, prefix []byte) {
	switch i := (*tree).(type) {
	case HuffmanLeaf:
		if len(prefix) == 0 {
			(*codes)[i.value] = string('0')
		} else {
			(*codes)[i.value] = string(prefix)
		}
	case HuffmanNode:
		prefix = append(prefix, '0')
		buildCodes(&i.left, codes, prefix)
		prefix = prefix[:len(prefix)-1]

		prefix = append(prefix, '1')
		buildCodes(&i.right, codes, prefix)
		prefix = prefix[:len(prefix)-1]
	}
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

	var n, value int
	_, err = fmt.Fscanln(reader, &n)
	checkError(err)

	frequencies := make(map[rune]int)
	for i := 0; i < n; i++ {
		_, err = fmt.Fscanln(reader, &value)
		frequencies[rune(i)] = value
		checkError(err)
	}

	tree := buildTree(frequencies)

	codes := make(map[rune]string)
	buildCodes(&tree, &codes, []byte{})

	maximumLength := math.MinInt64
	minimumLength := math.MaxInt64

	for _, code := range codes {
		n := len(code)

		if n < minimumLength {
			minimumLength = n
		}

		if n > maximumLength {
			maximumLength = n
		}
	}

	fmt.Fprintln(writer, maximumLength)
	fmt.Fprintln(writer, minimumLength)

	writer.Flush()
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
