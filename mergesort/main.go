package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	startTime := time.Now()
	setSize := 10000000
	list := rand.Perm(setSize)

	sort(list)
	fmt.Println("Time elapsed: ", time.Since(startTime))
	fmt.Println("Set size: ", setSize)
}

func splitArray(source []int) ([]int, []int) {
	return source[0 : len(source)/2], source[len(source)/2:]
}

func sort(source []int) []int {
	if len(source) <= 1 {
		return source
	}

	left, right := splitArray(source)
	left = sort(left)
	right = sort(right)

	return merge(left, right)
}

func merge(left, right []int) []int {
	destination := make([]int, len(left)+len(right))

	j, k := 0, 0

	for i := 0; i < len(destination); i++ {
		if j >= len(left) {
			destination[i] = right[k]
			k++
			continue
		} else if k >= len(right) {
			destination[i] = left[j]
			j++
			continue
		}

		if left[j] > right[k] {
			destination[i] = right[k]
			k++
		} else {
			destination[i] = left[j]
			j++
		}
	}

	return destination
}
