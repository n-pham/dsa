package main

import (
	"fmt"
	// "math"
	// "slices"
)

func maxScore1422(s string) int {
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

func main() {
	fmt.Println(vowelStrings2559([]string{"aba", "bcb", "ece", "aa", "e"}, [][]int{{0, 2}, {1, 4}, {1, 1}})) //
	// fmt.Println(maxScore1422("011101")) // 5
	// fmt.Println(maxScore1422("1111")) // 3
	// fmt.Println(maxScore1422("00")) // 1
	// fmt.Println(maxScore1422("11100")) // 2
}
