package main

import (
	"./geometry"
	"fmt"
	// "math"
)

func main() {
	fmt.Println(sum())
	fmt.Println(sum(3))
	fmt.Println(sum(1, 2, 3, 4))
	values := []int{1, 2, 3, 4, 5}
	fmt.Println(sum(values...))
	fmt.Printf("%T\n", sum)
	p := geometry.Point{1, 2}
	q := geometry.Point{4, 5}
	fmt.Println(geometry.Distance(p, q))
	fmt.Println(p.Distance(q))
}

func sum(values ...int) int {
	total := 0
	for _, val := range values {
		total += val
	}
	return total
}

// type Point struct{ X, Y float64 }

// func Distance(p, q Point) float64 {
// 	return math.Hypot(q.X-p.X, q.Y-p.Y)
// }
// func (p Point) Distance(q Point) float64 {
// 	fmt.Println(p.X, p.Y)
// 	return math.Hypot(q.X-p.X, q.Y-p.Y)
// }
