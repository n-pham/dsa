package main

import (

	"fmt"
	// "math"
	// "slices"
	// "strconv"
	// "strings"
)
func applyOperations2460(nums []int) []int {
	// 2460
	// 1,2,2,1,1,0
	// 1 4 0 1 1 0
	// 1 4 0 2 0 0
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			nums[i] *= 2
			nums[i+1] = 0
		}
	}
	result := []int{}
	for _, num := range nums {
		if num != 0 {
			result = append(result, num)
		}
	}
	for len(result) < len(nums) {
		result = append(result, 0)
	}
	return result
}

func applyOperations2460_fail(nums []int) []int {
	// 2460
	// 1,2,2,1,1,0
	// 1 4 0 1 1 0
	// 1 4 0 2 0 0
	prev, newLen := nums[0], 0
	for i := 1; i < len(nums); {
		num := nums[i]
		fmt.Println(prev, num)
		if num == prev {
			nums[newLen] = 2*prev
			i += 2
		} else if prev > 0 {
			nums[newLen] = prev
			i++
		}
		newLen++
		prev = num
	}
	return nums[:newLen]
}

func main() {
	fmt.Println(applyOperations2460([]int{1,2,2,1,1,0}))
}