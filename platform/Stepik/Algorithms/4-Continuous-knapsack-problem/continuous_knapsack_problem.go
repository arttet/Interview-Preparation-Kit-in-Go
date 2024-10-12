package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type knapsack struct {
	cost   int
	weight int
}

func continuousKnapsackProblem(arr []knapsack, weight int) float64 {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i].cost*arr[j].weight > arr[j].cost*arr[i].weight
	})

	var sum float64

	for i := range arr {
		item := &arr[i]
		if weight > item.weight {
			sum += float64(item.cost)
			weight -= item.weight
		} else {
			sum += float64(weight*item.cost) / float64(item.weight)

			break
		}
	}

	return sum
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

	var n, weight int

	_, err = fmt.Fscan(reader, &n, &weight)
	checkError(err)

	arr := make([]knapsack, n)
	for i := range n {
		_, err = fmt.Fscan(reader, &arr[i].cost, &arr[i].weight)
		checkError(err)
	}

	result := continuousKnapsackProblem(arr, weight)
	fmt.Fprintf(writer, "%0.3f", result)

	err = writer.Flush()
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
