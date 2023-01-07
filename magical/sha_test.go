package magical_test

import (
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"testing"
)

func BenchmarkSHA1(b *testing.B) {
	input := []byte("Benchmark tests in Go")

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		sha1.Sum(input)
	}
}

func BenchmarkSHA256(b *testing.B) {
	input := []byte("Benchmark tests in Go")

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		sha256.Sum256(input)
	}
}

func BenchmarkSHA256WithMemoryAllocation(b *testing.B) {
	input := []byte("Benchmark tests in Go")

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		sha256 := sha256.New()
		sha256.Sum(input)
	}
}

func BenchmarkSHA512(b *testing.B) {
	input := []byte("Benchmark tests in Go")

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		sha512.Sum512(input)
	}
}
