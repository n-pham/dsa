package main

import (

	"fmt"
	"github.com/mxschmitt/golang-combinations"
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

func mergeArrays2570(nums1 [][]int, nums2 [][]int) [][]int {
	// 2570
	i1, i2 := 0, 0
	rs, rsIdx := [][]int{}, 0
	for i1 < len(nums1) && i2 < len(nums2) {
		idx1, idx2 := nums1[i1][0], nums2[i2][0]
		if idx1 == idx2 {
			rs = append(rs, []int{idx2, nums1[i1][1] + nums2[i2][1]})
			rsIdx++
			i1++
			i2++
		} else if idx1 > idx2 {
			rs = append(rs, []int{idx2, nums2[i2][1]})
			i2++
		} else {
			rs = append(rs, []int{idx1, nums1[i1][1]})
			i1++
		}
	}
	for i1 < len(nums1) {
		rs = append(rs, []int{nums1[i1][0], nums1[i1][1]})
		i1++
	}
	for i2 < len(nums2) {
		rs = append(rs, []int{nums2[i2][0], nums2[i2][1]})
		i2++
	}
	return rs
}

func pivotArray26ms(nums []int, pivot int) []int {
	panic("not implemented")
}

func checkPowersOfThree1780(n int) bool {
	// 1780
	powers := []int{}
	for i := 0; i < 15; i++ {
		power := int(math.Pow(3, float64(i)))
		if power > n {
			break
		}
		powers = append(powers, power)
		combinations := combinations.All(powers)
		for _, combination := range combinations {
			total := 0
			for _, num := range combination {
				total += num
			}
			if total == n {
				return true
			}
		}
	}
	return false
}

func main() {
	fmt.Println(applyOperations2460([]int{1,2,2,1,1,0}))
	fmt.Println(mergeArrays2570([][]int{{1,2},{2,3},{4,5}}, [][]int{{1,4},{3,2},{4,1}}))
}