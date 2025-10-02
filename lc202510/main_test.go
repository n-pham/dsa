package main

import "fmt"

func ExampleMaxBottlesDrunk() {
	fmt.Println(MaxBottlesDrunk(13, 6))
	fmt.Println(MaxBottlesDrunk(10, 3))
	// Unordered output:
	// 15
	// 13
}

func ExampleNumWaterBottles() {
	fmt.Println(NumWaterBottles(9, 3))
	fmt.Println(NumWaterBottles(15, 4))
	// Unordered output:
	// 13
	// 19
}
