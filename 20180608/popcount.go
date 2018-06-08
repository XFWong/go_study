package main

import (
	"fmt"
)

var pc [256]byte

func init() {
	// fmt.Println("enter init")
	for i, _ := range pc {
		// fmt.Println(i)
		pc[i] = pc[i/2] + byte(i&1)
		// fmt.Println("****begin****")
		// fmt.Println(pc[i])
		// fmt.Println("****end****")

	}
	// fmt.Println("out init", pc)
}
func popcount(x uint64) int {
	fmt.Println(x>>(0*8), x>>(1*8), x>>(2*8), x>>(3*8), x>>(4*8), x>>(5*8), x>>(6*8), x>>(7*8))
	fmt.Println(byte(x>>(0*8)), byte(x>>(1*8)), byte(x>>(2*8)), byte(x>>(3*8)), byte(x>>(4*8)),
		byte(x>>(5*8)), byte(x>>(6*8)), byte(x>>(7*8)))
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

// func main() {
// 	fmt.Println(popcount(3))
// }
func main() {
	x := "hello!"
	for i := 0; i < len(x); i++ {
		fmt.Println(len(x), i, &x)
		x := x[i]
		if x != '!' {
			fmt.Printf("in if condition x:%p\n", &x)
			x := x + 'A' - 'a'
			fmt.Printf("reget:%p\n", &x)
			fmt.Printf("%c\n", x) // "HELLO" (one letter per iteration)
		}
	}
}
