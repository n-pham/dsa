package main

//lint:file-ignore U1000 Ignore all unused code, it's generated

import (
	"container/heap"
	"log"
	"os"
	"runtime"
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
