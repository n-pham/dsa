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
