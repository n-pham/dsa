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
	// 3170 LTE
	// aaba* --> aab
	// Stack to store indices of characters
	stack := make([]int, 0, len(s))
	// Array to track if character at index i should be included
	keep := make([]bool, len(s))

	// Process each character
	for i := range s {
		if s[i] != '*' {
			stack = append(stack, i)
			keep[i] = true
		} else if len(stack) > 0 {
			// Find leftmost smallest character in stack
			smallestIdx := stack[0]
			smallestPos := 0

			// Check all positions in stack for smallest character
			for j := 1; j < len(stack); j++ {
				if s[stack[j]] <= s[smallestIdx] {
					smallestIdx = stack[j]
					smallestPos = j
				}
			}

			// Remove the character by marking it as not kept
			keep[smallestIdx] = false
			// Remove the position from stack
			stack = append(stack[:smallestPos], stack[smallestPos+1:]...)
		}
	}

	// Build result string
	rs := make([]byte, 0, len(s))
	for i := range s {
		if keep[i] {
			rs = append(rs, s[i])
		}
	}

	return string(rs)
}
