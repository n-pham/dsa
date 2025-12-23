package main

import (
	"fmt"
)

func ExampleMinDeletionSize() {
	fmt.Println(MinDeletionSize([]string{"cba","daf","ghi"}))
	fmt.Println(MinDeletionSize([]string{"a","b"}))
	fmt.Println(MinDeletionSize([]string{"zyx","wvu","tsr"}))
	// Unordered output:
	// 1
	// 0
	// 3
}

func ExampleGetDescentPeriods() {
	fmt.Println(GetDescentPeriods([]int{3, 2, 1, 4}))
	fmt.Println(GetDescentPeriods([]int{8, 6, 7, 7}))
	fmt.Println(GetDescentPeriods([]int{1}))
	// Unordered output:
	// 7
	// 4
	// 1
}

func ExampleCountSpecialTriplets() {
	fmt.Println(CountSpecialTriplets([]int{6, 3, 6}))
	fmt.Println(CountSpecialTriplets([]int{0, 1, 0, 0}))
	fmt.Println(CountSpecialTriplets([]int{8, 4, 2, 8, 4}))
	// Unordered output:
	// 1
	// 1
	// 2
}

func ExampleCountSpecialTriplets_largeZeros() {
	nums := make([]int, 10000) // 10000 zeros
	fmt.Println(CountSpecialTriplets(nums))
	// Output:
	// 616668838
}

func ExampleCountTriples() {
	fmt.Println(CountTriples(5))
	fmt.Println(CountTriples(10))
	fmt.Println(CountTriples(18))
	// Unordered output:
	// 2
	// 4
	// 10
}

func ExampleCountOdds() {
	fmt.Println(CountOdds(3, 7))
	fmt.Println(CountOdds(8, 10))
	// Unordered output:
	// 3
	// 1
}

func ExampleCountPartitions() {
	fmt.Println(CountPartitions([]int{10, 10, 3, 7, 6}))
	fmt.Println(CountPartitions([]int{1, 2, 2}))
	fmt.Println(CountPartitions([]int{2, 4, 6, 8}))
	// Unordered output:
	// 4
	// 0
	// 3
}

func ExampleValidateCoupons() {
	codes := []string{"code1", "code2", "code3", "code4", "code5", "invalid-code", "code6", "code7", ""}
	businessLines := []string{"grocery", "electronics", "pharmacy", "restaurant", "grocery", "grocery", "electronics", "food", "pharmacy"}
	isActives := []bool{true, true, true, true, false, true, true, true, true}
	fmt.Println(ValidateCoupons(codes, businessLines, isActives))
	// Output: [code2 code6 code1 code3 code4]
}
