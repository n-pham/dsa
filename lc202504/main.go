package main

//lint:file-ignore U1000 Ignore all unused code, it's generated
import (
	"fmt"
	"math"
	"slices"
	"strconv"
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
		m[index] = max(solve0, skip0)
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
	var bigger int64 = math.MinInt64
	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			for k := j + 1; k < len(nums); k++ {
				bigger = max(bigger, int64(nums[i]-nums[j])*int64(nums[k]))
			}
		}
	}
	return max(0, bigger)
}

func maximumTripletValue2(nums []int) int64 {
	// 2874
	// 6ms
	n := len(nums)
	prefixMax, suffixMax := make([]int, n), make([]int, n)
	prefixMax[0] = nums[0]
	for i := 1; i < n; i++ {
		prefixMax[i] = max(nums[i], prefixMax[i-1])
	}
	suffixMax[n-1] = nums[n-1]
	for i := n - 2; i >= 0; i-- {
		suffixMax[i] = max(nums[i], suffixMax[i+1])
	}
	var bigger int64 = math.MinInt64
	for j := 1; j < len(nums)-1; j++ {
		bigger = max(bigger, int64(prefixMax[j-1]-nums[j])*int64(suffixMax[j+1]))
	}
	return max(0, bigger)
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

func minimumSum2(nums []int) int {
	// 2909
	n := len(nums)
	prefixMin, suffixMin := make([]int, n), make([]int, n)
	prefixMin[0] = nums[0]
	for i := 1; i < n; i++ {
		prefixMin[i] = nums[i]
		if prefixMin[i] > prefixMin[i-1] {
			prefixMin[i] = prefixMin[i-1]
		}
	}
	suffixMin[n-1] = nums[n-1]
	for i := n - 2; i >= 0; i-- {
		suffixMin[i] = nums[i]
		if suffixMin[i] > suffixMin[i+1] {
			suffixMin[i] = suffixMin[i+1]
		}
	}
	fmt.Println(prefixMin)
	fmt.Println(nums)
	fmt.Println(suffixMin)
	rs := math.MaxInt
	for j := 1; j < n-1; j++ {
		if prefixMin[j-1] < nums[j] && nums[j] > suffixMin[j+1] {
			tmp := prefixMin[j-1] + nums[j] + suffixMin[j+1]
			if rs > tmp {
				rs = tmp
			}
		}
	}
	if rs == math.MaxInt {
		return -1
	}
	return rs
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	// 226
	if root == nil {
		return nil
	}
	tmp := root.Left
	root.Left = root.Right
	root.Right = tmp
	invertTree(root.Left)
	invertTree(root.Right)
	return root
}

func maxDepth(root *TreeNode) int {
	// 104
	if root == nil {
		return 0
	}
	leftDepth := maxDepth(root.Left)
	rightDepth := maxDepth(root.Right)
	if leftDepth > rightDepth {
		return leftDepth + 1
	}
	return rightDepth + 1
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	// 100
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	return p.Val == q.Val && isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

func subtreeWithAllDeepest(root *TreeNode) *TreeNode {
	// 865
	var recur func(*TreeNode) (*TreeNode, int)
	recur = func(root *TreeNode) (*TreeNode, int) {
		if root == nil {
			return nil, 0
		}
		leftCandidate, leftDepth := recur(root.Left)
		rightCandidate, rightDepth := recur(root.Right)
		if leftDepth == rightDepth {
			return root, leftDepth + 1
		}
		if leftDepth > rightDepth {
			return leftCandidate, leftDepth + 1
		}
		return rightCandidate, rightDepth + 1
	}
	node, _ := recur(root)
	return node
}

func lcaDeepestLeaves(root *TreeNode) *TreeNode {
	// 1123
	// terrible explanation, duplicated with subtreeWithAllDeepest
	var recur func(*TreeNode) (*TreeNode, int)
	recur = func(root *TreeNode) (*TreeNode, int) {
		if root == nil {
			return nil, 0
		}
		leftCandidate, leftDepth := recur(root.Left)
		rightCandidate, rightDepth := recur(root.Right)
		if leftDepth == rightDepth {
			return root, leftDepth + 1
		}
		if leftDepth > rightDepth {
			return leftCandidate, leftDepth + 1
		}
		return rightCandidate, rightDepth + 1
	}
	node, _ := recur(root)
	return node
}

func subsetXORSum(nums []int) (rs int) {
	// 1863
	// 1, 3  --> 0 xor 1 + 1 xor 3 + 0 xor 3
	var recur func(int, int)
	recur = func(index int, prev_xor int) {
		fmt.Println(index, prev_xor)
		rs += prev_xor
		for i := index; i < len(nums); i++ { // backtrack
			next_xor := prev_xor ^ nums[i]
			recur(i+1, next_xor)
		}
	}
	recur(0, 0)
	return rs
}

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	// 572
	if root == nil {
		return false
	}
	return isSameTree(root, subRoot) || isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}

func largestDivisibleSubset(nums []int) []int {
	// 368
	// 1 2 3 4 9 81
	// 1036ms
	slices.Sort(nums)
	memoi := make(map[[2]int][]int)
	var recur func(int, int) []int
	recur = func(i int, j int) []int {
		if j == len(nums) {
			return []int{}
		}
		if result, exists := memoi[[2]int{i, j}]; exists {
			return result
		}
		rs := recur(i, j+1) // skip j
		if i == -1 || nums[j]%nums[i] == 0 {
			includeJ := append([]int{nums[j]}, recur(j, j+1)...)
			if len(includeJ) > len(rs) {
				rs = includeJ
			}
		}
		memoi[[2]int{i, j}] = rs
		return rs
	}
	t := recur(-1, 0)
	fmt.Println(memoi)
	return t
}

func levelOrder_notsogreat(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	rs := [][]int{{}}
	rsLen := 0
	currQ, nextQ := []*TreeNode{root}, []*TreeNode{}
	for len(currQ) > 0 {
		node := currQ[0]
		currQ = currQ[1:]
		rs[rsLen] = append(rs[rsLen], node.Val)
		if node.Left != nil {
			nextQ = append(nextQ, node.Left)
		}
		if node.Right != nil {
			nextQ = append(nextQ, node.Right)
		}
		if len(currQ) == 0 {
			currQ = nextQ
			nextQ = []*TreeNode{}
			rs = append(rs, []int{})
			rsLen++
		}
	}
	// Remove the last empty slice if it exists
	if len(rs[len(rs)-1]) == 0 {
		rs = rs[:len(rs)-1]
	}
	return rs
}

func levelOrder(root *TreeNode) [][]int {
	// 102
	if root == nil {
		return [][]int{}
	}
	rs := [][]int{}
	currQ, nextQ := []*TreeNode{root}, []*TreeNode{}
	for len(currQ) > 0 {
		level := []int{}
		for _, node := range currQ {
			level = append(level, node.Val)
			if node.Left != nil {
				nextQ = append(nextQ, node.Left)
			}
			if node.Right != nil {
				nextQ = append(nextQ, node.Right)
			}
		}
		rs = append(rs, level)
		currQ, nextQ = nextQ, []*TreeNode{}
	}
	return rs
}

func isValidBST(root *TreeNode) bool {
	// 98
	var recur func(node *TreeNode, min, max int) bool
	recur = func(node *TreeNode, min, max int) bool {
		if node == nil {
			return true
		}
		if node.Val <= min || node.Val >= max {
			return false
		}
		return recur(node.Left, min, node.Val) && recur(node.Right, node.Val, max)
	}
	return recur(root, math.MinInt, math.MaxInt)
}

func canPartition_fail(nums []int) bool {
	slices.Sort(nums)
	n := len(nums)
	prefixSum := make([]int, n+1)
	for i, num := range nums {
		prefixSum[i+1] = prefixSum[i] + num
	}
	fmt.Println(prefixSum)
	if prefixSum[n]%2 == 1 {
		return false
	}
	halfSum := prefixSum[n] / 2
	for i := n - 1; i > 0; i-- {
		fmt.Println(prefixSum[i], halfSum)
		if prefixSum[i] == halfSum {
			return true
		}
	}
	return false
}

func canPartition(nums []int) bool {
	// 416
	panic("not implemented")
}

func runningSum(nums []int) []int {
	// 1480
	rs := make([]int, len(nums)+1)
	for i, num := range nums {
		rs[i+1] = rs[i] + num
	}
	return rs[1:]
}

func maximumWealth(accounts [][]int) int {
	// 1672
	biggest := math.MinInt
	for _, amounts := range accounts {
		wealth := 0
		for _, amount := range amounts {
			wealth += amount
		}
		if wealth > biggest {
			biggest = wealth
		}
	}
	return biggest
}

func minimumOperations(nums []int) int {
	// 3396
	// 1,2,3,4,2,3,3,5,7
	// 3 3 2 1 1 1 0 0 0 suffixDup is overkill
	n := len(nums)
	iLastDup := 0
	iByNum := make(map[int]int)
	iByNum[nums[n-1]] = n - 1
	for i := n - 2; i >= 0; i-- {
		fmt.Println(i, iByNum)
		if iByNum[nums[i]] > 0 {
			iLastDup = i
			break
		}
		iByNum[nums[i]] = i
	}
	fmt.Println(nums[iLastDup], iLastDup, iByNum[nums[iLastDup]])
	if iLastDup == 0 && iByNum[nums[iLastDup]] == 0 {
		return 0
	}
	return (iLastDup + 3) / 3
}

func minOperations_11ms(nums []int, k int) int {
	m := make(map[int]struct{})
	for _, num := range nums {
		if num < k {
			return -1
		}
		m[num] = struct{}{}
	}
	if _, found := m[k]; found {
		return len(m) - 1
	}
	return len(m)
}

func minOperations(nums []int, k int) int {
	// 3375
	slices.Sort(nums)
	if nums[0] < k {
		return -1
	}
	rs := 0
	for i := len(nums) - 1; i >= 1; i-- {
		if nums[i] != nums[i-1] {
			rs++
		}
	}
	if nums[0] > k {
		rs++
	}
	return rs
}

func numberOfPowerfulInt_memory(start int64, finish int64, limit int, s string) int64 {
	// 2999
	generateNumers := func(digitCnt, limit int) []int {
		var result []int

		// Start from 0, iteratively generate numbers of length 'digitCnt'
		start := 0
		end := 1
		for i := 0; i < digitCnt; i++ {
			start = start * 10 // Expand range start
			end = end * 10     // Expand range end
		}

		// Generate numbers within the range, respecting the limit
		for i := start; i < end; i++ {
			valid := true
			num := i
			for num > 0 {
				digit := num % 10
				if digit > limit {
					valid = false
					break
				}
				num /= 10
			}
			if valid {
				result = append(result, i)
			}
		}

		return result
	}

	IntPow := func(base, exp int) int64 {
		result := int64(1)
		for {
			if exp&1 == 1 {
				result *= int64(base)
			}
			exp >>= 1
			if exp == 0 {
				break
			}
			base *= base
		}
		return result
	}

	sInt, _ := strconv.Atoi(s)
	sLen := 0
	for t := sInt; t > 0; t = t / 10 {
		sLen++
	}
	fLen := 0
	for t := finish; t > 0; t = t / 10 {
		fLen++
	}
	digitCnt := fLen - sLen
	nums := generateNumers(digitCnt, limit)
	cnt := int64(0)
	for _, num := range nums {
		t := int64(num)*IntPow(10, sLen) + int64(sInt)
		if t >= start && t <= finish {
			cnt++
		}
	}
	return cnt
}

func numberOfPowerfulInt(start int64, finish int64, limit int, s string) int64 {
	// 2999
	panic("not implemented")
}

func countSymmetricIntegers(low int, high int) (rs int) {
	// 2843
	// 4ms
	for i := low; i <= high; {
		length := 0
		for t := i; t > 0; t /= 10 {
			length++
		}
		if length%2 == 0 {
			t, rightSum, leftSum := i, 0, 0
			for j := 0; j < length/2; j++ {
				rightSum += t % 10
				t /= 10
			}
			for j := 0; j < length/2; j++ {
				leftSum += t % 10
				t /= 10
			}
			if leftSum == rightSum {
				rs++
				i += 9 // math: next 8 numbers will not be equal
				continue
			}
		}
		i++
	}
	return rs
}

func main() {
	fmt.Println(largestDivisibleSubset([]int{1, 2, 3, 4, 9, 81}))
	// fmt.Println(numberOfPowerfulInt(1, 6000, 4, "124"))
	// fmt.Println(minimumOperations([]int{5, 5}))
	// fmt.Println(minimumOperations([]int{6, 7, 8, 9}))
	// fmt.Println(minimumOperations([]int{1, 2, 3, 4, 2, 3, 3, 5, 7}))
	// fmt.Println(canPartition([]int{2, 2, 1, 1}))
	// fmt.Println(subsetXORSum([]int{1, 3}))
	// fmt.Println(minimumSum2([]int{6, 5, 4, 3, 4, 5}))  // -1
	// fmt.Println(minimumSum2([]int{5, 4, 8, 7, 10, 2})) // 13
	// fmt.Println(arithmeticTriplets([]int{0, 1, 4, 6, 7, 10}, 3))
	// fmt.Println(maximumTripletValue([]int{12, 6, 1, 2, 7}))
	// fmt.Println(mostPoints([][]int{{3, 2}, {4, 3}, {4, 4}, {2, 5}}))
	// fmt.Println(climbStairs(3))
}
