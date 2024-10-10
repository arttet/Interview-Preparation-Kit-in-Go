package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

// An Integer represents a signed multi-precision integer.
type Integer []int

func multiplication(s1, s2 string) string {
	n := getPowerOfTwo(max(len(s1), len(s2)))

	x := newInteger(s1, n)
	y := newInteger(s2, n)
	result := make(Integer, 6*n)

	karatsuba(x, y, result, n)
	result = result[:n<<1]

	result.doCarry()
	return result.String()
}

// Input: two n-digit positive integers x and y.
// Output: the product x * y.
// Assumption: n is a power of 2.
//
// |   b * d   |   a * c   |   p * q   | lower-recursion space |   p   |   q   |
// | n digits  |  n digits |  n digits |         2n digits     |  n/2  |   n/2 |
//
// The result must have space for 6n digits.
// The result will be in only the first 2n digits.
func karatsuba(x Integer, y Integer, result Integer, n int) {
	const cutOff = 4

	if n <= cutOff {
		gradeSchoolMultiplication(x, y, result, n)
		return
	}

	m := n / 2

	a, b := x[m:], x[:m] // x := 10^{n / 2} * a + b
	c, d := y[m:], y[:m] // y := 10^{n / 2} * c + d
	p := result[n*5:]    // p := a + b
	q := result[n*5+m:]  // q := c + d

	var i int
	for i = 0; i < m; i++ {
		p[i] = a[i] + b[i]
		q[i] = c[i] + d[i]
	}

	bd := result[0*n:]
	ac := result[1*n:]
	adbc := result[2*n:]

	karatsuba(b, d, bd, m)
	karatsuba(a, c, ac, m)
	karatsuba(p, q, adbc, m)

	for i = 0; i < n; i++ {
		adbc[i] = adbc[i] - ac[i] - bd[i] // adbc := pq − ac − bd
	}

	// ac^{10 * n} + adbc ^{10 * n/2} + bd
	for i = 0; i < n; i++ {
		result[i+m] += adbc[i]
	}
}

func gradeSchoolMultiplication(x Integer, y Integer, result Integer, n int) {
	var i, j int

	for i = 0; i < n<<1; i++ {
		result[i] = 0
	}

	for i = 0; i < n; i++ {
		for j = 0; j < n; j++ {
			result[i+j] += x[i] * y[j]
		}
	}
}

func newInteger(s string, n int) Integer {
	arr := make(Integer, n)

	for i, j := 0, len(s)-1; j >= 0; j-- {
		arr[i] = int(s[j] - '0')
		i++
	}

	return arr
}

func (x Integer) doCarry() {
	const base = 10

	var c int

	for i := range x {
		x[i] += c
		if x[i] >= 0 {
			c = x[i] / base
		} else {
			c = -(-(x[i]+1)/base + 1)
		}

		x[i] -= c * base
	}

	if c != 0 {
		panic(fmt.Sprintf("Overflow %d\n", c))
	}
}

func (x Integer) String() string {
	n := len(x)

	var i int
	for i = n - 1; i >= 0; i-- {
		if x[i] != 0 {
			break
		}
		n--
	}

	s := make([]byte, n)

	for j := 0; i >= 0; i-- {
		s[j] = byte(x[i]) + '0'
		j++
	}

	return string(s)
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

	var left, right string

	_, err = fmt.Fscan(reader, &left, &right)
	checkError(err)

	result := multiplication(left, right)
	fmt.Fprint(writer, result)

	writer.Flush()
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func getPowerOfTwo(n int) int {
	if n&(n-1) == 0 {
		return n
	}
	return 2 << int(math.Log(float64(n))/math.Log(2))
}
