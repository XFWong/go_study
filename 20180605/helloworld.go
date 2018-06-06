package main

import (
	p "fmt"
	// "os"
	// "runtime"
)

func main() {
	// var goos string = runtime.GOOS
	// fmt.Printf("The operating system is: %s\n", goos)
	// path := os.Getenv("PATH")
	// fmt.Printf("Path is %s\n", path)
	const (
		Str = "the word"
	)

	p.Println("hello 世界" + Str)

	println("hello world")
}
