package main

import "testing"

func TestBubbleSort(t *testing.T) {
	hundred := prepareTest(100)
	BubbleSort(hundred)

	for i, val := range hundred {
		if i != val {
			t.Fatalf("at index %d, expected %d, got %d.", i, i, val)
		}
	}
}

func BenchmarkBubbleSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BubbleSort(prepareTest(1000))
	}
}
