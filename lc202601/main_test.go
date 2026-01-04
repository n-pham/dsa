package main

import (
	"fmt"
)

func ExampleSumFourDivisors() {
	fmt.Println(SumFourDivisors([]int{21, 4, 7}))
	fmt.Println(SumFourDivisors([]int{21, 21}))
	fmt.Println(SumFourDivisors([]int{1, 2, 3, 4, 5}))
	// Unordered output:
	// 32
	// 64
	// 0
}

func ExampleReverseVowels() {
	fmt.Println(ReverseVowels("IceCreAm"))
	fmt.Println(ReverseVowels("leetcode"))
	// Unordered output:
	// AceCreIm
	// leotcede
}
