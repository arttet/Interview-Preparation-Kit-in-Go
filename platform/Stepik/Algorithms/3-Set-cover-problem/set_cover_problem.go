package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type segment struct {
	left  int
	right int
}

type pair struct {
	index  int
	value  int
	isLeft bool
}

func setCoverProblem(arr []segment) []int {
	n := len(arr)
	points := make([]pair, 2*n)
	for i := range arr {
		index := 2 * i
		points[index].index = i
		points[index].value = arr[i].left
		points[index].isLeft = true

		points[index+1].index = i
		points[index+1].value = arr[i].right
		points[index+1].isLeft = false
	}

	sort.SliceStable(points, func(i, j int) bool {
		return points[i].value < points[j].value || points[i].value == points[j].value && points[i].isLeft
	})

	coverage := make([]bool, n)
	stack := make([]int, 0, n)

	var result []int

	for i := range points {
		if points[i].isLeft {
			stack = append(stack, points[i].index)
		} else if !coverage[points[i].index] {
			result = append(result, points[i].value)
			for len(stack) != 0 {
				n := len(stack) - 1
				coverage[stack[n]] = true
				stack = stack[:n]
			}
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

	var n int

	_, err = fmt.Fscan(reader, &n)
	checkError(err)

	arr := make([]segment, n)
	for i := 0; i < n; i++ {
		_, err = fmt.Fscan(reader, &arr[i].left, &arr[i].right)
		checkError(err)
	}

	result := setCoverProblem(arr)
	fmt.Fprintln(writer, len(result))

	for i := range result {
		fmt.Fprintf(writer, "%d ", result[i])
	}

	writer.Flush()
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
