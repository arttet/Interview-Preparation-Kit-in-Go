package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func minimumBribes(queue []int) string {
	var i, j int
	var moves int

	for i = range queue {
		if queue[i]-i-1 > 2 {
			return "Too chaotic"
		}

		for j = max(queue[i]-2, 0); j < i; j++ {
			if queue[j] > queue[i] {
				moves++
			}
		}
	}

	return strconv.Itoa(moves)
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

	var t, n int
	_, err = fmt.Fscan(reader, &t)
	checkError(err)

	for ; t > 0; t-- {
		_, err = fmt.Fscan(reader, &n)
		checkError(err)
		queue := make([]int, n)
		for i := range n {
			_, err = fmt.Fscan(reader, &queue[i])
			checkError(err)
		}
		result := minimumBribes(queue)
		fmt.Fprintln(writer, result)
	}

	err = writer.Flush()
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
