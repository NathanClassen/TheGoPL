package main

import "testing"

func BenchmarkForEcho1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		main()
	}
}