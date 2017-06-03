package main

import "fmt"

func main() {
	fmt.Printf("Normal: %s\nReverse:%s\n\n", "Basketball", reverseString("Basketball"))
}

func reverseString(s string) string {
	chars := []rune(s)

	for i := 0; i < (len(s) / 2); i++ {
		chars[i], chars[len(s)-1-i] = chars[len(s)-1-i], chars[i]
	}

	return string(chars)
}
