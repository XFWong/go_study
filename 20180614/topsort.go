package main

import (
	"fmt"
	"sort"
)

// prereqs记录了每个课程的前置课程
var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},
	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},
	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func main() {
	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}
func topoSort(m map[string][]string) []string {
	var order []string
	fmt.Println("****33****")
	seen := make(map[string]bool)
	var visitAll func(items []string)
	visitAll = func(items []string) {
		fmt.Println("****37****")
		for i, tmp := range items {
			fmt.Printf("i:%d, tmp:%s\n", i, tmp)
		}
		var p *int
		var da int = 1
		p = &da
		fmt.Printf("line41: p:%p, *p:%d\n", p, *p)
		for _, item := range items {
			fmt.Println("line43: " + item)
			if !seen[item] {
				seen[item] = true
				fmt.Println("*****line46******")
				visitAll(m[item])
				order = append(order, item)
				for _, str := range order {
					fmt.Println("line50: " + str)
				}
			}
		}
		fmt.Println(p)
	}
	var keys []string
	for key := range m {
		keys = append(keys, key)
	}
	fmt.Println("line:60 before sort.String")
	sort.Strings(keys)
	fmt.Println("line:62 before visitAll keys")
	for _, key := range keys {
		fmt.Println("line64" + key)
	}
	visitAll(keys)
	fmt.Println("line:67 before return order")
	return order
}
