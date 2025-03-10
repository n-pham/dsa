package main

import (
	"fmt"
	// "github.com/mxschmitt/golang-combinations"
	"math"
	"math/bits"
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
	// "aba", "baba"
	panic("not implemented - Trie")
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

func stoneGame(piles []int) bool {
	// 877
	// 5,3,4,5
	// DP
	panic("not implemented")
}

func shipWithinDays_solution(weights []int, days int) int {
	// 1011
	left, right := 0, 0
	for _, weight := range weights {
		if weight > left {
			left = weight
		}
		right += weight
	}

	for left < right {
		mid := (left + right) / 2
		currentWeight, requiredDays := 0, 1
		for _, weight := range weights {
			if currentWeight+weight > mid {
				requiredDays++
				currentWeight = 0
			}
			currentWeight += weight
		}
		if requiredDays > days {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return left
}

func partition(s string) [][]string {
	// 131
	// aab --> a a b    aa b
	panic("not implemented")
}

func Combinations[T any](set []T, n int) (subsets [][]T) {
	length := uint(len(set))

	if n > len(set) {
		n = len(set)
	}

	// Go through all possible combinations of objects
	// from 1 (only first object in subset) to 2^length (all objects in subset)
	for subsetBits := 1; subsetBits < (1 << length); subsetBits++ {
		if n > 0 && bits.OnesCount(uint(subsetBits)) != n {
			continue
		}

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

func combinationSum3(k int, n int) [][]int {
	// 216
	// 3,9 --> 1,2,6 1,3,5 2,3,4
	rs := [][]int{}
	combinations := Combinations([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, k)
	for _, combination := range combinations {
		total := 0
		for _, num := range combination {
			total += num
		}
		if total == n {
			rs = append(rs, combination)
		}
	}
	return rs
}

func minimumRecolors(blocks string, k int) int {
	// 2379
	// 0123456789
	// .BB..BB.B.  7
	// >     4
	//  >     4
	//   >     4
	//    >     3
	countB := 0
	for i := 0; i < k; i++ {
		if blocks[i] == 'B' {
			countB++
		}
	}
	maxB := countB
	for i := k; i < len(blocks); i++ {
		if blocks[i-k] == 'B' {
			countB--
		}
		if blocks[i] == 'B' {
			countB++
		}
		if countB > maxB {
			maxB = countB
		}
	}
	return k - maxB
}

func countSubstrings_580ms(s string) int {
	// 647
	combinations := [][]byte{}
	for i := 0; i < len(s); i++ {
		for j := i + 1; j <= len(s); j++ {
			combinations = append(combinations, []byte(s[i:j]))
		}
	}
	fmt.Println(combinations)
	cnt := len(combinations)
	for _, combination := range combinations {
		for i := 0; i < len(combination)/2; i++ {
			if combination[i] != combination[len(combination)-1-i] {
				cnt--
				break
			}
		}
	}
	return cnt
}

func maxDistance(position []int, m int) int {
	// 1552
	// 1234  7    3
	// x  x  x
	// 12345   1000000000    2
	// x       x
	panic("not implemented")
}

func pancakeSort(arr []int) []int {
	// 969
	panic("not understood")
}

func singleNumber(nums []int) []int {
	// 260
	// 1,2,1,3,2,5
	m := make(map[int]byte)
	for _, num := range nums {
		m[num] ^= 1
	}
	rs := []int{}
	for num, val := range m {
		if val == 1 {
			rs = append(rs, num)
		}
	}
	return rs
}

func singleNumber_solution(nums []int) []int {
	// 260
	// 1,2,1,3,2,5
	xor := 0
	for _, num := range nums {
		xor ^= num
	}

	// Get the rightmost set bit
	rightmostSetBit := xor & -xor

	num1, num2 := 0, 0
	for _, num := range nums {
		if num&rightmostSetBit == 0 {
			num1 ^= num
		} else {
			num2 ^= num
		}
	}

	return []int{num1, num2}
}

func equalPairs(grid [][]int) int {
	// 2352
	// 3,1,2,2
	// 1,4,4,5
	// 2,4,2,2
	// 2,4,2,2
	panic("not implemented")
}

func numberOfAlternatingGroups(colors []int, k int) int {
	// 3208
	// 0,1,0,0,1,0,1    6
	// 0 1 0   1 0 1
	// 0 1   0 1 0 1
	// 0,1,0,1,0        3
	// 0 1 0
	//   1 0 1
	//     0 1 0
	count := 0
	for i := 0; i < len(colors)-k+1; i++ {
		valid := true
		for j := 1; j < k; j++ {
			if colors[i+j] == colors[i+j-1] {
				valid = false
				break
			}
		}
		if valid {
			count++
		}
	}
	return count
}

func countOfSubstrings(word string, k int) int64 {
	// 3306
	panic("not implemented")
}

func countOfSubstrings_time(word string, k int) int64 {
	// 3306
	// ieaouqqieaouqq    1
	// ieqouq
	//       qieaou
	//        ieqouq
	cnt := int64(0)
	for l := 0; l <= len(word)-5-k; l++ {
		missingVowels := map[byte]struct{}{'a': struct{}{}, 'e': struct{}{}, 'i': struct{}{}, 'o': struct{}{}, 'u': struct{}{}}
		consonantCnt := k
		for r := l; r < len(word); r++ {
			if word[r] == 'a' || word[r] == 'e' || word[r] == 'i' || word[r] == 'o' || word[r] == 'u' {
				delete(missingVowels, word[r])
			} else {
				consonantCnt--
			}
			if len(missingVowels) == 0 && consonantCnt == 0 {
				fmt.Println(word[l:r+1])
				cnt++
			}
		}
	}
	return cnt
}

func main() {
	fmt.Println(countOfSubstrings("aeiou", 0))
	fmt.Println(countOfSubstrings("ieaouqqieaouqq", 1))
	// fmt.Println(singleNumber([]int{1,2,1,3,2,5}))
	// fmt.Println(countSubstrings("fdsklf"))
	// fmt.Println(countSubstrings("aaa"))
	// fmt.Println(stoneGame([]int{5, 3, 4, 5}))
	// fmt.Println(findMissingAndRepeatedValues2965([][]int{{9, 1, 7}, {8, 7, 2}, {3, 4, 6}}))
	// fmt.Println(longestCommonPrefix_14([]string{"flower", "flow", "flight"}))
	// fmt.Println(twoSum1([]int{3, 2, 4}, 6))
	// fmt.Println(checkPowersOfThree1780(91))
	// fmt.Println(applyOperations2460([]int{1,2,2,1,1,0}))
	// fmt.Println(mergeArrays2570([][]int{{1,2},{2,3},{4,5}}, [][]int{{1,4},{3,2},{4,1}}))
}
