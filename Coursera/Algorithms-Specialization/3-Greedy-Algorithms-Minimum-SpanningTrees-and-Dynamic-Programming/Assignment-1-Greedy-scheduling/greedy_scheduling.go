package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type job struct {
	weight int
	length int
}

func greedyScheduling(jobs []job, comparator func(i, j int) bool) int {
	sort.SliceStable(jobs, comparator)

	var completedTime, accumulatedTime int

	for i := range jobs {
		accumulatedTime += jobs[i].length
		completedTime += accumulatedTime * jobs[i].weight
	}

	return completedTime
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
	_, err = fmt.Fscan(reader, &n)
	checkError(err)

	jobs := make([]job, n)
	for i := 0; i < n; i++ {
		_, err = fmt.Fscan(reader, &jobs[i].weight, &jobs[i].length)
		checkError(err)
	}

	task1 := greedyScheduling(jobs, func(i, j int) bool {
		di := jobs[i].weight - jobs[i].length
		dj := jobs[j].weight - jobs[j].length
		return di > dj || di == dj && jobs[i].weight > jobs[j].weight
	})
	fmt.Fprintln(writer, task1)

	task2 := greedyScheduling(jobs, func(i, j int) bool {
		return jobs[i].weight*jobs[j].length > jobs[j].weight*jobs[i].length
	})
	fmt.Fprintln(writer, task2)

	writer.Flush()
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
