package main

import "fmt"

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
