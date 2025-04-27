package main

//lint:file-ignore U1000 Ignore all unused code, it's generated
import (
	"fmt"
	"maps"
	"math"
	"slices"
	"sort"
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

func minOperations2033(grid [][]int, x int) (rs int) {
	// 2033
	// 30ms
	arr, iNext := make([]int, len(grid)*len(grid[0])), 0
	remainder := grid[0][0] % x
	for _, r := range grid {
		for _, e := range r {
			if e%x != remainder {
				return -1
			}
			arr[iNext] = e
			iNext++
		}
	}
	slices.Sort(arr)
	median := arr[len(arr)/2]
	for _, num := range arr {
		diff := num - median
		if diff < 0 {
			diff = -diff
		}
		rs += diff / x
	}
	return rs
}

func countGoodTriplets(arr []int, a int, b int, c int) (cnt int) {
	// 1534
	// 6ms
	for i := 0; i < len(arr)-2; i++ {
		for j := i + 1; j < len(arr)-1; j++ {
			ij := arr[i] - arr[j]
			if ij <= a && ij >= -a {
				for k := j + 1; k < len(arr); k++ {
					jk, ik := arr[j]-arr[k], arr[i]-arr[k]
					if jk <= b && jk >= -b && ik <= c && ik >= -c {
						cnt++
					}
				}
			}

		}
	}
	return cnt
}

func countGoodNumbers(n int64) int {
	// 1922
	panic("Modular Exponentiation")
}

func goodTriplets_fail(nums1 []int, nums2 []int) int64 {
	// 2 0 1   3
	//   0 1 2 3
	//   x y   z
	currentM := make(map[int]struct{})
	nums1BeforeY := make([]map[int]struct{}, len(nums1))
	for i, num := range nums1 {
		nums1BeforeY[i] = maps.Clone(currentM)
		currentM[num] = struct{}{}
	}
	fmt.Println(nums1BeforeY)
	currentM = make(map[int]struct{})
	nums2BeforeY := make([]map[int]struct{}, len(nums2))
	for i, num := range nums2 {
		nums2BeforeY[i] = maps.Clone(currentM)
		currentM[num] = struct{}{}
	}
	fmt.Println(nums2BeforeY)
	currentM = make(map[int]struct{})
	nums1AfterY := make([]map[int]struct{}, len(nums1))
	for i := len(nums1) - 1; i > -1; i-- {
		nums1AfterY[i] = maps.Clone(currentM)
		currentM[nums1[i]] = struct{}{}
	}
	fmt.Println(nums1AfterY)
	currentM = make(map[int]struct{})
	nums2AfterY := make([]map[int]struct{}, len(nums2))
	for i := len(nums2) - 1; i > -1; i-- {
		nums2AfterY[i] = maps.Clone(currentM)
		currentM[nums2[i]] = struct{}{}
	}
	fmt.Println(nums2AfterY)
	cnt := int64(0)
	for i := range nums1 {
		for numBefore := range nums1BeforeY[i] {
			if _, exists := nums2BeforeY[i][numBefore]; !exists {
				continue
			}
			for numAfter := range nums1AfterY[i] {
				if _, exists := nums2AfterY[i][numAfter]; !exists {
					continue
				}
				cnt++
				fmt.Println(nums1[i])
			}
		}
	}
	return cnt
}

func goodTriplets_solution(nums1 []int, nums2 []int) int64 {
	// not implemented 2179
	indexMap := make(map[int]int)
	for i, num := range nums2 {
		indexMap[num] = i
	}

	bit := make([]int, len(nums1)+1)
	update := func(idx, val int) {
		for idx < len(bit) {
			bit[idx] += val
			idx += idx & -idx
		}
	}
	query := func(idx int) int {
		sum := 0
		for idx > 0 {
			sum += bit[idx]
			idx -= idx & -idx
		}
		return sum
	}

	cnt := int64(0)
	for _, num := range nums1 {
		pos := indexMap[num] + 1
		leftCount := query(pos - 1)
		rightCount := int64(query(len(nums1)) - query(pos))
		cnt += int64(leftCount) * rightCount
		update(pos, 1)
	}
	return cnt
}

func countGood(nums []int, k int) (rs int64) {
	// 2537
	// 38ms
	// 1 2 2 1                   k = 2
	// 1       [1]=1       sum=0
	//   2     [1]=1 [2]=1 sum=0
	//     2   [1]=1 [2]=2 sum=1
	//       1 [1]=2 [2]=2 sum=2
	left, sum := 0, 0
	cntByNum := make(map[int]int)
	for _, num := range nums {
		sum += cntByNum[num]
		cntByNum[num]++
		for sum-cntByNum[nums[left]]+1 >= k {
			cntByNum[nums[left]] -= 1
			sum -= cntByNum[nums[left]]
			left++
		}
		if sum >= k {
			rs += int64(left + 1)
		}
	}
	return rs
}

func countPairs(nums []int, k int) (cnt int) {
	// 2176
	iByNum := map[int][]int{}
	for j := 0; j < len(nums); j++ {
		if indices, exists := iByNum[nums[j]]; exists {
			for _, i := range indices {
				if (i*j)%k == 0 {
					cnt++
				}
			}
		}
		iByNum[nums[j]] = append(iByNum[nums[j]], j)
	}
	return cnt
}

func freqFromIntStr(s string) (rs [][]byte) {
	cnt, prev := byte(1), s[0]
	for i := 1; i < len(s); i++ {
		c := s[i]
		if c == prev {
			cnt++
			continue
		}
		rs = append(rs, []byte{prev - '0', cnt})
		cnt = byte(1)
		prev = c
	}
	rs = append(rs, []byte{prev - '0', cnt})
	return rs
}

func intStrFromFreq(freq [][]byte) string {
	s := ""
	for _, pair := range freq {
		s += fmt.Sprintf("%d%d", pair[1], pair[0])
	}
	return s
}

func countAndSay(n int) string {
	// 38
	// 15ms
	if n == 1 {
		return "1"
	}
	current := "1"
	for i := 1; i < n; i++ {
		freq := freqFromIntStr(current)
		current = intStrFromFreq(freq)
	}
	return current
}

func countFairPairs_time(nums []int, lower int, upper int) (cnt int64) {
	// 0,1,4,4,5,7  lower 3 upper 6
	// 0 4,4,5
	// 1 4,4,5
	slices.Sort(nums)
	for i, numi := range nums {
		// if numi > upper {
		// 	break
		// }
		for j, numj := range nums[i+1:] {
			sum := numi + numj
			fmt.Println(i, j)
			// if sum > upper {
			// 	break
			// }
			if sum >= lower && sum <= upper {
				cnt++
			}
		}
	}
	return cnt
}

func countFairPairs(nums []int, lower int, upper int) (cnt int64) {
	// 2563
	// 63ms
	slices.Sort(nums)
	for i := 0; i < len(nums); i++ {
		low := lower - nums[i]
		high := upper - nums[i]
		left := sort.Search(len(nums)-i-1, func(j int) bool { return nums[i+1+j] >= low })
		right := sort.Search(len(nums)-i-1, func(j int) bool { return nums[i+1+j] > high })
		cnt += int64(right - left)
	}
	return cnt
}

func numRabbits(answers []int) (smallest int) {
	// 781
	// 1,1,2     [1]=2 [2]=1
	// 10,10,10  [10]=3
	// 1,0,1,0,0 [1]=2 [0]=3
	// 0,0,1,1,1 [0]=2 [1]=3
	mapAnswerCount := make(map[int]int)
	for _, otherCount := range answers {
		mapAnswerCount[otherCount]++
	}
	for otherCount, answerCount := range mapAnswerCount {
		groupSize := otherCount + 1
		groups := (answerCount + groupSize - 1) / groupSize // Calculate the number of groups
		smallest += groups * groupSize
	}
	return smallest
}

func numberOfArrays(differences []int, lower int, upper int) int {
	// 2145
	// 0 1 -3 4
	current, mn, mx := 0, 0, 0
	for _, d := range differences {
		current += d
		// fmt.Print(" ", current)
		if mn > current {
			mn = current
		}
		if mx < current {
			mx = current
		}
	}
	rs := upper - lower - (mx - mn) + 1
	if rs < 0 {
		return 0
	}
	return rs
}

func idealArrays_fail(n int, maxValue int) int {
	// fail 2338
	memoi := make(map[[2]int]int)
	var recur func(int, int) int
	recur = func(length int, maxVal int) int {
		if length == 1 {
			return maxVal
		}
		if result, exists := memoi[[2]int{length, maxVal}]; exists {
			return result
		}
		count := 0
		for i := 1; i <= maxVal; i++ {
			count = (count + recur(length-1, i)) % 1_000_000_007
		}
		memoi[[2]int{length, maxVal}] = count
		return count
	}
	var result int
	for length := 1; length <= n; length++ {
		result = (result + recur(length, maxValue)) % 1_000_000_007
	}
	return result
}

func countLargestGroup_fail(n int) int {
	// 1399
	// 13 [1,10], [2,11], [3,12], [4,13], [5], [6], [7], [8], [9]
	// 46 1,10  2,11,20  3,12,21,30  4,13,22,31,40  5,14,23,32,41  6,15,24,33,42  7,16,25,34,43  8,17,26,35,44  9,18,27,36,45  19,28,37,46  29,38  39
	// +9 to get the next number in group if still <= n
	mapCnt := [36]int{} // 10^4 --> max 9+9+9+9
	for i := 0; i < 36; i++ {
		num, cnt := i+1, 0
		for num <= n {
			cnt++
			num += 9
		}
		mapCnt[i] = cnt
	}
	fmt.Println(mapCnt)
	return 0
}

func countLargestGroup(n int) (rs int) {
	// 1399
	maxCnt, mapCnt := 0, make(map[int]int)
	for i := 1; i <= n; i++ {
		sum := 0
		for tmp := i; tmp > 0; {
			sum += tmp % 10
			tmp /= 10
		}
		mapCnt[sum]++
		if mapCnt[sum] > maxCnt {
			maxCnt = mapCnt[sum]
		}
	}
	for _, count := range mapCnt {
		if count == maxCnt {
			rs++
		}
	}
	return rs
}

func countCompleteSubarrays(nums []int) (cnt int) {
	// 2799
	mapDistinct := make(map[int]struct{})
	for _, num := range nums {
		mapDistinct[num] = struct{}{}
	}
	totalDistinct := len(mapDistinct)
	// Sliding window to count subarrays with all distinct elements
	distinctCount := make(map[int]int)
	left := 0
	for right := 0; right < len(nums); right++ {
		distinctCount[nums[right]]++
		for len(distinctCount) == totalDistinct {
			cnt += len(nums) - right
			distinctCount[nums[left]]--
			if distinctCount[nums[left]] == 0 {
				delete(distinctCount, nums[left])
			}
			left++
		}
	}
	return cnt
}

func countInterestingSubarrays(nums []int, modulo int, k int) (rs int64) {
	// 2845
	//   3 1 9 6
	// 0 1 1 2 3  prefixIndexCount
	prefixIndexCount := make([]int, len(nums)+1)
	for i, num := range nums {
		prefixIndexCount[i+1] = prefixIndexCount[i]
		if num%modulo == k {
			prefixIndexCount[i+1]++
		}
	}
	countByRemainder := map[int]int{0: 1}
	countByRemainder[0] = 1
	for _, count := range prefixIndexCount[1:] {
		target := (count - k + modulo) % modulo
		rs += int64(countByRemainder[target])
		countByRemainder[count%modulo]++
	}
	return rs
}

func countSubarrays(nums []int, minK int, maxK int) (rs int64) {
	// 2444
	subNums, start := [][]int{}, -1
	for i, num := range nums {
		if num < minK || num > maxK {
			if start != -1 {
				subNums = append(subNums, nums[start:i])
				start = -1
			}
		} else if start == -1 {
			start = i
		}
	}
	if start != -1 {
		subNums = append(subNums, nums[start:])
	}
	fmt.Println(subNums)
	for _, subArray := range subNums {
		minIndex, maxIndex := -1, -1
		for i := 0; i < len(subArray); i++ {
			if subArray[i] == minK {
				minIndex = i
			}
			if subArray[i] == maxK {
				maxIndex = i
			}
			leftBound := minIndex
			if maxIndex < minIndex {
				leftBound = maxIndex
			}
			rs += int64(leftBound + 1)
		}
	}
	return rs
}

func countSubarrays3392(nums []int) (cnt int) {
	// 3392
	first, mid := nums[0], nums[1]
	for _, third := range nums[2:] {
		if 2*(first+third) == mid {
			cnt++
		}
		first, mid = mid, third
	}
	return cnt
}

func main() {
	fmt.Println(countSubarrays([]int{4, 3}, 3, 3))
	fmt.Println(countSubarrays([]int{1, 3, 5, 2, 7, 5}, 1, 5))
	// fmt.Println(countInterestingSubarrays([]int{3, 1, 9, 6}, 3, 0))
	// fmt.Println(countLargestGroup(46))
	// fmt.Println(numberOfArrays([]int{-40}, -46, 53))
	// fmt.Println(numberOfArrays([]int{1, -3, 4}, 1, 6))
	// fmt.Println(countFairPairs([]int{-5, -7, -5, -7, -5}, -12, -12))
	// fmt.Println(countAndSay(4))
	// fmt.Println(freqFromInt(223314444411))
	// fmt.Println(intFromFreq(freqFromInt(223314444411)))
	// fmt.Println(goodTriplets_solution([]int{2, 0, 1, 3}, []int{0, 1, 2, 3}))
	// fmt.Println(countGoodNumbers(806166225460393))
	// fmt.Println(minOperations2033([][]int{{2, 4}, {6, 8}}, 2))
	// fmt.Println(largestDivisibleSubset([]int{1, 2, 3, 4, 9, 81}))
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
