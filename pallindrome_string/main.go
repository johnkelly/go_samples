package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("Has a Pallindrome Permutation: %t\n\n", isPallindrome(os.Args[1]))
}

func isPallindrome(input string) bool {
	charCount := map[string]int{}
	length := 0

	for _, char := range input {
		char := string(char)

		charCount[char] += 1
		length += 1
	}

	if odd(length) {
		oddCount := 0

		for _, v := range charCount {
			if odd(v) {
				oddCount += 1
			}
		}
		return oddCount == 1
	} else {
		for _, v := range charCount {
			if odd(v) {
				return false
			}
		}
		return true
	}
}

func even(x int) bool {
	return x%2 == 0
}

func odd(x int) bool {
	return x%2 != 0
}
