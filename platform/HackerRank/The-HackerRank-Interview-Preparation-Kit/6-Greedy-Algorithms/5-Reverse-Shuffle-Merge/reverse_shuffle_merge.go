package main

import (
	"bufio"
	"fmt"
	"os"
)

func reverseShuffleMerge(s string) string {
	unused := [26]int{}
	used := [26]int{}
	required := [26]int{}

	for i := range s {
		unused[s[i]-'a']++
	}

	for i := range unused {
		required[i] = unused[i] / 2
	}

	n := len(s)
	result := make([]byte, n/2)

	for i, j := n-1, 0; i >= 0; i-- {
		ch := s[i] - 'a'

		if i == n-1 || used[ch] < required[ch] {
			for j > 0 && ch < result[j-1] && used[result[j-1]]-1+unused[result[j-1]] >= required[result[j-1]] {
				j--
				used[result[j]]--
			}

			result[j] = ch
			j++
			used[ch]++
		}

		unused[ch]--
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
	}
	defer stdin.Close()

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	if err != nil {
		stdout = os.Stdout
	}
	defer stdout.Close()

	reader := bufio.NewReaderSize(stdin, 1024*1024)
	writer := bufio.NewWriterSize(stdout, 1024*1024)

	var s string
	_, err = fmt.Fscan(reader, &s)
	checkError(err)

	result := reverseShuffleMerge(s)
	fmt.Fprint(writer, result)
	writer.Flush()
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
