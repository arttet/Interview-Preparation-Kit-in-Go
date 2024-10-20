package main

import (
	"bufio"
	"fmt"
	"os"
)

func twoStrings(lhs, rhs string) string {
	dict := make(map[rune]bool)

	var char rune
	for _, char = range lhs {
		dict[char] = true
	}

	for _, char = range rhs {
		if dict[char] {
			return "YES"
		}
	}

	return "NO"
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

	err = writer.Flush()
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
