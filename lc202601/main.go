package main

//lint:file-ignore U1000 Ignore all unused code, it's generated

import (
	"dsa/kit"
	"math"
)

//lint:file-ignore U1000 Ignore all unused code, it's generated

func FirstUniqChar(s string) int {
	// 387
	// a     a:1
	// ab    a:1,b:2
	// abc   a:1,b:2,c:3
	// abca  a:max,b:2,c:3
	firstIndexByChar := [26]int{} // array of 26 chars is faster than map
	for i, char := range s {
		if firstIndexByChar[char-'a'] == 0 {
			firstIndexByChar[char-'a'] = i + 1 // use 1-based index to avoid default value 0
		} else {
			firstIndexByChar[char-'a'] = math.MaxInt
		}
	}
	firstIndex := math.MaxInt
	for _, index := range firstIndexByChar {
		if index > 0 && index < firstIndex {
			firstIndex = index
		}
	}
	if firstIndex == math.MaxInt {
		return -1
	}
	return firstIndex - 1
}

func CanConstruct(ransomNote string, magazine string) bool {
	// 383
	countByChar := [26]int{} // array of 26 chars is faster than map
	for _, char := range magazine {
		countByChar[char-'a']++
	}
	for _, char := range ransomNote {
		countByChar[char-'a']--
		if countByChar[char-'a'] < 0 {
			return false
		}
	}
	return true
}

func SumFourDivisors(nums []int) (result int) {
	// 1390
	for _, num := range nums {
		maxDivisor := int(math.Sqrt(float64(num)))
		divisorCount, divisorSum := 2, 1+num // 1 and num are always divisors
		for i := 2; i <= maxDivisor; i++ {
			if num%i == 0 {
				if i*i == num {
					divisorCount += 1
					divisorSum += i
				} else {
					divisorCount += 2
					divisorSum += i + num/i
				}
			}
		}
		if divisorCount == 4 {
			result += divisorSum
		}
	}
	return
}

func RepeatedNTimes(nums []int) int {
	// 961
	n := len(nums) / 2
	numMap := make(map[int]struct{}, n+1)
	for i := 0; i < n+2; i++ {
		if _, found := numMap[nums[i]]; found {
			return nums[i]
		}
		numMap[nums[i]] = struct{}{}
	}
	return 0
}

func ReverseVowels(s string) string {
	// 345
	vowels := map[rune]bool{
		'a': true, 'e': true, 'i': true, 'o': true, 'u': true,
		'A': true, 'E': true, 'I': true, 'O': true, 'U': true,
	}
	runeSlice := []rune(s)
	i1, i2 := 0, len(runeSlice)-1
	kit.DebugLog("Initial:", string(runeSlice), "i1:", i1, "i2:", i2)
	for i1 < i2 {
		kit.DebugLog("Before iteration: ", "runeSlice:", string(runeSlice), "i1:", i1, "i2:", i2)
		if !vowels[runeSlice[i1]] {
			kit.DebugLog("runeSlice[i1] not a vowel:", string(runeSlice[i1]))
			i1++
		} else if !vowels[runeSlice[i2]] {
			kit.DebugLog("runeSlice[i2] not a vowel:", string(runeSlice[i2]))
			i2--
		} else {
			kit.DebugLog("Swapping:", string(runeSlice[i1]), string(runeSlice[i2]))
			runeSlice[i1], runeSlice[i2] = runeSlice[i2], runeSlice[i1]
			i1++
			i2--
		}
		kit.DebugLog("After iteration: ", "runeSlice:", string(runeSlice), "i1:", i1, "i2:", i2)
	}
	kit.DebugLog("Final:", string(runeSlice))
	return string(runeSlice)
}
