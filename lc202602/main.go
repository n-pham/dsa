package main

import (
	"container/heap"
	"math"
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

func minimumCost(nums []int, k int, dist int) int64 {

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
