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