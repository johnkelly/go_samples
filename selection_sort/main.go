package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func prepareTest(size int) []int {
	rand.Seed(time.Now().UnixNano())

	list := make([]int, size)

	for i := 0; i <= size-1; i++ {
		list[i] = i
	}

	shuffle(list)
	return list
}

func main() {
	val, err := strconv.Atoi(os.Args[1])

	if err != nil {
		fmt.Println("You messed up the input!")
	}

	list := prepareTest(val)
	// fmt.Println(list)

	start := time.Now()

	SelectionSort(list)

	// fmt.Println(list)
	fmt.Println(time.Since(start))
}

func shuffle(a []int) {
	for i := range a {
		j := rand.Intn(i + 1)
		a[i], a[j] = a[j], a[i]
	}
}

func SelectionSort(list []int) []int {
	var minimum int = 9999999999
	var minIndex int

	for j := 0; j < len(list); j++ {
		for i := j; i < len(list); i++ {
			if list[i] < minimum {
				minIndex = i
				minimum = list[i]
			}
		}
		list[j], list[minIndex] = list[minIndex], list[j]
		minimum = 9999999999
	}

	return list
}
