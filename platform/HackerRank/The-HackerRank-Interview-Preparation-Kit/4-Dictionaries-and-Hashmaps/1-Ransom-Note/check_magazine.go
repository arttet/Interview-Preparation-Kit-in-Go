package main

import (
	"bufio"
	"fmt"
	"os"
)

func checkMagazine(magazine []string, note []string) {
	dict := make(map[string]int)

	var i int
	for i = range magazine {
		dict[magazine[i]]++
	}

	for i = range note {
		if dict[note[i]] > 0 {
			dict[note[i]]--
		} else {
			fmt.Print("No")
			return
		}
	}

	fmt.Print("Yes")
}

func main() {
	stdin, err := os.Open(os.Getenv("INPUT_PATH"))
	if err != nil {
		stdin = os.Stdin
	}
	defer stdin.Close()

	reader := bufio.NewReaderSize(stdin, 1024*1024)

	var m, n, i int
	_, err = fmt.Fscan(reader, &m, &n)
	checkError(err)

	magazine := make([]string, m)
	for i = 0; i < m; i++ {
		_, err = fmt.Fscan(reader, &magazine[i])
		checkError(err)
	}

	note := make([]string, n)
	for i = 0; i < n; i++ {
		_, err = fmt.Fscan(reader, &note[i])
		checkError(err)
	}

	checkMagazine(magazine, note)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
