package main

import (
	"fmt"
	"os"
	"strconv"
)

//Input is a string AaabbbcccCdd
//Out is A1a2b3c3C1d2
func main() {
	input := os.Args[1]
	current := rune(input[0])
	counter := 0
	output := ""

	for _, codepoint := range input {
		if codepoint == current {
			counter++
		} else {
			output += string(current)
			output += strconv.Itoa(counter)

			current = codepoint
			counter = 1
		}
	}
	output += string(current)
	output += strconv.Itoa(counter)

	fmt.Println("Output: " + output)
}
