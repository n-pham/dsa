package main

//lint:file-ignore U1000 Ignore all unused code, it's generated

import "fmt"

func maxTaskAssign(tasks []int, workers []int, pills int, strength int) int {
	// 2071
	// 1 2 3
	// 0 3 3
	fmt.Println(tasks, workers, pills, strength)
	return 0
}

func main() {
	fmt.Println(maxTaskAssign([]int{3, 2, 1}, []int{0, 3, 3}, 1, 1))
}
