package main

import (
	"bufio"
	"fmt"
	"os"
)

func alternatingCharacters(s string) int {
	var result int

	var i, j int
	for i = 1; i < len(s); i++ {
		if s[i] == s[j] {
			result++
		} else {
			j = i
		}
	}

	// var previousCharacter rune
	// for _, ch := range s {
	// 	if ch == previousCharacter {
	// 		result++
	// 	}
	// 	previousCharacter = ch
	// }

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

	var q int
	_, err = fmt.Fscan(reader, &q)
	checkError(err)

	var s string
	for ; q > 0; q-- {
		_, err = fmt.Fscan(reader, &s)
		checkError(err)

		result := alternatingCharacters(s)
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
