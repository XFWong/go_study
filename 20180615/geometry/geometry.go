package geometry

import (
	"fmt"
	"math"
)

type Point struct{ X, Y float64 }

func Distance(p, q Point) float64 {
	fmt.Println("line11")
	fmt.Println(q.X-p.X, q.Y-p.Y)
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}
func (p Point) Distance(q Point) float64 {
	fmt.Println("line16")
	fmt.Println(q.X-p.X, q.Y-p.Y)
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

type Path []Point

// Distance returns the distance traveled along the path.
func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		}
	}
	return sum
}
