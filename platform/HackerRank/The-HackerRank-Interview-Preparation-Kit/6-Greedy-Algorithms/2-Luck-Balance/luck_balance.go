package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func luckBalance(k int, contests [][]int) int {
	const L, T = 0, 1
	sort.Slice(contests, func(i, j int) bool {
		return contests[i][T] > contests[j][T] || contests[i][T] == contests[j][T] && contests[i][L] > contests[j][L]
	})

	var result int
	for i := 0; i < len(contests); i++ {
		if i < k || contests[i][T] == 0 {
			result += contests[i][L]
		} else {
			result -= contests[i][L]
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

	var n, k int
	_, err = fmt.Fscan(reader, &n, &k)
	checkError(err)

	contests := make([][]int, n)
	for i := 0; i < n; i++ {
		contests[i] = make([]int, 2)
		_, err = fmt.Fscan(reader, &contests[i][0], &contests[i][1])
		checkError(err)
	}

	result := luckBalance(k, contests)
	fmt.Fprint(writer, result)
	writer.Flush()
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
