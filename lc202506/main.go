package main

import (
	"fmt"
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
		fmt.Println(val, digits[i])
	}
	if val == 9 {
		return append([]int{1}, digits...)
	}
	return digits
}
