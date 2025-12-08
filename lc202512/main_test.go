package main

import (
	"fmt"
)

func ExampleCountTriples() {
	fmt.Println(CountTriples(5))
	fmt.Println(CountTriples(10))
	fmt.Println(CountTriples(18))
	// Unordered output:
	// 2
	// 4
	// 10
}

func ExampleCountOdds() {
	fmt.Println(CountOdds(3, 7))
	fmt.Println(CountOdds(8, 10))
	// Unordered output:
	// 3
	// 1
}

func ExampleCountPartitions() {
	fmt.Println(CountPartitions([]int{10, 10, 3, 7, 6}))
	fmt.Println(CountPartitions([]int{1, 2, 2}))
	fmt.Println(CountPartitions([]int{2, 4, 6, 8}))
	// Unordered output:
	// 4
	// 0
	// 3
}
