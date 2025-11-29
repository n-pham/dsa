package main

import (
	"fmt"
	"strings"
)


func ExampleMinOperations() {
	fmt.Println(MinOperations([]int{3,9,7}, 5))
	fmt.Println(MinOperations([]int{4,1,3}, 4))
	fmt.Println(MinOperations([]int{3,2}, 6))
	// Unordered output:
	// 4
	// 0
	// 5
}

func ExampleCountPalindromicSubsequence() {
	fmt.Println(CountPalindromicSubsequence("aabca"))
	fmt.Println(CountPalindromicSubsequence("adc"))
	fmt.Println(CountPalindromicSubsequence("bbcbaba"))
	// Unordered output:
	// 3
	// 0
	// 4
}

func ExampleKLengthApart() {
	fmt.Println(KLengthApart([]int{1, 0, 0, 0, 1, 0, 0, 1}, 2))
	fmt.Println(KLengthApart([]int{1, 0, 0, 1, 0, 1}, 2))
	// Unordered output:
	// true
	// false
}

// Helper to create a linked list from a slice
func makeList(vals []int) *ListNode {
	if len(vals) == 0 {
		return nil
	}
	head := &ListNode{Val: vals[0]}
	curr := head
	for i := 1; i < len(vals); i++ {
		curr.Next = &ListNode{Val: vals[i]}
		curr = curr.Next
	}
	return head
}

// Helper to print a linked list
func printList(head *ListNode) {
	if head == nil {
		fmt.Println("nil")
		return
	}
	var parts []string
	curr := head
	for curr != nil {
		parts = append(parts, fmt.Sprintf("%d", curr.Val))
		curr = curr.Next
	}
	fmt.Println(strings.Join(parts, " "))
}

func ExampleModifiedList() {
	// Test case 1: Basic removal
	head1 := makeList([]int{1, 2, 3, 4})
	nums1 := []int{1, 3}
	printList(ModifiedList(nums1, head1))

	// Test case 2: Removing the head
	head2 := makeList([]int{1, 2})
	nums2 := []int{1}
	printList(ModifiedList(nums2, head2))

	// Test case 3: Removing consecutive nodes
	head3 := makeList([]int{1, 2, 3})
	nums3 := []int{1, 2}
	printList(ModifiedList(nums3, head3))

	// Test case 4: No nodes to remove
	head4 := makeList([]int{1, 2, 3})
	nums4 := []int{5, 6}
	printList(ModifiedList(nums4, head4))

	// Test case 5: Empty list
	var head5 *ListNode
	nums5 := []int{1}
	printList(ModifiedList(nums5, head5))

	// Test case 6: All nodes removed
	head6 := makeList([]int{1, 2, 3})
	nums6 := []int{1, 2, 3}
	printList(ModifiedList(nums6, head6))

	// Unordered output:
	// 2 4
	// 2
	// 3
	// 1 2 3
	// nil
	// nil
}

func ExampleCountUnguarded() {
	// Test case 1
	m1 := 4
	n1 := 6
	guards1 := [][]int{{0, 0}, {1, 1}, {2, 3}}
	walls1 := [][]int{{0, 1}, {2, 2}, {1, 4}}
	fmt.Println(CountUnguarded(m1, n1, guards1, walls1))

	// Test case 2
	m2 := 3
	n2 := 3
	guards2 := [][]int{{1, 1}}
	walls2 := [][]int{{0, 1}, {1, 0}, {1, 2}, {2, 1}}
	fmt.Println(CountUnguarded(m2, n2, guards2, walls2))

	// Unordered output:
	// 7
	// 4
}

func ExampleFindXSum() {
	fmt.Println(FindXSum([]int{1, 1, 2, 2, 3, 4, 2, 3}, 6, 2))
	// Unordered output:
	// [6 10 12]
}

func ExampleGuessNumber() {
	var pickedNumber int
	guess = func(num int) int {
		if num > pickedNumber {
			return -1
		} else if num < pickedNumber {
			return 1
		}
		return 0
	}

	// Test case 1: Basic guess
	pickedNumber = 6
	fmt.Println(GuessNumber(10))

	// Test case 2: Another guess
	pickedNumber = 3
	fmt.Println(GuessNumber(5))

	// Unordered output:
	// 6
	// 3
}

func ExampleMaxSumDivThree() {
	fmt.Println(maxSumDivThree([]int{3, 6, 5, 1, 8}))
	fmt.Println(maxSumDivThree([]int{1, 2, 3, 4, 5}))
	fmt.Println(maxSumDivThree([]int{1}))
	fmt.Println(maxSumDivThree([]int{}))
	fmt.Println(maxSumDivThree([]int{4, 5, 6}))
	// Unordered output:
	// 18
	// 15
	// 0
	// 0
	// 15
}

func ExamplePrefixesDivBy5() {

	fmt.Println(PrefixesDivBy5([]int{0, 1, 1}))

	fmt.Println(PrefixesDivBy5([]int{1, 1, 1}))

	fmt.Println(PrefixesDivBy5([]int{0, 1, 1, 1, 1, 1, 0}))

	// Unordered output:

	// [true false false]

	// [false false false]

	// [true false false false true false false]

}

func ExampleSmallestRepunitDivByK() {

	fmt.Println(smallestRepunitDivByK(1))

	fmt.Println(smallestRepunitDivByK(2))

	fmt.Println(smallestRepunitDivByK(3))

	fmt.Println(smallestRepunitDivByK(7))

	fmt.Println(smallestRepunitDivByK(17))

	// Unordered output:

	// 1

	// -1

	// 3

	// 6

	// 16

}
