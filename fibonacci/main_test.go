package main

import "testing"

var tests = []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55}

func TestIterativeFib(t *testing.T) {
	for i, v := range tests {
		if val := IterativeFib(i); val != v {
			t.Fatalf("at index %d, expected %d, got %d.", i, v, val)
		}
	}
}

func TestRecursiveFib(t *testing.T) {
	for i, v := range tests {
		if val := RecursiveFib(i); val != v {
			t.Fatalf("at index %d, expected %d, got %d.", i, v, val)
		}
	}
}

func BenchmarkIterativeFib(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IterativeFib(70)
	}
}

func BenchmarkRecursiveFib(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RecursiveFib(70)
	}
}
