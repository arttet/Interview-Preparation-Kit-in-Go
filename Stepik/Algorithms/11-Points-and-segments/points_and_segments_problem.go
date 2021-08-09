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

	for i := 0; i < n; i++ {
		arr = append(arr, point{
			value:  segments[i][0],
			index:  i,
			border: left,
		})

		arr = append(arr, point{
			value:  segments[i][1],
			index:  i,
			border: right,
		})
	}

	for i := 0; i < m; i++ {
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
		if arr[i].border == left {
			counter++
		} else if arr[i].border == right {
			counter--
		} else {
			result[arr[i].index] = counter
		}
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

	var n, m int
	_, err = fmt.Fscanln(reader, &n, &m)
	checkError(err)

	segments := make([][]int, n)
	for i := 0; i < n; i++ {
		segments[i] = make([]int, 2)
		_, err = fmt.Fscanln(reader, &segments[i][0], &segments[i][1])
		checkError(err)
	}

	points := make([]int, m)
	for i := 0; i < m; i++ {
		_, err = fmt.Fscan(reader, &points[i])
		checkError(err)
	}

	result := pointsAndSegmentsProblem(segments, points)
	for _, counter := range result {
		fmt.Fprintf(writer, "%d ", counter)
	}

	writer.Flush()
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
