package main

import (
	"fmt"
)

var pc [256]byte

func init() {
	fmt.Println("enter init")
	for i := range pc {
		fmt.Println(i)
		pc[i] = pc[i/2] + byte(i&1)
		fmt.Println("****begin****")
		fmt.Println(pc[i])
		fmt.Println("****end****")

	}
	fmt.Println("out init")
}
func popcount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}
func main() {
	fmt.Println(popcount(4))
}
