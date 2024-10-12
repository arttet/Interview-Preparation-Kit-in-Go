package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	InsertCase = iota + 1
	DeleteCase
	CheckCase
)

func frequencyQueries(queries [][]int) []int {
	var result []int

	number := make(map[int]int)    // number:frequency
	frequency := make(map[int]int) // frequency:quantity

	for _, query := range queries {
		switch action, value := query[0], query[1]; action {
		case InsertCase:
			freq := number[value]
			number[value]++
			frequency[number[value]]++
			if freq > 0 {
				frequency[freq]--
			}
		case DeleteCase:
			if freq, ok := number[value]; ok {
				number[value]--
				if number[value] == 0 {
					delete(number, value)
				} else {
					frequency[number[value]]++
				}

				frequency[freq]--
				if frequency[freq] == 0 {
					delete(frequency, freq)
				}
			}
		case CheckCase:
			item := 0
			if freq, ok := frequency[value]; ok && freq > 0 {
				item = 1
			}
			result = append(result, item)
		}
	}

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

	queries := make([][]int, q)
	for i := range q {
		queries[i] = make([]int, 2)
		_, err = fmt.Fscan(reader, &queries[i][0], &queries[i][1])
		checkError(err)
	}

	result := frequencyQueries(queries)
	for i := range result {
		fmt.Fprintln(writer, result[i])
	}

	err = writer.Flush()
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
