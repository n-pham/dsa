package main

import "fmt"

func ExampleTrapRainWater() {
	fmt.Println(TrapRainWater([][]int{{1, 4, 3, 1, 3, 2}, {3, 2, 1, 3, 2, 4}, {2, 3, 3, 2, 3, 1}}))
	fmt.Println(TrapRainWater([][]int{{3, 3, 3, 3, 3}, {3, 2, 2, 2, 3}, {3, 2, 1, 2, 3}, {3, 2, 2, 2, 3}, {3, 3, 3, 3, 3}}))
	// Unordered output:
	// 4
	// 10
}

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
