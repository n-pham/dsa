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

func lengthAfterTransformations(s string, t int) (ln int) {
	// 3335
	// a  b  c  y  y    2
	// [1 1 1 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 2 0]
	// [0 1 1 1 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 2]
	// [2 2 1 1 1 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]
	const MOD = 1_000_000_007
	prev := [26]int{}
	for _, c := range s {
		prev[c-'a']++
	}
	for step := 0; step < t; step++ {
		curr := [26]int{}
		for i := 0; i < 25; i++ {
			curr[i+1] = prev[i] % MOD
		}
		// z --> a + b
		curr[0] = prev[25]
		curr[1] = (curr[1] + prev[25]) % MOD
		prev = curr
	}
	for _, cnt := range prev {
		ln += cnt
	}
	return ln % MOD
}

func lengthAfterTransformations3337_failed(s string, t int, nums []int) (ln int) {
	// time --> maybe to use matrix
	const MOD = 1_000_000_007
	prev := [26]int{}
	for _, c := range s {
		prev[c-'a']++
	}
	for step := 0; step < t; step++ {
		curr := [26]int{}
		for i := 0; i < 26; i++ {
			if prev[i] == 0 {
				continue
			}
			for j := 0; j < nums[i]; j++ {
				next := (i + j + 1) % 26
				curr[next] = (curr[next] + prev[i]) % MOD
			}
		}
		prev = curr
	}
	for _, count := range prev {
		ln = (ln + count) % MOD
	}
	return ln
}

func lengthAfterTransformations3337(s string, t int, nums []int) (ln int) {
	// 3337
	const MOD = 1_000_000_007
	// Create transition matrix where each element (i,j) represents how many times
	// character i transforms into character j in one step
	matrix := make([][]int, 26)
	for i := range matrix {
		matrix[i] = make([]int, 26)
		for j := 0; j < nums[i]; j++ {
			next := (i + j + 1) % 26
			matrix[i][next] = (matrix[i][next] + 1) % MOD
		}
	}
	state := make([]int, 26)
	for _, c := range s {
		state[c-'a']++
	}
	// Fast matrix power to calculate state after t steps
	for t > 0 {
		if t&1 == 1 {
			// Multiply state by matrix
			newState := make([]int, 26)
			for i := 0; i < 26; i++ {
				for j := 0; j < 26; j++ {
					newState[j] = (newState[j] + (state[i]*matrix[i][j])%MOD) % MOD
				}
			}
			state = newState
		}
		newMatrix := make([][]int, 26)
		for i := range newMatrix {
			newMatrix[i] = make([]int, 26)
			for j := 0; j < 26; j++ {
				for k := 0; k < 26; k++ {
					newMatrix[i][j] = (newMatrix[i][j] + (matrix[i][k]*matrix[k][j])%MOD) % MOD
				}
			}
		}
		matrix = newMatrix
		t >>= 1
	}
	for _, count := range state {
		ln = (ln + count) % MOD
	}
	return ln
}

func getLongestSubsequence(words []string, groups []int) (rs []string) {
	// 2900
	// 1 0 1 1
	prev := groups[0]
	rs = append(rs, words[0])
	for i, g := range groups[1:] {
		if g != prev {
			rs = append(rs, words[i+1])
		}
		prev = g
	}
	return rs
}

func getWordsInLongestSubsequence(words []string, groups []int) []string {
	// 2901
	// dp[i] represent the length of the longest subsequence ending with words[i] that satisfies the conditions.
	// dp[i] = (maximum value of dp[j]) + 1 for indices j < i, where groups[i] != groups[j], words[i] and words[j] are equal in length, and the hamming distance between words[i] and words[j] is exactly 1.
	// Keep track of the j values used to achieve the maximum dp[i] for each index i.
	// The expected array's length is max(dp[0:n]), and starting from the index having the maximum value in dp, we can trace backward to get the words.
	panic("not implemented")
}

func sortColors(nums []int) {
	// 75
	mapCnt := [3]int{0, 0, 0}
	for _, num := range nums {
		mapCnt[num]++
	}
	i := 0
	for color, v := range mapCnt {
		for j := 0; j < v; j++ {
			nums[i] = color
			i++
		}
	}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) (rs int) {
	// 230
	var recur func(*TreeNode)
	recur = func(node *TreeNode) {
		if node == nil {
			return
		}
		recur(node.Left)
		k--
		if k == 0 {
			rs = node.Val
			return
		}
		recur(node.Right)
	}
	recur(root)
	return rs
}

func triangleType_fail(nums []int) string {
	mapLen := make(map[int]struct{}, 3)
	namesByLen := map[int]string{1: "equilateral", 2: "isosceles", 3: "scalene"}
	for _, num := range nums {
		mapLen[num] = struct{}{}
	}
	return namesByLen[len(mapLen)]
}

func triangleType(nums []int) string {
	// 3024
	a, b, c := nums[0], nums[1], nums[2]
	if a+b <= c || a+c <= b || c+b <= a {
		return "none"
	} else if a == b && b == c && c == a {
		return "equilateral"
	} else if a == b || b == c || c == a {
		return "isosceles"
	}
	return "scalene"
}

func isZeroArray_fail(nums []int, queries [][]int) bool {
	// 3355
	//     1   3
	//   0   2
	// 0 1 2 2 1 prefixSum
	//   4 3 2 1 nums
	//   3 1 0 0
	prefixSum := make([]int, len(nums)+1)
	for i := 0; i < len(nums); i++ {
		prefixSum[i+1] = prefixSum[i] + nums[i]
	}
	fmt.Println(prefixSum)
	for _, q := range queries {
		l, r := q[0], q[1]
		if prefixSum[r+1]-prefixSum[l] != 0 {
			return false
		}
	}
	return true
}

func isZeroArray(nums []int, queries [][]int) bool {
	// 3355
	//     1   3
	//   0   2
	//   1 1 0-1-1 diffs
	// 0 1 2 2 1   prefixDiff
	//   4 3 2 1   nums
	//   3 1 0 0
	diffs := make([]int, len(nums)+1)
	for _, q := range queries {
		diffs[q[0]]++
		diffs[q[1]+1]--
	}
	prefixDiff := 0
	for i, num := range nums {
		prefixDiff += diffs[i]
		if num-prefixDiff > 0 {
			return false
		}
	}
	return true
}

func setZeroes_fail(matrix [][]int) {
	// 73
	for i, r := range matrix {
		for j, v := range r {
			if v == 0 {
				matrix[0][j] = 0
				matrix[i][0] = 0
			}
		}
	}
	fmt.Println(matrix)
	for i, r := range matrix[1:] {
		if r[0] != 0 {
			continue
		}
		for j := range r[1:] {
			fmt.Println(i, j+1)
			matrix[i][j+1] = 0
		}
	}
	for j, v := range matrix[0] {
		if v != 0 {
			continue
		}
		for i := range matrix[1:] {
			fmt.Println(i+1, j)
			matrix[i+1][j] = 0
		}
	}
	fmt.Println(matrix)
}

func setZeroes(matrix [][]int) {
	// 73
	rows, cols := len(matrix), len(matrix[0])
	rowZero, colZero := false, false

	// Determine if the first row or column should be zero
	for i := 0; i < rows; i++ {
		if matrix[i][0] == 0 {
			colZero = true
			break
		}
	}
	for j := 0; j < cols; j++ {
		if matrix[0][j] == 0 {
			rowZero = true
			break
		}
	}

	// Use first row and column to mark zeroes
	for i := 1; i < rows; i++ {
		for j := 1; j < cols; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}

	// Set matrix cells to zero based on markers
	for i := 1; i < rows; i++ {
		for j := 1; j < cols; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	// Handle the first row and column separately
	if colZero {
		for i := 0; i < rows; i++ {
			matrix[i][0] = 0
		}
	}
	if rowZero {
		for j := 0; j < cols; j++ {
			matrix[0][j] = 0
		}
	}
}

func maxRemoval(nums []int, queries [][]int) int {
	// 3362
	panic("not implemented")
}

func findWordsContaining(words []string, x byte) (indices []int) {
	// 2942
	for i, word := range words {
		for j := 0; j < len(word); j++ {
			if word[j] == x {
				indices = append(indices, i)
				break
			}
		}
	}
	return indices
}

func longestPalindrome(words []string) int {
	// 2131
	// ab ty yt lc cl ab aa aa bb
	// ba yt    cl    ba aa    bb cntByWord
	//        1     2        3    palindromePairCnt
	//                    1  0  1 sameCnt
	// tylcaabbaaclyt             2*sameCnt + 4*palindromePairCnt
	cntByWord := make(map[[2]byte]bool, len(words))
	palindromePairCnt, sameCnt := 0, 0
	for _, w := range words {
		first, second := w[0], w[1]
		if cntByWord[[2]byte{first, second}] {
			palindromePairCnt++
			if first == second {
				sameCnt--
			}
		} else {
			cntByWord[[2]byte{second, first}] = true
			if first == second {
				sameCnt++
			}
		}
		fmt.Println(w, cntByWord, sameCnt, palindromePairCnt)
	}
	if sameCnt > 1 {
		sameCnt = 1
	}
	return 2*sameCnt + 4*palindromePairCnt
}

func differenceOfSums(n int, m int) (rs int) {
	// 2894
	for num := 1; num <= n; num++ {
		if num%m != 0 {
			rs += num
		} else {
			rs -= num
		}
	}
	return rs
}

func removeDuplicates(nums []int) (newLen int) {
	// 26
	prev := nums[0]
	for _, num := range nums[1:] {
		if num == prev {
			continue
		}
		nums[newLen] = prev
		newLen++
		prev = num
	}
	nums[newLen] = prev
	nums = nums[:newLen+1]
	fmt.Println(nums)
	return newLen + 1
}

func main() {
	fmt.Println(removeDuplicates([]int{1, 1, 2}))
	// setZeroes([][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}})
	// fmt.Println(getLongestSubsequence([]string{"a", "b", "c", "d"}, []int{1, 0, 1, 1}))
	// fmt.Println(lengthAfterTransformations3337("abcyy", 2, []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2}))
	// fmt.Println(findEvenNumbers([]int{2, 1, 3, 0}))
	// fmt.Println(findEvenNumbers([]int{2, 2, 8, 8, 2}))
	// fmt.Println(minSum([]int{0, 16, 28, 12, 10, 15, 25, 24, 6, 0, 0}, []int{20, 15, 19, 5, 6, 29, 25, 8, 12}))
	// fmt.Println(minSum([]int{9, 5}, []int{15, 12, 5, 21, 4, 26, 27, 9, 6, 29, 0, 18, 16, 0, 0, 0, 20}))
	// fmt.Println(romanToInt("LVIII"), romanToInt("MCMXCIV"))
	// fmt.Println(numEquivDominoPairs([][]int{{1, 2}, {2, 1}, {1, 1}, {1, 2}, {2, 2}, {2, 2}}))
	// fmt.Println(minDominoRotations([]int{2, 1, 2, 4, 2, 2}, []int{5, 2, 6, 2, 3, 2}))
	// fmt.Println(minDominoRotations([]int{3, 5, 1, 2, 3}, []int{3, 6, 3, 3, 4}))
	// fmt.Println(pushDominoes(".L.R...LR..L.."))
	// fmt.Println(maxTaskAssign([]int{3, 2, 1}, []int{0, 3, 3}, 1, 1))
}
