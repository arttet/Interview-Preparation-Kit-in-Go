package utility

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestCase TODO
type TestCase struct {
	In  string
	Out string

	Input    string
	Output   string
	Expected string
}

// RunTest TODO
func (test *TestCase) RunTest(t *testing.T, run func()) {
	ast := assert.New(t)

	stdin, err := os.Open(test.In)
	checkError(err)
	defer stdin.Close()

	stdout, err := ioutil.TempFile("", "output.*.txt")
	checkError(err)
	defer os.Remove(stdout.Name())

	os.Stdin, os.Stdout = stdin, stdout

	run()

	content, err := ioutil.ReadFile(test.In)
	checkError(err)
	test.Input = strings.TrimSpace(string(content))
	ast.NotEmpty(test.Input)

	content, err = ioutil.ReadFile(stdout.Name())
	checkError(err)
	test.Output = strings.TrimSpace(string(content))
	ast.NotEmpty(test.Output)

	content, err = ioutil.ReadFile(test.Out)
	if err == nil {
		test.Expected = strings.TrimSpace(string(content))
	}
	ast.NotEmpty(test.Expected)

	ast.Equal(test.Expected, test.Output, "Test Case: %v %v", test.In, test.Out)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
