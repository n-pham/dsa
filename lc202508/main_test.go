package main

import "fmt"

func ExampleReverseString() {
	var s []byte
	s = []byte{'h', 'e', 'l', 'l', 'o'}
	ReverseString(s)
	fmt.Println(s)
	s = []byte{'H', 'a', 'n', 'n', 'a', 'h'}
	ReverseString(s)
	fmt.Println(s)
	// Unordered output:
	// [111 108 108 101 104]
	// [104 97 110 110 97 72]
}

func ExampleLargestGoodInteger() {
	fmt.Println(LargestGoodInteger("6777133339"))
	fmt.Println(LargestGoodInteger("2300019"))
	// Unordered output:
	// 777
	// 000
}
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
