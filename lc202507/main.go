package main

//lint:file-ignore U1000 Ignore all unused code, it's generated

import (
	"log"
	"os"
	"strings"
)

var (
	debugEnabled = os.Getenv("DEBUG") == "true"
	debugLogger  = log.New(os.Stdout, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)
)

func debugLog(v ...any) {
	if debugEnabled {
		debugLogger.Println(v...)
	}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
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
