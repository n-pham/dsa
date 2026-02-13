package main

import (
	"container/heap"
	"dsa/kit"
	"math"
	"slices"
)

// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
func (h *IntHeap) Peek() int {
	if h.Len() == 0 {
		return math.MaxInt64 // Sentinel for empty min-heap
	}
	return (*h)[0]
}

// A MaxIntHeap is a max-heap of ints.
type MaxIntHeap struct{ IntHeap }

func (h MaxIntHeap) Less(i, j int) bool { return h.IntHeap[i] > h.IntHeap[j] }
func (h *MaxIntHeap) Peek() int {
	if h.Len() == 0 {
		return math.MinInt64 // Sentinel for empty max-heap
	}
	return (*h).IntHeap[0]
}

// LazyHeap wraps a standard heap.Interface to provide lazy deletion.
type LazyHeap struct {
	h            heap.Interface
	removed      map[int]int // Count of elements logically removed
	physicalLen  int         // Actual number of elements in the underlying heap
	removedCount int         // Number of elements marked for removal (sum of values in 'removed' map)
}

func newLazyMinHeap() *LazyHeap {
	h := &IntHeap{}
	heap.Init(h)
	return &LazyHeap{h: h, removed: make(map[int]int)}
}

func newLazyMaxHeap() *LazyHeap {
	h := &MaxIntHeap{}
	heap.Init(h)
	return &LazyHeap{h: h, removed: make(map[int]int)}
}

// Push adds an element to the heap.
func (lh *LazyHeap) Push(x int) {
	heap.Push(lh.h, x)
	lh.physicalLen++
}

// Pop removes and returns the smallest/largest element, skipping logically removed ones.
func (lh *LazyHeap) Pop() int {
	lh.clean()
	val := heap.Pop(lh.h).(int)
	lh.physicalLen--
	return val
}

// Peek returns the smallest/largest element without removing it, skipping logically removed ones.
func (lh *LazyHeap) Peek() int {
	lh.clean()
	var val int
	switch hT := lh.h.(type) {
	case *IntHeap:
		val = hT.Peek()
	case *MaxIntHeap:
		val = hT.Peek()
	default:
		panic("unknown heap type")
	}
	return val
}

// Remove marks an element for lazy deletion.
func (lh *LazyHeap) Remove(x int) {
	lh.removed[x]++
	lh.removedCount++
}

// Len returns the logical size of the heap (number of non-removed elements).
func (lh *LazyHeap) Len() int {
	return lh.physicalLen - lh.removedCount
}

// clean removes elements from the physical heap that are marked for lazy deletion
// until the true top element is found.
func (lh *LazyHeap) clean() {
	for lh.physicalLen > 0 { // Check physical length
		var topVal int
		switch hT := lh.h.(type) {
		case *IntHeap:
			topVal = hT.Peek()
		case *MaxIntHeap:
			topVal = hT.Peek()
		default:
			panic("unknown heap type")
		}

		if lh.removed[topVal] > 0 {
			heap.Pop(lh.h)
			lh.removed[topVal]--
			lh.physicalLen--
			lh.removedCount--
		} else {
			break
		}
	}
}

func minimumCost3013(nums []int, k int, dist int) int64 {

	small := newLazyMaxHeap() // Max-heap for k-1 smallest elements

	large := newLazyMinHeap() // Min-heap for the rest

	var smallSum int64

	minSum := int64(math.MaxInt64)

	m := k - 1 // Number of additional elements to choose (excluding nums[0])

	if m == 0 {

		return int64(nums[0])

	}

	// We need to choose m elements from a window of size dist+1. This is only possible if m <= dist+1.

	if m > dist+1 {

		return int64(math.MaxInt64)

	}

	if len(nums)-1 < m { // Not enough elements in nums[1:] to choose m elements

		return int64(math.MaxInt64)

	}

	subNums := nums[1:] // Elements to consider for i_2 to i_k

	for i := 0; i < len(subNums); i++ {

		val := subNums[i]

		// Add new element 'val'

		if small.Len() < m {

			small.Push(val)

			smallSum += int64(val)

		} else if val < small.Peek() {

			oldSmallTop := small.Pop()

			smallSum -= int64(oldSmallTop)

			large.Push(oldSmallTop)

			small.Push(val)

			smallSum += int64(val)

		} else {

			large.Push(val)

		}

		// Rebalance heaps if needed (after adding)

		for small.Len() > m {

			topSmall := small.Pop()

			smallSum -= int64(topSmall)

			large.Push(topSmall)

		}

		for small.Len() < m && large.Len() > 0 {

			topLarge := large.Pop()

			small.Push(topLarge)

			smallSum += int64(topLarge)

		}

		// A valid window has size dist+1. The first such window is subNums[0...dist],

		// which is formed at loop iteration i = dist.

		if i >= dist {

			// At this point, the heaps contain elements from the window subNums[i-dist...i].

			// This window has size dist+1.

			if small.Len() == m {

				if smallSum < minSum {

					minSum = smallSum

				}

			}

			// Now, remove the element that's sliding out of the window

			toRemove := subNums[i-dist]

			if toRemove <= small.Peek() {

				small.Remove(toRemove)

				smallSum -= int64(toRemove)

			} else {

				large.Remove(toRemove)

			}

			// Rebalance after removal

			for small.Len() > m {

				topSmall := small.Pop()

				smallSum -= int64(topSmall)

				large.Push(topSmall)

			}

			for small.Len() < m && large.Len() > 0 {

				topLarge := large.Pop()

				small.Push(topLarge)

				smallSum += int64(topLarge)

			}

		}

	}

	if minSum == math.MaxInt64 { // If no valid combination found

		// This can happen if len(subNums) < dist+1, so the main check `i >= dist` is never met.

		// In this case, we need to calculate the sum of the smallest `m` elements from all of `subNums`.

		// The heaps at the end of the loop will contain exactly that.

		if small.Len() == m {

			minSum = smallSum

		} else {

			return math.MaxInt64 // Should not happen given initial checks

		}

	}

	return int64(nums[0]) + minSum

}

func IsTrionic(nums []int) bool {
	// 3637
	n := len(nums)
	i := 0
	for ; i < n-1 && nums[i] < nums[i+1]; i++ {
	}
	p := i
	kit.DebugLog("p ", p)
	if p == 0 || p == n-2 {
		return false
	}
	for ; i < n-1 && nums[i] > nums[i+1]; i++ {
	}
	q := i
	kit.DebugLog("q ", q)
	if q == p || q == n-1 {
		return false
	}
	for ; i < n-1; i++ {
		if nums[i] >= nums[i+1] {
			return false
		}
	}
	return true
}

func MaxSumTrionic(nums []int) int64 {
	// 3640
	n := len(nums)
	// DP arrays
	leftIncGe2 := make([]int64, n)
	leftIncDec := make([]int64, n)

	for i := 0; i < n; i++ {
		leftIncGe2[i] = math.MinInt64
		leftIncDec[i] = math.MinInt64
	}

	// Left-to-right pass
	for i := 1; i < n; i++ {
		// max sum of strictly increasing subarray of length >= 2 ending at i
		if nums[i] > nums[i-1] {
			sumLen2 := int64(nums[i-1]) + int64(nums[i])
			sumExtend := int64(math.MinInt64)
			if i > 1 && leftIncGe2[i-1] != math.MinInt64 {
				sumExtend = leftIncGe2[i-1] + int64(nums[i])
			}
			if sumLen2 > sumExtend {
				leftIncGe2[i] = sumLen2
			} else {
				leftIncGe2[i] = sumExtend
			}
		}

		// max sum of inc-dec subarray ending at i
		if nums[i] < nums[i-1] {
			cand1 := int64(math.MinInt64)
			if i > 1 && leftIncDec[i-1] != math.MinInt64 {
				cand1 = leftIncDec[i-1] + int64(nums[i])
			}
			cand2 := int64(math.MinInt64)
			if i > 1 && leftIncGe2[i-1] != math.MinInt64 {
				cand2 = leftIncGe2[i-1] + int64(nums[i])
			}

			if cand1 > cand2 {
				leftIncDec[i] = cand1
			} else if cand2 != math.MinInt64 {
				leftIncDec[i] = cand2
			}
		}
	}

	// Right-to-left pass
	rightIncGe2 := make([]int64, n)
	for i := 0; i < n; i++ {
		rightIncGe2[i] = math.MinInt64
	}
	for i := n - 2; i >= 0; i-- {
		// max sum of strictly increasing subarray of length >= 2 starting at i
		if nums[i] < nums[i+1] {
			sumLen2 := int64(nums[i]) + int64(nums[i+1])
			sumExtend := int64(math.MinInt64)
			if i < n-2 && rightIncGe2[i+1] != math.MinInt64 {
				sumExtend = rightIncGe2[i+1] + int64(nums[i])
			}
			if sumLen2 > sumExtend {
				rightIncGe2[i] = sumLen2
			} else {
				rightIncGe2[i] = sumExtend
			}
		}
	}

	maxSum := int64(0)
	found := false
	for q := 0; q < n; q++ {
		if q > 0 && q < n-1 && leftIncDec[q] != math.MinInt64 && rightIncGe2[q] != math.MinInt64 {
			currentSum := leftIncDec[q] + rightIncGe2[q] - int64(nums[q])
			if !found || currentSum > maxSum {
				maxSum = currentSum
				found = true
			}
		}
	}

	return maxSum
}

func ConstructTransformedArray(nums []int) []int {
	// 3379
	n := len(nums)
	newNums := make([]int, n)
	for i, num := range nums {
		newIndex := (i + num) % n
		if newIndex < 0 {
			newIndex += n
		}
		newNums[i] = nums[newIndex]
	}
	return newNums
}

func MinimumCost(source string, target string, original []string, changed []string, cost []int) int64 {
	if len(source) != len(target) {
		return -1 // Impossible if lengths are different
	}

	// 1. Pre-computation (Floyd-Warshall)
	stringToID := make(map[string]int)
	var idCount int

	getOrCreateID := func(s string) int {
		if id, ok := stringToID[s]; ok {
			return id
		}
		stringToID[s] = idCount
		idCount++
		return idCount - 1
	}

	// Collect all unique strings and assign IDs
	for i := 0; i < len(original); i++ {
		getOrCreateID(original[i])
		getOrCreateID(changed[i])
	}

	// Initialize dist matrix
	dist := make([][]int64, idCount)
	for i := range dist {
		dist[i] = make([]int64, idCount)
		for j := range dist[i] {
			if i == j {
				dist[i][j] = 0
			} else {
				dist[i][j] = math.MaxInt64
			}
		}
	}

	// Populate direct conversion costs
	for i := 0; i < len(original); i++ {
		u := getOrCreateID(original[i])
		v := getOrCreateID(changed[i])
		if dist[u][v] > int64(cost[i]) {
			dist[u][v] = int64(cost[i])
		}
	}

	// Floyd-Warshall
	for k := 0; k < idCount; k++ {
		for i := 0; i < idCount; i++ {
			for j := 0; j < idCount; j++ {
				if dist[i][k] != math.MaxInt64 && dist[k][j] != math.MaxInt64 {
					if dist[i][j] > dist[i][k]+dist[k][j] {
						dist[i][j] = dist[i][k] + dist[k][j]
					}
				}
			}
		}
	}

	// 2. Dynamic Programming
	N := len(source)
	dp := make([]int64, N+1)
	dp[0] = 0
	for i := 1; i <= N; i++ {
		dp[i] = math.MaxInt64
	}

	// Collect all unique lengths of original/changed strings for efficient iteration
	lengths := make(map[int]struct{})
	for _, s := range original {
		lengths[len(s)] = struct{}{}
	}

	for i := 0; i < N; i++ {
		if dp[i] == math.MaxInt64 {
			continue
		}

		if source[i] == target[i] {
			if dp[i+1] > dp[i] {
				dp[i+1] = dp[i]
			}
		}

		for L := range lengths {
			if i+L <= N {
				subSource := source[i : i+L]
				subTarget := target[i : i+L]

				u, ok1 := stringToID[subSource]
				v, ok2 := stringToID[subTarget]

				if ok1 && ok2 {
					conversionCost := dist[u][v]
					if conversionCost != math.MaxInt64 {
						if dp[i+L] > dp[i]+conversionCost {
							dp[i+L] = dp[i] + conversionCost
						}
					}
				}
			}
		}
	}

	if dp[N] == math.MaxInt64 {
		return -1
	}
	return dp[N]
}

func MinRemoval(nums []int, k int) int {
	// 3634
	slices.Sort(nums)
	n := len(nums)
	maxLen := 0
	j := 0
	// i is the left pointer (minimum element of the window)
	// j is the right pointer (potential maximum element of the window)
	for i := 0; i < n; i++ {
		// Expand j while the balanced condition holds: max <= k * min
		for j < n && int64(nums[j]) <= int64(k)*int64(nums[i]) {
			j++
		}
		// Current balanced window is nums[i...j-1], length is j - i
		if j-i > maxLen {
			maxLen = j - i
		}
	}
	return n - maxLen
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
    var getBalancedAndDepth func(*TreeNode) (bool, int)
    getBalancedAndDepth = func(root *TreeNode) (bool, int) {
        if root == nil {
            return true, 0
        }
        leftBalanced, leftDepth := getBalancedAndDepth(root.Left)
        rightBalanced, rightDepth := getBalancedAndDepth(root.Right)
        diffDepth := leftDepth-rightDepth
        if diffDepth < 0 {
            diffDepth = - diffDepth
        }
        resultBalanced := leftBalanced && rightBalanced && diffDepth <= 1
        if leftDepth > rightDepth {
            return resultBalanced, 1+leftDepth
        }
        return resultBalanced, 1+rightDepth
    }
    result, _ := getBalancedAndDepth(root)
    return result
}

func inOrder(root *TreeNode, nums *[]int) {
    if root == nil {
        return
    }
    inOrder(root.Left, nums)
    *nums = append(*nums, root.Val)
    inOrder(root.Right, nums)
}
func buildBalancedBST(nums []int, left, right int) *TreeNode {
    if left > right {
        return nil
    }
    mid := left + (right-left)/2
    node := &TreeNode{Val: nums[mid]}
    node.Left = buildBalancedBST(nums, left, mid-1)
    node.Right = buildBalancedBST(nums, mid+1, right)
    return node
}
func balanceBST(root *TreeNode) *TreeNode {
    var nums []int
    inOrder(root, &nums)
    return buildBalancedBST(nums, 0, len(nums)-1)
}

func longestBalanced(nums []int) int {
	// 3719
	maxLength := 0
	n := len(nums)

	// Iterate over all possible starting indices, i
	for i := 0; i < n; i++ {
		distinctEvens := make(map[int]bool)
		distinctOdds := make(map[int]bool)
		evenCount := 0
		oddCount := 0

		// Iterate over all possible ending indices, j
		for j := i; j < n; j++ {
			num := nums[j]
			if num%2 == 0 {
				if !distinctEvens[num] {
					distinctEvens[num] = true
					evenCount++
				}
			} else {
				if !distinctOdds[num] {
					distinctOdds[num] = true
					oddCount++
				}
			}

			// Check if the current subarray is balanced
			if evenCount == oddCount {
				currentLength := j - i + 1
				if currentLength > maxLength {
					maxLength = currentLength
				}
			}
		}
	}

	return maxLength
}

func longestBalancedSubstring(s string) int {
	// 3713
	n := len(s)
	maxLen := 0

	// k is the number of distinct characters
	for k := 1; k <= 26; k++ {
		// c is the frequency of each distinct character
		for c := 1; k*c <= n; c++ {
			winSize := k * c
			counts := make([]int, 26)
			distinctCount := 0
			validCount := 0

			// Initialize first window
			for i := 0; i < winSize; i++ {
				charIdx := s[i] - 'a'
				if counts[charIdx] == 0 {
					distinctCount++
				}
				counts[charIdx]++
				if counts[charIdx] == c {
					validCount++
				} else if counts[charIdx] == c+1 {
					validCount--
				}
			}

			if distinctCount == k && validCount == k {
				maxLen = max(maxLen, winSize)
			}

			// Slide the window
			for i := winSize; i < n; i++ {
				// Add new character
				inIdx := s[i] - 'a'
				if counts[inIdx] == 0 {
					distinctCount++
				}
				counts[inIdx]++
				if counts[inIdx] == c {
					validCount++
				} else if counts[inIdx] == c+1 {
					validCount--
				}

				// Remove old character
				outIdx := s[i-winSize] - 'a'
				if counts[outIdx] == c {
					validCount--
				} else if counts[outIdx] == c+1 {
					validCount++
				}
				counts[outIdx]--
				if counts[outIdx] == 0 {
					distinctCount--
				}

				if distinctCount == k && validCount == k {
					maxLen = max(maxLen, winSize)
				}
			}
		}
	}
	return maxLen
}

func longestBalancedSubstring2(s string) int {
	// 3714
	n := len(s)
    maxLen := 0

	for c := 1; c <= n; c++ {
        for k := 1; k <= 3; k++ {
            winSize := k * c
            if winSize > n || winSize <= maxLen {
                continue
            }

            counts := [3]int{}
            distinct := 0
            valid := 0

            for i := 0; i < winSize; i++ {
                idx := s[i] - 'a'
                if counts[idx] == 0 {
                    distinct++
                }
                counts[idx]++
                if counts[idx] == c {
                    valid++
                } else if counts[idx] == c+1 {
                    valid--
                }
            }

            if distinct == k && valid == k {
                maxLen = winSize
            }

            for i := winSize; i < n; i++ {
                inIdx := s[i] - 'a'
                if counts[inIdx] == 0 {
                    distinct++
                }
                counts[inIdx]++
                if counts[inIdx] == c {
                    valid++
                } else if counts[inIdx] == c+1 {
                    valid--
                }

                outIdx := s[i-winSize] - 'a'
                if counts[outIdx] == c {
                    valid--
                } else if counts[outIdx] == c+1 {
                    valid++
                }
                counts[outIdx]--
                if counts[outIdx] == 0 {
                    distinct--
                }

                if distinct == k && valid == k {
                    maxLen = winSize
                    break
                }
            }
        }
    }
    return maxLen
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

