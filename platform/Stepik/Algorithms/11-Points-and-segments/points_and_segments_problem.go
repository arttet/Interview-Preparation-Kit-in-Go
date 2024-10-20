package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

const (
	left  = iota
	none  = iota
	right = iota
)

type point struct {
	index  int
	value  int
	border int
}

func pointsAndSegmentsProblem(segments [][]int, points []int) []int {
	n, m := len(segments), len(points)
	arr := make([]point, 0, n+2*m)

	for i := range n {
		arr = append(arr,
			point{
				value:  segments[i][0],
				index:  i,
				border: left,
			},
			point{
				value:  segments[i][1],
				index:  i,
				border: right,
			},
		)
	}

	for i := range m {
		arr = append(arr, point{
			value:  points[i],
			index:  i,
			border: none,
		})
	}

	sort.SliceStable(arr, func(i, j int) bool {
		return arr[i].value < arr[j].value || arr[i].value == arr[j].value && arr[i].border < arr[j].border
	})

	var counter int
	result := make([]int, m)

	for i := range arr {
		switch arr[i].border {
		case left:
			counter++
		case right:
			counter--
		default:
			result[arr[i].index] = counter
		}
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

	var n, m int
	_, err = fmt.Fscanln(reader, &n, &m)
	checkError(err)

	segments := make([][]int, n)
	for i := range n {
		segments[i] = make([]int, 2)
		_, err = fmt.Fscanln(reader, &segments[i][0], &segments[i][1])
		checkError(err)
	}

	points := make([]int, m)
	for i := range m {
		_, err = fmt.Fscan(reader, &points[i])
		checkError(err)
	}

	result := pointsAndSegmentsProblem(segments, points)
	for _, counter := range result {
		fmt.Fprintf(writer, "%d ", counter)
	}

	err = writer.Flush()
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
