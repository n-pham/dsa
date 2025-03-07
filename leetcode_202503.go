package main

import (
	"fmt"
	// "github.com/mxschmitt/golang-combinations"
	"math"
	// "math/bits"
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
			nums[newLen] = 2 * prev
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

func All[T any](set []T) (subsets [][]T) {
	length := uint(len(set))

	// Go through all possible combinations of objects
	// from 1 (only first object in subset) to 2^length (all objects in subset)
	for subsetBits := 1; subsetBits < (1 << length); subsetBits++ {
		var subset []T

		for object := uint(0); object < length; object++ {
			// checks if object is contained in subset
			// by checking if bit 'object' is set in subsetBits
			if (subsetBits>>object)&1 == 1 {
				// add object to subset
				subset = append(subset, set[object])
			}
		}
		// add subset to subsets
		subsets = append(subsets, subset)
	}
	return subsets
}

func checkPowersOfThree1780_15ms(n int) bool {
	// 1780
	powers := []int{}
	for i := 0; i < 15; i++ {
		power := int(math.Pow(3, float64(i)))
		if power > n {
			break
		}
		powers = append(powers, power)
		combinations := All(powers)
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

func twoSum1_23ms(nums []int, target int) []int {
	// 1
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

func longestCommonPrefix_14(strs []string) string {
	// 14
	i := 0
	minLen := len(strs[0])
	for _, s := range strs[1:] {
		if len(s) < minLen {
			minLen = len(s)
		}
	}
	for i = 0; i < minLen; i++ {
		c := strs[0][i]
		for j := 1; j < len(strs); j++ {
			if strs[j][i] != c {
				return strs[0][:i]
			}
		}
	}
	return strs[0][:i]
}

func coloredCells2579(n int) int64 {
	// 2579
	if n == 1 {
		return 1
	}
	return int64(1 + 2*n*(n-1))
}

func countSubstrings1638(s string, t string) int {
	// 1638
	panic("not implemented")
}

func findMissingAndRepeatedValues2965(grid [][]int) []int {
	// 2965
	// 1+2+3+4+5+6+7+8+9
	// 1 2 3 4   6 7 8 9 7
	n := len(grid)
	m := make([]byte, n*n+1)
	dup, total := 0, 0
	for _, r := range grid {
		for _, v := range r {
			total += v
			if dup > 0 {
				continue
			}
			if m[v] == 1 {
				dup = v
			} else {
				m[v] = 1
			}
		}
	}
	diff := total - (n * n * (n*n + 1) / 2) // dup - mis
	return []int{dup, dup - diff}
}

func closestPrimes_93ms(left int, right int) []int {
	// 2523
	isPrime := make([]bool, right+1)
	for i := 2; i <= right; i++ {
		isPrime[i] = true
	}
	for i := 2; i*i <= right; i++ {
		if isPrime[i] {
			for j := i * i; j <= right; j += i {
				isPrime[j] = false
			}
		}
	}

	primes := []int{}
	for i := left; i <= right; i++ {
		if isPrime[i] {
			primes = append(primes, i)
		}
	}

	if len(primes) < 2 {
		return []int{-1, -1}
	}

	minDiff := right - left
	result := []int{primes[0], primes[1]}
	for i := 1; i < len(primes); i++ {
		diff := primes[i] - primes[i-1]
		if diff < minDiff {
			minDiff = diff
			result = []int{primes[i-1], primes[i]}
		}
	}

	return result
}

func main() {
	fmt.Println(findMissingAndRepeatedValues2965([][]int{{9, 1, 7}, {8, 7, 2}, {3, 4, 6}}))
	fmt.Println(longestCommonPrefix_14([]string{"flower", "flow", "flight"}))
	// fmt.Println(twoSum1([]int{3, 2, 4}, 6))
	// fmt.Println(checkPowersOfThree1780(91))
	// fmt.Println(applyOperations2460([]int{1,2,2,1,1,0}))
	// fmt.Println(mergeArrays2570([][]int{{1,2},{2,3},{4,5}}, [][]int{{1,4},{3,2},{4,1}}))
}
