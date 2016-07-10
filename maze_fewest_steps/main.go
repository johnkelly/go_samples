package main

import "fmt"

// given a maze with a start and end
// find the fewest number of steps to get to the end
// X X X START X
// X O X O     X
// O O O O     X
// O X O X     X
// O O O END   X

type Point struct {
	x, y int
}

func main() {
	adjacency_list := adjacency_list()
	steps := fewest_steps(Point{3, 0}, Point{3, 4}, map[Point]bool{}, adjacency_list)
	fmt.Printf("Fewest steps from start to end: %d", steps)
}

func adjacency_list() map[Point][]Point {
	return map[Point][]Point{
		Point{0, 2}: []Point{Point{1, 2}, Point{0, 3}},
		Point{0, 3}: []Point{Point{0, 2}, Point{0, 4}},
		Point{0, 4}: []Point{Point{0, 3}, Point{1, 4}},
		Point{1, 1}: []Point{Point{1, 2}},
		Point{1, 2}: []Point{Point{2, 2}, Point{0, 2}, Point{1, 1}},
		Point{1, 4}: []Point{Point{0, 4}, Point{2, 4}},
		Point{2, 2}: []Point{Point{3, 2}, Point{2, 3}, Point{1, 2}},
		Point{2, 3}: []Point{Point{2, 4}, Point{2, 2}},
		Point{2, 4}: []Point{Point{1, 4}, Point{2, 3}, Point{3, 4}},
		Point{3, 0}: []Point{Point{3, 1}},
		Point{3, 1}: []Point{Point{3, 0}, Point{3, 2}},
		Point{3, 2}: []Point{Point{3, 1}, Point{2, 2}},
		Point{3, 4}: []Point{Point{2, 4}},
	}
}

// Modified BFS to track steps - Returns 10. Expect 6
func fewest_steps(start Point, end Point, visited map[Point]bool, adj_list map[Point][]Point) int {
	visited[start] = true
	queue := []Point{start}
	steps := 0

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current == end {
			return steps
		}
		steps += 1

		for _, point := range adj_list[current] {
			if !visited[point] {
				visited[point] = true
				queue = append(queue, point)
			}
		}
	}
	return steps
}
