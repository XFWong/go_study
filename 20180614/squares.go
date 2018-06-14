package main

import (
	"fmt"
)

func squares() func() int {
	var x int
	fmt.Printf("line:9 x:%d\n", x)
	return func() int {
		fmt.Printf("line:11 x:%d\n", x)
		x++
		return x * x
	}
}

func main() {
	fmt.Println("******19******")
	f := squares()
	fmt.Println("******20******")
	fmt.Println(f()) // "1"
	fmt.Println("******22******")
	fmt.Println(f()) // "4"
	fmt.Println("******24******")
	fmt.Println(f()) // "9"
	fmt.Println("******26******")
	fmt.Println(f()) // "16"
	fmt.Println("******28******")

	f = squares()
	fmt.Println("******31******")
	fmt.Println(f()) // "1"
	fmt.Println("******33******")
}
