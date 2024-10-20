package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func triplets(a, b, c []int) int64 {
	a = removeDuplicates(a)
	b = removeDuplicates(b)
	c = removeDuplicates(c)

	sort.Ints(a)
	sort.Ints(c)

	var result int64

	for _, item := range b {
		result += int64(equalOrGreaterThan(item, a) * equalOrGreaterThan(item, c))
	}

	return result
}

func equalOrGreaterThan(q int, arr []int) int {
	return sort.Search(len(arr), func(i int) bool { return arr[i] > q })
}

func removeDuplicates(arr []int) []int {
	dict := make(map[int]bool)

	for _, item := range arr {
		dict[item] = true
	}

	result := make([]int, 0, len(dict))
	for k := range dict {
		result = append(result, k)
	}

	return result
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

	var lenA, lenB, lenC int
	_, err = fmt.Fscan(reader, &lenA, &lenB, &lenC)
	checkError(err)

	a := make([]int, lenA)
	for i := range lenA {
		_, err = fmt.Fscan(reader, &a[i])
		checkError(err)
	}

	b := make([]int, lenB)
	for i := range lenB {
		_, err = fmt.Fscan(reader, &b[i])
		checkError(err)
	}

	c := make([]int, lenC)
	for i := range lenC {
		_, err = fmt.Fscan(reader, &c[i])
		checkError(err)
	}

	result := triplets(a, b, c)
	fmt.Fprint(writer, result)

	err = writer.Flush()
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
