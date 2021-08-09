package main

import (
	"bufio"
	"container/heap"
	"errors"
	"fmt"
	"os"
	"sort"
	"strings"
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

func huffmanEncode(s string) (string, map[rune]string, error) {
	if len(s) == 0 {
		return "", nil, errors.New("Illegal argument: s cannot be empty")
	}

	frequencies := make(map[rune]int)
	for _, ch := range s {
		frequencies[ch]++
	}

	tree := buildTree(frequencies)
	codes := make(map[rune]string)

	buildCodes(&tree, &codes, []byte{})

	var encoded strings.Builder
	for _, ch := range s {
		fmt.Fprint(&encoded, codes[ch])
	}

	return encoded.String(), codes, nil
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

	var s string

	_, err = fmt.Fscan(reader, &s)
	checkError(err)

	encoded, codes, err := huffmanEncode(s)
	checkError(err)

	fmt.Fprintf(writer, "%d %d\n", len(codes), len(encoded))
	for ch, code := range codes {
		fmt.Fprintf(writer, "%c: %s\n", ch, code)
	}
	fmt.Fprintf(writer, "%s", encoded)

	writer.Flush()
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
