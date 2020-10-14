package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHello(t *testing.T) {
	ast := assert.New(t)
	ast.Equal("Hello, world.", hello())

	main()
}

func BenchmarkHello(b *testing.B) {
	for i := 0; i < b.N; i++ {
		hello()
	}
}
