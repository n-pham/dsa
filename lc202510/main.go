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
		if i - start >= k {
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
