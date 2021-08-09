package main

import (
	"bufio"
	"fmt"
	"os"
)

func activityNotifications(expenditures []int, d int) int {
	var notifications int

	const maxExpenditure = 201
	histogram := make([]int, maxExpenditure)

	var i, j int
	for i = 0; i < d; i++ {
		histogram[expenditures[i]]++
	}

	for n := len(expenditures); i < n; i++ {
		doubleMedian := 0

		cursor := 0
		left := -1

		for j = 0; j < maxExpenditure; j++ {
			cursor += histogram[j]
			if d%2 == 1 {
				// Odd -> Pick middle one for median
				if cursor >= d/2+1 {
					doubleMedian = 2 * j
					break
				}
			} else {
				// Even -> Pick average of two middle values for median
				if cursor == d/2 {
					left = j
				}

				if cursor > d/2 && left != -1 {
					right := j
					doubleMedian = left + right
					break
				}

				if cursor > d/2 && left == -1 {
					doubleMedian = 2 * j
					break
				}
			}
		}

		if expenditures[i] >= doubleMedian {
			notifications++
		}

		// Update histogram: slide window 1 index to right
		histogram[expenditures[i-d]]--
		histogram[expenditures[i]]++
	}

	return notifications
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

	var n, d int
	_, err = fmt.Fscan(reader, &n, &d)
	checkError(err)

	expenditures := make([]int, n)

	for i := range expenditures {
		_, err = fmt.Fscan(reader, &expenditures[i])
		checkError(err)
	}

	result := activityNotifications(expenditures, d)
	fmt.Fprint(writer, result)
	writer.Flush()
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
