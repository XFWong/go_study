package main

import (
	//"bufio"
	"fmt"
	//"os"
)

func main() {
	// 1.测试map
	// ages := make(map[string]int)
	// ages["wang"] = 31
	// ages["li"] = 32
	// for name := range ages {
	// 	fmt.Println(name)
	// }
	// age, ok := ages["bob"]
	// if !ok {
	// 	fmt.Println(age, ok)
	// }
	//2.
	// seen := make(map[string]bool)
	// input := bufio.NewScanner(os.Stdin)
	// for input.Scan() {
	// 	line := input.Text()
	// 	if !seen[line] {
	// 		seen[line] = true
	// 		fmt.Println(line)
	// 	}
	// }

	// if err := input.Err(); err != nil {
	// 	fmt.Fprintf(os.Stderr, "dedup:%v\n", err)
	// 	os.Exit(1)
	// }

	type address struct {
		hostname string
		port     int
	}
	hits := make(map[address]int)
	hits[address{"golang.org", 443}] = 0
	for i, j := range hits {
		fmt.Println(i, j)
	}
}
