package main

//lint:file-ignore U1000 Ignore all unused code, it's generated

import (
	"cmp"
	"log"
	"math"
	"os"
	"runtime"
	"slices"
	"strings"
	"unicode"
)

var (
	debugEnabled = os.Getenv("DEBUG") == "true"
	debugLogger  = log.New(os.Stdout, "DEBUG: ", log.Ldate|log.Ltime)
)

func debugLog(v ...any) {
	if debugEnabled {
		pc, _, _, ok := runtime.Caller(1)
		if !ok {
			debugLogger.Println(v...)
			return
		}
		if fn := runtime.FuncForPC(pc); fn != nil {
			name := fn.Name()
			if lastDot := strings.LastIndex(name, "."); lastDot != -1 {
				name = name[lastDot+1:]
			}
			args := append([]any{name + ":"}, v...)
			debugLogger.Println(args...)
		}
	}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func PossibleStringCount_fail(word string, k int) (cnt int) {
	// 3333
	const MOD = 1e9 + 7
	n := len(word)

	if n == 0 {
		if k == 0 {
			return 1
		}
		return 0
	}

	var groupLengths []int
	count := 1
	for i := 1; i < n; i++ {
		if word[i] == word[i-1] {
			count++
		} else {
			groupLengths = append(groupLengths, count)
			count = 1
		}
	}
	groupLengths = append(groupLengths, count)

	dp := make([]int64, n+1)
	dp[0] = 1

	for _, l := range groupLengths {
		nextDp := make([]int64, n+1)
		currentWindowSum := int64(0)

		// Calculate nextDp[1]
		// nextDp[1] = dp[0] (if l >= 1)
		if l >= 1 && dp[0] > 0 {
			currentWindowSum = (currentWindowSum + dp[0]) % MOD
		}
		nextDp[1] = currentWindowSum

		// Calculate nextDp for i from 2 to n
		for i := 2; i <= n; i++ {
			// Add dp[i-1] to window
			currentWindowSum = (currentWindowSum + dp[i-1]) % MOD
			// Remove dp[i-l-1] from window if it's outside the window
			if i-l-1 >= 0 {
				currentWindowSum = (currentWindowSum - dp[i-l-1] + MOD) % MOD
			}
			nextDp[i] = currentWindowSum
		}
		dp = nextDp
	}

	var totalCnt int64
	for j := k; j <= n; j++ {
		totalCnt = (totalCnt + dp[j]) % MOD
	}

	return int(totalCnt)
}

func KthCharacter(k int) byte {
	// 3304
	// b  ab  abbc  abbcbccd
	word := []int{0}
	for len(word) < k {
		m := len(word)
		for i := 0; i < m; i++ {
			word = append(word, (word[i]+1)%26)
		}
	}
	return 'a' + byte(word[k-1])
}

func KthCharacterII(k int64, operations []int) byte {
	// 3307
	k--
	inc := 0
	for i := len(operations) - 1; i >= 0; i-- {
		if i >= 63 {
			// k is at most 10^18, which is less than 2^63.
			// 2^i for i>=63 will be larger than k, so k >= 2^i is false.
			// We can skip these iterations.
			continue
		}
		length := int64(1) << i
		if k >= length {
			k -= length
			inc = (inc + operations[i]) % 26
		}
	}
	return 'a' + byte(inc)
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// 235
	if root == nil || root == p || root == q {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left != nil && right != nil {
		return root
	}
	if left != nil {
		return left
	}
	return right
}

func canFinish(numCourses int, prerequisites [][]int) bool {
	// 207
	graph := make([][]int, numCourses)
	for _, pre := range prerequisites {
		to, from := pre[0], pre[1]
		graph[from] = append(graph[from], to)
	}
	visited := make([]int, numCourses) // 0=unvisited, 1=visiting, 2=visited

	var hasCycle func(int) bool
	hasCycle = func(node int) bool {
		if visited[node] == 1 {
			return true // found a cycle
		}
		if visited[node] == 2 {
			return false // already checked, no cycle
		}
		visited[node] = 1 // mark as visiting
		for _, nei := range graph[node] {
			if hasCycle(nei) {
				return true
			}
		}
		visited[node] = 2 // mark as visited
		return false
	}

	for i := 0; i < numCourses; i++ {
		if hasCycle(i) {
			return false
		}
	}
	return true
}

func canFinish_slower(numCourses int, prerequisites [][]int) bool {
	// 207
	// Topological Sort (Kahn's Algorithm)
	// Build the adjacency list and in-degree array.
	graph := make([][]int, numCourses)
	inDegree := make([]int, numCourses)
	for _, pre := range prerequisites {
		to, from := pre[0], pre[1]
		graph[from] = append(graph[from], to)
		inDegree[to]++
	}
	// Collect nodes with in-degree 0.
	queue := make([]int, 0)
	for i := 0; i < numCourses; i++ {
		if inDegree[i] == 0 {
			queue = append(queue, i)
		}
	}
	visited := 0
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		visited++
		for _, neighbor := range graph[curr] {
			inDegree[neighbor]--
			if inDegree[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}
	return visited == numCourses
}

func WordPattern(pattern string, s string) bool {
	// 290
	words := strings.Fields(s)
	if len(words) != len(pattern) {
		return false
	}
	charToWord := make(map[byte]string)
	wordToChar := make(map[string]byte)

	for i := 0; i < len(pattern); i++ {
		pChar := pattern[i]
		word := words[i]

		if mappedWord, ok := charToWord[pChar]; ok {
			if mappedWord != word {
				return false
			}
		} else {
			if _, ok := wordToChar[word]; ok {
				return false
			}
			charToWord[pChar] = word
			wordToChar[word] = pChar
		}
	}

	return true
}

func ScoreOfString(s string) (score int) {
	// 3110
	prevChar := s[0]
	for i := 1; i < len(s); i++ {
		char := s[i]
		if char > prevChar {
			score += int(char - prevChar)
		} else {
			score += int(prevChar - char)
		}
		debugLog(char, prevChar, char-prevChar, score)
		prevChar = char
	}
	return score
}

func FindLucky(arr []int) int {
	// 1394
	cnts := [501]int{}
	for _, num := range arr {
		cnts[num]++
	}
	for num := 500; num > 0; num-- {
		if cnts[num] == num {
			return num
		}
	}
	return -1
}

// 1865
type FindSumPairs struct {
	nums1 []int
	nums2 []int
	cnt2  map[int]int
}

func Constructor(nums1 []int, nums2 []int) FindSumPairs {
	cnt2 := make(map[int]int)
	for _, v := range nums2 {
		cnt2[v]++
	}
	return FindSumPairs{
		nums1: nums1,
		nums2: nums2,
		cnt2:  cnt2,
	}
}

func (obj *FindSumPairs) Add(index int, val int) {
	oldVal := obj.nums2[index]
	obj.cnt2[oldVal]--
	obj.nums2[index] += val
	obj.cnt2[obj.nums2[index]]++
}

func (obj *FindSumPairs) Count(tot int) int {
	count := 0
	for _, v := range obj.nums1 {
		need := tot - v
		count += obj.cnt2[need]
	}
	return count
}

func MaxEvents(events [][]int) int {
	// 1353
	slices.SortFunc(events, func(a, b []int) int { return cmp.Compare(a[0], b[0]) })
	debugLog(events)
	return 0
}

func MatchPlayersAndTrainers(players []int, trainers []int) (cnt int) {
	// 2410
	// 1 2 5 8 8  trainers
	//     4 7 9  players
	slices.Sort(players)
	slices.Sort(trainers)

	for i, j := len(players)-1, len(trainers)-1; i >= 0 && j >= 0; i-- {
		debugLog(j, i, trainers[j], players[i])
		if trainers[j] >= players[i] {
			j--
			cnt++
		}
	}
	return cnt
}

func GetDecimalValue(head *ListNode) (num int) {
	// 1290
	for ; head != nil; head = head.Next {
		num = 2*num + head.Val
	}
	return num
}

func IsValid(word string) bool {
	// 3136
	if len(word) < 3 {
		return false
	}
	hasVowel := false
	hasConsonant := false
	vowels := "aeiouAEIOU"
	for _, ch := range word {
		if !unicode.IsLetter(ch) && !unicode.IsDigit(ch) {
			return false
		}
		if unicode.IsLetter(ch) {
			if strings.ContainsRune(vowels, ch) {
				hasVowel = true
			} else {
				hasConsonant = true
			}
		}
	}
	return hasVowel && hasConsonant
}

func MaximumLength_fail(nums []int) (maxLen int) {
	// 3201
	// all even elements, all odd elements, alternate even odd, or alternate odd even elements
	allEvenLen, allOddLen := 0, 0
	prev := nums[0]
	if prev%2 == 0 {
		allEvenLen = 1
	} else {
		allOddLen = 1
	}
	debugLog(prev, allEvenLen, allOddLen, maxLen)
	for _, num := range nums[1:] {
		debugLog(num, allEvenLen, allOddLen, maxLen)
		if num%2 == 0 {
			allOddLen = 0
			allEvenLen++
		} else {
			allEvenLen = 0
			allOddLen++
		}
		maxLen = max(maxLen, allEvenLen, allOddLen)
		prev = num
	}
	return maxLen
}

func RemoveSubfolders(folder []string) []string {
	// 1233
	return []string{}
}

func MakeFancyString(s string) string {
	// 1957
	if len(s) < 3 {
		return s
	}
	result := make([]byte, len(s))
	first, second := s[0], s[1]
	result = []byte{first, second}
	for i := 2; i < len(s); i++ {
		third := s[i]
		if first != second || second != third {
			result = append(result, third)
		}
		first, second = second, third
	}
	return string(result)
}
func MaxSum(nums []int) int {
	// 3487
	// If the maximum element in the array is less than zero, the answer is the maximum element.
	// Otherwise, the answer is the sum of all unique values that are greater than or equal to zero.
	uniques := [101]bool{} // nums[i] <= 100
	maxNum := math.MinInt
	for _, num := range nums {
		if num > maxNum {
			maxNum = num
		}
		if num > 0 {
			uniques[num] = true
		}
	}
	if maxNum < 0 {
		return maxNum
	}
	sumUniques := 0
	for num := 0; num <= 100; num++ {
		if uniques[num] {
			sumUniques += num
		}
	}
	return sumUniques
}
