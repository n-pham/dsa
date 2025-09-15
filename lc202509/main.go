package main

//lint:file-ignore U1000 Ignore all unused code, it's generated

import (
	"log"
	"math"
	"os"
	"runtime"
	"strings"
)

var (
	debugEnabled = os.Getenv("DEBUG") == "true"
	debugLogger  = log.New(os.Stdout, "DEBUG: ", log.Ldate|log.Ltime)
)

func debugLog(v ...any) {
	if debugEnabled {
		pc, _, _, ok := runtime.Caller(1)
		if !ok {
			debugLogger.Println(v...)
			return
		}
		if fn := runtime.FuncForPC(pc); fn != nil {
			name := fn.Name()
			if lastDot := strings.LastIndex(name, "."); lastDot != -1 {
				name = name[lastDot+1:]
			}
			args := append([]any{name + ":"}, v...)
			debugLogger.Println(args...)
		}
	}
}

func Intersection(nums1 []int, nums2 []int) []int {
	// 349
	num1Set := make([]bool, 1001)
	for _, num := range nums1 {
		num1Set[num] = true
	}
	commonMap := make(map[int]struct{}, 1001)
	for _, num := range nums2 {
		if num1Set[num] {
			commonMap[num] = struct{}{}
		}
	}
	commonNums := make([]int, 0, 1001)
	for num := range commonMap {
		commonNums = append(commonNums, num)
	}
	return commonNums
}

func FindClosest(x int, y int, z int) int {
	// 3516
	d1, d2 := math.Abs(float64(x-z)), math.Abs(float64(y-z))
	if d1 == d2 {
		return 0
	} else if d1 > d2 {
		return 2
	}
	return 1
}

func SumZero(n int) []int {
	// 1304
	// either 0  1  2 -1 -2
	// or    -2  1  2 -1
	maxNum := n / 2
	result := make([]int, n)
	for num := 1; num <= maxNum; num++ {
		result[num] = num
		if maxNum+num < n {
			result[maxNum+num] = -num
		} else {
			result[0] = -num
		}
		debugLog(num, result)
	}

	return result
}

func hasZero(n int) bool {
	has := false
	for t := n; t > 0; t /= 10 {
		if t%10 == 0 {
			has = true
			break
		}
	}
	return has
}

func GetNoZeroIntegers(n int) []int {
	// 1317
	for a := 1; a < n; a++ {
		if hasZero(a) {
			continue
		}
		if !hasZero(n - a) {
			return []int{a, n - a}
		}
	}
	return []int{}
}

func MaxFreqSum(s string) int {
	// 3541
	charCount := [26]int{}
	for _, c := range s {
		charCount[c-'a']++
	}
	maxVowelCount, maxConsonantCount := 0, 0
	for i, cnt := range charCount {
		if i == 0 || i == 4 || i == 8 || i == 14 || i == 20 {
			maxVowelCount = max(maxVowelCount, cnt)
		} else {
			maxConsonantCount = max(maxConsonantCount, cnt)
		}
	}
	return maxVowelCount + maxConsonantCount
}

func CanBeTypedWords(text string, brokenLetters string) (cnt int) {
	// 1935
	brokenLetterMap := [26]bool{}
	for _, letter := range brokenLetters {
		brokenLetterMap[letter-'a'] = true
	}
	canTypeWord := true
	for _, c := range text {
		if c == ' ' {
			if canTypeWord {
				cnt++
			}
			canTypeWord = true
		} else if canTypeWord && brokenLetterMap[c-'a'] {
			canTypeWord = false
		}
	}
	if canTypeWord {
		cnt++
	}
	return
}
