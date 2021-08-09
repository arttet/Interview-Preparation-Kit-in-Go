package main

import (
	"bufio"
	"fmt"
	"os"
)

type Vertice struct {
	// Key is the unique identifier of the vertex
	key    int
	weight int
}

type Edge struct {
	vertice Vertice
}

type Graph struct {
	nodes map[Vertice][]Edge // adjacency list
}

func newUndirectedGraph() *Graph {
	return &Graph{
		nodes: make(map[Vertice][]Edge),
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

	var n, m int
	_, err = fmt.Fscan(reader, &n, &m)
	checkError(err)

	// graph := newUndirectedGraph()

	for i := 0; i < m; i++ {
		var v1, v2, weight int
		_, err = fmt.Fscan(reader, &v1, &v2, &weight)
		checkError(err)

		// graph.addEdge(x, y, w)
	}

	writer.Flush()
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
