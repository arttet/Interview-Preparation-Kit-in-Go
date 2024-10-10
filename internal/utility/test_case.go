package utility

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	In  string
	Out string

	Input    string
	Output   string
	Expected string
}

const INPUT_PATH_ENV = "INPUT_PATH"
const OUTPUT_PATH_ENV = "OUTPUT_PATH"

func (test *TestCase) RunTest(t *testing.T, run func()) {
	t.Helper()
	ast := assert.New(t)

	err := os.Setenv(INPUT_PATH_ENV, test.In)
	checkError(err)

	dir, err := os.MkdirTemp("", "test_*_dir")
	checkError(err)
	defer os.RemoveAll(dir)

	fileName := fmt.Sprintf("output.%d.txt", time.Now().UnixNano())
	tempFileName := filepath.Join(dir, fileName)
	err = os.Setenv(OUTPUT_PATH_ENV, tempFileName)
	checkError(err)

	run()

	content, err := os.ReadFile(test.In)
	checkError(err)
	test.Input = strings.TrimSpace(string(content))
	ast.NotEmpty(test.Input)

	content, err = os.ReadFile(tempFileName)
	checkError(err)
	test.Output = strings.TrimSpace(string(content))

	content, err = os.ReadFile(test.Out)
	if err == nil {
		test.Expected = strings.TrimSpace(string(content))
	}

	err = os.Unsetenv(INPUT_PATH_ENV)
	checkError(err)

	err = os.Unsetenv(OUTPUT_PATH_ENV)
	checkError(err)

	ast.Equal(test.Expected, test.Output, "Test Case: %v %v", test.In, test.Out)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
