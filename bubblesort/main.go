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
	if len(os.Args) < 2 {
		fmt.Println("You must pass an argument!")
		os.Exit(1)
	}

	val, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("You messed up the input!")
		os.Exit(1)
	}

	list := prepareTest(val)
	// fmt.Println(list)

	start := time.Now()

	BubbleSort(list)

	// fmt.Println(list)
	fmt.Println(time.Since(start))
}

func shuffle(a []int) {
	for i := range a {
		j := rand.Intn(i + 1)
		a[i], a[j] = a[j], a[i]
	}
}

func BubbleSort(list []int) []int {
	for k := 0; k < len(list); k++ {
		for i := 0; i < len(list)-1-k; i++ {
			j := i + 1

			if list[i] > list[j] {
				list[i], list[j] = list[j], list[i]
			}
		}
	}

	return list
}
