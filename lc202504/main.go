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

func maximumTripletValue2(nums []int) int64 {
	// 2874
	// 7ms
	n := len(nums)
	prefixMax, suffixMax := make([]int, n), make([]int, n)
	prefixMax[0] = nums[0]
	for i := 1; i < n; i++ {
		prefixMax[i] = nums[i]
		if prefixMax[i] < prefixMax[i-1] {
			prefixMax[i] = prefixMax[i-1]
		}
	}
	suffixMax[n-1] = nums[n-1]
	for i := n - 2; i >= 0; i-- {
		suffixMax[i] = nums[i]
		if suffixMax[i] < suffixMax[i+1] {
			suffixMax[i] = suffixMax[i+1]
		}
	}
	var max int64 = math.MinInt64
	for j := 1; j < len(nums)-1; j++ {
		maxJ := int64(prefixMax[j-1]-nums[j]) * int64(suffixMax[j+1])
		if maxJ > max {
			max = maxJ
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

func main() {
	fmt.Println(subsetXORSum([]int{1, 3}))
	// fmt.Println(minimumSum2([]int{6, 5, 4, 3, 4, 5}))  // -1
	// fmt.Println(minimumSum2([]int{5, 4, 8, 7, 10, 2})) // 13
	// fmt.Println(arithmeticTriplets([]int{0, 1, 4, 6, 7, 10}, 3))
	// fmt.Println(maximumTripletValue([]int{12, 6, 1, 2, 7}))
	// fmt.Println(mostPoints([][]int{{3, 2}, {4, 3}, {4, 4}, {2, 5}}))
	// fmt.Println(climbStairs(3))
}
