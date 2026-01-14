package main

import (
	"dsa/kit"
	"fmt"
	"math"
)

func ThirdMax(nums []int) int {
	// 414
	max1 := math.MinInt
	max2 := math.MinInt
	max3 := math.MinInt
	for _, n := range nums {
		if n > max1 {
			max3 = max2
			max2 = max1
			max1 = n
		} else if n > max2 && n != max1 {
			max3 = max2
			max2 = n
		} else if n > max3 && n != max1 && n != max2 {
			max3 = n
		}
	}
	if max3 == math.MinInt {
		return max1
	}
	return max3
}

func FizzBuzz(n int) []string {
	// 412
	result := make([]string, n)
    for i := 1; i <= n; i++ {
		if i%3 == 0 {
			if i%5 == 0 {
				result[i-1] = "FizzBuzz"
			} else {
				result[i-1] = "Fizz"
			}
		} else {
			if i%5 == 0 {
				result[i-1] = "Buzz"
			} else {
				result[i-1] = fmt.Sprintf("%d", i)
			}
		}
	}
	return result
}

func MinTimeToVisitAllPoints(points [][]int) (totalTime int) {
	// 1266
	// diagonal is faster
	for i := 0; i < len(points)-1; i++ {
		dx := points[i+1][0] - points[i][0]
		if dx < 0 {
			dx = -dx
		}
		dy := points[i+1][1] - points[i][1]
		if dy < 0 {
			dy = -dy
		}
		if dx > dy {
			totalTime += dx
		} else {
			totalTime += dy
		}
	}
	return totalTime
}

func LongestPalindrome(s string) (longest int) {
	// 409
	countByChar := [58]int{} // A to z, faster than map
	for _, char := range s {
		countByChar[char-'A']++
	}
	middle := 0
	for _, cnt := range countByChar {
		if cnt%2 == 0 {
			longest += cnt
		} else {
			middle = 1
			if cnt >= 3 {
				longest += cnt - 1
			}
		}
	}
	return longest + middle
}

func MinimumDeleteSum(s1 string, s2 string) int {
	// 712
	sum := 0
	for _, c := range s1 {
		sum += int(c)
	}
	for _, c := range s2 {
		sum += int(c)
	}

	n := len(s2)
	dp := make([]int, n+1)

	for _, c1 := range s1 {
		diag := 0
		for j, c2 := range s2 {
			temp := dp[j+1]
			if c1 == c2 {
				dp[j+1] = diag + int(c1)
			} else {
				if dp[j+1] < dp[j] {
					dp[j+1] = dp[j]
				}
			}
			diag = temp
		}
	}
	return sum - 2*dp[n]
}

func IsSubsequence(s string, t string) bool {
	// 392
	j, lent := 0, len(t)
	for _, char := range s {
		for ; j < lent && rune(t[j]) != char; j++ {
		}
		if j >= len(t) {
			return false
		}
		j++
	}
	return true
}

func FindTheDifference(s string, t string) byte {
	// 389
	countByChar := [26]int{} // array of 26 chars is faster than map
	for _, char := range s {
		countByChar[char-'a']++
	}
	for _, char := range t {
		countByChar[char-'a']--
		if countByChar[char-'a'] < 0 {
			return byte(char)
		}
	}
	return 0
}

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
