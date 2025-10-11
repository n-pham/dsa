package main

import "fmt"

func ExampleMaximumEnergy() {
	fmt.Println(MaximumEnergy([]int{5,2,-10,-5,1}, 3))
	fmt.Println(MaximumEnergy([]int{-2,-3,-1}, 2))
	// Unordered output:
	// 3
	// -1
}

func ExampleMinTime() {
	fmt.Println(MinTime([]int{1,5,2,4}, []int{5,1,4,2}))
	fmt.Println(MinTime([]int{1,1,1}, []int{1,1,1}))
	fmt.Println(MinTime([]int{1,2,3,4}, []int{1, 2}))
	// Unordered output:
	// 110
	// 5
	// 21
}

func ExampleAvoidFlood() {
	fmt.Println(AvoidFlood([]int{1, 2, 3, 4}))
	fmt.Println(AvoidFlood([]int{1,2,0,0,2,1}))
	fmt.Println(AvoidFlood([]int{1,2,0,1,2}))
	fmt.Println(AvoidFlood([]int{69,0,0,0,69}))
	// Unordered output:
	// [-1 -1 -1 -1]
	// [-1 -1 2 1 -1 -1]
	// []
	// [-1 69 1 1 -1]
}

func ExampleSwimInWater() {
	fmt.Println(SwimInWater([][]int{{0, 2}, {1, 3}}))
	fmt.Println(SwimInWater([][]int{{0,1,2,3,4}, {24,23,22,21,5}, {12,13,14,15,16}, {11,17,18,19,20}, {10,9,8,7,6}}))
	// Unordered output:
	// 3
	// 16
}


func ExamplePacificAtlantic() {
	fmt.Println(PacificAtlantic([][]int{{1,2,2,3,5},{3,2,3,4,4},{2,4,5,3,1},{6,7,1,4,5},{5,1,1,2,4}}))
	fmt.Println(PacificAtlantic([][]int{{1}}))
	// Unordered output:
	// [[0 4] [1 3] [1 4] [2 2] [3 0] [3 1] [4 0]]
	// [[0 0]]
}

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

func ExampleMaximumTotalDamage() {
	fmt.Println(maximumTotalDamage([]int{1, 1, 4, 5, 5, 6}))
	fmt.Println(maximumTotalDamage([]int{7, 1, 6, 6}))
	fmt.Println(maximumTotalDamage([]int{2, 3, 5}))
	// Unordered output:
	// 12
	// 13
	// 7
}
