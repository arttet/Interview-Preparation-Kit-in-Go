package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type HuffmanNode struct {
	value rune
	left  *HuffmanNode
	right *HuffmanNode
}

func buildTree(codes map[rune]string) *HuffmanNode {
	root := &HuffmanNode{}

	var current *HuffmanNode
	for ch := range codes {
		current = root

		for _, code := range codes[ch] {
			if code == '0' {
				if current.left == nil {
					current.left = &HuffmanNode{}
				}
				current = current.left
			} else {
				if current.right == nil {
					current.right = &HuffmanNode{}
				}
				current = current.right
			}
		}
		current.value = ch
	}

	return root
}

func huffmanDecode(codes map[rune]string, encoded string) string {
	root := buildTree(codes)

	var result strings.Builder
	result.Grow(64)

	current := root
	for _, code := range encoded {
		if code == '0' {
			current = current.left
		} else {
			current = current.right
		}

		if current.value != 0 {
			result.WriteRune(current.value)
			current = root
		}
	}

	return result.String()
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

	var k, l int
	_, err = fmt.Fscanln(reader, &k, &l)
	checkError(err)

	codes := make(map[rune]string)
	var letter rune
	var code string
	for ; k > 0; k-- {
		_, err = fmt.Fscanf(reader, "%c: %s\n", &letter, &code)
		checkError(err)

		codes[letter] = code
	}

	var encoded string
	_, err = fmt.Fscan(reader, &encoded)
	checkError(err)

	str := huffmanDecode(codes, encoded)
	fmt.Fprint(writer, str)

	writer.Flush()
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
