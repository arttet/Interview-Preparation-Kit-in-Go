package main

import (
	"bufio"
	"fmt"
	"os"
)

func knapsack(weights []uint, weight uint) uint {
	var sum uint
	for _, w := range weights {
		sum += w
	}

	if sum <= weight {
		return sum
	}

	dp := [2][]uint{}
	for i := 0; i != 2; i++ {
		dp[i] = make([]uint, weight+1)
	}

	var w, wi uint
	for i := range weights {
		wi = weights[i]

		for w = 1; w <= weight; w++ {
			dp[1][w] = dp[0][w]
			if wi <= w && dp[0][w-wi]+wi <= w {
				dp[1][w] = max(dp[1][w], dp[0][w-wi]+wi)
			}
		}

		dp[0], dp[1] = dp[1], make([]uint, weight+1)
	}

	return dp[0][weight]
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

	weights := make([]uint, n)
	for i := range weights {
		_, err = fmt.Fscan(reader, &weights[i])
		checkError(err)
	}

	result := knapsack(weights, weight)
	fmt.Fprintln(writer, result)

	writer.Flush()
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
