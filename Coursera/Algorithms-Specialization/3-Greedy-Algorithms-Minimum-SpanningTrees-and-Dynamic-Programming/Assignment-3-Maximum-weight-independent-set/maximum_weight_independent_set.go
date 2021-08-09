package main

import (
	"bufio"
	"fmt"
	"os"
)

func maximumWeightIndependentSet(maxWeights []int) map[int]int {
	i := len(maxWeights) - 1
	set := make(map[int]int)

	for i >= 1 {
		if maxWeights[i] == maxWeights[i-1] {
			i--
		} else {
			set[i] = 1
			i -= 2
		}
	}

	return set
}

func maximumWeightCache(weights []int) []int {
	for i := 2; i < len(weights); i++ {
		weights[i] = max(weights[i-1], weights[i-2]+weights[i])
	}
	return weights
}

func max(lhs, rhs int) int {
	if lhs > rhs {
		return lhs
	}
	return rhs
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
	_, err = fmt.Fscanln(reader, &n)
	checkError(err)

	weights := make([]int, n+1)
	for i := 1; i <= n; i++ {
		_, err = fmt.Fscanln(reader, &weights[i])
	}

	result := maximumWeightIndependentSet(maximumWeightCache(weights))
	for _, vertex := range []int{1, 2, 3, 4, 17, 117, 517, 997} {
		fmt.Fprint(writer, result[vertex])
	}

	writer.Flush()
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
