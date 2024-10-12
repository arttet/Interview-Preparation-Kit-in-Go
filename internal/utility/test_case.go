package utility

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const (
	InputPathEnv  = "INPUT_PATH"
	OutputPathEnv = "OUTPUT_PATH"
)

type TestCase struct {
	In  string
	Out string

	Input    string
	Output   string
	Expected string
}

func (test *TestCase) RunTest(t *testing.T, run func()) {
	t.Helper()
	ast := assert.New(t)

	t.Setenv(InputPathEnv, test.In)

	dir, err := os.MkdirTemp("", "test_*_dir")
	checkError(err)
	defer os.RemoveAll(dir)

	fileName := fmt.Sprintf("output.%d.txt", time.Now().UnixNano())
	tempFileName := filepath.Join(dir, fileName)
	t.Setenv(OutputPathEnv, tempFileName)

	run()

	inputContent, err := os.ReadFile(test.In)
	checkError(err)
	test.Input = strings.TrimSpace(string(inputContent))
	ast.NotEmpty(test.Input)

	expectedContent, err := os.ReadFile(test.Out)
	if err == nil {
		test.Expected = strings.TrimSpace(string(expectedContent))
	}

	contentOutput, err := os.ReadFile(tempFileName) // #nosec G304
	checkError(err)
	test.Output = strings.TrimSpace(string(contentOutput))

	ast.Equal(test.Expected, test.Output, "Test Case: %v %v", test.In, test.Out)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
