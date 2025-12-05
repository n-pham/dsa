package main

import (
	"fmt"
)


func ExampleCountPartitions() {
	fmt.Println(CountPartitions([]int{10,10,3,7,6}))
	fmt.Println(CountPartitions([]int{1,2,2}))
	fmt.Println(CountPartitions([]int{2,4,6,8}))
	// Unordered output:
	// 4
	// 0
	// 3
}