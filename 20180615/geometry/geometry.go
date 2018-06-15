package geometry

import (
	"fmt"
	"math"
)

type Point struct{ X, Y float64 }

func Distance(p, q Point) float64 {
	fmt.Println(q.X-p.X, q.Y-p.Y)
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}
func (p Point) Distance(q Point) float64 {
	fmt.Println(q.X-p.X, q.Y-p.Y)
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}
