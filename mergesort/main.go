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
	fmt.Println(MergeSort(values))
	fmt.Printf("Duration: %s\n", time.Since(start))
}

func MergeSort(list []int) []int {
	if len(list) == 1 {
		return list
	}

	left := list[:len(list)/2]
	right := list[len(list)/2:]

	return merge(MergeSort(left), MergeSort(right))
}

func merge(left []int, right []int) (result []int) {
	iterations := len(left) + len(right)

	left_index := 0
	right_index := 0

	for i := 0; i < iterations; i++ {
		if left_index >= len(left) {
			result = append(result, right[right_index:]...)
			break
		}

		if right_index >= len(right) {
			result = append(result, left[left_index:]...)
			break
		}

		if left[left_index] > right[right_index] {
			result = append(result, right[right_index])
			right_index++
		} else {
			result = append(result, left[left_index])
			left_index++
		}
	}

	return result
}

//[]int{8, 6, 1, 7, 5, 9, 3, 2, 4}
// BenchmarkMergeSort-8	 1000000	      1282 ns/op
// ok  	github.com/johnkelly/mergesort	1.304s

// 1...1 million
// BenchmarkMergeSort-8	       3	 426463593 ns/op (.42 sec)
// ok  	github.com/johnkelly/mergesort	2.608s
