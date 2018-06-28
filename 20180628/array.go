package main

import (
	"fmt"
)

func main() {
	// 1.测试数组长度
	var (
		a [30]int = [30]int{1, 2, 3}
	)
	fmt.Println(a)
	fmt.Println(len(a))

}
