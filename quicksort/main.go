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
	fmt.Println(InPlaceQuickSort(values))
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

var elements []int

func InPlaceQuickSort(list []int) []int {
	elements = list
	low := 0
	high := len(list) - 1

	InPlaceHelper(low, high)
	return list
}

func InPlaceHelper(low int, high int) {
	if high-low <= 1 {
		if elements[low] > elements[high] {
			elements[low], elements[high] = elements[high], elements[low]
		} else if high <= low {
		} else {
			pivotIndex := low
			partitionIndex := partition(low, high, pivotIndex)
			InPlaceHelper(low, partitionIndex-1)
			InPlaceHelper(partitionIndex+1, high)
		}
	}
}

func partition(low, high, pivotIndex int) int {
	i := low

	for j := low; j <= high; j++ {
		if elements[j] < elements[pivotIndex] {
			i++
			elements[i], elements[j] = elements[j], elements[i]
		}
	}

	elements[i], elements[pivotIndex] = elements[pivotIndex], elements[i]
	return i
}

// A million
// BenchmarkQuickSort-8	       2	 621360143 ns/op (.62sec)
// ok  	github.com/johnkelly/quicksort	1.978s
