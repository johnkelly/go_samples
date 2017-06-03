package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	input := []string{"I", "am", "happy", "so", "I", "sing", "happy", "happy", "songs"}
	fmt.Printf("Result: %v", uniqueSorted(input))
}

func uniqueSorted(input []string) []string {
	lookup := map[string]int{}
	result := [][]string{}

	for _, element := range input {
		lookup[element] += 1
	}

	for key, value := range lookup {
		stringValue := strconv.Itoa(value)
		result = append(result, []string{stringValue, key})
	}

	sort.Sort(ByFrequency(result))

	sortedWords := []string{}
	for _, element := range result {
		sortedWords = append(sortedWords, element[1])
	}

	return sortedWords
}

type ByFrequency [][]string

func (x ByFrequency) Len() int { return len(x) }
func (x ByFrequency) Swap(i, j int) {
	x[i], x[j] = x[j], x[i]
}
func (x ByFrequency) Less(i, j int) bool {
	return x[i][0] > x[j][0]
}
