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
	iZCount := []int {-1}
	zTotal := 0
	for i, c := range s {
		if c == '0' {
			zTotal += 1
			iZCount = append(iZCount, i)
		}
	}
	fmt.Println(iZCount)
	if len(iZCount) == 1 {
		return len(s)-1
	}
	maxS := max(0, iZCount[1]) // "1"s before first 0
	for zCount := len(iZCount)-1; zCount > 0; zCount-- {
		// ignore if end is 0
		if iZCount[zCount] == len(s)-1 {
			continue
		}
		oCount := len(s)-1 - iZCount[zCount] - (len(iZCount)-1-zCount)
		fmt.Println("maxS", maxS, "zCount", zCount, "oCount", oCount)
		maxS = max(maxS, zCount + oCount)
	}
	return maxS
}

func main() {
	// fmt.Println(maxScore1422("011101")) // 5
	// fmt.Println(maxScore1422("1111")) // 3
	fmt.Println(maxScore1422("00")) // 1
	fmt.Println(maxScore1422("11100")) // 2
}