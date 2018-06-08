//Ftoc test f 2 c conversions
package main

import (
	"fmt"
)

func main() {
	const (
		freezingF, boilingF = 32.0, 100.0
	)
	fmt.Printf("%g°F=%g°C\n", freezingF, fToC(freezingF))
	fmt.Printf("%g°F=%g°C\n", boilingF, fToC(boilingF))
	var (
		s string = "12345"
	)
	s, b := "2345", "123"
	fmt.Printf("s:%s, b:%s\n", s, b)
	var (
		x, y int
	)
	fmt.Println(&x == &y, &x == &x, &y == &y, nil)
	fmt.Println(gcd(5, 3))
}

func fToC(f float64) float64 {
	return (f - 32) * 5 / 9
}

func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
		fmt.Println(x, y)
	}

	return x
}
