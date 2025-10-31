package main

//lint:file-ignore U1000 Ignore all unused code, it's generated

import (
	"container/heap"
	"log"
	"math"
	"math/bits"
	"os"
	"runtime"
	"sort"
	"strings"
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

func GetSneakyNumbers(nums []int) []int {
	// 3289
	exists := [100]bool{}
	first := -1
	for _, num := range nums {
		if exists[num] {
			if first == -1 {
				first = num
			} else {
				return []int{first, num}
			}
		} else {
			exists[num] = true
		}
	}
	return []int{0, 0}
}

func MinNumberOperations(target []int) int {
	// 1526
	result := target[0]
	for i := 1; i < len(target); i++ {
		if target[i] > target[i-1] {
			result += target[i] - target[i-1]
		}
	}
	return result
}

func CountValidSelections(nums []int) (cnt int) {
	// 3354
	n, totalSum := len(nums), 0
	prefixSum := make([]int, n+1)
	for i := 0; i < n; i++ {
		prefixSum[i+1] = prefixSum[i] + nums[i]
		totalSum += nums[i]
	}
	for s := 0; s < n; s++ {
		if nums[s] == 0 {
			l := prefixSum[s]
			diff := 2*l - totalSum
			switch diff {
			case 0:
				cnt += 2
			case 1, -1:
				cnt += 1
			}
		}
	}
	return
}

// 2043
type Bank struct {
	balance []int64
}

func Constructor(balance []int64) Bank {
	return Bank{balance: balance}
}

func (bank *Bank) Transfer(account1 int, account2 int, money int64) bool {
	n := len(bank.balance)
	if account1 < 1 || account1 > n || account2 < 1 || account2 > n {
		return false
	}
	if bank.balance[account1-1] < money {
		return false
	}
	bank.balance[account1-1] -= money
	bank.balance[account2-1] += money
	return true
}

func (bank *Bank) Deposit(account int, money int64) bool {
	n := len(bank.balance)
	if account < 1 || account > n {
		return false
	}
	bank.balance[account-1] += money
	return true
}

func (bank *Bank) Withdraw(account int, money int64) bool {
	n := len(bank.balance)
	if account < 1 || account > n {
		return false
	}
	if bank.balance[account-1] < money {
		return false
	}
	bank.balance[account-1] -= money
	return true
}

func TotalMoney(n int) (total int) {
	// 1716
	fullWeekCount := n / 7
	lastWeekDayCount := n % 7
	for weekNumber := 1; weekNumber <= fullWeekCount; weekNumber++ {
		total += 7*weekNumber + 1 + 2 + 3 + 4 + 5 + 6
	}
	for dayNumber := 1; dayNumber <= lastWeekDayCount; dayNumber++ {
		total += fullWeekCount + dayNumber
	}
	return
}

func isPerfectSquare(num int) bool {
	// 367 Newton-Raphson
	for guess, nextGuess := num, 0; ; guess = nextGuess {
		nextGuess = (guess + num/guess) / 2
		if nextGuess >= guess {
			return guess*guess == num
		}
	}
}

func FinalValueAfterOperations(operations []string) (X int) {
	// 2011
	valByOp := map[string]int{"X++": 1, "++X": 1, "X--": -1, "--X": -1}
	for _, op := range operations {
		X += valByOp[op]
	}
	return
}

func MaxDistinctElements(nums []int, k int) int {
	// 3397
	// Loop from smallest try to select the smallest possible new value that is
	// greater than the previously selected value. This maximizes the available
	// space for subsequent numbers.
	sort.Ints(nums)
	count := 1
	// For the first number, we select the smallest possible value.
	prev := nums[0] - k
	for i := 1; i < len(nums); i++ {
		num := nums[i]
		// The next value must be at least `prev + 1` to be distinct.
		// It also must be at least `num - k` to be reachable from `num`.
		// So we pick the larger of these two as our target.
		target := prev + 1
		if num-k > target {
			target = num - k
		}
		// If the target is reachable (within `num + k`), we've found a new distinct number.
		if target <= num+k {
			count++
			prev = target
		}
	}
	return count
}

func MaxIncreasingSubarrays(nums []int) int {
	// 3350
	n := len(nums)
	// leftRun[i] = length of strictly increasing subarray ending at i
	leftRun := make([]int, n)
	leftRun[0] = 1
	for i := 1; i < n; i++ {
		if nums[i] > nums[i-1] {
			leftRun[i] = leftRun[i-1] + 1
		} else {
			leftRun[i] = 1
		}
	}
	maxK := 0
	// rightRun is the length of the increasing subarray starting at i+1
	rightRun := 1
	// Loop from n-2 down to 0. The split is between i and i+1.
	for i := n - 2; i >= 0; i-- {
		// For the split between i and i+1:
		// The left part is an increasing subarray ending at i. Max length is leftRun[i].
		// The right part is an increasing subarray starting at i+1. Max length is rightRun.
		// Calculate potential k for this split.
		k := leftRun[i]
		if rightRun < k {
			k = rightRun
		}
		if k > maxK {
			maxK = k
		}
		// Update rightRun for the next iteration. It will become the run starting at i.
		if nums[i] < nums[i+1] {
			rightRun++
		} else {
			rightRun = 1
		}
	}
	return maxK
}

func HasIncreasingSubarrays(nums []int, k int) bool {
	// 3349
	n := len(nums)
	// `incRun[i]` will store the length of the strictly increasing subarray ending at index i.
	incRun := make([]int, n)
	incRun[0] = 1
	for i := 1; i < n; i++ {
		if nums[i] > nums[i-1] {
			incRun[i] = incRun[i-1] + 1
		} else {
			incRun[i] = 1
		}
	}
	// 1. `nums[a:a+k]` is increasing, which means `incRun[a+k-1] >= k`.
	// 2. `nums[a+k:a+2*k]` is increasing, which means `incRun[a+2*k-1] >= k`.
	for a := 0; a <= n-2*k; a++ {
		if incRun[a+k-1] >= k && incRun[a+2*k-1] >= k {
			return true
		}
	}
	return false
}

func HasIncreasingSubarrays_1(nums []int, k int) bool {
	// 3349
	n := len(nums)
	start := 0
	prev := nums[0]
	for i := 1; i < n; i++ {
		num := nums[i]
		if num <= prev {
			start = i
		}
		if i-start >= k {
			return true
		}
	}
	return false
}

func RemoveAnagrams(words []string) (result []string) {
	// 2273
	prev := words[0]
	prevCharCount := [26]int{}
	for _, c := range prev {
		prevCharCount[c-'a']++
	}
	result = append(result, prev)
	for _, word := range words {
		wordCharCount := [26]int{}
		for _, c := range word {
			wordCharCount[c-'a']++
		}
		if prevCharCount != wordCharCount {
			result = append(result, word)
			prev = word
			prevCharCount = wordCharCount
		}
	}
	return
}

const MOD = 1e9 + 7

var memo [51][51][51]map[int]int64
var C [51][51]int64
var fact, invFact [51]int64
var p [51][]int64
var n int
var Nums []int

func power_m(base, exp int64) int64 {
	var res int64 = 1
	base %= MOD
	for exp > 0 {
		if exp%2 == 1 {
			res = (res * base) % MOD
		}
		base = (base * base) % MOD
		exp /= 2
	}
	return res
}

func precompute(m int) {
	fact[0] = 1
	invFact[0] = 1
	for i := 1; i <= m; i++ {
		fact[i] = (fact[i-1] * int64(i)) % MOD
		invFact[i] = power_m(fact[i], MOD-2)
	}

	for i := 0; i <= m; i++ {
		C[i][0] = 1
		for j := 1; j <= i; j++ {
			C[i][j] = (C[i-1][j-1] + C[i-1][j]) % MOD
		}
	}

	for i := 0; i < n; i++ {
		p[i] = make([]int64, m+1)
		p[i][0] = 1
		for j := 1; j <= m; j++ {
			p[i][j] = (p[i][j-1] * int64(Nums[i])) % MOD
		}
	}
}

func solve(i, m, carry int) map[int]int64 {
	if i == n {
		if m == 0 {
			return map[int]int64{bits.OnesCount(uint(carry)): 1}
		} else {
			return map[int]int64{}
		}
	}
	if memo[i][m][carry] != nil {
		return memo[i][m][carry]
	}

	res := make(map[int]int64)
	for j := 0; j <= m; j++ {
		sub := solve(i+1, m-j, (carry+j)>>1)
		bit := (carry + j) & 1
		term := (p[i][j] * invFact[j]) % MOD
		for popcnt, v := range sub {
			res[popcnt+bit] = (res[popcnt+bit] + v*term) % MOD
		}
	}
	memo[i][m][carry] = res
	return res
}

func MagicalSum(m int, k int, nums []int) int {
	// 3539
	n = len(nums)
	Nums = nums
	for i := 0; i <= n; i++ {
		for j := 0; j <= m; j++ {
			for l := 0; l <= m; l++ {
				memo[i][j][l] = nil
			}
		}
	}
	for i := 0; i < 51; i++ {
		p[i] = make([]int64, 51)
	}

	precompute(m)
	ansMap := solve(0, m, 0)
	return int((ansMap[k] * fact[m]) % MOD)
}

func MaximumEnergy(energy []int, k int) int {
	// 3147
	n, maxEnergy := len(energy), math.MinInt
	for i := n - 1; i >= 0; i-- {
		if i+k < n {
			energy[i] += energy[i+k]
		}
		if energy[i] > maxEnergy {
			maxEnergy = energy[i]
		}
	}
	return maxEnergy
}

func MaximumEnergy_fail(energy []int, k int) int {
	// 3147
	sums := make([]int, k)
	for i, val := range energy {
		sums[i%k] += val
	}
	maxSum := sums[0]
	for _, sum := range sums[1:] {
		if sum > maxSum {
			maxSum = sum
		}
	}
	return maxSum
}

func MinTime(skill []int, mana []int) int64 {
	// 3494
	n := len(skill)
	m := len(mana)
	// C_prev stores the completion times of the previous potion (j-1) for each wizard
	C_prev := make([]int64, n)
	C_prev[0] = int64(skill[0]) * int64(mana[0])
	for i := 1; i < n; i++ {
		C_prev[i] = C_prev[i-1] + int64(skill[i])*int64(mana[0])
	}
	for j := 1; j < m; j++ {
		C_curr := make([]int64, n)
		// Calculate the start time for the current potion j on the first wizard.
		// This is the core of the "no-wait" logic, ensuring wizards are free.
		var startTime int64 = 0
		var p_sum_prefix int64 = 0
		for i := 0; i < n; i++ {
			term := C_prev[i] - p_sum_prefix
			if term > startTime {
				startTime = term
			}
			p_sum_prefix += int64(skill[i]) * int64(mana[j])
		}
		// Calculate completion times for the current potion j
		C_curr[0] = startTime + int64(skill[0])*int64(mana[j])
		for i := 1; i < n; i++ {
			C_curr[i] = C_curr[i-1] + int64(skill[i])*int64(mana[j])
		}
		C_prev = C_curr
	}
	return C_prev[n-1]
}

func successfulPairs(spells []int, potions []int, success int64) []int {
	// 2300 binary search and re-use
	sort.Ints(potions)
	for i := 0; i < len(spells); i++ {
		start, end := 0, len(potions)
		for start < end {
			mid := (start + end) / 2
			if int64(potions[mid])*int64(spells[i]) >= success {
				end = mid
			} else {
				start = mid + 1
			}
		}
		spells[i] = len(potions) - start
	}
	return spells
}

func SuccessfulPairs_1(spells []int, potions []int, success int64) []int {
	// 2300
	potionsLen := len(potions)
	sort.Ints(potions)
	result := make([]int, len(spells))
	for i, spell := range spells {
		if spell > 0 {
			minPotion := (success + int64(spell) - 1) / int64(spell) // better way to calculate ceiling
			idx := sort.Search(potionsLen, func(j int) bool {
				return int64(potions[j]) >= minPotion
			})
			result[i] = potionsLen - idx
		}
	}
	return result
}

func SuccessfulPairs_2(spells []int, potions []int, success int64) []int {
	// 2300
	type spellWithIndex struct {
		value int
		index int
	}
	sortedSpells := make([]spellWithIndex, len(spells))
	for i, s := range spells {
		sortedSpells[i] = spellWithIndex{value: s, index: i}
	}
	sort.Slice(sortedSpells, func(i, j int) bool {
		return sortedSpells[i].value < sortedSpells[j].value
	})
	sort.Ints(potions)
	result := make([]int, len(spells))
	potionsLen := len(potions)
	potionIndex := potionsLen - 1
	for _, s := range sortedSpells {
		// Move the potion pointer to the left as long as the current potion is strong enough.
		for potionIndex >= 0 && int64(s.value)*int64(potions[potionIndex]) >= success {
			potionIndex--
		}
		result[s.index] = potionsLen - (potionIndex + 1)
	}
	return result
}

func AvoidFlood(rains []int) []int {
	// 1488
	// The problem is to assign for each dry day (rains[i] == 0), which lake to dry.
	// A lake needs to be dried if it's full and it's going to rain on it again.
	// To make an optimal decision, when it's about to rain on a full lake,
	// we should choose the earliest possible dry day that occurred *after* the lake was last filled.
	// This leaves later dry days available for future potential floods.

	ans := make([]int, len(rains))
	// `fullLakes` maps a lake number to the index of the day it was last filled.
	fullLakes := make(map[int]int)
	// `dryDays` stores the indices of days with no rain, kept sorted.
	dryDays := []int{}

	for i, lake := range rains {
		if lake == 0 {
			// This is a dry day, add its index to our list of available dry days.
			dryDays = append(dryDays, i)
			// We'll decide what to do with this day later, so we tentatively set it to 1.
			ans[i] = 1
		} else {
			// This is a rainy day.
			ans[i] = -1
			if lastRainDay, ok := fullLakes[lake]; ok {
				// The lake is already full. We need to find a dry day to empty it.
				// The dry day must have occurred after the lake was last filled.
				// We use binary search to find the first available dry day after `lastRainDay`.
				dryDayIndex := sort.SearchInts(dryDays, lastRainDay+1)

				if dryDayIndex == len(dryDays) {
					// No available dry day found after the lake was filled. A flood is inevitable.
					return []int{}
				}

				// We found a suitable dry day. Let's use it.
				dayToDry := dryDays[dryDayIndex]
				ans[dayToDry] = lake

				// This dry day is now used, so we remove it from our list of available dry days.
				dryDays = append(dryDays[:dryDayIndex], dryDays[dryDayIndex+1:]...)
			}
			// Update the last rain day for this lake to the current day.
			fullLakes[lake] = i
		}
	}

	return ans
}

func SwimInWater(grid [][]int) int {
	// 778
	n := len(grid)
	pq := &MinHeap{}
	heap.Init(pq)
	visited := make([][]bool, n)
	for i := range visited {
		visited[i] = make([]bool, n)
	}
	// The "height" of a cell in the priority queue will be the maximum elevation
	// on the path to that cell.
	heap.Push(pq, &Cell{row: 0, col: 0, height: grid[0][0]})
	visited[0][0] = true
	directions := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	for pq.Len() > 0 {
		cell := heap.Pop(pq).(*Cell)
		r, c, t := cell.row, cell.col, cell.height
		if r == n-1 && c == n-1 {
			return t
		}
		for _, dir := range directions {
			nr, nc := r+dir[0], c+dir[1]
			if nr >= 0 && nr < n && nc >= 0 && nc < n && !visited[nr][nc] {
				visited[nr][nc] = true
				// The time (max elevation) to reach the neighbor is the maximum of the
				// current path's time and the neighbor's own elevation.
				newTime := t
				if grid[nr][nc] > newTime {
					newTime = grid[nr][nc]
				}
				heap.Push(pq, &Cell{row: nr, col: nc, height: newTime})
			}
		}
	}
	return -1
}

func PacificAtlantic(heights [][]int) [][]int {
	// 417
	rows, cols := len(heights), len(heights[0])
	pacificReachable := make([][]bool, rows)
	atlanticReachable := make([][]bool, rows)
	for i := range pacificReachable {
		pacificReachable[i] = make([]bool, cols)
		atlanticReachable[i] = make([]bool, cols)
	}

	var dfs func(r, c int, reachable [][]bool)
	dfs = func(r, c int, reachable [][]bool) {
		if reachable[r][c] {
			return
		}
		reachable[r][c] = true

		// up
		if r > 0 && heights[r-1][c] >= heights[r][c] {
			dfs(r-1, c, reachable)
		}
		// down
		if r < rows-1 && heights[r+1][c] >= heights[r][c] {
			dfs(r+1, c, reachable)
		}
		// left
		if c > 0 && heights[r][c-1] >= heights[r][c] {
			dfs(r, c-1, reachable)
		}
		// right
		if c < cols-1 && heights[r][c+1] >= heights[r][c] {
			dfs(r, c+1, reachable)
		}
	}

	// DFS from Pacific borders
	for r := 0; r < rows; r++ {
		dfs(r, 0, pacificReachable)
	}
	for c := 0; c < cols; c++ {
		dfs(0, c, pacificReachable)
	}

	// DFS from Atlantic borders
	for r := 0; r < rows; r++ {
		dfs(r, cols-1, atlanticReachable)
	}
	for c := 0; c < cols; c++ {
		dfs(rows-1, c, atlanticReachable)
	}

	var result [][]int
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if pacificReachable[r][c] && atlanticReachable[r][c] {
				result = append(result, []int{r, c})
			}
		}
	}

	return result
}

func TrapRainWater(heightMap [][]int) int {
	// 407
	m, n := len(heightMap), len(heightMap[0])
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}

	pq := &MinHeap{}
	heap.Init(pq)

	// Add border cells to the min-heap
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 || i == m-1 || j == 0 || j == n-1 {
				heap.Push(pq, &Cell{row: i, col: j, height: heightMap[i][j]})
				visited[i][j] = true
			}
		}
	}

	var water int
	directions := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	for pq.Len() > 0 {
		cell := heap.Pop(pq).(*Cell)

		for _, dir := range directions {
			r, c := cell.row+dir[0], cell.col+dir[1]

			if r >= 0 && r < m && c >= 0 && c < n && !visited[r][c] {
				visited[r][c] = true
				if cell.height > heightMap[r][c] {
					water += cell.height - heightMap[r][c]
				}
				newHeight := heightMap[r][c]
				if cell.height > newHeight {
					newHeight = cell.height
				}
				heap.Push(pq, &Cell{row: r, col: c, height: newHeight})
			}
		}
	}

	return water
}

func maxFrequency(nums []int, k int, numOperations int) int {
	// 3346
	// max(s_i) - min(s_i) <= 2k.
	// The number of operations is the number of `s_i` not equal to maxFrequency.
	sort.Ints(nums)
	n := len(nums)
	maxFreq := 0
	// Case 1: maxFrequency is not one of the original numbers in the chosen group.
	// This requires `m` operations. So, we need `m <= numOperations`.
	// We find the longest subarray `nums[left...right]` where `nums[right] - nums[left] <= 2k`
	// and its length is at most `numOperations`.
	left := 0
	for right := 0; right < n; right++ {
		for nums[right]-nums[left] > 2*k {
			left++
		}
		m := right - left + 1
		if m <= numOperations {
			if m > maxFreq {
				maxFreq = m
			}
		}
	}
	// Case 2: maxFrequency is one of the numbers from the original array.
	// We iterate through each unique number in `nums` and treat it as a potential target.
	counts := make(map[int]int)
	for _, num := range nums {
		counts[num]++
	}
	for i := 0; i < n; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		target := nums[i]
		// Find the window of all elements that can be converted to `target`.
		// These are numbers `num` such that `target-k <= num <= target+k`.
		startIdx := sort.SearchInts(nums, target-k)
		endIdx := sort.Search(n, func(j int) bool { return nums[j] > target+k }) - 1
		if startIdx > endIdx {
			continue
		}
		windowSize := endIdx - startIdx + 1
		countOfTarget := counts[target]
		others := windowSize - countOfTarget
		achievableFreq := countOfTarget
		if others < numOperations {
			achievableFreq += others
		} else {
			achievableFreq += numOperations
		}
		if achievableFreq > maxFreq {
			maxFreq = achievableFreq
		}
	}
	return maxFreq
}

func maxFrequency2(nums []int, k int, numOperations int) int {
	// 3347
	sort.Ints(nums)
	n := len(nums)
	counts := make(map[int]int)
	maxFreq := 0
	if n > 0 {
		// maxFreq is at least 1 if there are elements.
		maxFreq = 1
		for _, num := range nums {
			counts[num]++
			if counts[num] > maxFreq {
				maxFreq = counts[num]
			}
		}
	}
	if numOperations == 0 {
		return maxFreq
	}
	// Case 1: The target value `T` is chosen optimally, but not necessarily from `nums`.
	// If `T` is not in `nums`, its original count is 0.
	// The achievable frequency is `min(|S_T|, numOperations)`, where `S_T` are
	// elements convertible to `T`. We need to find `T` that maximizes `|S_T|`.
	// `|S_T|` is the number of elements in `[T-k, T+k]`, a window of size `2k`.
	// The max number of elements in any `2k` window is found with a sliding window.
	max_m := 0
	left := 0
	for right := 0; right < n; right++ {
		for nums[right]-nums[left] > 2*k {
			left++
		}
		m := right - left + 1
		if m > max_m {
			max_m = m
		}
	}
	freqCase1 := max_m
	if freqCase1 > numOperations {
		freqCase1 = numOperations
	}
	if freqCase1 > maxFreq {
		maxFreq = freqCase1
	}
	// Case 2: The target value `T` is one of the numbers from the original array.
	for i := 0; i < n; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue // Iterate unique numbers only
		}
		target := nums[i]
		// Find window of elements convertible to `target`: `[target-k, target+k]`
		startIdx := sort.SearchInts(nums, target-k)
		endIdx := sort.Search(n, func(j int) bool { return nums[j] > target+k }) - 1

		if startIdx > endIdx {
			continue
		}
		windowSize := endIdx - startIdx + 1
		countOfTarget := counts[target]
		others := windowSize - countOfTarget
		achievableFreq := countOfTarget
		ops_to_use := others
		if ops_to_use > numOperations {
			ops_to_use = numOperations
		}
		achievableFreq += ops_to_use
		if achievableFreq > maxFreq {
			maxFreq = achievableFreq
		}
	}
	return maxFreq
}

func MaximumTotalDamage(power []int) int64 {
	// 3186
	counts := make(map[int]int)
	for _, p := range power {
		counts[p]++
	}
	uniquePowers := make([]int, 0, len(counts))
	for p := range counts {
		uniquePowers = append(uniquePowers, p)
	}
	sort.Ints(uniquePowers)
	n := len(uniquePowers)
	dp := make([]int64, n)
	dp[0] = int64(uniquePowers[0]) * int64(counts[uniquePowers[0]])
	for i := 1; i < n; i++ {
		p_i := uniquePowers[i]
		// Option 2: Do not cast spells with damage p_i.
		damage_if_not_cast := dp[i-1]
		// Option 1: Cast spells with damage p_i.
		damage_i := int64(p_i) * int64(counts[p_i])
		var prev_compatible_damage int64 = 0
		// Use binary search to find largest `j < i` where `uniquePowers[j] < p_i - 2`.
		target := p_i - 2
		j := sort.Search(i, func(x int) bool {
			return uniquePowers[x] >= target
		})
		if j > 0 {
			prev_compatible_damage = dp[j-1]
		}
		damage_if_cast := damage_i + prev_compatible_damage
		if damage_if_cast > damage_if_not_cast {
			dp[i] = damage_if_cast
		} else {
			dp[i] = damage_if_not_cast
		}
	}
	return dp[n-1]
}

type Cell struct {
	row, col, height int
}

type MinHeap []*Cell

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].height < h[j].height }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x any) {
	*h = append(*h, x.(*Cell))
}

func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func NumWaterBottles(numBottles int, numExchange int) (result int) {
	// 1518
	result = numBottles
	for numBottles >= numExchange {
		result += numBottles / numExchange
		numBottles = numBottles/numExchange + numBottles%numExchange
	}
	return
}

func MaxBottlesDrunk(numBottles int, numExchange int) (result int) {
	// 3100
	result = numBottles
	emptyBottles := numBottles
	for emptyBottles >= numExchange {
		// exchange 1 bottle
		emptyBottles -= numExchange
		result++
		emptyBottles++
		numExchange++
	}
	return
}

func FindSmallestInteger(nums []int, value int) int {
	// 2598
	counts := make([]int, value)
	for _, num := range nums {
		rem := num % value
		if rem < 0 {
			rem += value
		}
		counts[rem]++
	}
	for mex := 0; ; mex++ {
		rem := mex % value
		if counts[rem] > 0 {
			counts[rem]--
		} else {
			return mex
		}
	}
}

func maxPartitionsAfterOperations(s string, k int) int {
	// 3003
	n := len(s)
	type State struct {
		index, current_chars, can_change int
	}
	dfs := func(self func(State) int, key State) int {
		index, current_chars, can_change := key.index, key.current_chars, key.can_change
		if index >= n {
			return 1
		}
		char_bit := 1 << (s[index] - 'a')
		next_chars := current_chars | char_bit
		var result int
		if bits.OnesCount(uint(next_chars)) > k {
			// Current character creates a new partition.
			result = self(State{index: index + 1, current_chars: char_bit, can_change: can_change}) + 1
		} else {
			// Continue the current partition.
			result = self(State{index: index + 1, current_chars: next_chars, can_change: can_change})
		}
		if can_change == 1 {
			for letter_index := 0; letter_index < 26; letter_index++ {
				changed_char_bit := 1 << letter_index
				next_chars_changed := current_chars | changed_char_bit
				var current_res int
				if bits.OnesCount(uint(next_chars_changed)) > k {
					// The change forces a new partition.
					// The new partition starts with the changed character.
					// `can_change` becomes 0 for the subsequent recursive calls.
					current_res = self(State{index: index + 1, current_chars: changed_char_bit, can_change: 0}) + 1
				} else {
					// The change does not force a new partition.
					// Continue the current partition with the changed character included.
					current_res = self(State{index: index + 1, current_chars: next_chars_changed, can_change: 0})
				}
				if current_res > result {
					result = current_res
				}
			}
		}
		return result
	}
	memoized_dfs := cachedFunction(dfs)
	return memoized_dfs(State{index: 0, current_chars: 0, can_change: 1})
}

func cachedFunction[K comparable, V any](f func(self func(K) V, key K) V) func(K) V {
	cache := make(map[K]V)
	var recursive_call func(K) V
	recursive_call = func(key K) V {
		if v, ok := cache[key]; ok {
			return v
		}
		res := f(recursive_call, key)
		cache[key] = res
		return res
	}
	return recursive_call
}

func FindLexSmallestString(s string, a int, b int) string {
	// 1625
	q := []string{s}
	visited := make(map[string]struct{})
	visited[s] = struct{}{}
	minS := s
	for len(q) > 0 {
		curr := q[0]
		q = q[1:]
		// Add operation
		buf := []byte(curr)
		for i := 1; i < len(buf); i += 2 {
			buf[i] = byte(((int(buf[i]-'0') + a) % 10) + '0')
		}
		added := string(buf)
		if _, seen := visited[added]; !seen {
			visited[added] = struct{}{}
			q = append(q, added)
			if added < minS {
				minS = added
			}
		}
		// Rotate operation
		n := len(curr)
		rotated := curr[n-b:] + curr[:n-b]
		if _, seen := visited[rotated]; !seen {
			visited[rotated] = struct{}{}
			q = append(q, rotated)
			if rotated < minS {
				minS = rotated
			}
		}
	}
	return minS
}

func SmallestNumber(n int) int {
	// 3370
	// Return the smallest number x greater than or equal to n, such that the
	// binary representation of x contains only set bits.
	// A number with all set bits is of the form (1<<k - 1).
	// We need to find the smallest k such that (1<<k - 1) >= n.
	// Let l = bits.Len(uint(n)). For n > 0, this means 2^(l-1) <= n < 2^l.
	// The number (1<<l - 1) is 2^l - 1. Since n < 2^l, we have n <= 2^l - 1.
	// So (1<<l - 1) is >= n.
	// The next smaller number with all bits set is (1<<(l-1) - 1) = 2^(l-1) - 1.
	// Since n >= 2^(l-1), this smaller number is always < n.
	// Therefore, (1<<l - 1) is the smallest such number >= n.
	// For n=0, bits.Len(0) is 0, and (1<<0 - 1) is 0.
	return (1 << bits.Len(uint(n))) - 1
}