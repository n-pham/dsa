package main

import "fmt"

func ExampleReplaceNonCoprimes() {
	fmt.Println(ReplaceNonCoprimes([]int{6, 4, 3, 2, 7, 6, 2}))
	fmt.Println(ReplaceNonCoprimes([]int{2, 2, 1, 1, 3, 3, 3}))
	// Unordered output:
	// [12 7 6]
	// [2 1 1 3]
}

func ExampleCanBeTypedWords() {
	fmt.Println(CanBeTypedWords("hello world", "ad"))
	fmt.Println(CanBeTypedWords("leet code", "lt"))
	fmt.Println(CanBeTypedWords("leet code", "e"))
	// Unordered output:
	// 1
	// 1
	// 0
}

func ExampleMaxFreqSum() {
	fmt.Println(MaxFreqSum("successes"))
	fmt.Println(MaxFreqSum("aeiaeia"))
	// Unordered output:
	// 6
	// 3
}

func ExampleSumZero() {
	fmt.Println(SumZero(5))
	fmt.Println(SumZero(4))
	// Unordered output:
	// [0 1 2 -1 -2]
	// [-2 1 2 -1]
}

func ExampleIntersection() {
	result1 := Intersection([]int{1, 2, 2, 1}, []int{2, 2})
	for _, value := range result1 {
		fmt.Println(value)
	}
	result2 := Intersection([]int{4, 9, 5}, []int{9, 4, 9, 8, 4})
	for _, value := range result2 {
		fmt.Println(value)
	}
	// Unordered output:
	// 2
	// 9
	// 4
}
