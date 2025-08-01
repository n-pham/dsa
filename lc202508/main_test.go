package main

import "fmt"

func ExampleGenerate() {
	fmt.Println(Generate(1))
	fmt.Println(Generate(2))
	fmt.Println(Generate(3))
	// Unordered output:
	// [[1]]
	// [[1] [1 1]]
	// [[1] [1 1] [1 2 1]]
}
