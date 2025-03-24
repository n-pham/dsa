package main

import (
	"fmt"
	// "github.com/mxschmitt/golang-combinations"
	"maps"
	"math"
	"math/bits"
	"slices"
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

func twoSum_23ms(nums []int, target int) []int {
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

func twoSum(nums []int, target int) []int {
	// 1
	iByNum := make(map[int]int, len(nums))
	for i, num := range nums {
		if j, found := iByNum[target-num]; found {
			if i < j {
				return []int{i, j}
			}
			return []int{j, i}
		}
		iByNum[num] = i
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
	// not implemented 1638
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
	// not implemented 877
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
	// not implemented 131
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
	// not implemented 1552
	// 1234  7    3
	// x  x  x
	// 12345   1000000000    2
	// x       x
	panic("not implemented")
}

func pancakeSort(arr []int) []int {
	// not implemented 969
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
	// not implemented 2352
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
	// not implemented 3306
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
				fmt.Println(word[l : r+1])
				cnt++
			}
		}
	}
	return cnt
}

func numberOfSubstrings(s string) int {
	// not implemented 1358
	// abcabc
	// ...     4
	//  ...    3
	//   ...   2
	//    ...  1
	panic("not implemented")
}

func maximumCount(nums []int) int {
	// 2529
	posCnt, negCnt := 0, 0
	for _, num := range nums {
		if num > 0 {
			posCnt++
		} else if num < 0 {
			negCnt++
		}
	}
	if posCnt > negCnt {
		return posCnt
	}
	return negCnt
}

func containsDuplicate_8ms_struct(nums []int) bool {
	// leetcode 217, neetcode
	m := make(map[int]struct{}, len(nums))
	for _, num := range nums {
		if _, found := m[num]; found {
			return true
		}
		m[num] = struct{}{}
	}
	return false
}

func containsDuplicate_10ms_or_more(nums []int) bool {
	// leetcode 217, neetcode,
	m := make(map[int]bool, len(nums))
	for _, num := range nums {
		if m[num] {
			return true
		}
		m[num] = true
	}
	return false
}

func containsDuplicate_8ms_byte(nums []int) bool {
	// leetcode 217, neetcode
	m := make(map[int]byte, len(nums))
	for _, num := range nums {
		if m[num] == 1 {
			return true
		}
		m[num] = 1
	}
	return false
}

func groupAnagrams_9ms(strs []string) [][]string {
	iByM := make(map[[26]int]int)
	rs, rsLen := [][]string{}, -1
	for _, s := range strs {
		m := [26]int{}
		for _, c := range s {
			m[c-'a']++
		}
		if p, found := iByM[m]; found {
			rs[p] = append(rs[p], s)
		} else {
			rsLen++
			rs = append(rs, []string{s})
			iByM[m] = rsLen
		}
	}
	return rs
}

func groupAnagrams(strs []string) [][]string {
	// 49
	// 5ms
	groupsByM := make(map[[26]int][]string)
	for _, s := range strs {
		m := [26]int{}
		for _, c := range s {
			m[c-'a']++
		}
		groupsByM[m] = append(groupsByM[m], s)
	}
	fmt.Println(groupsByM)
	rs := make([][]string, 0, len(groupsByM))
	for _, g := range groupsByM {
		rs = append(rs, g)
	}
	return rs
}

func topKFrequent(nums []int, k int) []int {
	// 347
	frequencyMap := make(map[int]int)
	for _, num := range nums {
		frequencyMap[num]++
	}
	buckets := make([][]int, len(nums)+1)
	for num, freq := range frequencyMap {
		buckets[freq] = append(buckets[freq], num)
	}
	rs := []int{}
	for i := len(buckets) - 1; i >= 0 && len(rs) < k; i-- {
		if len(buckets[i]) > 0 {
			rs = append(rs, buckets[i]...)
		}
	}
	return rs[:k]
}

func isPalindrome(s string) bool {
	// 125
	// A man, a plan, a canal: Panama
	// amanaplanacanalpanama
	l, r := 0, len(s)-1
	for l <= r {
		lc, rc := s[l], s[r]
		fmt.Println(string(lc), string(rc))
		if !(lc >= 'a' && lc <= 'z') && !(lc >= 'A' && lc <= 'Z') && !(lc >= '0' && lc <= '9') {
			l++
			continue
		}
		if !(rc >= 'a' && rc <= 'z') && !(rc >= 'A' && rc <= 'Z') && !(rc >= '0' && rc <= '9') {
			r--
			continue
		}
		if lc >= 'A' && lc <= 'Z' {
			lc += 32
		}
		if rc >= 'A' && rc <= 'Z' {
			rc += 32
		}
		if lc != rc {
			return false
		}
		l++
		r--
	}
	return true
}

func minZeroArray_time(nums []int, queries [][]int) int {
	// 3356
	m := make(map[int]int, len(nums))
	for i, num := range nums {
		if num > 0 {
			m[i] = num
		}
	}
	if len(m) == 0 {
		return 0
	}
	for i, q := range queries {
		for j := q[0]; j <= q[1]; j++ {
			if m[j] == 0 {
				continue
			}
			m[j] -= q[2]
			if m[j] <= 0 {
				delete(m, j)
			}
		}
		if len(m) == 0 {
			return i + 1
		}
	}
	return -1
}

func findMin(nums []int) int {
	// 153 sorted --> binary search
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return nums[left]
}

func maximumCandies_time(candies []int, k int64) int {
	// 2226
	// 5 8 6    3 --> 5
	// 5 8 6    3 --> 5
	// 5 8 6    5 --> 8? 7? 6? 5? 4? 3
	rs := math.MinInt
	for _, num := range candies {
		if num > rs {
			rs = num
		}
	}
	for rs > 0 {
		cnt := int64(0)
		for _, num := range candies {
			cnt += int64(num / rs)
		}
		if cnt >= k {
			return rs
		}
		rs--
	}
	return 0
}

func search(nums []int, target int) int {
	// 704 sorted --> binary search
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	if nums[left] == target {
		return left
	}
	return -1
}

func searchMatrix(matrix [][]int, target int) bool {
	// 74 sorted --> binary search
	left, right := 0, len(matrix)*len(matrix[0])-1
	for left < right {
		mid := left + (right-left)/2
		midR, midC := mid/len(matrix[0]), mid%len(matrix[0])
		if matrix[midR][midC] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	leftR, leftC := left/len(matrix[0]), left%len(matrix[0])
	if matrix[leftR][leftC] == target {
		return true
	}
	return false
}

func minEatingSpeed(piles []int, h int) int {
	// 875 max speed = max pile, min speed = 1 --> binary search
	//    3,6,7,11 27 > 8
	//  2 2 3 4  6
	//  3 1 2 3  4
	//  4 1 2 2  3  8 = 8
	// ...
	// 11 1 1 1  1  4 < 8
	left := 1
	right := math.MinInt
	for _, num := range piles {
		if num > right {
			right = num
		}
	}
	for left < right {
		mid := (left + right) / 2
		hours := 0
		for _, num := range piles {
			hours += (num + mid - 1) / mid // ceiling of num / mid
		}
		if hours > h {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}

func maximumCandies_25ms(candies []int, k int64) int {
	// 2226  max = max num, min = 1 --> binary search
	// rs    1 8 8   4
	// 1     1 8 8
	// 2     0 4 4
	// ...
	// 4     0 2 2
	// 8     0 1 1
	left, right, total := 1, math.MinInt, int64(0)
	for _, num := range candies {
		total += int64(num)
		if num > right {
			right = num
		}
	}
	if total < k {
		return 0
	}
	for left < right {
		mid := (left + right + 1) / 2
		cnt := int64(0)
		for _, num := range candies {
			cnt += int64(num / mid)
		}
		if cnt >= k {
			left = mid
		} else {
			right = mid - 1
		}
	}
	return left
}

func maximumCandies(candies []int, k int64) int {
	// 2226  max = total/k, min = 1 --> binary search
	// rs    1 8 8   4
	// 1     1 8 8
	// 2     0 4 4
	// ...
	// 4     0 2 2
	// 8     0 1 1
	left, total := 1, int64(0)
	for _, num := range candies {
		total += int64(num)
	}
	if total < k {
		return 0
	}
	right := int(total / k)
	for left < right {
		mid := left + 1 + (right-left)/2
		cnt := int64(0)
		for _, num := range candies {
			cnt += int64(num / mid)
		}
		if cnt >= k {
			left = mid
		} else {
			right = mid - 1
		}
	}
	return left
}

func minCapability(nums []int, k int) int {
	// not implemented 2560
	// 2,7,9,3,1    2
	// 1+x to 9+y
	panic("how to find minimum of any k non-consecutive elements?")
}

func threeSum_dup(nums []int) [][]int {
	// 15
	rs := [][]int{}
	for iFirst, first := range nums {
		if first > 0 {
			continue
		}
		numM := make(map[int]struct{}, len(nums))
		for iSecond, second := range nums {
			if iSecond == iFirst {
				continue
			}
			if _, found := numM[-first-second]; found {
				rs = append(rs, []int{first, second, -first - second})
			}
			numM[first+second] = struct{}{}

		}
	}
	return rs
}

func threeSum_solution(nums []int) [][]int {
	// 15
	// -1,0,1,2,-1,-4
	slices.Sort(nums)
	rs := [][]int{}
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left, right := i+1, len(nums)-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == 0 {
				rs = append(rs, []int{nums[i], nums[left], nums[right]})
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			} else if sum < 0 {
				left++
			} else {
				right--
			}
		}
	}
	return rs
}

func maxProfit(prices []int) int {
	// 121
	// 7,1,5,3,6,4
	maxProfit := math.MinInt
	minBuyBefore := prices[0]
	for sellAt := 1; sellAt < len(prices); sellAt++ {
		profit := prices[sellAt] - minBuyBefore
		if profit > maxProfit {
			maxProfit = profit
		}
		if prices[sellAt] < minBuyBefore {
			minBuyBefore = prices[sellAt]
		}
	}
	if maxProfit > 0 {
		return maxProfit
	}
	return 0
}

func characterReplacement(s string, k int) int {
	// not implemented 424
	// A.A.  2 --> window with 2 chars
	cntMax, cntByChar := 0, [26]int{}
	for _, c := range s {
		cntByChar[c-'a']++
		if cntByChar[c-'a'] > cntMax {
			cntMax = cntByChar[c-'a']
		}
	}
	panic("not implemented")
}

func repairCars(ranks []int, cars int) int64 {
	// not implemented 2594
	// 4 2 3 1  cars=10
	// 4*2*2 , 2*2*2 , 3*2*2 , 1*4*4
	panic("not implemented")
}

func divideArray_5ms(nums []int) bool {
	// 2206
	oddM := make(map[int]struct{})
	for _, num := range nums {
		if _, found := oddM[num]; found {
			delete(oddM, num)
		} else {
			oddM[num] = struct{}{}
		}
	}
	return len(oddM) == 0
}

func divideArray(nums []int) bool {
	// 2206
	oddM := make(map[int]byte)
	for _, num := range nums {
		if oddM[num] > 0 {
			delete(oddM, num)
		} else {
			oddM[num] = 1
		}
	}
	return len(oddM) == 0
}

func checkInclusion_fail(s1 string, s2 string) bool {
	// 567
	// abc  lecabee  true
	// abc  lecaabee false
	m1 := make(map[rune]int)
	for _, c := range s1 {
		m1[c]++
	}
	m2 := maps.Clone(m1)
	for _, c := range s2 {
		if len(m2) == 0 {
			return true
		}
		cnt := m2[c]
		if cnt == 0 {
			m2 = maps.Clone(m1)
			continue
		} else if cnt == 1 {
			delete(m2, c)
		} else {
			m2[c] = cnt - 1
		}
	}
	return false
}

func checkInclusion(s1 string, s2 string) bool {
	// 567
	// abc  lecabee  true
	// abc  lecaabee false
	// Sliding window approach
	if len(s1) > len(s2) {
		return false
	}

	s1Count := [26]int{}
	s2Count := [26]int{}

	for i := 0; i < len(s1); i++ {
		s1Count[s1[i]-'a']++
		s2Count[s2[i]-'a']++
	}

	matches := 0
	for i := 0; i < 26; i++ {
		if s1Count[i] == s2Count[i] {
			matches++
		}
	}

	for i := 0; i < len(s2)-len(s1); i++ {
		if matches == 26 {
			return true
		}

		leftChar := s2[i] - 'a'
		rightChar := s2[i+len(s1)] - 'a'

		s2Count[rightChar]++
		if s2Count[rightChar] == s1Count[rightChar] {
			matches++
		} else if s2Count[rightChar] == s1Count[rightChar]+1 {
			matches--
		}

		s2Count[leftChar]--
		if s2Count[leftChar] == s1Count[leftChar] {
			matches++
		} else if s2Count[leftChar] == s1Count[leftChar]-1 {
			matches--
		}
	}

	return matches == 26
}

func longestNiceSubarray(nums []int) int {
	// not implemented 2401
	// 1,3,8,48,10
	//     11
	//   1000
	// 110000
	panic("backtrack?")
}

func isValid_fail(s string) bool {
	// 20
	var pCnt, bCnt, cCnt int
	var prev rune
	for _, c := range s {
		switch c {
		case '(':
			pCnt++
		case '[':
			bCnt++
		case '{':
			cCnt++
		case ')':
			if prev == '[' || prev == '{' {
				return false
			}
			pCnt--
		case ']':
			if prev == '(' || prev == '{' {
				return false
			}
			bCnt--
		case '}':
			if prev == '[' || prev == '(' {
				return false
			}
			cCnt--
		}
		if pCnt < 0 || bCnt < 0 || cCnt < 0 {
			return false
		}
		prev = c
	}
	if pCnt != 0 || bCnt != 0 || cCnt != 0 {
		return false
	}
	return true
}

func isValid(s string) bool {
	// 20
	lenStack, stack := 0, make([]rune, len(s))
	for _, c := range s {
		switch c {
		case '(', '[', '{':
			stack[lenStack] = c
			lenStack++
		case ')':
			if lenStack == 0 || stack[lenStack-1] != '(' {
				return false
			}
			lenStack--
		case ']':
			if lenStack == 0 || stack[lenStack-1] != '[' {
				return false
			}
			lenStack--
		case '}':
			if lenStack == 0 || stack[lenStack-1] != '{' {
				return false
			}
			lenStack--
		}
	}
	if lenStack > 0 {
		return false
	}
	return true
}

func minOperations(nums []int) int {
	// 3191 5ms
	rs := 0
	for i := 0; i < len(nums)-2; i++ {
		if nums[i] == 0 {
			// nums[i] = 1
			nums[i+1] ^= 1
			nums[i+2] ^= 1
			rs++
		}
		fmt.Println(i, nums)
	}
	if nums[len(nums)-2] == 1 && nums[len(nums)-1] == 1 {
		return rs
	}
	return -1
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	// 206
	if head == nil {
		return nil
	}
	current := head
	next := current.Next
	head.Next = nil // Set the original head's Next pointer to nil
	for next != nil {
		next2 := next.Next
		next.Next = current
		current = next
		next = next2
	}
	return current
}

func minimumCost_fail(n int, edges [][]int, query [][]int) []int {
	// 3108
	rs := []int{}
	for _, q := range query {
		fmt.Println(q)
		visited := make(map[int]bool)
		stack := []int{q[0]}
		min := math.MaxInt
		for len(stack) > 0 {
			v := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if !visited[v] {
				visited[v] = true
			}
			for _, e := range edges {
				if e[0] == v && !visited[e[1]] {
					stack = append(stack, e[1])
				}
				if e[0] == v || e[1] == v {
					if e[2] < min {
						min = e[2]
					}
				}
			}
		}
		if min < math.MaxInt {
			rs = append(rs, min)
		} else {
			rs = append(rs, -1)
		}
	}
	return rs
}

func findAllRecipes_fail(recipes []string, ingredients [][]string, supplies []string) []string {
	// 2115
	m := make(map[string]bool)
	mContainsAll := func(ingredients []string) bool {
		for _, i := range ingredients {
			if !m[i] {
				return false
			}
		}
		return true
	}
	for _, s := range supplies {
		m[s] = true
	}
	rs := []string{}
	found := true
	for found {
		found = false
		for i := 0; i < len(recipes); i++ {
			if mContainsAll(ingredients[i]) {
				m[recipes[i]] = true
				rs = append(rs, recipes[i])
				found = true
				recipes = append(recipes[:i], recipes[i+1:]...)
				break
			}
		}
	}
	return rs
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	// 21 needed Copilot
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	head := list2
	thisCurrent, otherCurrent := list2, list1
	if list1.Val < list2.Val {
		head = list1
		thisCurrent, otherCurrent = list1, list2
	}
	for otherCurrent != nil {
		if thisCurrent.Next == nil || thisCurrent.Next.Val > otherCurrent.Val {
			tmp := otherCurrent.Next
			otherCurrent.Next = thisCurrent.Next
			thisCurrent.Next = otherCurrent
			otherCurrent = tmp
		} else {
			thisCurrent = thisCurrent.Next
		}
	}
	// Append the remaining nodes of the other list
	if otherCurrent != nil {
		thisCurrent.Next = otherCurrent
	}
	return head
}

func countCompleteComponents(n int, edges [][]int) int {
	// not implemented 2685
	panic("not implemented")
}

func countPaths(n int, roads [][]int) int {
	// not implemented 1976
	panic("not implemented")
}

func numIslands(grid [][]byte) int {
	// 200
	// recur single data structure
	height, width := len(grid), len(grid[0])
	var recur func(int, int)
	recur = func(i, j int) {
		if i < 0 || j < 0 || i >= height || j >= width || grid[i][j] == '0' {
			return
		}
		grid[i][j] = '0'
		recur(i-1, j)
		recur(i+1, j)
		recur(i, j-1)
		recur(i, j+1)
	}
	cnt := 0
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if grid[i][j] == '1' {
				cnt++
				recur(i, j)
			}
		}
	}
	return cnt
}

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph_3ms(node *Node) *Node {
	m := make(map[*Node]*Node)
	var recur func(*Node) *Node
	recur = func(node *Node) *Node {
		if node == nil {
			return nil
		}
		neighbors := make([]*Node, len(node.Neighbors))
		clone := &Node{Val: node.Val, Neighbors: neighbors}
		m[node] = clone
		for i, x := range node.Neighbors {
			if m[x] == nil {
				neighbors[i] = recur(x)
			} else {
				neighbors[i] = m[x]
			}
		}
		return clone
	}
	return recur(node)
}

func cloneGraph(node *Node) *Node {
	// 133
	// 0ms using queue and map Val to *Node
	if node == nil {
		return nil
	}
	m := make(map[int]*Node)
	m[node.Val] = &Node{Val: node.Val}
	queue := []*Node{node}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		for _, neighbor := range node.Neighbors {
			if clone := m[neighbor.Val]; clone == nil {
				clone = &Node{Val: neighbor.Val}
				m[neighbor.Val] = clone
				queue = append(queue, neighbor)
			}
			m[node.Val].Neighbors = append(m[node.Val].Neighbors, m[neighbor.Val])
		}
	}
	return m[node.Val]
}

func merge(intervals [][]int) [][]int {
	// 56
	slices.SortFunc(intervals, func(a, b []int) int {
		return a[0] - b[0]
	})
	newIntervals := [][]int{}
	prevInterval := intervals[0]
	for _, interval := range intervals[1:] {
		if prevInterval[1] < interval[0] {
			newIntervals = append(newIntervals, prevInterval)
			prevInterval = interval
		} else if prevInterval[1] < interval[1] {
			prevInterval[1] = interval[1]
		}
	}
	newIntervals = append(newIntervals, prevInterval) // Add the last interval
	return newIntervals
}

func main() {
	fmt.Println(merge([][]int{{1, 3}, {8, 10}, {15, 18}, {2, 6}}))
	fmt.Println(merge([][]int{{1, 3}, {8, 15}, {15, 18}, {2, 6}}))
	fmt.Println(merge([][]int{{1, 4}, {1, 4}}))

	// fmt.Println(numIslands([][]byte{{'1','1','0','0','0'},{'1','1','0','0','0'},{'0','0','1','0','0'},{'0','0','0','1','1'}}))
	// fmt.Println(countPaths(7, [][]int{{0,6,7},{0,1,2},{1,2,3},{1,3,3},{6,3,3},{3,5,1},{6,5,1},{2,5,1},{0,4,5},{4,6,2}}))
	// fmt.Println(countCompleteComponents(6, [][]int{{0,1},{0,2},{1,2},{3,4}}))
	// fmt.Println(countCompleteComponents(6, [][]int{{0,1},{0,2},{1,2},{3,4},{3,5}}))
	// fmt.Println(findAllRecipes([]string{"ju","fzjnm","x","e","zpmcz","h","q"}, [][]string{{"cpivl","zpmcz","h","e","fzjnm","ju"}, {"cpivl","hveml","zpmcz","ju","h"}, {"h","fzjnm","e","q","x"}, {"d","hveml","cpivl","q","zpmcz","ju","e","x"}, {"f","hveml","cpivl"}}, []string{"f","hveml","cpivl","d"})) // "ju","fzjnm","q"
	// fmt.Println(findAllRecipes([]string{"bread", "sandwich", "burger"}, [][]string{{"yeast", "flour"}, {"bread", "meat"}, {"sandwich", "meat", "bread"}}, []string{"yeast", "flour", "meat"}))
	// fmt.Println(minOperations([]int{0,1,1,1,0,0}))
	// fmt.Println(isValid("([)]"))
	// fmt.Println(isValid("([])"))
	// fmt.Println(isValid("[([]])"))
	// fmt.Println(checkInclusion("adc", "dcda"))
	// fmt.Println(checkInclusion("abc", "lecabee"))
	// fmt.Println(checkInclusion("abc", "lecaabee"))
	// fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
	// fmt.Println(threeSum([]int{-1,0,1,2,-1,-4}))
	// fmt.Println(maximumCandies([]int{5,6,4,10,10,1,1,2,2,2}, 9))
	// fmt.Println(maximumCandies([]int{2,5}, 11))
	// fmt.Println(maximumCandies([]int{1,8,8}, 4))
	// fmt.Println(minEatingSpeed([]int{3,6,7,11}, 8))
	// fmt.Println(searchMatrix([][]int{{1,3,5,7},{10,11,16,20},{23,30,34,60}}, 3))
	// fmt.Println(search([]int{-1,0,3,5,9,12}, 9))
	// fmt.Println(maximumCandies([]int{1,2,3,4,10}, 5))
	// fmt.Println(maximumCandies([]int{5,8,6}, 3))
	// fmt.Println(maximumCandies([]int{5,8,6}, 4))
	// fmt.Println(maximumCandies([]int{5,8,6}, 5))
	// fmt.Println(minZeroArray([]int{2,0,2}, [][]int{{0,2,1},{0,2,1},{1,1,3}}))
	// fmt.Println(isPalindrome("0P"))
	// fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
	// fmt.Println(groupAnagrams([]string{"eat","tea","tan","ate","nat","bat"}))
	// fmt.Println(countOfSubstrings("aeiou", 0))
	// fmt.Println(countOfSubstrings("ieaouqqieaouqq", 1))
	// fmt.Println(singleNumber([]int{1,2,1,3,2,5}))
	// fmt.Println(countSubstrings("fdsklf"))
	// fmt.Println(countSubstrings("aaa"))
	// fmt.Println(stoneGame([]int{5, 3, 4, 5}))
	// fmt.Println(findMissingAndRepeatedValues2965([][]int{{9, 1, 7}, {8, 7, 2}, {3, 4, 6}}))
	// fmt.Println(longestCommonPrefix_14([]string{"flower", "flow", "flight"}))
	// fmt.Println(twoSum([]int{3, 2, 4}, 6))
	// fmt.Println(twoSum([]int{2,7,11,15}, 9))
	// fmt.Println(checkPowersOfThree1780(91))
	// fmt.Println(applyOperations2460([]int{1,2,2,1,1,0}))
	// fmt.Println(mergeArrays2570([][]int{{1,2},{2,3},{4,5}}, [][]int{{1,4},{3,2},{4,1}}))
}
