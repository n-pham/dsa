package main

//lint:file-ignore U1000 Ignore all unused code, it's generated

import (
	"log"
	"math"
	"os"
	"runtime"
	"strings"
)

//lint:file-ignore U1000 Ignore all unused code, it's generated

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
	for i1 < i2 {
		if !vowels[runeSlice[i1]] {
			i1++
		} else if !vowels[runeSlice[i2]] {
			i2--
		} else {
			runeSlice[i1], runeSlice[i2] = runeSlice[i2], runeSlice[i1]
			i1++
			i2--
		}
	}
	return string(runeSlice)
}
