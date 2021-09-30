package main

import (
	"bufio"
	"fmt"
	"os"
)

func isValid(s string) string {
	counter := countLetters(s)

	var max int
	var removed bool
	for _, occurrence := range counter {
		if occurrence == 0 || occurrence == max {
			continue
		}

		if max == 0 {
			max = occurrence
			continue
		}

		if !removed && (occurrence == max+1 || occurrence == 1) {
			removed = true
			continue
		}
		return "NO"
	}

	return "YES"
}

func countLetters(s string) []int {
	counts := make([]int, 26)
	for _, r := range s {
		counts[r-'a']++
	}

	return counts
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

	result := isValid(s)
	fmt.Fprintln(writer, result)

	writer.Flush()
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
