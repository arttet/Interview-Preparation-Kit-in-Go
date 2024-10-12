package main

import (
	"bufio"
	"fmt"
	"os"
)

func commonChild(lhs, rhs string) int {
	var i, j int
	m, n := len(lhs), len(rhs)

	longestCommonSubsequence := make([][]int, m+1)
	for i = range longestCommonSubsequence {
		longestCommonSubsequence[i] = make([]int, n+1)
	}

	for i = 1; i <= m; i++ {
		for j = 1; j <= n; j++ {
			if lhs[i-1] == rhs[j-1] {
				longestCommonSubsequence[i][j] = longestCommonSubsequence[i-1][j-1] + 1
			} else {
				longestCommonSubsequence[i][j] = max(
					longestCommonSubsequence[i-1][j],
					longestCommonSubsequence[i][j-1])
			}
		}
	}

	return longestCommonSubsequence[m][n]
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

	var s1, s2 string
	_, err = fmt.Fscan(reader, &s1, &s2)
	checkError(err)

	result := commonChild(s1, s2)
	fmt.Fprintln(writer, result)

	err = writer.Flush()
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
