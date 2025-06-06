package main

import "fmt"

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
