package main

import (
	"bufio"
	"fmt"
	"os"
)

func countSwaps(arr []int) {
	var i, j int
	n := len(arr)

	var numSwaps int
	for i = 0; i < n; i++ {
		for j = 0; j < n-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				numSwaps++
			}
		}
	}

	fmt.Printf("Array is sorted in %d swaps.\n", numSwaps)
	fmt.Println("First Element:", arr[0])
	fmt.Println("Last Element:", arr[n-1])
}

func main() {
	stdin, err := os.Open(os.Getenv("INPUT_PATH"))
	if err != nil {
		stdin = os.Stdin
	}
	defer stdin.Close()

	reader := bufio.NewReaderSize(stdin, 1024*1024)

	var n int
	_, err = fmt.Fscan(reader, &n)
	checkError(err)

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		_, err = fmt.Fscan(reader, &arr[i])
		checkError(err)
	}

	countSwaps(arr)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
