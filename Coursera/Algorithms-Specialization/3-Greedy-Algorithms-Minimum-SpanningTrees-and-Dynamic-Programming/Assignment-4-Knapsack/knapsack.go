package main

import (
	"bufio"
	"fmt"
	"os"
)

type item struct {
	value  uint
	weight uint
}

func knapsackWithoutRepetition(items []item, weight uint) uint {
	dp := [2][]uint{}
	for i := 0; i != 2; i++ {
		dp[i] = make([]uint, weight+1)
	}

	var w, wi, vi uint
	for i := range items {
		wi = items[i].weight
		vi = items[i].value
		for w = 1; w <= weight; w++ {
			dp[1][w] = dp[0][w]
			if wi <= w {
				dp[1][w] = max(dp[1][w], dp[0][w-wi]+vi)
			}
		}

		dp[0], dp[1] = dp[1], make([]uint, weight+1)
	}

	return dp[0][weight]
}

func max(lhs, rhs uint) uint {
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

	var weight, n uint
	_, err = fmt.Fscanln(reader, &weight, &n)
	checkError(err)

	items := make([]item, n)
	for i := range items {
		_, err = fmt.Fscanln(reader, &items[i].value, &items[i].weight)
	}

	result := knapsackWithoutRepetition(items, weight)
	fmt.Fprint(writer, result)

	writer.Flush()
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
