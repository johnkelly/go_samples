package main

import "testing"

func TestSelectionSort(t *testing.T) {
	hundred := prepareTest(100)
	SelectionSort(hundred)

	for i, val := range hundred {
		if i != val {
			t.Fatalf("at index %d, expected %d, got %d.", i, i, val)
		}
	}
}

func BenchmarkSelectionSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SelectionSort(prepareTest(1000))
	}
}
