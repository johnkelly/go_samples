package main

import (
	"fmt"
	"os"
	"strconv"
)

// input is a sorted array that needs to be stored as a tree
// INPUT:  [1 2 3 4 5 6 7 8 9]
// OUTPUT: &{0xc82000e2c0 0xc82000e360 5}

type Node struct {
	left  *Node
	right *Node
	value int
}

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
	root := &Node{
		left:  buildTree(0, midIndex),
		right: buildTree(midIndex+1, len(numbers)),
		value: midValue,
	}
	fmt.Printf("Tree:\t%v\n", root)
	fmt.Printf("Left:\t%v\n", root.left)
	fmt.Printf("Left Left:\t%v\n", root.left.left)
	fmt.Printf("Right:\t%v\n", root.right)
	fmt.Printf("Right Left:\t%v\n", root.right.left)
}

func buildTree(start, end int) *Node {
	if start == end {
		return nil
	}

	midIndex := (end + start) / 2
	midValue := numbers[midIndex]

	return &Node{
		left:  buildTree(start, midIndex),
		right: buildTree(midIndex+1, end),
		value: midValue,
	}
}
