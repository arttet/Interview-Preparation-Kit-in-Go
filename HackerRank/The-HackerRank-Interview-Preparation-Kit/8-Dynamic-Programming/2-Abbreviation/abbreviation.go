package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func abbreviation(a string, b string) string {
	if a == b {
		return "YES"
	}

	runeA, runeB := []rune(a), []rune(b)
	m, n := len(runeA), len(runeB)

	if m == 0 && n != 0 || m != 0 && n == 0 {
		return "NO"
	}

	previous := make([]bool, m+1)
	current := make([]bool, m+1)

	previous[0] = true
	for j := 1; j <= m; j++ {
		previous[j] = unicode.IsLower(runeA[j-1]) && previous[j-1]
	}

	var i, j int
	for i = 1; i <= n; i++ {
		for j = 1; j <= m; j++ {
			if j == 1 {
				current[0] = false
			}

			if unicode.IsLower(runeA[j-1]) {
				current[j] = previous[j-1] && runeB[i-1] == unicode.ToUpper(runeA[j-1]) || current[j-1]
			} else {
				current[j] = previous[j-1] && runeB[i-1] == runeA[j-1]
			}
		}

		previous, current = current, previous
	}

	if previous[m] {
		return "YES"
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

	var q int
	fmt.Fscanln(reader, &q)

	var a, b string
	for i := 0; i < q; i++ {
		_, err = fmt.Fscanln(reader, &a)
		checkError(err)
		_, err = fmt.Fscanln(reader, &b)
		checkError(err)

		answer := abbreviation(a, b)
		fmt.Fprintln(writer, answer)
	}

	writer.Flush()
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}