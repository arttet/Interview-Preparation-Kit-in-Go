package main

import (
	"bufio"
	"fmt"
	"os"
)

func makeAnagram(a, b string) int {
	return isAnagram(countLetters(a), countLetters(b))
}

func isAnagram(countsA, countsB []int) int {
	var result int

	for i := range countsA {
		if countsA[i] != countsB[i] {
			result += abs(countsA[i] - countsB[i])
		}
	}

	return result
}

func countLetters(s string) []int {
	const alphabetSize = 26

	counts := make([]int, alphabetSize)
	for _, ch := range s {
		counts[ch-'a']++
	}

	return counts
}

func abs(num int) int {
	if num < 0 {
		return -num
	}

	return num
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

	var s1, s2 string
	_, err = fmt.Fscan(reader, &s1, &s2)
	checkError(err)

	result := makeAnagram(s1, s2)
	fmt.Fprintln(writer, result)

	err = writer.Flush()
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
