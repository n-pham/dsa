package main

//lint:file-ignore U1000 Ignore all unused code, it's generated
import (
	"fmt"
	"math"
	// "github.com/mxschmitt/golang-combinations"
	// "strconv"
	// "strings"
)

func mostPoints_time(questions [][]int) int64 {
	if len(questions) == 0 {
		return 0
	}
	skip0 := mostPoints(questions[1:])
	solve0 := int64(questions[0][0])
	if questions[0][1]+1 < len(questions) {
		solve0 += mostPoints(questions[questions[0][1]+1:])
	}
	if solve0 > skip0 {
		return solve0
	}
	return skip0
}

func mostPoints(questions [][]int) int64 {
	// 2140
	// 43ms
	m := make([]int64, len(questions))
	var recur func(int) int64
	recur = func(index int) int64 {
		if index >= len(questions) {
			return 0
		}
		if val := m[index]; val > 0 {
			return val
		}
		skip0 := recur(index + 1)
		solve0 := int64(questions[index][0])
		if index+questions[index][1]+1 < len(questions) {
			solve0 += recur(index + questions[index][1] + 1)
		}
		if solve0 > skip0 {
			m[index] = solve0
		} else {
			m[index] = skip0
		}
		return m[index]
	}
	return recur(0)
}

func climbStairs(n int) int {
	// 70
	m := map[int]int{}
	var recur func(int) int
	recur = func(n int) int {
		if n <= 2 {
			return n
		}
		if val, exists := m[n]; exists {
			return val
		}
		m[n] = recur(n-1) + recur(n-2)
		return m[n]
	}
	return recur(n)
}

func maximumTripletValue(nums []int) int64 {
	// 2873
	var max int64 = math.MinInt64
	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			for k := j + 1; k < len(nums); k++ {
				val := int64(nums[i]-nums[j]) * int64(nums[k])
				if max < val {
					max = val
				}
			}
		}
	}
	if max < 0 {
		return 0
	}
	return max
}

func arithmeticTriplets(nums []int, diff int) (rs int) {
	// 2367
	// i-j j-k
	for j := 1; j < len(nums)-1; j++ {
		for i := 0; i < j; i++ {
			if nums[j]-nums[i] == diff {
				for k := j + 1; k < len(nums); k++ {
					if nums[k]-nums[j] == diff {
						rs++
					}
				}
			}
		}
	}
	return rs
}

func minimumSum(nums []int) int {
	// 2908
	rs := math.MaxInt
	for j := 1; j < len(nums)-1; j++ {
		for i := 0; i < j; i++ {
			if nums[i] < nums[j] {
				kMin := math.MaxInt
				for k := j + 1; k < len(nums); k++ {
					if nums[j] > nums[k] && kMin > nums[k] {
						kMin = nums[k]
					}
				}
				fmt.Println(nums[i], nums[j], kMin)
				if kMin != math.MaxInt {
					tmp := nums[i] + nums[j] + kMin
					if rs > tmp {
						rs = tmp
					}
				}
			}
		}
	}
	if rs == math.MaxInt {
		return -1
	}
	return rs
}

func main() {
	fmt.Println(minimumSum([]int{6, 5, 4, 3, 4, 5}))  // -1
	fmt.Println(minimumSum([]int{5, 4, 8, 7, 10, 2})) // 13
	// fmt.Println(arithmeticTriplets([]int{0, 1, 4, 6, 7, 10}, 3))
	// fmt.Println(maximumTripletValue([]int{12, 6, 1, 2, 7}))
	// fmt.Println(mostPoints([][]int{{3, 2}, {4, 3}, {4, 4}, {2, 5}}))
	// fmt.Println(climbStairs(3))
}
