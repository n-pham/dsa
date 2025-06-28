package main

import (
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
)

//lint:file-ignore U1000 Ignore all unused code, it's generated

func AnswerString(word string, numFriends int) string {
	// 3403 1208ms
	stringLen := len(word) - numFriends + 1
	maxSubstr := ""
	for length := 1; length <= stringLen; length++ {
		for start := 0; start <= len(word)-length; start++ {
			curr := word[start : start+length]
			if curr > maxSubstr {
				maxSubstr = curr
			}
		}
	}
	return maxSubstr
}

func RobotWithString_fail(s string) string {
	// s: bdda
	// q: abcd
	//    1102
	// stack if not a --> q: bdd s: b
	// stack if not b --> q: bdd s: _
	// write stack (reverse order)
	rs, rsLen := make([]byte, len(s)), 0
	q, qLen, countsByChar := make([]byte, len(s)), 0, [26]int{}
	for _, c := range s {
		countsByChar[c-'a']++
	}
	smallestCharIndex := byte(0)
	for i := 0; i < len(s); i++ {
		c := s[i]
		for ; smallestCharIndex < 26 && countsByChar[smallestCharIndex] == 0; smallestCharIndex++ {
		}
		if smallestCharIndex == 26 {
			break
		}
		if c-'a' == smallestCharIndex {
			rs[rsLen] = c
			rsLen++
			countsByChar[c-'a']--
		} else {
			q[qLen] = c
			qLen++
			countsByChar[c-'a']--
		}
	}
	q = q[:qLen]
	fmt.Println(q)
	for i := qLen - 1; i > -1; i-- {
		rs[rsLen] = q[i]
		rsLen++
	}
	rs = rs[:rsLen]
	fmt.Println(rs)
	return string(rs)
}

func RobotWithString(s string) string {
	// 2434
	rs := make([]byte, 0, len(s))
	stack := make([]byte, 0, len(s))
	countsByChar := [26]int{}
	for i := range s {
		countsByChar[s[i]-'a']++
	}
	minChar := byte('a')
	for i := range s {
		// Update minChar if current character count becomes 0
		countsByChar[s[i]-'a']--
		// Push current character to stack
		stack = append(stack, s[i])
		// Find next smallest remaining character
		for minChar <= 'z' && countsByChar[minChar-'a'] == 0 {
			minChar++
		}
		// Pop from stack while we can get smaller characters
		for len(stack) > 0 && stack[len(stack)-1] <= minChar {
			rs = append(rs, stack[len(stack)-1])
			stack = stack[:len(stack)-1]
		}
	}
	return string(rs)
}

func ClearStars_fail(s string) string {
	rs := make([]byte, 0, len(s))
	countsByChar := [26]int{}
	for i := range s {
		c := s[i]
		if c != '*' {
			countsByChar[c-'a']++
		} else { // remove smallest char
			smallestCharIndex := 0
			for ; smallestCharIndex < 26 && countsByChar[smallestCharIndex] == 0; smallestCharIndex++ {
			}
			if smallestCharIndex == 26 {
				continue
			}
			countsByChar[smallestCharIndex]--
		}
	}
	for i := range s {
		c := s[i]
		if c != '*' && countsByChar[c-'a'] > 0 {
			rs = append(rs, c)
			countsByChar[c-'a']--
		}
	}
	return string(rs)
}

func ClearStars(s string) string {
	// 3170
	// aaba* --> aab
	indicesByChar := [26][]int{}
	removes := make([]bool, len(s))
	for i := range s {
		c := s[i]
		if c == '*' {
			removes[i] = true // remove *
			for j, indices := range indicesByChar {
				if len(indices) > 0 {
					removes[indices[len(indices)-1]] = true
					indicesByChar[j] = indices[:len(indices)-1]
					break
				}
			}
		} else {
			indicesByChar[c-'a'] = append(indicesByChar[c-'a'], i)
		}
	}
	rs := make([]byte, 0, len(s))
	for i := range s {
		if !removes[i] {
			rs = append(rs, s[i])
		}
	}
	return string(rs)
}

func LexicalOrder(n int) (nums []int) {
	// 386
	var recur func(curr int)
	recur = func(curr int) {
		if curr > n {
			return
		}
		nums = append(nums, curr)
		for i := 0; i <= 9; i++ {
			recur(curr*10 + i)
		}
	}
	for i := 1; i <= 9; i++ {
		recur(i)
	}
	return nums
}

func FindKthNumber_time(n int, k int) int {
	var recur func(curr int) int
	recur = func(curr int) int {
		if curr > n {
			return 0
		}
		k--
		if k == 0 {
			return curr
		}
		for i := 0; i <= 9; i++ {
			if num := recur(curr*10 + i); num > 0 {
				return num
			}
		}
		return 0
	}
	for i := 1; i <= 9; i++ {
		if num := recur(i); num > 0 {
			return num
		}
	}
	return 0
}

func FindKthNumber(n int, k int) int {
	// 440
	// Helper function to calculate the number of steps between two prefixes
	steps := func(prefix, n int) int {
		curr, next := prefix, prefix+1
		count := 0
		for curr <= n {
			count += min(next, n+1) - curr
			curr *= 10
			next *= 10
		}
		return count
	}
	curr := 1
	k--
	for k > 0 {
		count := steps(curr, n)
		if k >= count {
			curr++
			k -= count
		} else {
			curr *= 10
			k--
		}
	}
	return curr
}

func SearchInsert(nums []int, target int) int {
	// 35
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}

func PlusOne(digits []int) []int {
	// 66
	var val int
	for i := len(digits) - 1; i > -1; i-- {
		val = digits[i]
		if val < 9 {
			digits[i] = val + 1
			return digits
		} else {
			digits[i] = 0
		}
	}
	if val == 9 {
		return append([]int{1}, digits...)
	}
	return digits
}

func MaxDifference(s string) int {
	// 3442
	countsByChar := [26]int{}
	for i := range s {
		countsByChar[s[i]-'a']++
	}
	maxOdd, minEven := math.MinInt, math.MaxInt
	for _, cnt := range countsByChar {
		if cnt > 0 && cnt%2 == 0 {
			if minEven > cnt {
				minEven = cnt
			}
		} else {
			if maxOdd < cnt {
				maxOdd = cnt
			}
		}
	}
	return maxOdd - minEven
}

func MaxDifference3445(s string, k int) int {
	// LTE
	// s consists only of digits '0' to '4'
	n := len(s)
	prefixFreq := make([][5]int, n+1)
	for i := 0; i < n; i++ {
		for d := 0; d < 5; d++ {
			prefixFreq[i+1][d] = prefixFreq[i][d]
		}
		prefixFreq[i+1][int(s[i]-'0')]++
	}

	maxDiff := math.MinInt
	for l := 0; l < n; l++ {
		for r := l + k - 1; r < n; r++ {
			freq := [5]int{}
			for d := 0; d < 5; d++ {
				freq[d] = prefixFreq[r+1][d] - prefixFreq[l][d]
			}
			for a := 0; a < 5; a++ {
				if freq[a] > 0 && freq[a]%2 == 1 {
					for b := 0; b < 5; b++ {
						if freq[b] > 0 && freq[b]%2 == 0 {
							diff := freq[a] - freq[b]
							if diff > maxDiff {
								maxDiff = diff
							}
						}
					}
				}
			}
		}
	}
	return maxDiff
}

func AddBinary(a string, b string) string {
	// 67
	aLen, bLen := len(a), len(b)
	minLen, maxLen := len(a), len(b)
	if minLen > bLen {
		minLen = bLen
		maxLen = aLen
	}
	rs := make([]byte, maxLen)
	var prev byte = 0
	for i := 1; i <= minLen; i++ {
		aDigit, bDigit := a[aLen-i]-'0', b[bLen-i]-'0'
		total := aDigit + bDigit + prev
		rs[maxLen-i] = '0' + (total % 2)
		prev = total / 2
	}
	for i := maxLen - minLen - 1; i >= 0; i-- {
		var digit byte
		if aLen > bLen {
			digit = a[i] - '0'
		} else {
			digit = b[i] - '0'
		}
		total := digit + prev
		rs[i] = '0' + (total % 2)
		prev = total / 2
	}
	if prev > 0 {
		rs = append([]byte{'1'}, rs...)
	}
	return string(rs)
}

func MaxAdjacentDistance(nums []int) int {
	// 3423
	maxDiff := math.MinInt
	prev := nums[len(nums)-1]
	for _, num := range nums {
		diff := num - prev
		if diff < 0 {
			diff = -diff
		}
		if maxDiff < diff {
			maxDiff = diff
		}
		prev = num
	}
	return maxDiff
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func DeleteDuplicates(head *ListNode) *ListNode {
	// 83
	prev, curr := head, head
	for ; curr != nil; curr = curr.Next {
		if curr.Val == prev.Val {
			prev.Next = curr.Next
		} else {
			prev = curr
		}
	}
	return head
}

func SingleNumber(nums []int) (num int) {
	// 136
	for _, n := range nums {
		num ^= n
	}
	return num
}

func ConvertToTitle(columnNumber int) string {
	// 168
	panic("not implemented")
}

func MinimizeMax(nums []int, p int) int {
	// 2616
	slices.Sort(nums)
	n := len(nums)

	// Helper function to check if we can form `p` pairs with max difference `maxDiff`
	canFormPairs := func(maxDiff int) bool {
		dp := make([]int, n+1)
		for i := 2; i <= n; i++ {
			dp[i] = dp[i-1]
			if nums[i-1]-nums[i-2] <= maxDiff {
				dp[i] = max(dp[i], dp[i-2]+1)
			}
		}
		return dp[n] >= p
	}

	// Binary search for the minimum possible max difference
	left, right := 0, nums[n-1]-nums[0]
	for left < right {
		mid := left + (right-left)/2
		if canFormPairs(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func MinMaxDifference(num int) int {
	// 2566
	digits := make([]int, 0, 9)
	for t := num; t > 0; t /= 10 {
		digits = append([]int{t % 10}, digits...)
	}
	leftMost := digits[0]
	leftMostUnder9 := 9
	for _, digit := range digits {
		if digit < 9 {
			leftMostUnder9 = digit
			break
		}
	}
	maxVal, minVal := 0, 0
	for _, digit := range digits {
		minVal *= 10
		if digit != leftMost {
			minVal += digit
		}
		maxVal *= 10
		if digit != leftMostUnder9 {
			maxVal += digit
		} else {
			maxVal += 9
		}
	}
	return maxVal - minVal
}

func MaxDiff(num int) int {
	// 1432
	// Convert number to string for easier digit manipulation
	numStr := strconv.Itoa(num)

	// Find first non-9 digit for maximum number
	maxNum := num
	for i := 0; i < len(numStr); i++ {
		if numStr[i] != '9' {
			// Replace all occurrences of this digit with 9
			maxNum, _ = strconv.Atoi(strings.ReplaceAll(numStr, string(numStr[i]), "9"))
			break
		}
	}

	// Find first non-1 digit for minimum number
	minNum := num
	if numStr[0] != '1' {
		// If first digit is not 1, replace all occurrences with 1
		minNum, _ = strconv.Atoi(strings.ReplaceAll(numStr, string(numStr[0]), "1"))
	} else {
		// If first digit is 1, find first non-0 digit and replace with 0
		for i := 1; i < len(numStr); i++ {
			if numStr[i] != '0' && numStr[i] != '1' {
				minNum, _ = strconv.Atoi(strings.ReplaceAll(numStr, string(numStr[i]), "0"))
				break
			}
		}
	}

	return maxNum - minNum
}

func MaximumDifference_fail(nums []int) int {
	n := len(nums)
	prefixMin := make([]int, math.MaxInt, n+1)
	for i, num := range nums {
		prefixMin[i+1] = prefixMin[i]
		if num < prefixMin[i] {
			prefixMin[i+1] = num
		}
	}
	maxDiff := 0
	for i, num := range nums {
		diff := num - prefixMin[i+1]
		if diff < 0 {
			diff = -diff
		}
		if maxDiff < diff {
			maxDiff = diff
		}
	}
	return maxDiff
}

func MaximumDifference(nums []int) int {
	// 2016
	minVal, maxDiff := nums[0], -1
	for i := 1; i < len(nums); i++ {
		if nums[i] > minVal {
			maxDiff = max(maxDiff, nums[i]-minVal)
		} else {
			minVal = nums[i]
		}
	}
	return maxDiff
}

func MajorityElement(nums []int) int {
	// 169
	candidate, cnt := nums[0], 1
	for _, num := range nums[1:] {
		if num == candidate {
			cnt++
			continue
		}
		cnt--
		if cnt == 0 {
			candidate = num
			cnt = 1
		}
	}
	return candidate
}

func DivideArray(nums []int, k int) (result [][]int) {
	// 2966
	slices.Sort(nums)
	group := []int{}
	for _, num := range nums {
		group = append(group, num)
		if len(group) == 3 {
			if group[len(group)-1]-group[0] <= k {
				result = append(result, group)
				group = []int{}
			} else {
				return [][]int{}
			}
		}
	}
	if len(group) > 0 {
		return [][]int{}
	}
	return result
}

func PartitionArray(nums []int, k int) int {
	// 2294
	slices.Sort(nums)
	result := 1
	minNum := nums[0]
	for _, num := range nums[1:] {
		if num > minNum+k {
			result++
			minNum = num
		}
	}
	return result
}

func HammingWeight(n int) (cnt int) {
	// 191
	for t := n; t > 0; t /= 2 {
		cnt += t % 2
	}
	return cnt
}

func MinimumDeletions(word string, k int) int {
	// 3085
	// dabdcbdcdcd
	// a:1 b:2 c:3 d:5
	counts := [26]int{}
	for i := 0; i < len(word); i++ {
		counts[word[i]-'a']++
	}
	freqs := []int{}
	for _, cnt := range counts {
		if cnt > 0 {
			freqs = append(freqs, cnt)
		}
	}
	slices.Sort(freqs)
	minDel := math.MaxInt
	for i := 0; i < len(freqs); i++ {
		target := freqs[i]
		del := 0
		for j := 0; j < len(freqs); j++ {
			if freqs[j] < target {
				del += freqs[j]
			} else if freqs[j] > target+k {
				del += freqs[j] - (target + k)
			}
		}
		if del < minDel {
			minDel = del
		}
	}
	return minDel
}

func DivideString(s string, k int, fill byte) []string {
	// 2138
	strings := []string{}
	n := len(s)
	start := 0
	for ; start <= n-k; start += k {
		strings = append(strings, s[start:start+k])
	}
	if start < n {
		last := s[start:]
		for len(last) < k {
			last += string(fill)
		}
		strings = append(strings, last)
	}
	return strings
}

func CountGoodNumbers(n int64) int {
	// 1922 - fast pow
	const MOD = 1_000_000_007
	pow := func(a, b int64) int64 {
		res := int64(1)
		a = a % MOD
		for b > 0 {
			if b%2 == 1 {
				res = (res * a) % MOD
			}
			a = (a * a) % MOD
			b /= 2
		}
		return res
	}
	odd := n / 2
	even := n - odd
	return int((pow(5, even) * pow(4, odd)) % MOD)
}

func FindKDistantIndices_52ms(nums []int, key int, k int) []int {
	m := make(map[int]struct{})
	for i, num := range nums {
		if num == key {
			for j := max(0, i-k); j <= min(len(nums)-1, i+k); j++ {
				m[j] = struct{}{}
			}
		}
	}
	result := make([]int, 0, len(m))
	for idx := range m {
		result = append(result, idx)
	}
	slices.Sort(result)
	return result
}

func FindKDistantIndices(nums []int, key int, k int) (res []int) {
	// 2200
	right := 0
	n := len(nums)
	for i, num := range nums {
		if num == key {
			left := max(right, i-k)
			right = min(n-1, i+k) + 1
			for j := left; j < right; j++ {
				res = append(res, j)
			}
		}
	}
	return res
}

func ContainsNearbyDuplicate(nums []int, k int) bool {
	// 219
	lastIndicesByNum := make(map[int]int)
	for j, num := range nums {
		if i := lastIndicesByNum[num]; i > 0 && j+1-i <= k {
			return true
		}
		lastIndicesByNum[num] = j + 1 // 1-index to know num exists or not
	}
	return false
}

func AddDigits(num int) int {
	// 258
	newNum := num
	for newNum > 9 {
		t := newNum
		newNum = 0
		for ; t > 0; t /= 10 {
			newNum += t % 10
		}
	}
	return newNum
}

func IsUgly(n int) bool {
	// 263
	if n == 0 {
		return false
	}
	for ; n%5 == 0; n /= 5 {
	}
	for ; n%3 == 0; n /= 3 {
	}
	for ; n%2 == 0; n /= 2 {
	}
	return n == 1
}

var isBadVersion func(int) bool

func FirstBadVersion(n int) int {
	// 278
	left, right := 1, n
	for left < right {
		mid := left + (right-left)/2
		if isBadVersion(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func MoveZeroes(nums []int) {
	// 283
	i := 0
	for _, num := range nums {
		if num == 0 {
			continue
		}
		nums[i] = num
		i++
	}
	for j := i; j < len(nums); j++ {
		nums[j] = 0
	}
}

func LongestSubsequence(s string, k int) int {
	// 2311
	oneCount, num, pow := 0, 0, 1
	for i := len(s) - 1; i > -1 && num+pow <= k; i-- {
		if s[i] == '1' {
			oneCount++
			num += pow
		}
		pow *= 2
	}
	return strings.Count(s, "0") + oneCount
}

func MaxSubsequenceNoOrder(nums []int, k int) []int {
	res := make([]int, 0, len(nums))
	for _, num := range nums {
		pos, _ := slices.BinarySearch(res, num)
		res = slices.Insert(res, pos, num)
		if len(res) > k {
			res = res[1:]
		}
	}
	return res
}

func MaxSubsequence(nums []int, k int) []int {
	// 2099
	type pair struct {
		val int
		idx int
	}
	n := len(nums)
	pairs := make([]pair, n)
	for i, v := range nums {
		pairs[i] = pair{v, i}
	}
	// Sort by value descending, then by index ascending
	slices.SortFunc(pairs, func(a, b pair) int {
		if a.val != b.val {
			return b.val - a.val
		}
		return a.idx - b.idx
	})
	selected := pairs[:k]
	// Sort selected by index to preserve order
	slices.SortFunc(selected, func(a, b pair) int {
		return a.idx - b.idx
	})
	res := make([]int, k)
	for i, p := range selected {
		res[i] = p.val
	}
	return res
}
