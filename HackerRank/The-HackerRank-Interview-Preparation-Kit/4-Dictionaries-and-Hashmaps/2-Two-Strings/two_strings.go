package main

import (
	"bufio"
	"fmt"
	"os"
)

func twoStrings(s1 string, s2 string) string {
	dict := make(map[rune]bool)

	var ch rune
	for _, ch = range s1 {
		dict[ch] = true
	}

	for _, ch = range s2 {
		if dict[ch] {
			return "YES"
		}
	}

	return "NO"
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

	var p int
	_, err = fmt.Fscan(reader, &p)
	checkError(err)

	var s1, s2 string
	for ; p > 0; p-- {
		_, err = fmt.Fscan(reader, &s1, &s2)
		checkError(err)

		result := twoStrings(s1, s2)
		fmt.Fprintln(writer, result)
	}

	writer.Flush()
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
