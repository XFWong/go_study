package main

import (
	"fmt"
	"time"
)

func main() {
	var pc [256]byte = func() (pc [256]byte) {
		for i := range pc {
			pc[i] = pc[i/2] + byte(i&1)
		}
		return
	}()
	for i := range pc {
		fmt.Println(i)
	}
	go spinner(100 * time.Millisecond)
	fmt.Println("go after spinner")
	const n = 2
	fibN := fib(n) // slow
	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
}
func spinner(delay time.Duration) {
	// fmt.Printf("delay:%d\n", delay)
	for {
		for _, r := range `-\|/` {
			fmt.Printf("%c\n", r)
			time.Sleep(delay)
		}
	}
}
func fib(x int) int {
	fmt.Printf("x:%d\n ", x)
	if x < 2 {
		return x
	}
	result := fib(x-1) + fib(x-2)
	// fmt.Printf(" result:%d\n", result)
	return result
}
