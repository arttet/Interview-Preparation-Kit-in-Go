package main

import (
	"bufio"
	"fmt"
	"os"
)

func minimumBribes(queue []int) {
	var i, j int
	var moves int

	for i = range queue {
		if queue[i]-i-1 > 2 {
			fmt.Println("Too chaotic")
			return
		}

		for j = max(queue[i]-2, 0); j < i; j++ {
			if queue[j] > queue[i] {
				moves++
			}
		}
	}

	fmt.Println(moves)
}

func max(lhs int, rhs int) int {
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

	reader := bufio.NewReaderSize(stdin, 1024*1024)

	var t, n int
	_, err = fmt.Fscan(reader, &t)
	checkError(err)

	for ; t > 0; t-- {
		_, err = fmt.Fscan(reader, &n)
		checkError(err)
		queue := make([]int, n)
		for i := 0; i < n; i++ {
			_, err = fmt.Fscan(reader, &queue[i])
			checkError(err)
		}
		minimumBribes(queue)
	}
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
