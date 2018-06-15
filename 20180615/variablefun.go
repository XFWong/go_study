package main

import (
	"fmt"
)

func main() {
	fmt.Println(sum())
	fmt.Println(sum(3))
	fmt.Println(sum(1, 2, 3, 4))
	values := []int{1, 2, 3, 4, 5}
	fmt.Println(sum(values...))
	fmt.Printf("%T\n", sum)
}

func sum(values ...int) int {
	total := 0
	for _, val := range values {
		total += val
	}
	return total
}
