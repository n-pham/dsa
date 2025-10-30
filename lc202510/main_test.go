package main

import "fmt"

func ExampleMinNumberOperations() {
	fmt.Println(MinNumberOperations([]int{1, 2, 3, 2, 1}))
	fmt.Println(MinNumberOperations([]int{3, 1, 5, 4, 2}))
	// Unordered output:
	// 3
	// 7
}

func ExampleTotalMoney() {
	fmt.Println(TotalMoney(4))
	fmt.Println(TotalMoney(10))
	fmt.Println(TotalMoney(20))
	// Unordered output:
	// 10
	// 37
	// 96
}

func ExampleMagicalSum() {
	fmt.Println(MagicalSum(5, 5, []int{1, 10, 100, 10000, 1000000}))
	fmt.Println(MagicalSum(2, 2, []int{5, 4, 3, 2, 1}))
	fmt.Println(MagicalSum(2, 2, []int{8}))
	fmt.Println(MagicalSum(4, 2, []int{41}))
	fmt.Println(MagicalSum(4, 3, []int{1}))
	// Unordered output:
	// 991600007
	// 170
	// 0
	// 0
	// 0
}

func ExampleMaximumEnergy() {
	fmt.Println(MaximumEnergy([]int{5, 2, -10, -5, 1}, 3))
	fmt.Println(MaximumEnergy([]int{-2, -3, -1}, 2))
	// Unordered output:
	// 3
	// -1
}

func ExampleMinTime() {
	fmt.Println(MinTime([]int{1, 5, 2, 4}, []int{5, 1, 4, 2}))
	fmt.Println(MinTime([]int{1, 1, 1}, []int{1, 1, 1}))
	fmt.Println(MinTime([]int{1, 2, 3, 4}, []int{1, 2}))
	// Unordered output:
	// 110
	// 5
	// 21
}

func ExampleAvoidFlood() {
	fmt.Println(AvoidFlood([]int{1, 2, 3, 4}))
	fmt.Println(AvoidFlood([]int{1, 2, 0, 0, 2, 1}))
	fmt.Println(AvoidFlood([]int{1, 2, 0, 1, 2}))
	fmt.Println(AvoidFlood([]int{69, 0, 0, 0, 69}))
	// Unordered output:
	// [-1 -1 -1 -1]
	// [-1 -1 2 1 -1 -1]
	// []
	// [-1 69 1 1 -1]
}

func ExampleSwimInWater() {
	fmt.Println(SwimInWater([][]int{{0, 2}, {1, 3}}))
	fmt.Println(SwimInWater([][]int{{0, 1, 2, 3, 4}, {24, 23, 22, 21, 5}, {12, 13, 14, 15, 16}, {11, 17, 18, 19, 20}, {10, 9, 8, 7, 6}}))
	// Unordered output:
	// 3
	// 16
}

func ExamplePacificAtlantic() {
	fmt.Println(PacificAtlantic([][]int{{1, 2, 2, 3, 5}, {3, 2, 3, 4, 4}, {2, 4, 5, 3, 1}, {6, 7, 1, 4, 5}, {5, 1, 1, 2, 4}}))
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
	fmt.Println(MaximumTotalDamage([]int{1, 1, 4, 5, 5, 6}))
	fmt.Println(MaximumTotalDamage([]int{7, 1, 6, 6}))
	fmt.Println(MaximumTotalDamage([]int{2, 3, 5}))
	// Unordered output:
	// 12
	// 13
	// 7
}

func ExampleFindLexSmallestString() {
	fmt.Println(FindLexSmallestString("5525", 9, 2))
	fmt.Println(FindLexSmallestString("74", 5, 1))
	fmt.Println(FindLexSmallestString("0011", 4, 2))
	fmt.Println(FindLexSmallestString("43987654", 7, 3))
	// Unordered output:
	// 2050
	// 24
	// 0011
	// 00553311
}

func ExampleCountValidSelections() {
	fmt.Println(CountValidSelections([]int{0, 1, 0}))
	fmt.Println(CountValidSelections([]int{0, 1, 1, 0}))
	fmt.Println(CountValidSelections([]int{0, 1, 0, 1, 0}))
	// Unordered output:
	// 2
	// 0
	// 2
}

func ExampleSmallestNumber() {
	fmt.Println(SmallestNumber(6))
	fmt.Println(SmallestNumber(8))
	fmt.Println(SmallestNumber(7))
	fmt.Println(SmallestNumber(0))
	fmt.Println(SmallestNumber(1))
	// Unordered output:
	// 7
	// 15
	// 7
	// 0
	// 1
}