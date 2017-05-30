package main

import (
	"math/rand"
	"testing"
	"time"
)

func TestFindRecursive(t *testing.T) {
	var assertions = []struct {
		input    int
		expected interface{}
	}{
		{100, 100},
		{200, 200},
		{30, 30},
		{3, 3},
		{-100, nil},
		{0, nil},
		{1000, nil},
	}
	input := []int{100, 200, 300, 400, 1, 2, 3, 4, 10, 20, 30, 40}

	tree := &Tree{}
	tree.InsertAll(input)

	for _, assertion := range assertions {
		result := tree.FindRecursive(assertion.input)
		if result != nil && result.Value != assertion.expected {
			t.Errorf("Finding %d - Expected to find node with %d, got %v.", assertion.input, assertion.expected, result)
		} else if result == nil && assertion.expected != nil {
			t.Errorf("Finding %d - Expected to find node with %d, got %v.", assertion.input, assertion.expected, result)
		}
	}

	tree = &Tree{}
	result := tree.FindRecursive(100)
	if result != nil {
		t.Error("Finding 100 in empty tree - Expected nil")
	}
}

func BenchmarkFindRecursive(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	size := 1000000
	tree := &Tree{}

	for i := 0; i <= size-1; i++ {
		tree.Insert(i)
	}

	for i := 0; i < b.N; i++ {
		tree.FindRecursive(rand.Intn(size))
	}
}

func TestFindIterative(t *testing.T) {
	var assertions = []struct {
		input    int
		expected interface{}
	}{
		{100, 100},
		{200, 200},
		{30, 30},
		{3, 3},
		{-100, nil},
		{0, nil},
		{1000, nil},
	}
	input := []int{100, 200, 300, 400, 1, 2, 3, 4, 10, 20, 30, 40}

	tree := &Tree{}
	tree.InsertAll(input)

	for _, assertion := range assertions {
		result := tree.FindIterative(assertion.input)
		if result != nil && result.Value != assertion.expected {
			t.Errorf("Finding %d - Expected to find node with %d, got %v.", assertion.input, assertion.expected, result)
		} else if result == nil && assertion.expected != nil {
			t.Errorf("Finding %d - Expected to find node with %d, got %v.", assertion.input, assertion.expected, result)
		}
	}

	tree = &Tree{}
	result := tree.FindIterative(100)
	if result != nil {
		t.Error("Finding 100 in empty tree - Expected nil")
	}
}

func BenchmarkFindIterative(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	size := 1000000
	tree := &Tree{}

	for i := 0; i <= size-1; i++ {
		tree.Insert(i)
	}

	for i := 0; i < b.N; i++ {
		tree.FindIterative(rand.Intn(size))
	}
}
