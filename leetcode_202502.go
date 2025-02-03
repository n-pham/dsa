package main

import (
	"fmt"
	// "math"
	// "slices"
	// "strconv"
	// "strings"
)

func isArraySpecial3151(nums []int) bool {
	// 3151
	firstValue := nums[0] & 1
	for i, num := range nums[1:] {
		if (i%2 == 0 && (num&1 == firstValue)) || (i%2 == 1 && (num&1 != firstValue)) {
			return false
		}
	}
	return true
}

func findRedundantConnection648_fail(edges [][]int) []int {
	// 648
	//         3
	//         |
	// 1 - 5 - 8 - 6 - 2
	//         |
	//        10 - 9 - 4
	//             |
	//             7
	// {4, 10}
	nodeMap := make([]int, 11) // n <= 1000
	errors := [][]int{}
	for _, edge := range edges {
		if (nodeMap[edge[0]] == 1) && (nodeMap[edge[1]] == 1) {
			errors = append(errors, edge)
		} else {
			nodeMap[edge[0]] = 1
			nodeMap[edge[1]] = 1
		}
		fmt.Println(nodeMap, errors)
	}
	return errors[len(errors)-1]
}

func findRedundantConnection648(edges [][]int) []int {
	// 648
	//         3
	//         |
	// 1 - 5 - 8 - 6 - 2
	//         |
	//        10 - 9 - 4
	//             |
	//             7
	// {4, 10}
	parent := make([]int, len(edges)+1)
	for i := range parent {
		parent[i] = i
	}

	var find func(int) int
	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}

	union := func(x, y int) bool {
		rootX := find(x)
		rootY := find(y)
		if rootX != rootY {
			parent[rootX] = rootY
			return true
		}
		return false
	}

	for _, edge := range edges {
		if !union(edge[0], edge[1]) {
			return edge
		}
	}
	return nil
}

func check1752_fail(nums []int) bool {
	// 1752
	maxNum, wrappedNum := 0, 101 // nums[i] <= 100
	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[i-1] {
			if wrappedNum != 101 {
				return false
			} else {
				wrappedNum = nums[i-1]
			}
		}
		if nums[i] > maxNum {
			maxNum = nums[i]
		}
		if nums[i] > wrappedNum {
			return false
		}
	}
	return true
}

func check1752(nums []int) bool {
	// 1752
	wrappedOnce := false
	for i := 0; i < len(nums); i++ {
		if nums[i] > nums[(i+1)%len(nums)] {
			if wrappedOnce {
				return false
			} else {
				wrappedOnce = true
			}
		}
	}
	return true
}

func longestMonotonicSubarray3105(nums []int) int {
	// 3105
	maxLen, prev, incLen, decLen := 1, nums[0], 1, 1
	for _, num := range nums[1:] {
		if num > prev {
			decLen, incLen = 1, incLen+1
			if incLen > maxLen {
				maxLen = incLen
			}
		} else if prev > num {
			decLen, incLen = decLen+1, 1
			if decLen > maxLen {
				maxLen = decLen
			}
		} else {
			decLen, incLen = 1, 1
		}
		prev = num
	}
	return maxLen
}

func checkIfPrerequisite1462(numCourses int, prerequisites [][]int, queries [][]int) []bool {
	// 1462
	panic("not implemented")
}

func lexicographicallySmallestArray2948(nums []int, limit int) []int {
	// 2948
	panic("not implemented")
}

func eventualSafeNodes802(graph [][]int) (rs []int) {
	// 802
	//           0     1     2   3   4   5  6
	//           [1,2],[2,3],[4],[0],[5],[],[]
	// terminal                          5  6
	panic("not implemented")
	m := make([]int, len(graph))
	g := graph[:]
	added := true
	for added {
		added = false
		for i, targets := range g {
			isContained := true
			for _, target := range targets {
				if m[target] == 0 {
					isContained = false
					break
				}
			}
			if isContained {
				m[i] = 1
				added = true
				g = append(g[:i], g[i+1:]...)
			}
		}
	}
	for i, v := range m {
		if v == 1 {
			rs = append(rs, i)
		}
	}
	return rs
}

func countServers1267(grid [][]int) (cnt int) {
	// 1267, 4ms
	positions := [][]int{}
	for i, row := range grid {
		for j, val := range row {
			if val == 1 {
				positions = append(positions, []int{i,j})
			}
		}
	}
	isolatedCnt := 0
	for i, pos1 := range positions {
		isolated := true
		for j, pos2 := range positions {
			if i != j && (pos1[0] == pos2[0] || pos1[1] == pos2[1]) {
				isolated = false
				break
			}
		}
		if isolated {
			isolatedCnt++
		}
	}
	return len(positions) - isolatedCnt
}

func main() {
	fmt.Println(countServers1267([][]int{{1, 1, 0, 0}, {0, 0, 1, 0}, {0, 0, 1, 0}, {0, 0, 0, 1}}))
	// fmt.Println(eventualSafeNodes802([][]int {{1,2},{2,3},{4},{0},{5},{},{}}))
	// fmt.Println(longestMonotonicSubarray3105([]int{1, 4, 3, 3, 2}))
	// fmt.Println(check1752([]int{2, 4, 1, 3}))    // false
	// fmt.Println(check1752([]int{2, 1, 3, 4}))    // false
	// fmt.Println(check1752([]int{3, 4, 5, 1, 2})) // true
	// fmt.Println(check1752([]int{6, 10, 6}))      // true
	// fmt.Println(findRedundantConnection648([][]int{{9, 10}, {5, 8}, {2, 6}, {1, 5}, {3, 8}, {4, 9}, {8, 10}, {4, 10}, {6, 8}, {7, 9}}))
	// fmt.Println(findRedundantConnection648([][]int{{1, 2}, {2, 3}, {3, 4}, {1, 4}, {1, 5}}))
	// fmt.Println(isArraySpecial3151([]int{1, 6, 2}))    // false
	// fmt.Println(isArraySpecial3151([]int{2, 1, 4}))    // true
	// fmt.Println(isArraySpecial3151([]int{4, 3, 1, 6})) // false
}
