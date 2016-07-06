package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

// ./router "search/:penguin" "homes" "homes/:id" "homes/:id/room" "homes/:id/room/:id" ":id"
// RANK MAP: %v map[search/:penguin:2 homes:0 homes/:id:2 homes/:id/room:2 homes/:id/room/:id:10 :id:1]

var routes []string
var rank map[string]int
var sortedRoutes map[int]string

func main() {
	rank = make(map[string]int)
	for _, route := range os.Args[1:] {
		score := 0
		levels := strings.Split(route, "/")
		for idx, level := range levels {
			first_rune := level[0]
			if first_rune == 58 {
				fmt.Printf("LEVEL: %d", idx)
				score += (1 * int(math.Pow(2, float64(idx))))
			} else {
				score += (0 * int(math.Pow(2, float64(idx))))
			}
		}
		fmt.Println(levels)
		rank[route] = score
	}

	fmt.Println("RANK MAP: %v", rank)

	//Then sort list
}
