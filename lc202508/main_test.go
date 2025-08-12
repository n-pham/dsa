package main

import "fmt"

func ExampleCountBits() {
	fmt.Println(CountBits(2))
	fmt.Println(CountBits(5))
	// Unordered output:
	// [0 1 1]
	// [0 1 1 2 1 2]
}

func ExampleNumOfUnplacedFruits() {
	fmt.Println(NumOfUnplacedFruits([]int{4, 2, 5}, []int{3, 5, 4}))
	fmt.Println(NumOfUnplacedFruits([]int{3, 6, 1}, []int{6, 4, 7}))
	// Unordered output:
	// 1
	// 0
}

func ExampleGenerate() {
	fmt.Println(Generate(1))
	fmt.Println(Generate(2))
	fmt.Println(Generate(3))
	// Unordered output:
	// [[1]]
	// [[1] [1 1]]
	// [[1] [1 1] [1 2 1]]
}
