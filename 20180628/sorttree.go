package main

import (
	"fmt"
)

type tree struct {
	value       int
	left, right *tree
}

func sort(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	// for _, data := range root {
	// 	fmt.Println(data.value)
	// }
	appendValues(values[:0], root)
}

func appendValues(values []int, t *tree) []int {
	fmt.Println(values, t)
	if t != nil {
		values = appendValues(values, t.left)
		fmt.Printf("line:27,t.value:%d,vlues:%v\n", t.value, values)
		values = append(values, t.value)
		fmt.Printf("line:29,t.value:%d,vlues:%v\n", t.value, values)
		values = appendValues(values, t.right)
	}
	fmt.Printf("line:32,vlues:%v\n", values)
	return values
}

func add(t *tree, value int) *tree {
	fmt.Printf("value:%d\n", value)
	fmt.Printf("t:%v,%T\n", t, t)

	if t == nil {
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}
func main() {
	values := []int{5, 4, 4, 3, 2, 1}
	sort(values)
	fmt.Println(values)
}
