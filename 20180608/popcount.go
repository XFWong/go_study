package main

import (
	"fmt"
	// "os"
)

var pc [256]byte

func init() {
	// fmt.Println("enter init")
	// s, _ := os.Getwd()
	// fmt.Println(s, &s)
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
	// fmt.Println(x>>(0*8), x>>(1*8), x>>(2*8), x>>(3*8), x>>(4*8), x>>(5*8), x>>(6*8), x>>(7*8))
	// fmt.Println(byte(x>>(0*8)), byte(x>>(1*8)), byte(x>>(2*8)), byte(x>>(3*8)), byte(x>>(4*8)),
	// byte(x>>(5*8)), byte(x>>(6*8)), byte(x>>(7*8)))
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
	// var s string
	// s = "23"
	// fmt.Println(&s, s)
	/*
		x := "hello!"
		fmt.Println(x[0:len(x)])
		fmt.Println(x[:])
		fmt.Println(x[1:])
		fmt.Println(x[:3])
		for i := 0; i < len(x); i++ {
			// fmt.Println(len(x), i, &x)
			x := x[i]
			if x != '!' {
				// fmt.Printf("in if condition x:%p\n", &x)
				x := x + 'A' - 'a'
				// fmt.Printf("reget:%p\n", &x)
				fmt.Printf("%c", x) // "HELLO" (one letter per iteration)
			}
		}
	*/
	// var s *string
	// s = "hello!"
	s := "hello"
	t := &s
	*t += "world"
	// *s += "world!"
	fmt.Println(s, *t)
}
