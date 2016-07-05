package main

import (
	"fmt"
	"os"
	"strconv"
)

// input is a sorted array that needs to be stored as a tree
// INPUT:  [1 2 3 4 5 6 7 8 9]
// OUTPUT: [5 3 8 2 4 7 9 1 0 0 0 6 0 0 0]

var tree [100]int
var numbers []int

func main() {
	for _, stringNumber := range os.Args[1:] {
		number, err := strconv.Atoi(stringNumber)
		if err != nil {
			panic("Non integer value inputted")
		}
		numbers = append(numbers, number)
	}
	fmt.Printf("Input:\t%v\n", numbers)

	midIndex := len(numbers) / 2
	midValue := numbers[midIndex]
	tree[0] = midValue
	buildTree(0, 0, 0, midIndex)
	buildTree(0, 1, midIndex+1, len(numbers))
	fmt.Printf("Tree:\t%v\n", tree)
}

func buildTree(parent, position, start, end int) {
	if start == end {
		return
	}

	midIndex := (end + start) / 2
	midValue := numbers[midIndex]
	index := 2*parent + position + 1
	tree[index] = midValue

	buildTree(index, 0, start, midIndex)
	buildTree(index, 1, midIndex+1, end)
}
