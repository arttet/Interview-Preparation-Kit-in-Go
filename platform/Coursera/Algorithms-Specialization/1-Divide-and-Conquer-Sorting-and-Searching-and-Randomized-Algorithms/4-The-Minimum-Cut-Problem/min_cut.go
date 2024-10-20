package main

import (
	"bufio"
	"crypto/rand"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"
	"sync"
)

/**
 * Karger's algorithm to compute the minimum cut of a connected graph.
 * Takes as input an undirected connected graph and computes a cut
 * with the fewest number of crossing edges. We need to repeat the
 * algorithm n * (n - 1) * ln(n) / 2 times to guarantee success.
 * The probability of not finding the minimum cut is 1 / n.
 * // https://en.wikipedia.org/wiki/Karger's_algorithm
 */

type vertice int

type edge struct {
	node vertice
}

type graph struct {
	nodes map[vertice][]edge // adjacency list
}

func newGraph() *graph {
	return &graph{
		nodes: make(map[vertice][]edge),
	}
}

func (g *graph) Len() int {
	return len(g.nodes)
}

func (g *graph) addEdge(origin, destiny vertice) {
	g.nodes[origin] = append(g.nodes[origin], edge{node: destiny})
}

func (g *graph) removeNode(node vertice) {
	delete(g.nodes, node)
}

func (g *graph) removeEdge(origin, destiny vertice) {
	edges := g.getEdges(origin)
	n := len(edges) - 1

	for i := range edges {
		vi := edges[i].node
		if vi == destiny {
			edges[i], edges[n] = edges[n], edges[i]
			g.nodes[origin] = edges[:n]

			break
		}
	}
}

func (g *graph) getEdges(node vertice) []edge {
	return g.nodes[node]
}

func (g *graph) clone() *graph {
	copyGraph := newGraph()

	for i, edges := range g.nodes {
		copyGraph.nodes[i] = make([]edge, len(edges))
		copy(copyGraph.nodes[i], edges)
	}

	return copyGraph
}

func minimumCut(g *graph, iterations int) int {
	cutSize := make([]int, iterations)

	var wg sync.WaitGroup
	for i := range iterations {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			cutSize[i] = g.clone().karger()
		}(i)
	}
	wg.Wait()

	minCut := cutSize[0]
	for i := range iterations {
		if cutSize[i] < minCut {
			minCut = cutSize[i]
		}
	}

	return minCut
}

func (g *graph) karger() int {
	var minCut int

	for v := g.Len(); v > 2; v-- {
		v1, v2 := g.randomNodes()

		// Adding the edges from the absorbed node:
		for _, edge := range g.getEdges(v2) {
			vi := edge.node
			if vi != v1 {
				g.addEdge(v1, vi)
			}
		}

		// Deleting the references to the absorbed node and
		// changing them to the source node:
		for _, edge := range g.getEdges(v2) {
			vi := edge.node
			g.removeEdge(vi, v2)
			if vi != v1 {
				g.addEdge(vi, v1)
			}
		}

		g.removeNode(v2)
	}

	for key := range g.nodes {
		minCut = len(g.nodes[key])

		break
	}

	return minCut
}

func (g *graph) randomNodes() (vertice, vertice) {
	var v1, v2 vertice

	n := g.Len()
	rnd, err := rand.Int(rand.Reader, big.NewInt(int64(n)))
	checkError(err)
	v1Index := int(rnd.Int64())

	var i int
	for node := range g.nodes {
		if i == v1Index {
			v1 = node

			break
		}
		i++
	}

	connectedV1 := g.getEdges(v1)

	m := len(connectedV1)
	rnd, err = rand.Int(rand.Reader, big.NewInt(int64(m)))
	checkError(err)

	v2Index := int(rnd.Int64())

	for i := range connectedV1 {
		if i == v2Index {
			v2 = connectedV1[i].node

			break
		}
	}

	return v1, v2
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

	reader := bufio.NewScanner(stdin)
	writer := bufio.NewWriterSize(stdout, 1024*1024)

	var origin, destiny int
	graph := newGraph()

	for reader.Scan() {
		edges := strings.Fields(reader.Text())
		origin, err = strconv.Atoi(edges[0])
		checkError(err)

		for i := 1; i < len(edges); i++ {
			destiny, err = strconv.Atoi(edges[i])
			checkError(err)

			graph.addEdge(vertice(origin), vertice(destiny))
		}
	}

	n := graph.Len()
	it := min(n*(n-1)/2, 1024)

	result := minimumCut(graph, it)
	fmt.Fprint(writer, result)

	err = writer.Flush()
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
