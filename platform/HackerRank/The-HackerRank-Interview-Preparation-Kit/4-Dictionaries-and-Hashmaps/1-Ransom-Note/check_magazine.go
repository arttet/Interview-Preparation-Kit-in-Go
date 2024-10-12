package main

import (
	"bufio"
	"fmt"
	"os"
)

func checkMagazine(magazine, note []string) string {
	dict := make(map[string]int)

	for i := range magazine {
		dict[magazine[i]]++
	}

	for i := range note {
		if dict[note[i]] > 0 {
			dict[note[i]]--
		} else {
			return "No"
		}
	}

	return "Yes"
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

	var m, n int
	_, err = fmt.Fscan(reader, &m, &n)
	checkError(err)

	magazine := make([]string, m)
	for i := range m {
		_, err = fmt.Fscan(reader, &magazine[i])
		checkError(err)
	}

	note := make([]string, n)
	for i := range n {
		_, err = fmt.Fscan(reader, &note[i])
		checkError(err)
	}

	result := checkMagazine(magazine, note)
	fmt.Fprintln(writer, result)

	err = writer.Flush()
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
