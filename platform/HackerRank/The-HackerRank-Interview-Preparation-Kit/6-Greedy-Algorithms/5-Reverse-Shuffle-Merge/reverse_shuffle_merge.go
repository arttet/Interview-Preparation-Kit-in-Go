package main

import (
	"bufio"
	"fmt"
	"os"
)

func reverseShuffleMerge(str string) string {
	unused := [26]int{}
	used := [26]int{}
	required := [26]int{}

	for i := range str {
		unused[str[i]-'a']++
	}

	for i := range unused {
		required[i] = unused[i] / 2
	}

	n := len(str)
	result := make([]byte, n/2)

	for i, j := n-1, 0; i >= 0; i-- {
		char := str[i] - 'a'

		if i == n-1 || used[char] < required[char] {
			for j > 0 && char < result[j-1] && used[result[j-1]]-1+unused[result[j-1]] >= required[result[j-1]] {
				j--
				used[result[j]]--
			}

			result[j] = char
			j++
			used[char]++
		}

		unused[char]--
	}

	for i := range result {
		result[i] += 'a'
	}

	return string(result)
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

	var s string
	_, err = fmt.Fscan(reader, &s)
	checkError(err)

	result := reverseShuffleMerge(s)
	fmt.Fprint(writer, result)

	err = writer.Flush()
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
