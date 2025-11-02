package main

//lint:file-ignore U1000 Ignore all unused code, it's generated

import (
	"log"
	"os"
	"runtime"
	"strings"
)

type GridState int

const (
	Empty GridState = iota
	Guard
	Wall
	Guarded
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

func CountUnguarded(m int, n int, guards [][]int, walls [][]int) (cnt int) {
	// 2257
	grid := make([][]GridState, m)
	for i := range grid {
		grid[i] = make([]GridState, n)
	}
	for _, wall := range walls {
		grid[wall[0]][wall[1]] = Wall
	}
	for _, guard := range guards {
		grid[guard[0]][guard[1]] = Guard
	}
	// Mark guarded cells
	for r := 0; r < m; r++ {
		// L-R
		for c := 0; c < n; {
			if grid[r][c] == Guard {
				c++ // move to next cell
				for c < n && grid[r][c] != Guard && grid[r][c] != Wall {
					grid[r][c] = Guarded
					c++
				}
			} else {
				c++
			}
		}
		// R-L
		for c := n - 1; c >= 0; {
			if grid[r][c] == Guard {
				c-- // move to next cell
				for c >= 0 && grid[r][c] != Guard && grid[r][c] != Wall {
					grid[r][c] = Guarded
					c--
				}
			} else {
				c--
			}
		}
	}

	for c := 0; c < n; c++ {
		// T-B
		for r := 0; r < m; {
			if grid[r][c] == Guard {
				r++ // move to next cell
				for r < m && grid[r][c] != Guard && grid[r][c] != Wall {
					grid[r][c] = Guarded
					r++
				}
			} else {
				r++
			}
		}
		// B-T
		for r := m - 1; r >= 0; {
			if grid[r][c] == Guard {
				r-- // move to next cell
				for r >= 0 && grid[r][c] != Guard && grid[r][c] != Wall {
					grid[r][c] = Guarded
					r--
				}
			} else {
				r--
			}
		}
	}

	for r := 0; r < m; r++ {
		for c := 0; c < n; c++ {
			if grid[r][c] == Empty {
				cnt++
			}
		}
	}
	return
}

type ListNode struct {
    Val int
    Next *ListNode
}

func ModifiedList(nums []int, head *ListNode) *ListNode {
	// 3217
	exists := make(map[int]bool)
	for _, num := range nums {
		exists[num] = true
	}
	parent := ListNode{Next: head}
	for curr := &parent; curr.Next != nil; {
		if exists[curr.Next.Val] {
			curr.Next = curr.Next.Next
		} else {
			curr = curr.Next
		}
	}
	return parent.Next
}