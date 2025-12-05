package main

//lint:file-ignore U1000 Ignore all unused code, it's generated

func CountPartitionsNaive(nums []int) (cnt int) {
	// 3432
	n := len(nums)
	leftSum := 0
	totalSum := 0
	for _, num := range nums {
		totalSum+= num
	}
	for i := 0; i < n-1; i++ {
		leftSum += nums[i]
		if (totalSum - leftSum - leftSum) % 2 == 0 {
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