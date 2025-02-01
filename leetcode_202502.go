package main

import (
	"fmt"
	// "math"
	// "slices"
	// "strconv"
	// "strings"
)

func isArraySpecial3151(nums []int) bool {
	// 3151
	firstValue := nums[0] & 1
	for i, num := range nums[1:] {
		if (i%2 == 0 && (num&1 == firstValue)) || (i%2 == 1 && (num&1 != firstValue)) {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isArraySpecial3151([]int{1, 6, 2}))    // false
	fmt.Println(isArraySpecial3151([]int{2, 1, 4}))    // true
	fmt.Println(isArraySpecial3151([]int{4, 3, 1, 6})) // false
}
