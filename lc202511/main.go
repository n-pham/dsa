package main

//lint:file-ignore U1000 Ignore all unused code, it's generated

import (
	"log"
	"os"
	"runtime"
	"sort"
	"strings"
)

func CountPalindromicSubsequence(s string) (result int) {
	// 1930
	// xyx --> find x with firstIndex and lastIndex, count unique y characters
	firstIndices, lastIndices := [26]int{}, [26]int{}  // 1-based
	for i, c := range s {
		if firstIndices[c-'a'] == 0 {
			firstIndices[c-'a'] = i+1
		} else {
			lastIndices[c-'a'] = i+1
		}
	}
	for i := 0; i < 26; i++ {
		firstIdx := firstIndices[i]
		lastIdx := lastIndices[i]
		if firstIdx > 0 && lastIdx > 0 {
			sub := s[firstIdx : lastIdx-1]
			uniqueChars := make(map[rune]struct{})
			for _, char := range sub {
				uniqueChars[char] = struct{}{}
			}
			result += len(uniqueChars)
		}
	}
	return
}

func findFinalValue(nums []int, original int) int {
	// 2154
	isInMap := make(map[int]bool, len(nums))
	for _, num := range nums {
		isInMap[num] = true
	}
	for ; isInMap[original]; original *= 2 {}
	return original
}

func isOneBitCharacter(bits []int) bool {
	// 717
	i := 0
	for i < len(bits)-1 {
		i += bits[i] + 1  // advance 1 if 0, 2 if 1
	}
	// If the pointer lands on the last index, the last character is 1-bit.
	return i == len(bits)-1
}


func KLengthApart(nums []int, k int) bool {
    // 1437
    lastIndex := -100_000
    for index, num := range nums {
        if num == 1 {
            if index <= lastIndex + k {
                return false
            }
            lastIndex = index
        }
    }
    return true
}


func countOperations(num1 int, num2 int) (cnt int) {
    // 2169
    for ; num1 > 0 && num2 > 0; cnt++ {
        if num1 >= num2 {
            num1 -= num2
        } else {
            num2 -= num1
        }
    }
    return
}


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

var guess func(num int) int

func GuessNumber(n int) int {
    // 374
    low, high := 1, n
    for low <= high {
        mid := low + (high-low)/2
        switch guess(mid) {
        case 0:
            return mid
        case -1:
            high = mid - 1
        default:
            low = mid + 1
        }
    }
    return -1
}

func FindXSum(nums []int, k int, x int) []int {
    // 3318
    // for sliding window of size k --> x elements with most frequencies --> sum
    n := len(nums)
    answer := make([]int, 0, n-k+1)
    freq := make(map[int]int)
    // counts[i] -> list of numbers with frequency i
    counts := make([][]int, k+1)
    calculateXSum := func() int {
        xSum := 0
        remainingX := x
        for i := k; i >= 1 && remainingX > 0; i-- {
            if len(counts[i]) > 0 {
                numbers := make([]int, len(counts[i]))
                copy(numbers, counts[i])
                sort.Sort(sort.Reverse(sort.IntSlice(numbers)))

                for _, num := range numbers {
                    if remainingX > 0 {
                        xSum += num * i
                        remainingX--
                    } else {
                        break
                    }
                }
            }
        }
        return xSum
    }

    updateBucket := func(num, oldCount, newCount int) {
        if oldCount > 0 {
            list := counts[oldCount]
            for i, v := range list {
                if v == num {
                    counts[oldCount] = append(list[:i], list[i+1:]...)
                    break
                }
            }
        }
        if newCount > 0 {
            counts[newCount] = append(counts[newCount], num)
        }
    }

    // Initial window
    for i := 0; i < k; i++ {
        num := nums[i]
        oldCount := freq[num]
        freq[num]++
        newCount := freq[num]
        updateBucket(num, oldCount, newCount)
    }
    answer = append(answer, calculateXSum())

    // Sliding window
    for i := 1; i <= n-k; i++ {
        // Outgoing
        outNum := nums[i-1]
        outOldCount := freq[outNum]
        freq[outNum]--
        outNewCount := freq[outNum]
        updateBucket(outNum, outOldCount, outNewCount)

        // Incoming
        inNum := nums[i+k-1]
        inOldCount := freq[inNum]
        freq[inNum]++
        inNewCount := freq[inNum]
        updateBucket(inNum, inOldCount, inNewCount)

        answer = append(answer, calculateXSum())
    }

    return answer
}