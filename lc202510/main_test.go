package main

import "fmt"

func ExampleNumWaterBottles() {
	fmt.Println(NumWaterBottles(9, 3))
	fmt.Println(NumWaterBottles(15, 4))
	// Unordered output:
	// 13
	// 19
}
