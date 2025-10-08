package main

//lint:file-ignore U1000 Ignore all unused code, it's generated

import (
	"container/heap"
	"log"
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

func SuccessfulPairs(spells []int, potions []int, success int64) []int {
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
