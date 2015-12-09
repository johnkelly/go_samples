package geometry

import "math"

//Point has an x and y coordinate.
type Point struct{ X, Y float64 }

//Distance calculates the distance between two points.
func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

//Distance calculates the distance between two points.
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

//A Path is a journey connecting the points with straight lines.
type Path []Point

//Distance returns the distance traveled along the path.
func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		}
	}
	return sum
}
