package main

import (
	"math/rand"
	"testing"
	"time"
)

func TestQuickSort(t *testing.T) {
	values := []int{8, 6, 1, 7, 5, 9, 3, 2, 4}
	sortedValues := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	values = QuickSort(values)

	for i, val := range values {
		if (i + 1) != val {
			v := sortedValues[i]
			t.Fatalf("at index %d, expected %d, got %d.", i, v, val)
		}
	}
}

func BenchmarkQuickSort(b *testing.B) {
	rand.Seed(time.Now().UnixNano())

	size := 1000000
	million := make([]int, size)

	for i := 0; i <= size-1; i++ {
		million[i] = i
	}

	shuffle(million)

	for i := 0; i < b.N; i++ {
		QuickSort(million)
	}
}

func TestInPlaceQuickSort(t *testing.T) {
	values := []int{8, 6, 1, 7, 5, 9, 3, 2, 4}
	sortedValues := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	values = InPlaceQuickSort(values)

	for i, val := range values {
		if (i + 1) != val {
			v := sortedValues[i]
			t.Fatalf("at index %d, expected %d, got %d.", i, v, val)
		}
	}
}

func BenchmarkInPlaceQuickSort(b *testing.B) {
	rand.Seed(time.Now().UnixNano())

	size := 1000000
	million := make([]int, size)

	for i := 0; i <= size-1; i++ {
		million[i] = i
	}

	shuffle(million)

	for i := 0; i < b.N; i++ {
		InPlaceQuickSort(million)
	}
}
