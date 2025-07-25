package main

import "fmt"

func ExampleNumSubseq() {
	fmt.Println(NumSubseq([]int{3, 3, 6, 8}, 10))
	// Output: 6
}

func ExampleContainsNearbyDuplicate() {
	fmt.Println(ContainsNearbyDuplicate([]int{1, 2, 3, 1, 2, 3}, 2))
	// Output: false
}

func ExampleDivideString() {
	fmt.Println(DivideString("abcdefghi", 3, 'x'))
	fmt.Println(DivideString("abcdefghij", 3, 'x'))
	// Unordered output:
	// [abc def ghi]
	// [abc def ghi jxx]
}

func ExampleHammingWeight() {
	fmt.Println(HammingWeight(11))
	// Output: 3
}

func ExampleMinMaxDifference() {
	fmt.Println(MinMaxDifference(99999))
	// Output: 99999
}
func ExampleAddBinary() {
	fmt.Println(AddBinary("1010", "1011"))
	// Output: 10101
}

func ExampleMaxDifference() {
	fmt.Println(MaxDifference("aaaaabbc"))
	// Output: 3
}

func ExamplePlusOne() {
	fmt.Println(PlusOne([]int{9}))
	// Output: [1 0]
}

func ExampleClearStars() {
	fmt.Println(ClearStars("d*o*"))
	// Output:
}
func ExampleRobotWithString() {
	fmt.Println(RobotWithString("bdda"))
	fmt.Println(RobotWithString("bac"))
	// Unordered output:
	// abc
	// addb
}
