package main

import (
	"fmt"
	"os"
)

var stack []rune

func main() {
	phrase := os.Args[1]
	fmt.Println(phrase)

	for _, char := range phrase {
		if char == 40 {
			stack = append([]rune{char}, stack...)
		}
		if char == 41 {
			if len(stack) > 0 {
				_, stack = stack[0], stack[1:]
			} else {
				fmt.Println("NOT Balanced.")
				return
			}
		}
	}

	if len(stack) == 0 {
		fmt.Println("Balanced!")
	} else {
		fmt.Println("NOT Balanced.")
	}
}
