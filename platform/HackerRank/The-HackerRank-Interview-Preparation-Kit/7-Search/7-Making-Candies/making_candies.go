package main

import (
	"bufio"
	"fmt"
	"math"
	"math/big"
	"os"
)

func minimumPasses(machines int64, workers int64, price int64, target int64) int64 {
	if new(big.Int).Mul(big.NewInt(machines), big.NewInt(workers)).Cmp(big.NewInt(target)) >= 0 {
		return 1
	}

	var left int64 = 1
	var right int64 = math.MaxInt64

	for left < right {
		passes := left + (right-left)>>1
		if check(machines, workers, price, target, passes) {
			right = passes
		} else {
			left = passes + 1
		}
	}

	return left
}

func check(machines int64, workers int64, price int64, target int64, passes int64) bool {
	// if machines >= (target+workers-1)/workers {
	// 	return true
	// }

	candies := machines * workers
	passes--

	for {
		rounds := (target - candies + machines*workers - 1) / (machines * workers)
		if rounds <= passes {
			return true
		}

		if candies < price {
			rounds = (price - candies + machines*workers - 1) / (machines * workers)
			passes -= rounds
			if passes < 1 {
				break
			}
			candies += rounds * machines * workers
		}

		candies -= price
		if machines > workers {
			workers++
		} else {
			machines++
		}
	}

	return false
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

	var m, w, p, n int64
	_, _ = fmt.Fscan(reader, &m, &w, &p, &n)

	answer := minimumPasses(m, w, p, n)
	fmt.Fprint(writer, answer)

	writer.Flush()
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
