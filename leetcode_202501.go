package main

import (
	"fmt"
	// "math"
	// "slices"
)

func maxScore1422_fail(s string) int {
	// 011101
	// iZCount = -1,0,4
	// zCount 2: oCount = len-1 - iZCount[zCount]
	// zCount 1: oCount = len-1 - iZCount[zCount] - 1 // zCount == 2
	panic("not implemented")
	iZCount := []int{-1}
	zTotal := 0
	for i, c := range s {
		if c == '0' {
			zTotal += 1
			iZCount = append(iZCount, i)
		}
	}
	fmt.Println(iZCount)
	if len(iZCount) == 1 {
		return len(s) - 1
	}
	maxS := max(0, iZCount[1]) // "1"s before first 0
	for zCount := len(iZCount) - 1; zCount > 0; zCount-- {
		// ignore if end is 0
		if iZCount[zCount] == len(s)-1 {
			continue
		}
		oCount := len(s) - 1 - iZCount[zCount] - (len(iZCount) - 1 - zCount)
		fmt.Println("maxS", maxS, "zCount", zCount, "oCount", oCount)
		maxS = max(maxS, zCount+oCount)
	}
	return maxS
}

func maxScore1422(s string) int {
	// 1422
	maxS, forwardSums, backwardSums := 0, make([]int, len(s)+1), make([]int, len(s)+1)
	for i, c := range s {
		forwardSums[i+1] = forwardSums[i]
		if c == '0' {
			forwardSums[i+1] += 1
		}
	}
	for i := len(s) - 1; i > 0; i-- {
		backwardSums[i] = backwardSums[i+1]
		if s[i] == '1' {
			backwardSums[i] += 1
		}
	}
	for i := 1; i <= len(s)-1; i++ {
		maxS = max(maxS, forwardSums[i]+backwardSums[i])
	}
	return maxS
}

func vowelStrings2559_time(words []string, queries [][]int) []int {
	queriesByStart := make(map[int]map[int]struct{})
	queriesByEnd := make(map[int]map[int]struct{})
	for i, q := range queries {
		if m, found := queriesByStart[q[0]]; found {
			m[i] = struct{}{}
			// queriesByStart[q[0]] = m
		} else {
			m := make(map[int]struct{})
			m[i] = struct{}{}
			queriesByStart[q[0]] = m
		}
		if m, found := queriesByEnd[q[1]]; found {
			m[i] = struct{}{}
			// queriesByStart[q[0]] = m
		} else {
			m := make(map[int]struct{})
			m[i] = struct{}{}
			queriesByEnd[q[1]] = m
		}
	}
	// fmt.Println(queriesByStart, queriesByEnd)
	rs := make([]int, len(queries))
	currentQueryByIdx := make(map[int]struct{})
	for i, w := range words {
		// fmt.Println("i", i, "w", w, "currentQueryByIdx", currentQueryByIdx)
		if qIs, found := queriesByStart[i]; found {
			// fmt.Println("Start qIs", qIs)
			for qi := range qIs {
				currentQueryByIdx[qi] = struct{}{}
			}
		}
		if isVowel(w[0]) && isVowel(w[len(w)-1]) {
			for qi := range currentQueryByIdx {
				rs[qi] += 1
			}
		}
		if qIs, found := queriesByEnd[i]; found {
			// fmt.Println("End qIs", qIs)
			for qi := range qIs {
				delete(currentQueryByIdx, qi)
			}
		}
	}
	return rs
}

func vowelStrings2559_time_2(words []string, queries [][]int) []int {
	ones := make([]int, len(words))
	for i, w := range words {
		if isVowel(w[0]) && isVowel(w[len(w)-1]) {
			ones[i] = int(1)
		}
	}
	fmt.Println(ones)
	rs := make([]int, len(queries))
	for i, q := range queries {
		fmt.Println(ones[q[0]:q[1]])
		for _, val := range ones[q[0] : q[1]+1] {
			rs[i] += val
		}
	}
	return rs
}

func vowelStrings2559(words []string, queries [][]int) []int {
	//     aba bcb ece  aa   e
	//       1   0   1   1   1
	//   0   1   1   2   3   4
	// 2559
	prefixSums := make([]int, len(words)+1)
	for i, w := range words {
		prefixSums[i+1] = prefixSums[i]
		if isVowel(w[0]) && isVowel(w[len(w)-1]) {
			prefixSums[i+1] += 1
		}
	}
	rs := make([]int, len(queries))
	for i, q := range queries {
		rs[i] = prefixSums[q[1]+1] - prefixSums[q[0]]
	}
	return rs
}

func isVowel(c byte) bool {
	switch c {
	case
		'a',
		'e',
		'i',
		'o',
		'u':
		return true
	}
	return false
}

func waysToSplitArray220_12ms(nums []int) int {
	//     10  4 -8  7
	// --> 10 14  6 13
	//     13  3 -1  7 <--
	// 2270
	forwardSums := make([]int, len(nums)+1)
	for i, num := range nums {
		forwardSums[i+1] = forwardSums[i] + num
	}
	backwardSums := make([]int, len(nums)+1)
	for i := len(nums) - 1; i >= 0; i-- {
		backwardSums[i] = backwardSums[i+1] + nums[i]
	}
	// fmt.Println(forwardSums, backwardSums)
	cnt := 0
	for i := 0; i < len(nums)-1; i++ {
		// fmt.Println(i, forwardSums[i], backwardSums[i+1])
		if forwardSums[i+1] >= backwardSums[i+1] {
			cnt += 1
		}
	}
	return cnt
}

func waysToSplitArray220(nums []int) int {
	//     10  4 -8  7  --> sum 13
	//     10  3            cur, sum-cur
	//        14 -1         cur, sum-cur
	//            6  7      cur, sum-cur
	// 2270
	var cnt, curSum, total int
	for _, num := range nums {
		total += num
	}
	for i := 0; i < len(nums)-1; i++ {
		curSum += nums[i]
		if curSum >= total-curSum {
			cnt++
		}
	}
	return cnt
}

func executeInstructions2120(n int, startPos []int, s string) []int {
	//    R  R  D  D  L  U
	// dx 1  2  2  2  1  1
	// dy 0  0 -1 -2 -1  0
	// dx -1  0 -1 -1 -1  0 <--|
	// dy  1  1  1  0 -1 -1 <--|
	// 2120
	panic("not implemented")
}

func minOperations1551(n int) int {
	// 1551
	// 1 3 5        --> n 3
	// 2 3 4
	// 3 3 3
	// 2+0+2
	// 1 3 5 7 9 11 --> n 6
	// 5+3+1
	rs := 0
	for d := n-1; d > 0; d -= 2 {
		rs += d
	}
	return rs
}

func main() {
	// fmt.Println(executeInstructions2120(3, []int {0,1}, "RRDDLU")) // 
	// fmt.Println(waysToSplitArray220([]int {10,4,-8,7})) // 2
	// fmt.Println(waysToSplitArray220([]int {-2,-1})) // 0
	// fmt.Println(vowelStrings2559([]string{"aba", "bcb", "ece", "aa", "e"}, [][]int{{0, 2}, {1, 4}, {1, 1}})) //
	// fmt.Println(maxScore1422("00"))     // 1
	// fmt.Println(maxScore1422("010"))    // 2
	// fmt.Println(maxScore1422("011101")) // 5
	// fmt.Println(maxScore1422("1111"))   // 3
	// fmt.Println(maxScore1422("11100"))  // 2
}
