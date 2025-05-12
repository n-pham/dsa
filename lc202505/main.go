package main

//lint:file-ignore U1000 Ignore all unused code, it's generated

import (
	"fmt"
	"slices"
)

func maxTaskAssign(tasks []int, workers []int, pills int, strength int) int {
	// 2071
	// 1 2 3
	// 0 3 3
	fmt.Println(tasks, workers, pills, strength)
	return 0
}

func pushDominoes(dominoes string) string {
	// 838
	//  . L . R . . . L R . . L . .
	// [0 0 0 1 2 3 4 0 1 2 3 0 0 0]
	// [2 1 0 0 4 3 2 1 0 3 2 1 0 0]
	cnt := 0
	fallingRight, rightTimes := false, make([]int, len(dominoes))
	fallingLeft, leftTimes := false, make([]int, len(dominoes))
	for i, d := range dominoes {
		fmt.Print(" ", string(d))
		if d == 'L' {
			fallingRight = false
			continue
		}
		if d == 'R' {
			fallingRight = true
			rightTimes[i] = 1
			cnt = 2
		} else if d == '.' && fallingRight {
			rightTimes[i] = cnt
			cnt++
		}
	}
	fmt.Println()
	fmt.Println(rightTimes)
	cnt = 0
	for i := len(dominoes) - 1; i > -1; i-- {
		d := dominoes[i]
		if d == 'R' {
			fallingLeft = false
			continue
		}
		if d == 'L' {
			fallingLeft = true
			leftTimes[i] = 1
			cnt = 2
		} else if d == '.' && fallingLeft {
			leftTimes[i] = cnt
			cnt++
		}
	}
	fmt.Println(leftTimes)
	rs := make([]rune, len(dominoes))
	for i, rightTime := range rightTimes {
		leftTime := leftTimes[i]
		if leftTime == rightTime {
			rs[i] = '.'
			continue
		}
		if leftTime == 0 {
			rs[i] = 'R'
		} else if rightTime == 0 {
			rs[i] = 'L'
		} else if leftTime < rightTime {
			rs[i] = 'L'
		} else {
			rs[i] = 'R'
		}
	}
	return string(rs)
}

func minDominoRotations(tops []int, bottoms []int) int {
	// 1007
	// 2,1,2,4,2,2
	// 5,2,6,2,3,2
	num1, num2 := 0, 0
	for i := 0; i < len(tops); i++ {
		top := tops[i]
		bottom := bottoms[i]
		if num1 == 0 {
			num1 = top
		} else if num2 == 0 {
			num2 = top
		}
		if num1 == 0 {
			num1 = bottom
		} else if num2 == 0 {
			num2 = bottom
		}
		fmt.Println(num1, num2, top, bottom)
		if num1 != 0 && num2 != 0 && !((top == num1 || top == num2) && (bottom == num1 || bottom == num2)) {
			return -1
		}
	}
	return 0
}

func numEquivDominoPairs_1(dominoes [][]int) (pairCnt int) {
	mapCnt := make(map[[2]int]int)
	for _, d := range dominoes {
		smaller, bigger := d[0], d[1]
		if smaller > bigger {
			smaller, bigger = d[1], d[0]
		}
		mapCnt[[2]int{smaller, bigger}]++
	}
	fmt.Println(mapCnt)
	for _, cnt := range mapCnt {
		pairCnt += cnt * (cnt - 1) / 2
	}
	return pairCnt
}

func numEquivDominoPairs(dominoes [][]int) (pairCnt int) {
	// 1128
	mapCnt := make(map[[2]int]int, len(dominoes))
	for _, d := range dominoes {
		smaller, bigger := d[0], d[1]
		if smaller > bigger {
			smaller, bigger = d[1], d[0]
		}
		pairCnt += mapCnt[[2]int{smaller, bigger}]
		mapCnt[[2]int{smaller, bigger}]++
	}
	return pairCnt
}

func numTilings(n int) int {
	// 790
	// a  a a  a e  a e e  a i e  i i a  z z x  z x x  z z x x  z z e e x
	// a  e e  a e  a i i  a i e  e e a  z x x  z z x  z a a x  z a a x x
	const MOD = 1_000_000_007
	tilingWays := [4]int{1, 0, 0, 0}
	for i := 1; i <= n; i++ {
		newTilingWays := [4]int{0, 0, 0, 0}
		// Full cover is obtained by adding one 2x2 tile or two 2x1 tiles to any of the four previous states.
		newTilingWays[0] = (tilingWays[0] + tilingWays[1] + tilingWays[2] + tilingWays[3]) % MOD
		// Top row missing one can be obtained by adding a 2x1 tile to the previous state of bottom row missing one or both top and bottom missing one.
		newTilingWays[1] = (tilingWays[2] + tilingWays[3]) % MOD
		// Bottom row missing one can be obtained by adding a 2x1 tile to the previous state of top row missing one or both top and bottom missing one.
		newTilingWays[2] = (tilingWays[1] + tilingWays[3]) % MOD
		// Both top and bottom missing one can only be obtained by placing a 2x2 tile in the full cover state.
		newTilingWays[3] = tilingWays[0]

		// Update tilingWays array with the new computed values.
		tilingWays = newTilingWays
	}
	return tilingWays[0]
}

func buildArray(nums []int) []int {
	// 1920
	ans := make([]int, len(nums))
	for i, num := range nums {
		ans[i] = nums[num]
	}
	return ans
}

func romanToInt(s string) (num int) {
	// 13
	singleValue := map[string]int{"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000}
	doubleValue := map[string]int{"IV": 4, "IX": 9, "XL": 40, "XC": 90, "CD": 400, "CM": 900}
	prev := ""
	for i := 0; i < len(s); i++ {
		c := string(s[i])
		if prev == "" {
			prev = c
			continue
		}
		if dv := doubleValue[prev+c]; dv > 0 {
			num += dv
			prev = ""
		} else {
			num += singleValue[prev]
			prev = c
		}
	}
	num += singleValue[prev]
	return num
}

func minTimeToReach(moveTime [][]int) int {
	// 3341
	// 0 4  4↓ .
	// 4 4  5→ 6
	panic("not implemented")
}

func minSum(nums1 []int, nums2 []int) int64 {
	// 2918
	var sum1, sum2, zeroCnt1, zeroCnt2 int64
	for _, num := range nums1 {
		sum1 += int64(num)
		if num == 0 {
			zeroCnt1++
		}
	}
	for _, num := range nums2 {
		sum2 += int64(num)
		if num == 0 {
			zeroCnt2++
		}
	}
	sum1, sum2 = sum1+zeroCnt1, sum2+zeroCnt2
	fmt.Println(sum1, sum2)
	if sum1 == sum2 {
		return sum1
	} else if sum1 < sum2 {
		if zeroCnt1 > 0 {
			return sum2
		}
	} else if zeroCnt2 > 0 {
		return sum1
	}
	return -1
}

func threeConsecutiveOdds(arr []int) bool {
	// 1550
	consecutiveOddCnt := 0
	for _, num := range arr {
		if num%2 == 1 {
			consecutiveOddCnt++
			if consecutiveOddCnt == 3 {
				return true
			}
		} else {
			consecutiveOddCnt = 0
		}
	}
	return false
}

func findEvenNumbers_183ms(digits []int) (nums []int) {
	n := len(digits)
	for k := 0; k < n; k++ {
		if digits[k]%2 != 0 {
			continue
		}
		for i := 0; i < n; i++ {
			if i == k || digits[i] == 0 {
				continue
			}
			for j := 0; j < n; j++ {
				if j == k || j == i {
					continue
				}
				val := 100*digits[i] + 10*digits[j] + digits[k]
				if i, found := slices.BinarySearch(nums, val); !found {
					nums = slices.Insert(nums, i, val)
				}
			}
		}
	}
	return nums
}

func findEvenNumbers(digits []int) (nums []int) {
	// 2094
	freq := [10]int{}
	for _, d := range digits {
		freq[d]++
	}
	for i := 1; i < 10; i++ {
		if freq[i] == 0 {
			continue
		}
		freq[i]--
		for j := 0; j < 10; j++ {
			if freq[j] == 0 {
				continue
			}
			freq[j]--
			for k := 0; k < 10; k += 2 {
				if freq[k] > 0 {
					nums = append(nums, 100*i+10*j+k)
				}
			}
			freq[j]++
		}
		freq[i]++
	}
	slices.Sort(nums)
	return nums
}

func main() {
	fmt.Println(findEvenNumbers([]int{2, 1, 3, 0}))
	fmt.Println(findEvenNumbers([]int{2, 2, 8, 8, 2}))
	// fmt.Println(minSum([]int{0, 16, 28, 12, 10, 15, 25, 24, 6, 0, 0}, []int{20, 15, 19, 5, 6, 29, 25, 8, 12}))
	// fmt.Println(minSum([]int{9, 5}, []int{15, 12, 5, 21, 4, 26, 27, 9, 6, 29, 0, 18, 16, 0, 0, 0, 20}))
	// fmt.Println(romanToInt("LVIII"), romanToInt("MCMXCIV"))
	// fmt.Println(numEquivDominoPairs([][]int{{1, 2}, {2, 1}, {1, 1}, {1, 2}, {2, 2}, {2, 2}}))
	// fmt.Println(minDominoRotations([]int{2, 1, 2, 4, 2, 2}, []int{5, 2, 6, 2, 3, 2}))
	// fmt.Println(minDominoRotations([]int{3, 5, 1, 2, 3}, []int{3, 6, 3, 3, 4}))
	// fmt.Println(pushDominoes(".L.R...LR..L.."))
	// fmt.Println(maxTaskAssign([]int{3, 2, 1}, []int{0, 3, 3}, 1, 1))
}
