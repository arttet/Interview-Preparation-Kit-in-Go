package main

import (
	"bufio"
	"fmt"
	"os"
)

type Node struct {
	left  *Node
	right *Node
	data  int
	level int
}

func NewNode(data, level int) *Node {
	return &Node{
		data:  data,
		level: level,
		left:  nil,
		right: nil,
	}
}

func newTree(indexes [][]int) *Node {
	var queue []*Node

	root := NewNode(1, 1)
	queue = append(queue, root)

	for i := range indexes {
		left, right := indexes[i][0], indexes[i][1]
		current := queue[0]

		if left != -1 {
			leftNode := NewNode(left, current.level+1)
			current.left = leftNode
			queue = append(queue, leftNode)
		}

		if right != -1 {
			rightNode := NewNode(right, current.level+1)
			current.right = rightNode
			queue = append(queue, rightNode)
		}

		queue = queue[1:]
	}

	return root
}

func (tree *Node) traverseInorder(result *[]int) {
	if tree.left != nil {
		tree.left.traverseInorder(result)
	}

	*result = append(*result, tree.data)

	if tree.right != nil {
		tree.right.traverseInorder(result)
	}
}

func (tree *Node) swap(k int) {
	if tree.left != nil {
		tree.left.swap(k)
	}

	if tree.right != nil {
		tree.right.swap(k)
	}

	if tree.level%k == 0 {
		tree.left, tree.right = tree.right, tree.left
	}
}

func swapNodes(indexes [][]int, queries []int) [][]int {
	result := make([][]int, len(queries))
	tree := newTree(indexes)

	for i, k := range queries {
		tree.swap(k)
		tree.traverseInorder(&result[i])
	}

	return result
}

func main() {
	stdin, err := os.Open(os.Getenv("INPUT_PATH"))
	if err != nil {
		stdin = os.Stdin
	} else {
		defer stdin.Close()
	}

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	if err != nil {
		stdout = os.Stdout
	} else {
		defer stdout.Close()
	}

	reader := bufio.NewReaderSize(stdin, 1024*1024)
	writer := bufio.NewWriterSize(stdout, 1024*1024)

	var n int
	_, err = fmt.Fscan(reader, &n)
	checkError(err)

	indexes := make([][]int, n)
	for i := range n {
		indexes[i] = make([]int, 2)

		_, err = fmt.Fscan(reader, &indexes[i][0], &indexes[i][1])
		checkError(err)
	}

	var k int
	_, err = fmt.Fscan(reader, &k)
	checkError(err)

	queries := make([]int, k)
	for i := range k {
		_, err = fmt.Fscan(reader, &queries[i])
		checkError(err)
	}

	result := swapNodes(indexes, queries)

	for i := range result {
		for j := range result[i] {
			fmt.Fprint(writer, result[i][j])

			if j != len(result[i])-1 {
				fmt.Fprint(writer, " ")
			}
		}

		if i != len(result)-1 {
			fmt.Fprintln(writer)
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
