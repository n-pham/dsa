package main

//lint:file-ignore U1000 Ignore all unused code, it's generated

import (
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

func Generate(numRows int) [][]int {
	// 118
	// numRows <= 30
	if numRows == 1 {
		return [][]int{{1}}
	}
	rows := Generate(numRows - 1)
	rowsLen := len(rows)
	lenght := len(rows[rowsLen-1]) + 1
	newRow := make([]int, lenght)
	newRow[0] = 1
	prevNum := 1
	for i, num := range rows[rowsLen-1][1:] {
		newRow[i+1] = prevNum + num
		prevNum = num
	}
	newRow[lenght-1] = 1
	return append(rows, newRow)
}

func NumOfUnplacedFruits(fruits []int, baskets []int) (unplacedCnt int) {
	// 3477
	for _, fruitCnt := range fruits {
		unplaced := 1
		for i, basketCapacity := range baskets {
			if basketCapacity >= fruitCnt {
				baskets[i] = 0
				unplaced = 0
				break
			}
		}
		unplacedCnt += unplaced
	}
	return
}

func CountBits(n int) []int {
	// 338
	//  i & (i - 1) results in i with the least significant 1 bit turned off.
	counts := make([]int, n+1)
	for i := 1; i <= n; i++ {
		counts[i] = counts[i&(i-1)] + 1
	}
	return counts
}
