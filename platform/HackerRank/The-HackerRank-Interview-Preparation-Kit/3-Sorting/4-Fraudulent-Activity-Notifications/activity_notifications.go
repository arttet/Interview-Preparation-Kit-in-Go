package main

import (
	"bufio"
	"fmt"
	"os"
)

const maxExpenditure = 201

func activityNotifications(expenditures []int, d int) int {
	var notifications int
	histogram := buildHistogram(expenditures[:d])

	for n, i := len(expenditures), d; i < n; i++ {
		doubleMedian := calculateDoubleMedian(histogram, d)

		if expenditures[i] >= doubleMedian {
			notifications++
		}

		// Update histogram: slide window 1 index to right
		updateHistogram(histogram, expenditures[i-d], expenditures[i])
	}

	return notifications
}

func buildHistogram(expenditures []int) []int {
	histogram := make([]int, maxExpenditure)
	for _, exp := range expenditures {
		histogram[exp]++
	}

	return histogram
}

func calculateDoubleMedian(histogram []int, days int) int {
	var doubleMedian int

	cursor := 0
	left := -1

	for j := range maxExpenditure {
		cursor += histogram[j]

		if days%2 == 1 { //nolint: nestif
			// Odd
			if cursor >= days/2+1 {
				doubleMedian = 2 * j

				break
			}
		} else {
			// Even
			if cursor == days/2 {
				left = j
			}
			if cursor > days/2 {
				if left != -1 {
					doubleMedian = left + j
				} else {
					doubleMedian = 2 * j
				}

				break
			}
		}
	}

	return doubleMedian
}

func updateHistogram(histogram []int, outgoing, incoming int) {
	histogram[outgoing]--
	histogram[incoming]++
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

	err = writer.Flush()
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
