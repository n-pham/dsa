package main

import (
	"log"
	"math"
	"os"
	"runtime"
	"slices"
	"sort"
	"strconv"
	"strings"
)

//lint:file-ignore U1000 Ignore all unused code, it's generated

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

func NumMagicSquaresInside(grid [][]int) (magicSquareCount int) {
	// 840
	m, n := len(grid), len(grid[0])
	isMagic := func(r, c int) bool {
		if grid[r+1][c+1] != 5 {
			return false
		}
		// Rows
		if grid[r][c]+grid[r][c+1]+grid[r][c+2] != 15 {
			return false
		}
		if grid[r+1][c]+grid[r+1][c+1]+grid[r+1][c+2] != 15 {
			return false
		}
		if grid[r+2][c]+grid[r+2][c+1]+grid[r+2][c+2] != 15 {
			return false
		}
		// Cols
		if grid[r][c]+grid[r+1][c]+grid[r+2][c] != 15 {
			return false
		}
		if grid[r][c+1]+grid[r+1][c+1]+grid[r+2][c+1] != 15 {
			return false
		}
		if grid[r][c+2]+grid[r+1][c+2]+grid[r+2][c+2] != 15 {
			return false
		}
		// Diagonals
		if grid[r][c]+grid[r+1][c+1]+grid[r+2][c+2] != 15 {
			return false
		}
		if grid[r][c+2]+grid[r+1][c+1]+grid[r+2][c] != 15 {
			return false
		}
		return true
	}
	// moving window
	var counts [16]int8
	for i := 0; i <= m-3; i++ {
		counts = [16]int8{}
		distinct := 0

		// Fill first 3x3
		for r := i; r < i+3; r++ {
			for c := 0; c < 3; c++ {
				val := grid[r][c]
				if val >= 1 && val <= 9 {
					counts[val]++
					if counts[val] == 1 {
						distinct++
					}
				}
			}
		}

		if distinct == 9 && isMagic(i, 0) {
			magicSquareCount++
		}

		for j := 1; j <= n-3; j++ {
			// Remove col j-1
			for r := i; r < i+3; r++ {
				val := grid[r][j-1]
				if val >= 1 && val <= 9 {
					counts[val]--
					if counts[val] == 0 {
						distinct--
					}
				}
			}
			// Add col j+2
			for r := i; r < i+3; r++ {
				val := grid[r][j+2]
				if val >= 1 && val <= 9 {
					counts[val]++
					if counts[val] == 1 {
						distinct++
					}
				}
			}

			if distinct == 9 && isMagic(i, j) {
				magicSquareCount++
			}
		}
	}
	return
}

func CountNegatives(grid [][]int) (cnt int) {
	// 1351
	//  4  3  2 -1
	//  3  2  1 -1
	//  1  1 -1 -2
	// -1 -1 -2 -3
	m, n := len(grid), len(grid[0])
	for i := 0; i < m; i++ {
		if grid[i][0] < 0 {
			cnt += n * (m-i)
			break
		}
		for j := 0; j < n; j++ {
			if grid[i][j] < 0 {
				cnt += n-j
				break
			}
		}
	}
	return
}

func CanAttendMeetings(intervals [][]int) bool {
	// 252
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	for i := 0; i < len(intervals)-1; i++ {
		if intervals[i][1] > intervals[i+1][0] {
			return false
		}
	}
	return true
}

func BestClosingTime(customers string) int {
	// 2483
	//          YYNY
	// prefixN 00011
	// suffixY 32110
	//         32121
	n := len(customers)
	prefixN := make([]int, n+1)
	for i, c := range customers {
		if c == 'N' {
			prefixN[i+1] = prefixN[i] + 1
		} else {
			prefixN[i+1] = prefixN[i]
		}
	}
	suffixY := make([]int, n+1)
	for i := n - 1; i >= 0; i-- {
		if customers[i] == 'Y' {
			suffixY[i] = suffixY[i+1] + 1
		} else {
			suffixY[i] = suffixY[i+1]
		}
	}
	closingHour, minPenalty := 0, math.MaxUint32
	for i, nCount := range prefixN {
		if nCount+suffixY[i] < minPenalty {
			minPenalty = nCount + suffixY[i]
			closingHour = i
		}
	}
	return closingHour
}

func MinimumBoxes(apple []int, capacity []int) (boxCount int) {
	// 3074
	// apple = [5,5,5], capacity = [2,4,2,7] --> 4
	appleCount := 0
	for _, cnt := range apple {
		appleCount += cnt
	}
	slices.Sort(capacity)
	capacityLen := len(capacity)
	for ; appleCount > 0; boxCount++ {
		appleCount -= capacity[capacityLen-1-boxCount]
	}
	return
}

func MinDeletionSize(strs []string) (cnt int) {
	// 944
	// abc
	// bce
	// cae
	//  ^
	h, w := len(strs), len(strs[0])
	for cIndex := 0; cIndex < w; cIndex++ {
		sortedInt := 0
		for rIndex := 1; rIndex < h; rIndex++ {
			if strs[rIndex][cIndex] < strs[rIndex-1][cIndex] {
				sortedInt = 1
				break
			}
		}
		cnt += sortedInt
	}
	return
}

func GetDescentPeriods(prices []int) int64 {
	// 2110
	streak := int64(1)
	cnt := int64(1)
	for i := 1; i < len(prices); i++ {
		if prices[i-1]-prices[i] == 1 {
			streak++
		} else {
			streak = 1
		}
		cnt += streak
	}
	return cnt
}

func ValidateCoupons(code []string, businessLine []string, isActive []bool) []string {
	// 3606
	type Coupon struct {
		Code         string
		BusinessLine string
	}
	validBusinessLines := map[string]int{
		"electronics": 0,
		"grocery":     1,
		"pharmacy":    2,
		"restaurant":  3,
	}
	var validCoupons []Coupon
	for i := 0; i < len(code); i++ {
		if !isActive[i] {
			continue
		}
		if _, ok := validBusinessLines[businessLine[i]]; !ok {
			continue
		}
		if code[i] == "" {
			continue
		}
		isAlphanumeric := true
		for _, char := range code[i] {
			if !((char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || (char >= '0' && char <= '9') || char == '_') {
				isAlphanumeric = false
				break
			}
		}
		if !isAlphanumeric {
			continue
		}
		validCoupons = append(validCoupons, Coupon{Code: code[i], BusinessLine: businessLine[i]})
	}
	sort.Slice(validCoupons, func(i, j int) bool {
		if validBusinessLines[validCoupons[i].BusinessLine] != validBusinessLines[validCoupons[j].BusinessLine] {
			return validBusinessLines[validCoupons[i].BusinessLine] < validBusinessLines[validCoupons[j].BusinessLine]
		}
		return validCoupons[i].Code < validCoupons[j].Code
	})
	result := make([]string, len(validCoupons))
	for i, c := range validCoupons {
		result[i] = c.Code
	}
	return result
}

func CountCoveredBuildings(n int, buildings [][]int) int {
	// 3531
	return 0
}

type Event struct {
	Type      string
	Timestamp int
	Payload   string
}

func CountMentions(numberOfUsers int, events [][]string) []int {
	// 3433
	parsedEvents := make([]Event, len(events))
	for i, e := range events {
		ts, _ := strconv.Atoi(e[1])
		parsedEvents[i] = Event{Type: e[0], Timestamp: ts, Payload: e[2]}
	}
	sort.Slice(parsedEvents, func(i, j int) bool {
		if parsedEvents[i].Timestamp != parsedEvents[j].Timestamp {
			return parsedEvents[i].Timestamp < parsedEvents[j].Timestamp
		}
		return parsedEvents[i].Type == "OFFLINE" && parsedEvents[j].Type != "OFFLINE"
	})
	mentions := make([]int, numberOfUsers)
	offlineUntil := make(map[int]int)
	for _, event := range parsedEvents {
		if event.Type == "OFFLINE" {
			userID, _ := strconv.Atoi(event.Payload)
			if userID >= 0 && userID < numberOfUsers {
				offlineUntil[userID] = event.Timestamp + 60
			}
		} else if event.Type == "MESSAGE" {
			mentionsStr := event.Payload
			if mentionsStr == "ALL" {
				for i := 0; i < numberOfUsers; i++ {
					mentions[i]++
				}
			} else if mentionsStr == "HERE" {
				for i := 0; i < numberOfUsers; i++ {
					backOnlineTs, isOffline := offlineUntil[i]
					if !isOffline || event.Timestamp >= backOnlineTs {
						mentions[i]++
					}
				}
			} else {
				idStrs := strings.Split(mentionsStr, " ")
				for _, idStr := range idStrs {
					if len(idStr) > 2 {
						userID, err := strconv.Atoi(idStr[2:])
						if err == nil && userID >= 0 && userID < numberOfUsers {
							mentions[userID]++
						}
					}
				}
			}
		}
	}
	return mentions
}

func CountSpecialTriplets(nums []int) (cnt int) {
	// 3583
	// nums[i] == nums[j] * 2
	// nums[k] == nums[j] * 2
	const MOD = 1_000_000_007
	n := len(nums)
	prefixCounts := make(map[int]int)
	suffixCounts := make(map[int]int)
	for _, num := range nums {
		suffixCounts[num]++
	}
	for j := 0; j < n; j++ {
		suffixCounts[nums[j]]--
		if suffixCounts[nums[j]] == 0 {
			delete(suffixCounts, nums[j])
		}
		target := nums[j] * 2
		countLeft := prefixCounts[target]
		countRight := suffixCounts[target]
		cnt = (cnt + int((int64(countLeft)*int64(countRight))%MOD)) % MOD
		prefixCounts[nums[j]]++
	}
	return
}

func CountTriples(n int) (cnt int) {
	// 1925
	for i := 1; i < n-1; i++ {
		if i*i*2 > n*n {
			break
		}
		for j := i + 1; j < n; j++ {
			sqrInt := i*i + j*j
			if sqrInt > n*n {
				break
			}
			if sqrtInt := int(math.Sqrt(float64(sqrInt))); sqrtInt*sqrtInt == sqrInt {
				cnt += 2
			}
		}
	}
	return
}

func CountOdds(low int, high int) int {
	// 1523
	bothCount := high - low + 1
	if bothCount%2 == 0 {
		return bothCount / 2
	} else {
		return bothCount/2 + (low % 2)
	}
}

func CountPartitionsNaive(nums []int) (cnt int) {
	// 3432
	n := len(nums)
	leftSum := 0
	totalSum := 0
	for _, num := range nums {
		totalSum += num
	}
	for i := 0; i < n-1; i++ {
		leftSum += nums[i]
		if (totalSum-leftSum-leftSum)%2 == 0 {
			cnt++
		}
	}
	return
}

func CountPartitions(nums []int) (cnt int) {
	// 3432
	// A partition is valid if (sum of right part - sum of left part) is even.
	// Let totalSum = sum(nums). For a partition at index i, leftSum = sum(nums[:i+1]).
	// The condition is (totalSum - leftSum - leftSum) % 2 == 0.
	// This simplifies to totalSum % 2 == 0, as 2*leftSum is always even.
	// If totalSum is odd, no partition is valid. Count is 0.
	// If totalSum is even, all n-1 partitions are valid.
	// An array can be partitioned only if it has at least 2 elements.
	totalSum := 0
	for _, num := range nums {
		totalSum += num
	}
	if totalSum%2 != 0 {
		return 0
	}
	return len(nums) - 1
}
