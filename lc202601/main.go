package main

import (
	"container/heap"
	"dsa/kit"
	"fmt"
	"math"
)

type pairItem struct {
	sum  int
	u, v int
}

type priorityQueue []*pairItem

func (pq priorityQueue) Len() int { return len(pq) }

func (pq priorityQueue) Less(i, j int) bool {
	if pq[i].sum == pq[j].sum {
		return pq[i].u < pq[j].u
	}
	return pq[i].sum < pq[j].sum
}

func (pq priorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *priorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*pairItem))
}

func (pq *priorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	*pq = old[0 : n-1]
	return item
}

type node struct {
	val     int
	prev    int
	next    int
	removed bool
}

func MinimumPairRemoval(nums []int) int {
	// 3507
	n := len(nums)
	if n <= 1 {
		return 0
	}

	nodes := make([]node, n)
	for i := 0; i < n; i++ {
		nodes[i] = node{
			val:  nums[i],
			prev: i - 1,
			next: i + 1,
		}
	}
	nodes[n-1].next = -1

	descCount := 0
	for i := 0; i < n-1; i++ {
		if nums[i] > nums[i+1] {
			descCount++
		}
	}

	pq := &priorityQueue{}
	heap.Init(pq)

	for i := 0; i < n-1; i++ {
		heap.Push(pq, &pairItem{
			sum: nums[i] + nums[i+1],
			u:   i,
			v:   i + 1,
		})
	}

	ops := 0
	for descCount > 0 && pq.Len() > 0 {
		item := heap.Pop(pq).(*pairItem)
		u, v := item.u, item.v

		if nodes[u].removed || nodes[v].removed {
			continue
		}
		if nodes[u].next != v {
			continue
		}
		if nodes[u].val+nodes[v].val != item.sum {
			continue
		}

		ops++

		w := nodes[v].next
		p := nodes[u].prev

		if p != -1 && nodes[p].val > nodes[u].val {
			descCount--
		}
		if nodes[u].val > nodes[v].val {
			descCount--
		}
		if w != -1 && nodes[v].val > nodes[w].val {
			descCount--
		}

		nodes[u].val += nodes[v].val
		nodes[u].next = w
		nodes[v].removed = true
		if w != -1 {
			nodes[w].prev = u
		}

		if p != -1 && nodes[p].val > nodes[u].val {
			descCount++
		}
		if w != -1 && nodes[u].val > nodes[w].val {
			descCount++
		}

		if p != -1 {
			heap.Push(pq, &pairItem{
				sum: nodes[p].val + nodes[u].val,
				u:   p,
				v:   u,
			})
		}
		if w != -1 {
			heap.Push(pq, &pairItem{
				sum: nodes[u].val + nodes[w].val,
				u:   u,
				v:   w,
			})
		}
	}

	return ops
}

func CountConsistentStrings(allowed string, words []string) (cnt int) {
	// 1684
	allowedChar := [26]bool{}
	for _, char := range allowed {
		allowedChar[char-'a'] = true
	}
	for _, word := range words {
		containsAllowed := true
		for _, char := range word {
			if !allowedChar[char-'a'] {
				containsAllowed = false
				break
			}
		}
		if containsAllowed {
			cnt++
		}
	}
	return
}

func ReverseDegree(s string) (total int) {
	// 3498
	for i := range s {
		total += (i + 1) * (26 - int(s[i]-'a'))
	}
	return
}

func GetConcatenation(nums []int) []int {
	// 1929
	n := len(nums)
	result := make([]int, 2*n)
	for i, num := range nums {
		result[i] = num
		result[n+i] = num
	}
	return result
}

func RecoverOrder(order []int, friends []int) []int {
	// 3668
	// 1,3,4    3̲, 1̲, 2, 5, 4̲
	friendIds := [101]bool{} // faster than map
	for _, id := range friends {
		friendIds[id] = true
	}
	friendOrder := make([]int, 0, len(order))
	for _, id := range order {
		if friendIds[id] {
			friendOrder = append(friendOrder, id)
		}
	}
	return friendOrder
}

func TheMaximumAchievableX(num int, t int) int {
	// 2769
	return num + 2*t
}

func AddStrings(num1 string, num2 string) string {
	// 415
	i, j := len(num1)-1, len(num2)-1
	carry := 0
	var result []byte
	for i >= 0 || j >= 0 || carry > 0 {
		var n1, n2 int
		if i >= 0 {
			n1 = int(num1[i] - '0')
			i--
		}
		if j >= 0 {
			n2 = int(num2[j] - '0')
			j--
		}
		sum := n1 + n2 + carry
		carry = sum / 10
		result = append(result, byte(sum%10+'0'))
	}
	// Reverse
	n := len(result)
	for k := 0; k < n/2; k++ {
		result[k], result[n-1-k] = result[n-1-k], result[k]
	}
	return string(result)
}

func ThirdMax(nums []int) int {
	// 414
	max1 := math.MinInt
	max2 := math.MinInt
	max3 := math.MinInt
	for _, n := range nums {
		if n > max1 {
			max3 = max2
			max2 = max1
			max1 = n
		} else if n > max2 && n != max1 {
			max3 = max2
			max2 = n
		} else if n > max3 && n != max1 && n != max2 {
			max3 = n
		}
	}
	if max3 == math.MinInt {
		return max1
	}
	return max3
}

func FizzBuzz(n int) []string {
	// 412
	result := make([]string, n)
	for i := 1; i <= n; i++ {
		if i%3 == 0 {
			if i%5 == 0 {
				result[i-1] = "FizzBuzz"
			} else {
				result[i-1] = "Fizz"
			}
		} else {
			if i%5 == 0 {
				result[i-1] = "Buzz"
			} else {
				result[i-1] = fmt.Sprintf("%d", i)
			}
		}
	}
	return result
}

func MinTimeToVisitAllPoints(points [][]int) (totalTime int) {
	// 1266
	// diagonal is faster
	for i := 0; i < len(points)-1; i++ {
		dx := points[i+1][0] - points[i][0]
		if dx < 0 {
			dx = -dx
		}
		dy := points[i+1][1] - points[i][1]
		if dy < 0 {
			dy = -dy
		}
		if dx > dy {
			totalTime += dx
		} else {
			totalTime += dy
		}
	}
	return totalTime
}

func LongestPalindrome(s string) (longest int) {
	// 409
	countByChar := [58]int{} // A to z, faster than map
	for _, char := range s {
		countByChar[char-'A']++
	}
	middle := 0
	for _, cnt := range countByChar {
		if cnt%2 == 0 {
			longest += cnt
		} else {
			middle = 1
			if cnt >= 3 {
				longest += cnt - 1
			}
		}
	}
	return longest + middle
}

func MinimumDeleteSum(s1 string, s2 string) int {
	// 712
	sum := 0
	for _, c := range s1 {
		sum += int(c)
	}
	for _, c := range s2 {
		sum += int(c)
	}

	n := len(s2)
	dp := make([]int, n+1)

	for _, c1 := range s1 {
		diag := 0
		for j, c2 := range s2 {
			temp := dp[j+1]
			if c1 == c2 {
				dp[j+1] = diag + int(c1)
			} else {
				if dp[j+1] < dp[j] {
					dp[j+1] = dp[j]
				}
			}
			diag = temp
		}
	}
	return sum - 2*dp[n]
}

func IsSubsequence(s string, t string) bool {
	// 392
	j, lent := 0, len(t)
	for _, char := range s {
		for ; j < lent && rune(t[j]) != char; j++ {
		}
		if j >= len(t) {
			return false
		}
		j++
	}
	return true
}

func FindTheDifference(s string, t string) byte {
	// 389
	countByChar := [26]int{} // array of 26 chars is faster than map
	for _, char := range s {
		countByChar[char-'a']++
	}
	for _, char := range t {
		countByChar[char-'a']--
		if countByChar[char-'a'] < 0 {
			return byte(char)
		}
	}
	return 0
}

func FirstUniqChar(s string) int {
	// 387
	// a     a:1
	// ab    a:1,b:2
	// abc   a:1,b:2,c:3
	// abca  a:max,b:2,c:3
	firstIndexByChar := [26]int{} // array of 26 chars is faster than map
	for i, char := range s {
		if firstIndexByChar[char-'a'] == 0 {
			firstIndexByChar[char-'a'] = i + 1 // use 1-based index to avoid default value 0
		} else {
			firstIndexByChar[char-'a'] = math.MaxInt
		}
	}
	firstIndex := math.MaxInt
	for _, index := range firstIndexByChar {
		if index > 0 && index < firstIndex {
			firstIndex = index
		}
	}
	if firstIndex == math.MaxInt {
		return -1
	}
	return firstIndex - 1
}

func CanConstruct(ransomNote string, magazine string) bool {
	// 383
	countByChar := [26]int{} // array of 26 chars is faster than map
	for _, char := range magazine {
		countByChar[char-'a']++
	}
	for _, char := range ransomNote {
		countByChar[char-'a']--
		if countByChar[char-'a'] < 0 {
			return false
		}
	}
	return true
}

func SumFourDivisors(nums []int) (result int) {
	// 1390
	for _, num := range nums {
		maxDivisor := int(math.Sqrt(float64(num)))
		divisorCount, divisorSum := 2, 1+num // 1 and num are always divisors
		for i := 2; i <= maxDivisor; i++ {
			if num%i == 0 {
				if i*i == num {
					divisorCount += 1
					divisorSum += i
				} else {
					divisorCount += 2
					divisorSum += i + num/i
				}
			}
		}
		if divisorCount == 4 {
			result += divisorSum
		}
	}
	return
}

func RepeatedNTimes(nums []int) int {
	// 961
	n := len(nums) / 2
	numMap := make(map[int]struct{}, n+1)
	for i := 0; i < n+2; i++ {
		if _, found := numMap[nums[i]]; found {
			return nums[i]
		}
		numMap[nums[i]] = struct{}{}
	}
	return 0
}

func ReverseVowels(s string) string {
	// 345
	vowels := map[rune]bool{
		'a': true, 'e': true, 'i': true, 'o': true, 'u': true,
		'A': true, 'E': true, 'I': true, 'O': true, 'U': true,
	}
	runeSlice := []rune(s)
	i1, i2 := 0, len(runeSlice)-1
	kit.DebugLog("Initial:", string(runeSlice), "i1:", i1, "i2:", i2)
	for i1 < i2 {
		kit.DebugLog("Before iteration: ", "runeSlice:", string(runeSlice), "i1:", i1, "i2:", i2)
		if !vowels[runeSlice[i1]] {
			kit.DebugLog("runeSlice[i1] not a vowel:", string(runeSlice[i1]))
			i1++
		} else if !vowels[runeSlice[i2]] {
			kit.DebugLog("runeSlice[i2] not a vowel:", string(runeSlice[i2]))
			i2--
		} else {
			kit.DebugLog("Swapping:", string(runeSlice[i1]), string(runeSlice[i2]))
			runeSlice[i1], runeSlice[i2] = runeSlice[i2], runeSlice[i1]
			i1++
			i2--
		}
		kit.DebugLog("After iteration: ", "runeSlice:", string(runeSlice), "i1:", i1, "i2:", i2)
	}
	kit.DebugLog("Final:", string(runeSlice))
	return string(runeSlice)
}

// MinBitwiseArray returns an array ans such that ans[i] OR (ans[i] + 1) == nums[i]
// and ans[i] is minimized. If no such ans[i] exists, ans[i] = -1.
// nums consists of prime integers.
func MinBitwiseArray(nums []int) []int {
	// 3314
	ans := make([]int, len(nums))
	for i, num := range nums {
		if num == 2 {
			ans[i] = -1
		} else {
			// num is prime and > 2, so it's odd.
			// This means the least significant bit (bit 0) is 1.
			// We find the first zero bit starting from LSB.
			p := 0
			for ((num >> p) & 1) == 1 {
				p++
			}
			// The sequence of trailing ones is from bit 0 to p-1.
			// To minimize the result x, where x | (x+1) == num,
			// we set x = num - 2^(k) where k is the index of the most significant bit
			// in the trailing ones sequence. Here k = p-1.
			ans[i] = num - (1 << (p - 1))
		}
	}
	return ans
}
