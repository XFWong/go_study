package main

import (
	"fmt"
)

func squares() func() int {
	var x int
	fmt.Printf("line:9 x:%d\n", x)
	var fun func() int
	fun = func() int {
		fmt.Printf("line:12 x:%d\n", x)
		x++
		return x * x
	}
	return fun
}

func main() {
	fmt.Println("******20******")
	f := squares()
	fmt.Println("******22******")
	fmt.Println(f()) // "1"
	fmt.Println("******24******")
	fmt.Println(f()) // "4"
	fmt.Println("******26******")
	fmt.Println(f()) // "9"
	fmt.Println("******28******")
	fmt.Println(f()) // "16"
	fmt.Println("******30******")

	f = squares()
	fmt.Println("******33******")
	fmt.Println(f()) // "1"
	fmt.Println("******35******")
}
