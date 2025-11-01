package main

import (
	"fmt"
	"strings"
)

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
