package main

//lint:file-ignore U1000 Ignore all unused code, it's generated
import (
	"fmt"
	// "github.com/mxschmitt/golang-combinations"
	// "strconv"
	// "strings"
)

func mostPoints_time(questions [][]int) int64 {
	if len(questions) == 0 {
		return 0
	}
	skip0 := mostPoints(questions[1:])
	solve0 := int64(questions[0][0])
	if questions[0][1]+1 < len(questions) {
		solve0 += mostPoints(questions[questions[0][1]+1:])
	}
	if solve0 > skip0 {
		return solve0
	}
	return skip0
}

func mostPoints(questions [][]int) int64 {
	// 2140
	panic("Dynamic Programming")
}

func main() {
	fmt.Println(mostPoints([][]int{{3, 2}, {4, 3}, {4, 4}, {2, 5}}))
}
