package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func shuffle(a []int) {
	for i := range a {
		j := rand.Intn(i + 1)
		a[i], a[j] = a[j], a[i]
	}
}

func main() {
	fmt.Println("What size do you want?")
	size, err := strconv.Atoi(os.Args[1])

	if err != nil {
		fmt.Errorf("Invalid size to sort: %v!", size)
	}

	values := make([]int, size)

	for i := 0; i <= size-1; i++ {
		values[i] = i
	}

	rand.Seed(time.Now().UnixNano())
	shuffle(values)

	fmt.Println(values)
	start := time.Now()
	fmt.Println(QuickSort(values))
	fmt.Printf("Duration: %s\n", time.Since(start))
}

func QuickSort(list []int) []int {
	if len(list) <= 1 {
		return list
	}

	pivot := rand.Intn(len(list))
	pivotValue := list[pivot]

	left := make([]int, 0)
	right := make([]int, 0)

	for _, v := range list {
		if v < pivotValue {
			left = append(left, v)
		} else if v != pivotValue {
			right = append(right, v)
		}
	}

	return merge(QuickSort(left), pivotValue, QuickSort(right))
}

func merge(left []int, pivot int, right []int) (result []int) {
	result = append(result, left...)
	result = append(result, pivot)
	result = append(result, right...)
	return result
}

// A million
// BenchmarkQuickSort-8	       2	 621360143 ns/op (.62sec)
// ok  	github.com/johnkelly/quicksort	1.978s
