package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type packet struct {
	arrival  int
	duration int
}

func networkPacketProcessingSimulation(arr []packet, size int) []int {
	return nil
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

	nums := strings.Fields(readLine(reader))

	size, err := strconv.Atoi(nums[0])
	checkError(err)
	n, err := strconv.Atoi(nums[1])
	checkError(err)

	arr := make([]packet, int(n))
	for i := 0; i < int(n); i++ {
		packetStr := strings.Fields(readLine(reader))

		arr[i].arrival, err = strconv.Atoi(packetStr[0])
		checkError(err)
		arr[i].duration, err = strconv.Atoi(packetStr[1])
		checkError(err)
	}

	result := networkPacketProcessingSimulation(arr, int(size))
	for i := range result {
		fmt.Fprintf(writer, "%d ", result[i])
	}

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}
	checkError(err)
	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
