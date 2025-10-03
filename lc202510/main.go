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
