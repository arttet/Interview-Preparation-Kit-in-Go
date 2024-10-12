package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func getMinimumCost(k int, c []int) int {
	sort.Ints(c)

	var minCost int
	if n := len(c); k >= n {
		for i := range n {
			minCost += c[i]
		}
	} else {
		var previousPurchases int

		for i, count := n-1, 0; i >= 0; i-- {
			if count == k {
				count = 0
				previousPurchases++
			}

			minCost += (previousPurchases + 1) * c[i]
			count++
		}
	}

	return minCost
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

	var n, k int
	_, err = fmt.Fscan(reader, &n, &k)
	checkError(err)

	arr := make([]int, n)
	for i := range n {
		_, err = fmt.Fscan(reader, &arr[i])
		checkError(err)
	}

	result := getMinimumCost(k, arr)
	fmt.Fprint(writer, result)

	err = writer.Flush()
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
