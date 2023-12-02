package main

import "testing"

func BenchmarkPartOne(b *testing.B) {
	for i := 0; i < b.N; i++ {
		partOne(input, &config{r: 12, g: 13, b: 14})
	}
}

func BenchmarkPartTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		partTwo(input)
	}
}
