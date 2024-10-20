package main

import (
	"bufio"
	"fmt"
	"os"
)

func isValid(s string) string {
	counter := countLetters(s)

	var maxValue int
	var removed bool
	for _, occurrence := range counter {
		if occurrence == 0 || occurrence == maxValue {
			continue
		}

		if maxValue == 0 {
			maxValue = occurrence

			continue
		}

		if !removed && (occurrence == maxValue+1 || occurrence == 1) {
			removed = true

			continue
		}

		return "NO"
	}

	return "YES"
}

func countLetters(s string) []int {
	const alphabetSize = 26

	counts := make([]int, alphabetSize)
	for _, r := range s {
		counts[r-'a']++
	}

	return counts
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

	result := isValid(s)
	fmt.Fprintln(writer, result)

	err = writer.Flush()
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
