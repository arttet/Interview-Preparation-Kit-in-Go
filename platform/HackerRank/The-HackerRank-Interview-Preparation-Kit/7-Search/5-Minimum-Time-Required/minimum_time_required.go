package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func minimumTime(machines []int64, goal int64) int64 {
	sort.Slice(machines, func(i, j int) bool { return machines[i] < machines[j] })

	slowestMachine, fastestMachine, n := machines[0], machines[len(machines)-1], int64(len(machines))
	lowerBound := goal * slowestMachine / n
	upperBound := goal * fastestMachine / n

	for lowerBound < upperBound {
		numDays := (lowerBound + upperBound) >> 1
		if total := getNumItems(machines, numDays); total >= goal {
			upperBound = numDays
		} else {
			lowerBound = numDays + 1
		}
	}

	return lowerBound
}

func getNumItems(machines []int64, numDays int64) int64 {
	var result int64
	for _, machine := range machines {
		result += numDays / machine
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

	var n int
	var goal int64

	_, err = fmt.Fscan(reader, &n, &goal)
	checkError(err)

	machines := make([]int64, n)
	for i := range n {
		_, err = fmt.Fscan(reader, &machines[i])
		checkError(err)
	}

	answer := minimumTime(machines, goal)
	fmt.Fprint(writer, answer)

	err = writer.Flush()
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
