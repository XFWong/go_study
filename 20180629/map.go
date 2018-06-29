package main

import (
	"fmt"
)

func main() {
	ages := make(map[string]int)
	ages["wang"] = 31
	ages["li"] = 32
	for name := range ages {
		fmt.Println(name)
	}
	age, ok := ages["bob"]
	if !ok {
		fmt.Println(age, ok)
	}

}
