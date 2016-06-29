package main

import (
	"fmt"
	"os"
	"strings"
)

// go build .
//./reverse_sentence "the quick brown fox ate delicious sushi"
// sushi delicious ate fox brown quick the

func main() {
	sentence := os.Args[1]

	fmt.Println("Higer Level with split & join")
	fmt.Println(higherLevel(sentence))

	fmt.Println("Lower Level")
	fmt.Println(lowerLevel(sentence))
}

func higherLevel(sentence string) string {
	words := strings.Split(sentence, " ")

	for i := 0; i < len(words)/2; i++ {
		tempWord := words[len(words)-1-i]

		words[len(words)-1-i] = words[i]
		words[i] = tempWord
	}

	return strings.Join(words, " ")
}

func lowerLevel(sentence string) string {
	var words [][]rune
	var currentWord []rune
	codePoints := []rune(sentence)
	for i := 0; i < len(codePoints); i++ {
		codepoint := codePoints[i]
		if codepoint == 32 {
			words = append(words, currentWord)
			currentWord = []rune{}
		} else if i == len(codePoints)-1 {
			currentWord = append(currentWord, codePoints[len(codePoints)-1])
			words = append(words, currentWord)
		} else {
			currentWord = append(currentWord, codepoint)
		}
	}

	for i := 0; i < len(words)/2; i++ {
		tempWord := words[len(words)-i-1]
		words[len(words)-i-1] = words[i]
		words[i] = tempWord
	}

	reverse := ""
	for i := 0; i < len(words); i++ {
		reverse += string(words[i])
		if i != len(words)-1 {
			reverse += " "
		}
	}
	return reverse
}
