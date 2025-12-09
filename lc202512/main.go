package main

import (
	"math"
)

//lint:file-ignore U1000 Ignore all unused code, it's generated

func CountSpecialTriplets(nums []int) (cnt int) {
	// 3583
	// nums[i] == nums[j] * 2
	// nums[k] == nums[j] * 2
	const MOD = 1_000_000_007
	n := len(nums)
	prefixCounts := make(map[int]int)
	suffixCounts := make(map[int]int)
	for _, num := range nums {
		suffixCounts[num]++
	}
	for j := 0; j < n; j++ {
		suffixCounts[nums[j]]--
		if suffixCounts[nums[j]] == 0 {
			delete(suffixCounts, nums[j])
		}
		target := nums[j] * 2
		countLeft := prefixCounts[target]
		countRight := suffixCounts[target]
		cnt = (cnt + int((int64(countLeft)*int64(countRight))%MOD))%MOD
		prefixCounts[nums[j]]++
	}
	return
}

func CountTriples(n int) (cnt int) {
	// 1925
	for i := 1; i < n-1; i++ {
		if i*i*2 > n*n {
			break
		}
		for j := i + 1; j < n; j++ {
			sqrInt := i*i + j*j
			if sqrInt > n*n {
				break
			}
			if sqrtInt := int(math.Sqrt(float64(sqrInt))); sqrtInt*sqrtInt == sqrInt {
				cnt += 2
			}
		}
	}
	return
}

func CountOdds(low int, high int) int {
	// 1523
	bothCount := high - low + 1
	if bothCount%2 == 0 {
		return bothCount / 2
	} else {
		return bothCount/2 + (low % 2)
	}
}

func CountPartitionsNaive(nums []int) (cnt int) {
	// 3432
	n := len(nums)
	leftSum := 0
	totalSum := 0
	for _, num := range nums {
		totalSum += num
	}
	for i := 0; i < n-1; i++ {
		leftSum += nums[i]
		if (totalSum-leftSum-leftSum)%2 == 0 {
			cnt++
		}
	}
	return
}

func CountPartitions(nums []int) (cnt int) {
	// 3432
	// A partition is valid if (sum of right part - sum of left part) is even.
	// Let totalSum = sum(nums). For a partition at index i, leftSum = sum(nums[:i+1]).
	// The condition is (totalSum - leftSum - leftSum) % 2 == 0.
	// This simplifies to totalSum % 2 == 0, as 2*leftSum is always even.
	// If totalSum is odd, no partition is valid. Count is 0.
	// If totalSum is even, all n-1 partitions are valid.
	// An array can be partitioned only if it has at least 2 elements.
	totalSum := 0
	for _, num := range nums {
		totalSum += num
	}
	if totalSum%2 != 0 {
		return 0
	}
	return len(nums) - 1
}
