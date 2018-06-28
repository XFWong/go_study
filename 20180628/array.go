package main

import (
	"fmt"
)

func nonempty(strings []string) []string {
	i := 0
	for d, s := range strings {
		fmt.Println(d, i)
		if s != "" {
			strings[i] = s
			i++
		}
	}
	fmt.Printf("%d\n", i)
	return strings[:i]
}

// func nonempty2(strings []string) []string {
// 	out := strings[:0]
// 	for d, s := range strings {
// 		fmt.Println(d, i)
// 		if s != "" {
// 			strings[i] = s
// 			i++
// 		}
// 	}
// 	fmt.Printf("%d\n", i)
// 	return strings[:i]
// }

func main() {
	// 1.测试数组长度
	// var (
	// 	a [30]int = [30]int{1, 2, 3}
	// )
	// fmt.Println(a)
	// fmt.Println(len(a))
	date := []string{"one", "", "three"}
	fmt.Printf("%q\n", nonempty(date))
	fmt.Printf("%q\n", date)
}
