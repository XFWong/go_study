//test sha
package main

import (
	// "crypto/sha256"
	"fmt"
)

func main() {
	// c1 := sha256.Sum256([]byte("wang"))
	// c2 := sha256.Sum256([]byte("wang"))
	// fmt.Printf("c1:%x, %T\n", c1, c1)
	// fmt.Println("c1:", c1, "\n", "c2:", c2)
	// test := [...]byte{1, 2, 2}
	// zero(&test)
	//slice
	a := []int{0, 1, 2, 3, 4, 5}
	reverse(a[2:])
	fmt.Println(a)
	reverse(a[:2])
	fmt.Println(a)
	reverse(a)
	fmt.Println(a)
	//normal array
	b := [...]int{0, 1, 2}
	reverse3(b)
	fmt.Println(b)
}

func zero(ptr *[3]byte) {
	for i := range ptr {
		ptr[i] = 0
	}
	fmt.Println(*ptr)
}

// reverse reverses a slice of ints in place.
func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

// reverse reverses a slice of ints in place.
func reverse3(s [3]int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
