package parser

import (
	"testing"
)

func BenchmarkBitset(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Parser("ip.txt")
	}
}

var bytes = []byte{56, 46, 51, 52, 46, 53, 46, 50, 51}

func BenchmarkParsing(b *testing.B) {
	for i := 0; i < b.N; i++ {
		parseByte(bytes)
	}
}
