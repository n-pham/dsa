package main

import (
	"container/heap"
	"fmt"
	"math"
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
	// IIIDIDDD  result     stack
	// IIID      123        4
	// I         1235       4(5)
	// DDD       1235       467
	// end       1235764
	rs := make([]byte, len(pattern)+1)
	stack := []int{} // ascending
	for i := 0; i <= len(pattern); i++ {
		stack = append(stack, i+1)
		if i == len(pattern) || pattern[i] == 'I' {
			for len(stack) > 0 {
				rs[i-len(stack)+1] = byte(stack[len(stack)-1] + '0')
				stack = stack[:len(stack)-1]
			}
		}
	}
	return string(rs)
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
			if i != j &&
				len(s1)+len(s2) == len(target) &&
				s1 == target[:len(s1)] &&
				s2 == target[len(s1):] {
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
			first := string('a' + i)
			for _, rest := range rec(n-1, i) {
				rs = append(rs, first+rest)
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
		prev := strconv.Itoa(num - 1)
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
	panic("not implemented")
}

func partitionArray2294_fail(nums []int, k int) int {
	// 2294 slices.Sort(nums) time limit
	// 16      16,16
	// 8        8,16
	// 17       8,17
	// 0        8,17  0, 0
	// 3        8,17  0, 3
	// 20       8,17  0, 3 20,20
	// correct  0, 8 16,20
	partitions := [][2]int{} // partitions of [start, end]
	for _, num := range nums {
		found := false
		for i, p := range partitions {
			if p[1]-num > k || num-p[0] > k {
				continue
			} else if p[1]-num <= k && num < p[0] { // num p0 p1
				partitions[i][0] = num
			} else if num-p[0] <= k && p[1] < num { // p0 p1 num
				partitions[i][1] = num
			}
			found = true
		}
		if !found {
			partitions = append(partitions, [2]int{num, num})
		}
		fmt.Println(num, partitions)
	}
	return len(partitions)
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

func countBadPairs2364_time(nums []int) (cnt int64) {
	// 2364
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if j-i != nums[j]-nums[i] {
				cnt++
			}
		}
	}
	return cnt
}

func countBadPairs2364_solution(nums []int) (cnt int64) {
	// 2364
	diffMap := make(map[int]int)
	for i, num := range nums {
		diff := num - i
		if count, exists := diffMap[diff]; exists {
			cnt += int64(i - count)
		} else {
			cnt += int64(i)
		}
		diffMap[diff]++
	}
	return cnt
}

func countBadPairs2364_44ms(nums []int) (cnt int64) {
	// 2364
	// 2 3 6 5 6 9 8  nums
	// 0 1 2 3 4 5 6  i
	// 2 2 4 2 2 4 2  num-i if change --> count before (==i) + after (how?)
	// 0 0 2 1 1 4 2  expected
	// 0 1 2 3 4 5 6  i
	// 0 1 0 2 3 1 4  prevCnt[num-i]
	// 0 0 2 1 1 4 2  expected = i - prevCnt[num-i] WHOEVER thought of this?
	prevCnt := make(map[int]int)
	for i, num := range nums {
		count := 0
		count, _ = prevCnt[num-i]
		cnt += int64(i - count)
		prevCnt[num-i]++
	}
	return cnt
}

func countBadPairs2364(nums []int) int64 {
	// 2364 27ms j - i != nums[j] - nums[i] => nums[i] - i != nums[j] - j
	// 2 3 6 5 6 9 8  nums
	// 0 1 2 3 4 5 6  i
	// 2 2 4 2 2 4 2  num-i if change --> bad
	// 0 1 0 2 3 1 4  cntM --> goodCnt == 11
	cntM, goodCnt := map[int]int{}, 0
	for i := 0; i < len(nums); i++ {
		val := i - nums[i]   // access slice element once
		goodCnt += cntM[val] // +1 next time
		cntM[val]++
	}
	return int64((len(nums)*(len(nums)-1))/2 - goodCnt)
}

func compress443(chars []byte) int {
	// 443
	// aabccccccccccd
	// a2bc10d
	newLen, cnt := 0, 1
	for i, c := range chars {
		if i == len(chars)-1 || chars[i+1] != c {
			chars[newLen], newLen = c, newLen+1
			if cnt > 1 {
				cntStr := strconv.Itoa(cnt)
				for _, ch := range cntStr {
					chars[newLen], newLen = byte(ch), newLen+1
				}
			}
			cnt = 1
		} else {
			cnt++
		}
	}
	fmt.Println(string(chars[:newLen]))
	return newLen
}

func clearDigits3174(s string) string {
	// 3174
	// acb34d --> ac4 --> ad
	rs, newLen := make([]rune, len(s)), 0
	for _, c := range s {
		if c < '0' || c > '9' {
			rs[newLen], newLen = c, newLen+1
		} else {
			newLen--
		}
	}
	return string(rs[:newLen])
}

func canVisitAllRooms841(rooms [][]int) bool {
	// 841
	visited := make(map[int]byte)
	q := []int{}
	for _, k := range rooms[0] {
		q = append(q, k)
	}
	visited[0] = 1
	var curNum int
	for len(q) > 0 {
		curNum, q = q[0], q[1:]
		visited[curNum] = 1
		for _, k := range rooms[curNum] {
			if visited[k] == 0 {
				q = append(q, k)
			}
		}
	}
	return len(visited) == len(rooms)
}

func maximumSum2342_14ms(nums []int) int {
	// 2342
	// 18  9  18
	// 43  7  43
	// 36  9  36  18+36 = 54
	// 13  4
	//  7  7  43  43+ 7 = 50
	// 16  7  43  43+16 = 59
	rs, maxNumberBydigitSum := -1, make(map[int]int)
	for _, num := range nums {
		digitSum := 0
		for num2 := num; num2 > 0; digitSum, num2 = digitSum+num2%10, num2/10 {
		}
		if maxNum := maxNumberBydigitSum[digitSum]; maxNum > 0 {
			if numberSum := num + maxNum; numberSum > rs {
				rs = numberSum
			}
		}
		if num > maxNumberBydigitSum[digitSum] {
			maxNumberBydigitSum[digitSum] = num
		}
	}
	return rs
}

func maximumSum2342(nums []int) int {
	// 2342
	// 18  9  18
	// 43  7  43
	// 36  9  36  18+36 = 54
	// 13  4
	//  7  7  43  43+ 7 = 50
	// 16  7  43  43+16 = 59
	rs, maxNumberBydigitSum := -1, make([]int, 1+81) // 9..9 9 times
	for _, num := range nums {
		digitSum := 0
		for num2 := num; num2 > 0; digitSum, num2 = digitSum+num2%10, num2/10 {
		}
		if maxNum := maxNumberBydigitSum[digitSum]; maxNum > 0 {
			if numberSum := num + maxNum; numberSum > rs {
				rs = numberSum
			}
			if num > maxNum {
				maxNumberBydigitSum[digitSum] = num
			}
		} else {
			maxNumberBydigitSum[digitSum] = num
		}
	}
	return rs
}

func reconstructQueue406(people [][]int) [][]int {
	// 406
	//      0   1   2   3   4   5
	// 7,0
	// 4,4
	// 7,1
	// 5,0
	// 6,1
	// 5,2
	panic("not implemented")
}

func maxIceCream1833(costs []int, coins int) int {
	// 1833
	// count := make([]int, 100000+1) // slower
	maxCost := 0
	for _, cost := range costs {
		if cost > maxCost {
			maxCost = cost
		}
	}
	count := make([]int, maxCost+1) // make([]int, 100000+1) is slower
	for _, cost := range costs {
		count[cost]++
	}
	totalIceCreams := 0
	for cost := 1; cost < len(count); cost++ {
		iceCreamCnt := count[cost]
		if canBuyCnt := coins / cost; canBuyCnt < iceCreamCnt {
			iceCreamCnt = canBuyCnt
		}
		totalIceCreams += iceCreamCnt
		coins -= iceCreamCnt * cost
		if coins == 0 {
			return totalIceCreams
		}
	}
	return totalIceCreams
}

func frequencySort451(s string) string {
	// 451
	// Aabb --> bbAa
	count := [74 + 1]int{} // '0' to 'z'
	for _, c := range s {
		count[c-'0']++
	}
	type pair struct {
		char  rune
		count int
	}
	pairs := make([]pair, 0, len(count))
	for c, cnt := range count {
		if cnt > 0 {
			p := pair{rune(c + '0'), cnt}
			idx, _ := slices.BinarySearchFunc(pairs, p, func(a, b pair) int {
				if a.count == b.count {
					if a.char < b.char {
						return -1
					} else if a.char > b.char {
						return 1
					}
					return 0
				}
				if a.count > b.count {
					return -1
				}
				return 1
			})
			pairs = slices.Insert(pairs, idx, p)
		}
	}
	rs := make([]rune, 0, len(s))
	for _, p := range pairs {
		for i := 0; i < p.count; i++ {
			rs = append(rs, p.char)
		}
	}
	return string(rs)
}

func isValidSudoku36(board [][]byte) bool {
	// 36 still in progress
	var existedRow [9]int
	for _, row := range board {
		existedRow = [9]int{}
		for _, c := range row {
			if c != '.' {
				existedRow[c-'1']++
				if existedRow[c-'1'] > 1 {
					return false
				}
			}
		}
	}
	return true
}

func productExceptSelf238_3ms(nums []int) []int {
	// 238
	// 1  2*3*4     1  2*3*4
	// 2  1*3*4     1    3*4
	// 3  1*2*4     1*2    4
	// 4  1*2*3     1*2*3  1
	leftPrefixProduct, rightPrefixProduct := make([]int, len(nums)), make([]int, len(nums))
	leftPrefixProduct[0], rightPrefixProduct[len(nums)-1] = 1, 1
	leftPrefixProduct[1], rightPrefixProduct[len(nums)-2] = nums[0], nums[len(nums)-1]
	for i := 2; i < len(nums); i++ {
		leftPrefixProduct[i] = nums[i-1] * leftPrefixProduct[i-1]
	}
	for i := len(nums) - 3; i >= 0; i-- {
		rightPrefixProduct[i] = nums[i+1] * rightPrefixProduct[i+1]
	}
	for i := 0; i < len(nums); i++ { // in-place re-use leftPrefixProduct
		leftPrefixProduct[i] *= rightPrefixProduct[i]
	}
	return leftPrefixProduct
}

func productExceptSelf238(nums []int) []int {
	// 238
	// 1  2*3*4     1  2*3*4
	// 2  1*3*4     1    3*4
	// 3  1*2*4     1*2    4
	// 4  1*2*3     1*2*3  1
	leftPrefixProduct := make([]int, len(nums))
	leftPrefixProduct[0] = 1
	for i := 1; i < len(nums); i++ {
		leftPrefixProduct[i] = nums[i-1] * leftPrefixProduct[i-1]
	}
	rightProduct := 1
	for i := len(nums) - 1; i >= 0; i-- {
		leftPrefixProduct[i] *= rightProduct
		rightProduct *= nums[i]
	}
	return leftPrefixProduct
}

func minOperations3066_solution(nums []int, k int) int {
	panic("not implemented")
	// if len(nums) < 2 {
	// 	return 0
	// }
	// count := 0
	// h := MinHeap(nums)
	// h.buildHeap()
	// for h.length()>=2{
	//     f := h.pop()
	//     if f >= k {
	//         return count
	//     }
	//     s := h.pop()
	// 	val := min(f, s)*2 + max(f, s)
	// 	h.insert(val)
	//     count++
	// }
	// if h.peek() >= k {
	//     return count
	// }
	// return count + 1
}

func minOperations3066(nums []int, k int) int {
	// 3066
	h := &MinHeap{}
	heap.Init(h)
	for _, num := range nums {
		heap.Push(h, num)
	}
	operations := 0
	for h.Len() > 0 {
		fmt.Println(h)
		smallest := heap.Pop(h).(int)
		if smallest >= k {
			return operations
		}
		if h.Len() == 0 {
			return -1
		}
		secondSmallest := heap.Pop(h).(int)
		newNum := 2*smallest + secondSmallest
		heap.Push(h, newNum)
		operations++
	}
	return operations
}

type ProductOfNumbers struct {
	products []int
}

// func Constructor() ProductOfNumbers {
// 	return ProductOfNumbers{products: []int{1}}
// }

func (this *ProductOfNumbers) Add(num int) {
	if num == 0 {
		this.products = []int{1}
	} else {
		this.products = append(this.products, this.products[len(this.products)-1]*num)
	}
}

func (this *ProductOfNumbers) GetProduct(k int) int {
	n := len(this.products)
	if k >= n {
		return 0
	}
	return this.products[n-1] / this.products[n-1-k]
}

func punishmentNumber2698(n int) int {
	// 2698
	// 36** 1296
	// 36,1-296          36,12-96   36,129-6
	// 35,2-96 35,29-6   24,9-6     false
	// false   true      false
	var isValidPartition func(int, int) bool
	isValidPartition = func(num int, target int) bool {
		if target < 0 || num < target {
			return false
		}
		if num == target {
			return true
		}
		return isValidPartition(num/10, target-(num%10)) ||
			isValidPartition(num/100, target-(num%100)) ||
			isValidPartition(num/1000, target-(num%1000))
	}
	rs := 0
	for i := 1; i <= n; i++ {
		if i%9 < 2 && isValidPartition(i*i, i) {
			rs += i * i
		}
	}
	return rs
}

func longestConsecutive128_time(nums []int) int {
	// 128
	// 100  100:1
	// 4    100:1, 4:1
	// 200  100:1, 4:1, 200:1
	// 1    100:1, 4:1, 200:1, 1:1
	// 3    100:1, 4:1, 200:1, 1:1, 3:1
	// 2    for m[i-left] > 0; left++ {} for m[i+right] > 0; right++ {}
	maxLen, m := 0, make(map[int]byte)
	for _, num := range nums {
		m[num] = 1
		left, right := 0, 0
		for ; m[num-left-1] > 0; left++ {
		}
		for ; m[num+right+1] > 0; right++ {
		}
		if right+left+1 > maxLen {
			maxLen = right + left + 1
		}
	}
	return maxLen
}

func longestConsecutive128_time2(nums []int) int {
	// 128
	// num in firsts, for m[num+right+1] exists; right++ {}
	maxLen, m, firsts := 0, make(map[int]struct{}), make(map[int]struct{})
	for _, num := range nums {
		m[num] = struct{}{}
		if _, found := firsts[num-1]; !found {
			firsts[num] = struct{}{}
		}
		for right, found := 0, true; found; right++ {
			if _, found = firsts[num+right+1]; found {
				delete(firsts, num+right+1)
			}
		}
	}
	fmt.Println(firsts)
	for num, _ := range firsts {
		right := 0
		for found := true; found; right++ {
			_, found = m[num+right+1]
		}
		if right > maxLen {
			maxLen = right
		}
	}
	return maxLen
}

func longestConsecutive128_41ms(nums []int) int {
	// 128
	// first num = without predecessor, for m[num+right+1] exists; right++ {}
	maxLen, m := 0, make(map[int]struct{})
	for _, num := range nums {
		m[num] = struct{}{}
	}
	for num, _ := range m {
		if _, found := m[num-1]; found {
			continue
		}
		right := 0
		for found := true; found; right++ {
			_, found = m[num+right+1]
		}
		if right > maxLen {
			maxLen = right
		}
	}
	return maxLen
}

func longestConsecutive(nums []int) int {
	// 128
	// ignore O(n), just sort it :D
	if len(nums) == 0 {
		return 0
	}
	slices.Sort(nums)
	prev, curLen, maxLen := nums[0], 1, 1
	for i := 1; i < len(nums); i++ {
		curr := nums[i]
		if curr == prev+1 {
			curLen++
		} else if curr != prev {
			if curLen > maxLen {
				maxLen = curLen
			}
			curLen = 1
		}
		prev = curr
	}
	if curLen > maxLen {
		return curLen
	}
	return maxLen
}

func findPeakElement162(nums []int) int {
	// 162
	minus2, minus1 := math.MinInt, nums[0]
	for i := 1; i < len(nums); i++ {
		current := nums[i]
		if minus2 < minus1 && minus1 > current {
			return i - 1
		}
		minus2, minus1 = minus1, current
	}
	if minus2 < minus1 {
		return len(nums) - 1
	}
	return -1
}

func constructDistancedSequence1718(n int) []int {
	// 1718
	// . . . . .
	// 3 . . 3
	// 2 . 2
	// . . . . . . . . .
	// 5 . . . . 5 . . .
	// 5 3 . . 3 5 . . .
	// 5 3 4 . 3 5 4 . . cannot put 2
	// 5 3 1 4 3 5 2 4 2
	result := make([]int, 2*n-1)
	used := make([]bool, n+1)

	var backtrack func(int) bool
	backtrack = func(index int) bool {
		if index == len(result) {
			return true
		}
		if result[index] != 0 {
			return backtrack(index + 1)
		}
		for num := n; num > 0; num-- {
			if used[num] {
				continue
			}
			if num == 1 {
				result[index] = 1
				used[1] = true
				if backtrack(index + 1) {
					return true
				}
				result[index] = 0
				used[1] = false
			} else if index+num < len(result) && result[index+num] == 0 {
				result[index] = num
				result[index+num] = num
				used[num] = true
				if backtrack(index + 1) {
					return true
				}
				result[index] = 0
				result[index+num] = 0
				used[num] = false
			}
		}
		return false
	}

	backtrack(0)
	return result
}

func stoneGameII1140(piles []int) int {
	// 1140
	// 2,7,9,4,4
	// 2M=2
	// 2,9 9,4
	// 2M=2??
	panic("not understood")
}

func averageWaitingTime1701(customers [][]int) float64 {
	// 1701
	waitSum, available := 0, 0
	for _, c := range customers {
		start, duration := c[0], c[1]
		if start > available {
			available = start + duration
			waitSum += duration
		} else {
			available += duration
			waitSum += available - start
		}
	}
	return float64(waitSum) / float64(len(customers))
}

func minimumArea3195(grid [][]int) int {
	// 3195
	// 0 1 0
	// 1 0 1
	// 1 0
	// 0 0
	iMin, iMax, jMin, jMax := math.MaxInt, 0, math.MaxInt, 0
	for i, row := range grid {
		for j, cell := range row {
			if cell != 1 {
				continue
			}
			if i < iMin {
				iMin = i
			}
			if i > iMax {
				iMax = i
			}
			if j < jMin {
				jMin = j
			}
			if j > jMax {
				jMax = j
			}
		}
	}
	fmt.Println(iMin, iMax, jMin, jMax)
	return (jMax - jMin + 1) * (iMax - iMin + 1)
}

func appendCharacters2486(s string, t string) int {
	// 2486
	// accoachingd coding --> ac(co)aching(d)ing
	si, ti, tmax := 0, 0, 0
	for si < len(s) && ti < len(t) {
		if s[si] == t[ti] {
			si++
			ti++
			if ti > tmax {
				tmax = ti
			}
		} else {
			si++
		}
	}
	fmt.Println(si, ti, tmax)
	return len(t) - tmax
}

func occurrencesOfElement3159_9ms(nums []int, queries []int, x int) []int {
	// 3159
	// 1,3,1,7
	// x=1 1,3,2,4
	rs, pos, posLen := make([]int, len(queries)), make([]int, len(nums)+1), 1
	for i, num := range nums {
		if num == x {
			pos[posLen], posLen = i+1, posLen+1
		}
	}
	fmt.Println(pos)
	for i, q := range queries {
		if q >= len(pos) {
			rs[i] = -1
		} else if p := pos[q]; p > 0 {
			rs[i] = p - 1
		} else {
			rs[i] = -1
		}
	}
	return rs
}

func occurrencesOfElement3159(nums []int, queries []int, x int) []int {
	// 3159 re-use both nums and queries
	posLen := 0
	for i, num := range nums {
		if num == x {
			nums[posLen], posLen = i, posLen+1
		}
	}
	fmt.Println(nums)
	for i, q := range queries {
		if q <= posLen {
			queries[i] = nums[q-1]
		} else {
			queries[i] = -1
		}
	}
	return queries
}

func findWinners2225_527ms(matches [][]int) [][]int {
	// 2225
	win, lostOnce, lostCntById, allId := []int{}, []int{}, make(map[int]int), make(map[int]struct{})
	for _, m := range matches {
		lostCntById[m[1]]++
		allId[m[0]] = struct{}{}
	}
	for id, cnt := range lostCntById {
		if cnt == 1 {
			idx, _ := slices.BinarySearch(lostOnce, id)
			lostOnce = slices.Insert(lostOnce, idx, id)
		}
		delete(allId, id)
	}
	for id, _ := range allId {
		idx, _ := slices.BinarySearch(win, id)
		win = slices.Insert(win, idx, id)
	}
	return [][]int{win, lostOnce}
}

func findWinners2225_125ms(matches [][]int) [][]int {
	// 2225
	win, lostOnce, lostCntById, allId := []int{}, []int{}, make([]int, 100001), make(map[int]struct{})
	for _, m := range matches {
		lostCntById[m[1]]++
		allId[m[0]] = struct{}{}
	}
	fmt.Println(allId, lostCntById)
	for id := 1; id < len(lostCntById); id++ {
		cnt := lostCntById[id]
		if cnt == 0 {
			continue
		}
		if cnt == 1 {
			idx, _ := slices.BinarySearch(lostOnce, id)
			lostOnce = slices.Insert(lostOnce, idx, id)
		}
		delete(allId, id)
	}
	fmt.Println(allId)
	for id, _ := range allId {
		idx, _ := slices.BinarySearch(win, id)
		win = slices.Insert(win, idx, id)
	}
	return [][]int{win, lostOnce}
}

func findWinners2225(matches [][]int) [][]int {
	// 2225  increasing order --> array  55ms
	lostCntById := make([]int, 100001) // 0 no player 1+ loss count -1 no loss
	rs := [][]int{
		[]int{},
		[]int{},
	}
	for _, m := range matches {
		win, loss := m[0], m[1]
		if lostCntById[win] == 0 {
			lostCntById[win] = -1
		}
		if lostCntById[loss] == -1 {
			lostCntById[loss] = 1
		} else {
			lostCntById[loss]++
		}
	}
	for id, cnt := range lostCntById {
		if cnt == -1 {
			rs[0] = append(rs[0], id)
		} else if cnt == 1 {
			rs[1] = append(rs[1], id)
		}
	}
	return rs
}

func minSteps2186(s string, t string) int {
	// 2186
	// leetcode      c1e3d1l1o1t1
	// coats     a-1   e3d1l1    s-1
	count, m := 0, [26]int{}
	for _, c := range s {
		m[c-'a']++
	}
	for _, c := range t {
		m[c-'a']--
	}
	for _, c := range m {
		if c >= 0 {
			count += c
		} else {
			count -= c
		}
	}
	return count
}

func main() {
	fmt.Println(findWinners2225([][]int{{1, 3}, {2, 3}, {3, 6}, {5, 6}, {5, 7}, {4, 5}, {4, 8}, {4, 9}, {10, 4}, {10, 9}}))
	// fmt.Println(occurrencesOfElement3159([]int{1, 1, 3, 1, 1, 3, 2, 1}, []int{3}, 3))
	// fmt.Println(occurrencesOfElement3159([]int{1, 3, 1, 7}, []int{1, 3, 2, 4}, 1))
	// fmt.Println(occurrencesOfElement3159([]int{1, 2, 3}, []int{10}, 5))
	// fmt.Println(smallestNumber2375("IIIDIDDD"))
	// fmt.Println(appendCharacters2486("accoachingd", "coding"))
	// fmt.Println(minimumArea3195([][]int{{0}, {1}}))
	// fmt.Println(minimumArea3195([][]int{{0,1,0},{1,0,1}}))
	// fmt.Println(minimumArea3195([][]int{{1,0},{0,0}}))
	// fmt.Println(averageWaitingTime1701([][]int{{5,2},{5,4},{10,3},{20,1}}))
	// fmt.Println(averageWaitingTime1701([][]int{{1,2},{2,5},{4,3}}))
	// fmt.Println(constructDistancedSequence1718(5))
	// fmt.Println(findPeakElement162([]int{1}))
	// fmt.Println(findPeakElement162([]int{1,4}))
	// fmt.Println(findPeakElement162([]int{1,2,1,3,5,6,4}))
	// fmt.Println(longestConsecutive128([]int{100, 4, 200, 1, 3, 2}))
	// fmt.Println(longestConsecutive128([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}))
	// fmt.Println(punishmentNumber2698(37))
	// var isValidPartition_6ms func(string, int) bool
	// isValidPartition_6ms = func(s string, target int) bool {
	// 	if len(s) == 0 {
	// 		return target == 0
	// 	}
	// 	for i := 1; i <= len(s); i++ {
	// 		num, _ := strconv.Atoi(s[:i])
	// 		if target >= num && isValidPartition_6ms(s[i:], target-num) {
	// 			return true
	// 		}
	// 	}
	// 	return false
	// }
	// fmt.Println(isValidPartition_6ms("1296", 36))
	// ProductOfNumbers productOfNumbers = new ProductOfNumbers()
	// productOfNumbers.add(3)        // [3]
	// productOfNumbers.add(0)        // [3,0]
	// productOfNumbers.add(2)        // [3,0,2]
	// productOfNumbers.add(5)        // [3,0,2,5]
	// productOfNumbers.add(4)        // [3,0,2,5,4]
	// productOfNumbers.getProduct(2) // return 20. The product of the last 2 numbers is 5 * 4 = 20
	// productOfNumbers.getProduct(3) // return 40. The product of the last 3 numbers is 2 * 5 * 4 = 40
	// productOfNumbers.getProduct(4) // return 0. The product of the last 4 numbers is 0 * 2 * 5 * 4 = 0
	// productOfNumbers.add(8)        // [3,0,2,5,4,8]
	// productOfNumbers.getProduct(2) // return 32. The product of the last 2 numbers is 4 * 8 = 32
	// fmt.Println('0'-'A', '9'-'A', 'z'-'0', 'a'-'A', 'Z'-'A')
	// fmt.Println(productExceptSelf238([]int{1, 2, 3, 4}))
	// fmt.Println(isValidSudoku36([][]byte{
	// 	{'8', '3', '.', '.', '7', '.', '.', '.', '.'},
	// 	{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
	// 	{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
	// 	{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
	// 	{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
	// 	{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
	// 	{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
	// 	{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
	// 	{'.', '.', '.', '.', '8', '.', '.', '7', '9'}}))
	// fmt.Println(minOperations3066([]int{2,11,10,1,3}, 10)) // 2
	// fmt.Println(minOperations3066([]int{97,73,5,78}, 98)) // 3
	// fmt.Println(frequencySort451("Aabb"))                        // bbAa
	// fmt.Println(frequencySort451("2a554442f544asfasssffffasss")) //
	// fmt.Println(maxIceCream1833([]int{1,3,2,4,1}, 7))
	// fmt.Println(maximumSum2342([]int{18,43,36,13,7,16}))
	// fmt.Println(maximumSum2342([]int{10,12,19,14}))
	// fmt.Println(canVisitAllRooms841([][]int{{1}, {2}, {3}, {}}))
	// fmt.Println(canVisitAllRooms841([][]int{{1, 3}, {3, 0, 1}, {2}, {0}}))
	// fmt.Println(compress443([]byte("aabccccccccccd")))
	// fmt.Println(countBadPairs2364([]int{4, 1, 3, 3}))
	// fmt.Println(countBadPairs2364([]int{2, 3, 6, 5, 6, 9, 8}))
	// fmt.Println(partitionArray2294([]int{16,8,17,0,3,17,8,20}, 10)) // 0,3,8 16,17,20
	// fmt.Println(partitionArray2294([]int{3,6,1,2,5,4}, 2))
	// fmt.Println(partitionArray2294([]int{5,16,3,20,9,20,16,19,6}, 4))
	// nc := Constructor()
	// fmt.Println(nc.Find(10)) // There is no index that is filled with number 10. Therefore, we return -1.
	// fmt.Println(nc.Change(2, 10)) // Your container at index 2 will be filled with number 10.
	// fmt.Println(nc.Change(1, 10)) // Your container at index 1 will be filled with number 10.
	// fmt.Println(nc.Change(3, 10)) // Your container at index 3 will be filled with number 10.
	// fmt.Println(nc.Change(5, 10)) // Your container at index 5 will be filled with number 10.
	// fmt.Println(nc.Find(10)) // Number 10 is at the indices 1, 2, 3, and 5. Since the smallest index that is filled with 10 is 1, we return 1.
	// fmt.Println(nc.Change(1, 20)) // Your container at index 1 will be filled with number 20. Note that index 1 was filled with 10 and then replaced with 20.
	// fmt.Println(nc.Find(10)) // Number 10 is at the indices 2, 3, and 5. The smallest index that is filled with 10 is 2. Therefore, we return 2.
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
