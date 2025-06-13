package main

import (
	"fmt"
	"math"
	"sort"
)

//lint:file-ignore U1000 Ignore all unused code, it's generated

func answerString(word string, numFriends int) string {
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

func robotWithString_fail(s string) string {
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

func robotWithString(s string) string {
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

func clearStars_fail(s string) string {
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

func clearStars(s string) string {
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

func lexicalOrder(n int) (nums []int) {
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

func findKthNumber_time(n int, k int) int {
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

func findKthNumber(n int, k int) int {
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

func searchInsert(nums []int, target int) int {
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

func plusOne(digits []int) []int {
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

func maxDifference(s string) int {
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

func maxDifference3445(s string, k int) int {
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

func addBinary(a string, b string) string {
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

func maxAdjacentDistance(nums []int) int {
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

func deleteDuplicates(head *ListNode) *ListNode {
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

func singleNumber(nums []int) (num int) {
	// 136
	for _, n := range nums {
		num ^= n
	}
	return num
}

func convertToTitle(columnNumber int) string {
	// 168
	panic("not implemented")
}

func minimizeMax(nums []int, p int) int {
	// 2616
	sort.Ints(nums)
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
