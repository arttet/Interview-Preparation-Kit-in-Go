package main

import (
	"bufio"
	"fmt"
	"os"
)

func sherlockAndAnagrams(str string) int {
	n := len(str)
	substrings := make([][][]int, n-1)

	// Work on one small test case:
	// "abba":
	// "a" | "b" | "b" | "a"
	// "ab" | "bb" | "ba"
	// "abb" | "bba"
	for length := 1; length < n; length++ {
		for begin := 0; begin <= n-length; begin++ {
			substring := str[begin : begin+length]
			substrings[length-1] = append(substrings[length-1], countLetters(substring))
		}
	}

	var counter int
	for k := range substrings {
		length := len(substrings[k])
		for i := range length {
			for j := i + 1; j < length; j++ {
				if isAnagram(substrings[k][i], substrings[k][j]) {
					counter++
				}
			}
		}
	}

	return counter
}

func isAnagram(countsA, countsB []int) bool {
	for i := range countsA {
		if countsA[i] != countsB[i] {
			return false
		}
	}

	return true
}

func countLetters(s string) []int {
	const alphabetSize = 26

	counts := make([]int, alphabetSize)
	for _, ch := range s {
		counts[ch-'a']++
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

	var q int
	_, err = fmt.Fscan(reader, &q)
	checkError(err)

	var s string
	for ; q > 0; q-- {
		_, err = fmt.Fscan(reader, &s)
		checkError(err)

		result := sherlockAndAnagrams(s)
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
