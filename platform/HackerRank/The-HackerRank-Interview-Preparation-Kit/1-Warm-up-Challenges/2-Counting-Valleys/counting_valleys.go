package main

import (
	"bufio"
	"fmt"
	"os"
)

func countingValleys(_ int, s string) int {
	var valleys int
	var currentLevel int

	for _, chart := range s {
		if chart == 'U' {
			currentLevel++
		} else {
			currentLevel--
		}

		if currentLevel == 0 && chart == 'U' {
			valleys++
		}
	}

	return valleys
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

	var n int
	var str string

	_, err = fmt.Fscan(reader, &n)
	checkError(err)

	_, err = fmt.Fscan(reader, &str)
	checkError(err)

	result := countingValleys(n, str)
	fmt.Fprintln(writer, result)

	err = writer.Flush()
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
