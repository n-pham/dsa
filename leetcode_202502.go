package main

import (
	"container/heap"
	"fmt"
	// "math"
	"slices"
	"strconv"
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
				positions = append(positions, []int{i, j})
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

func maxAscendingSum1800(nums []int) int {
	// 1800
	maxSum, prev, curSum := nums[0], nums[0], nums[0]
	for _, num := range nums[1:] {
		if num > prev {
			curSum += num
			if curSum > maxSum {
				maxSum = curSum
			}
		} else {
			curSum = num
		}
		prev = num
	}
	return maxSum
}

func isAnagram242(s string, t string) bool {
	// 242
	m := [26]int{}
	for _, c := range s {
		m[c-'a']++
	}
	for _, c := range t {
		m[c-'a']--
	}
	for _, v := range m {
		if v != 0 {
			return false
		}
	}
	return true
}

func minChanges2914(s string) int {
	// 2914
	panic("not implemented")
	oddCnt, curCnt, prev := [2]int{}, 1, s[0]
	for i := 1; i < len(s); i++ {
		if prev == s[i] {
			curCnt++
		} else {
			if curCnt%2 == 1 {
				oddCnt[prev-'0']++
			}
			curCnt = 1
		}
		fmt.Println("curCnt", curCnt, "oddCnt", oddCnt)
		prev = s[i]
	}
	if curCnt%2 == 1 {
		oddCnt[prev-'0']++
	}
	fmt.Println("curCnt", curCnt, "oddCnt", oddCnt)
	maxCnt := oddCnt[0]
	if oddCnt[1] > oddCnt[0] {
		maxCnt = oddCnt[1]
	}
	return maxCnt
}

func smallestEquivalentString1061(s1 string, s2 string, baseStr string) string {
	// 1061
	// p m m:m, p:m
	// a o a:a, o:a
	// r r r:r
	// k r k:k, r:k
	// e i e:e, i:e
	// r s s:k
	//
	// l p p:l l:l
	// e r r:e e:e
	// e o o:e
	// t g t:g
	// c r r:c
	// o a a:a o:a AND e:a i:a r:a
	// d m d:d m:d
	// e s s:a
	m := [26]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25}
	for i := 0; i < len(s1); i++ {
		smallest := min(m[s1[i]-'a'], m[s2[i]-'a'], s1[i]-'a', s2[i]-'a')
		if prev := m[s1[i]-'a']; prev > smallest {
			for j := 0; j < 26; j++ {
				if m[j] == prev {
					m[j] = smallest
				}
			}
		}
		if prev := m[s2[i]-'a']; prev > smallest {
			for j := 0; j < 26; j++ {
				if m[j] == prev {
					m[j] = smallest
				}
			}
		}
		m[s1[i]-'a'], m[s2[i]-'a'] = smallest, smallest
		// fmt.Println(m)
	}
	rs := make([]byte, len(baseStr))
	for i := 0; i < len(baseStr); i++ {
		rs[i] = m[baseStr[i]-'a'] + 'a'
	}
	return string(rs)
}

func smallestNumber2375(pattern string) string {
	// 2375
	//  IIIDIDDD
	// 12343
	// 123546
	// 123549876
	panic("not implemented")
}

func numTilePossibilities1079(tiles string) int {
	// 1079
	// A B
	// 2 1
	m := [26]int{}
	for _, tile := range tiles {
		m[tile-'A']++
	}
	var dfs func() int
	dfs = func() int {
		cnt := 0
		for i := range m {
			if m[i] > 0 {
				cnt++
				m[i]--
				cnt += dfs()
				m[i]++
			}
		}
		return cnt
	}
	return dfs()
}

func findDuplicates_5ms(nums []int) (dups []int) {
	// 442
	m := make([]int, len(nums)+1)
	for _, num := range nums {
		if m[num] > 0 {
			dups = append(dups, num)
		} else {
			m[num] = 1
		}
	}
	return dups
}

func findDuplicates442_bitmask(nums []int) []int {
	// 442
	dupLen, bitmask := 0, make([]uint64, (len(nums)/64)+1)
	for _, num := range nums {
		if (bitmask[num/64] & (1 << (num % 64))) > 0 {
			nums[dupLen] = num
			dupLen++
		} else {
			bitmask[num/64] |= 1 << (num % 64)
		}
	}
	return nums[:dupLen]
}

func findDuplicates442(nums []int) []int {
	// 442
	dupLen, m := 0, make([]byte, len(nums)+1)
	for _, num := range nums {
		if m[num] > 0 {
			nums[dupLen] = num
			dupLen++
		} else {
			m[num] = 1
		}
	}
	return nums[:dupLen]
}

func areAlmostEqual1790(s1 string, s2 string) bool {
	// 1790
	diffCount, diff1, diff2 := 0, byte(0), byte(0)
	for i := range s1 {
		if s1[i] != s2[i] {
			if diff1 != 0 {
				if diff1 != s2[i] || diff2 != s1[i] {
					return false
				}
				diffCount--
			} else {
				diff1, diff2, diffCount = s1[i], s2[i], diffCount+1
			}
		}
	}
	return diffCount == 0
}

func numOfPairs2023(nums []string, target string) (cnt int) {
	// 2023
	for i, s1 := range nums {
		for j, s2 := range nums {
			if i == j || len(s1)+len(s2) != len(target) {
				continue
			}
			if s1 == target[:len(s1)] && s2 == target[len(s1):] {
				cnt++
			}
		}
	}
	return cnt
}

func tupleSameProduct1726_time(nums []int) (cnt int) {
	// 1726
	for _, a := range nums {
		for _, b := range nums {
			if b == a {
				continue
			}
			for _, c := range nums {
				if c == b || c == a {
					continue
				}
				for _, d := range nums {
					if d == c || d == b || d == a {
						continue
					}
					if a*b == c*d {
						cnt++
					}
				}
			}
		}
	}
	return cnt
}

func tupleSameProduct1726(nums []int) (cnt int) {
	// 1726
	productMap := make(map[int]int)
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			product := nums[i] * nums[j]
			cnt += productMap[product] // +1 from next time (c*d)
			productMap[product]++
		}
	}
	return cnt * 8 // 8 different ways
}

func getHappyString1415_226ms(n int, k int) string {
	// 1415
	var rec func(n int, prev int) []string
	rec = func(n int, prev int) []string {
		if n == 0 {
			return []string{""}
		}
		rs := []string{}
		for i := 0; i < 3; i++ {
			if i == prev {
				continue
			}
			first := string('a'+i)
			for _, rest := range rec(n-1, i) {
				rs = append(rs, first + rest)
			}
		}
		return rs
	}
	strings := []string{}
	for _, v := range rec(n, 1) {
		p, found := slices.BinarySearch(strings, v)
		if !found {
			strings = slices.Insert(strings, p, v)
		}
	}
	for _, v := range rec(n, 2) {
		p, found := slices.BinarySearch(strings, v)
		if !found {
			strings = slices.Insert(strings, p, v)
		}
	}
	for _, v := range rec(n, 0) {
		p, found := slices.BinarySearch(strings, v)
		if !found {
			strings = slices.Insert(strings, p, v)
		}
	}
	if k > len(strings) {
		return ""
	}
	return strings[k-1]
}

func getHappyString1415(n int, k int) string {
	// 1415
	panic("not implemented")
}

func queryResults3160_memory(limit int, queries [][]int) []int {
	// 3160
	//       0 0 0 0 0
	// 1,4   0 4 0 0 0 [4]++
	// 2,5   0 4̲ 5 0 0 [5]++
	// 1,3   0 3̲ 5 0 0 [4]-- [3]++
	// 3,4   0 3 5 4 0 [4]++
	rs, colorByLabel, colorCnt := make([]int, len(queries)), make([]int, limit+1), make(map[int]int)
	for i, q := range queries {
		if colorByLabel[q[0]] > 0 {
			if v, _ := colorCnt[colorByLabel[q[0]]]; v == 1 {
				delete(colorCnt, colorByLabel[q[0]])
			} else {
				colorCnt[colorByLabel[q[0]]]--
			}
		}
		colorByLabel[q[0]] = q[1]
		colorCnt[q[1]]++
		rs[i] = len(colorCnt)
	}
	return rs
}

func queryResults3160(limit int, queries [][]int) []int {
	// 3160
	// 1,4   1:4            [4]++
	// 2,5   1:4̲, 2:5       [5]++
	// 1,3   1:3̲, 2:5       [4]-- [3]++
	// 3,4   1:3, 2:5, 3:4  [4]++
	rs, colorByLabel, colorCnt := make([]int, len(queries)), make(map[int]int), make(map[int]int)
	for i, q := range queries {
		if color, found := colorByLabel[q[0]]; found {
			if cnt, _ := colorCnt[color]; cnt == 1 {
				delete(colorCnt, color)
			} else {
				colorCnt[color]--
			}
		}
		colorByLabel[q[0]] = q[1]
		colorCnt[q[1]]++
		rs[i] = len(colorCnt)
	}
	return rs
}

func findDifferentBinaryString1980(nums []string) string {
	// 1980
	panic("not implemented")
	slices.Sort(nums)
	for i := 1; i < len(nums); i++ {
		num, _ := strconv.Atoi(nums[i])
		prev := strconv.Itoa(num-1)
		if prev != nums[i-1] {
			return prev
		}
	}
	return "Error"
}

func restoreArray1743(adjacentPairs [][]int) []int {
	// 1743
	// 2,1   2̶1̶ 12
	// 3,4   34 4̶3̶
	// 3,2   3̶2̶ 23
	panic("not implemented")
	numMap := make([]int, len(adjacentPairs)+2)
	for _, pair := range adjacentPairs {
		if numMap[pair[0]] == 0 {
			numMap[pair[0]] = pair[1]
		} else if numMap[pair[1]] == 0 {
			numMap[pair[1]] = pair[0]
		}
		
	}
	fmt.Println(numMap)
	return []int{99}
}

func partitionArray2294(nums []int, k int) int {
	// 2294
	// 3 3
	// 6 3     6
	// 1 1,3   6
	// 2 1,3
	// 5 1,3   5,6
	// 4 1,3   4,6
	panic("not implemented")
}

// 2349 solved by Copilot
type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type NumberContainers struct {
	numberByIndex map[int]int
	indexByNumber map[int]*MinHeap
}

func Constructor() NumberContainers {
	return NumberContainers{
		numberByIndex: make(map[int]int),
		indexByNumber: make(map[int]*MinHeap),
	}
}

func (this *NumberContainers) Change(index int, number int) map[int]int {
	if oldNumber, exists := this.numberByIndex[index]; exists {
		heap.Remove(this.indexByNumber[oldNumber], findIndex(this.indexByNumber[oldNumber], index))
		if this.indexByNumber[oldNumber].Len() == 0 {
			delete(this.indexByNumber, oldNumber)
		}
	}
	this.numberByIndex[index] = number
	if _, exists := this.indexByNumber[number]; !exists {
		this.indexByNumber[number] = &MinHeap{}
		heap.Init(this.indexByNumber[number])
	}
	heap.Push(this.indexByNumber[number], index)
	return this.numberByIndex
}

func findIndex(h *MinHeap, index int) int {
	for i, v := range *h {
		if v == index {
			return i
		}
	}
	return -1
}

func (this *NumberContainers) Find(number int) int {
	if indices, exists := this.indexByNumber[number]; exists && indices.Len() > 0 {
		return (*indices)[0]
	}
	return -1
}

func main() {
	nc := Constructor()
	fmt.Println(nc.Find(10)) // There is no index that is filled with number 10. Therefore, we return -1.
	fmt.Println(nc.Change(2, 10)) // Your container at index 2 will be filled with number 10.
	fmt.Println(nc.Change(1, 10)) // Your container at index 1 will be filled with number 10.
	fmt.Println(nc.Change(3, 10)) // Your container at index 3 will be filled with number 10.
	fmt.Println(nc.Change(5, 10)) // Your container at index 5 will be filled with number 10.
	fmt.Println(nc.Find(10)) // Number 10 is at the indices 1, 2, 3, and 5. Since the smallest index that is filled with 10 is 1, we return 1.
	fmt.Println(nc.Change(1, 20)) // Your container at index 1 will be filled with number 20. Note that index 1 was filled with 10 and then replaced with 20. 
	fmt.Println(nc.Find(10)) // Number 10 is at the indices 2, 3, and 5. The smallest index that is filled with 10 is 2. Therefore, we return 2.
	// fmt.Println(restoreArray1743([][]int{{2,1},{3,4},{3,2}}))
	// fmt.Println(findDifferentBinaryString1980([]string{"111","011","001"}))
	// fmt.Println(queryResults3160(4, [][]int{{1,4},{2,5},{1,3},{3,4}}))
	// fmt.Println(getHappyString1415(3, 9))
	// fmt.Println(tupleSameProduct1726([]int{2, 3, 4, 6}))
	// fmt.Println(numOfPairs2023([]string{"123","4","12","34"}, "1234"))
	// fmt.Println(numOfPairs2023([]string{"777","7","77","77"}, "7777"))
	// fmt.Println(areAlmostEqual1790("bank", "kanb"))
	// fmt.Println(areAlmostEqual1790("aa", "ac"))
	// fmt.Println(areAlmostEqual1790("baaa", "abbb"))
	// fmt.Println(findDuplicates442([]int{4,3,2,7,8,2,3,1}))
	// fmt.Println(numTilePossibilities1079("AAB"))
	// fmt.Println(smallestEquivalentString1061("leetcode", "programs", "sourcecode"))
	// fmt.Println(smallestEquivalentString1061( "parker", "morris", "parser"))
	// fmt.Println(minChanges2914("01010000011001001101")) // 6
	// fmt.Println(minChanges2914("11000111"))
	// fmt.Println(minChanges2914("1001"))
	// fmt.Println(isAnagram242("anagram", "nagaram"))
	// fmt.Println(isAnagram242("rat", "car"))
	// fmt.Println(isAnagram242("a", "ab"))
	// fmt.Println(maxAscendingSum1800([]int{12,17,15,13,10,11,12}))
	// fmt.Println(countServers1267([][]int{{1, 1, 0, 0}, {0, 0, 1, 0}, {0, 0, 1, 0}, {0, 0, 0, 1}}))
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
