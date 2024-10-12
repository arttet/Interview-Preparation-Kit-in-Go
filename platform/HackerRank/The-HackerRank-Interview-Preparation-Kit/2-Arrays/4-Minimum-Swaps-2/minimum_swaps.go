package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type pair struct {
	index int
	value int
}

func minimumSwaps(arr []int) int {
	var i, j int
	arrayPosition := make([]pair, len(arr))

	for i = range arr {
		arrayPosition[i].index = i
		arrayPosition[i].value = arr[i]
	}

	sort.Slice(arrayPosition, func(i, j int) bool {
		return arrayPosition[i].value < arrayPosition[j].value
	})

	visited := make([]bool, len(arr))
	var swaps int

	for i = range visited {
		if visited[i] || arrayPosition[i].index == i {
			continue
		}

		var cycle int
		for j = i; !visited[j]; cycle++ {
			visited[j] = true
			j = arrayPosition[j].index
		}

		if cycle > 0 {
			swaps += cycle - 1
		}
	}

	return swaps
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

	arr := make([]int, n)
	for i := range n {
		_, err = fmt.Fscan(reader, &arr[i])
		checkError(err)
	}

	result := minimumSwaps(arr)
	fmt.Fprint(writer, result)

	err = writer.Flush()
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
