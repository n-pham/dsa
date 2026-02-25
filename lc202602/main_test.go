package main

import (
	"fmt"
	"math"
	"testing"
)

func TestMinimumCost3013(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		dist int
		want int64
	}{
		{
			name: "Example 01",
			nums: []int{1, 3, 2, 6, 4, 2},
			k:    3,
			dist: 3,
			want: 5,
		},
		{
			name: "Example 02",
			nums: []int{10, 1, 2, 2, 2, 1},
			k:    4,
			dist: 3,
			want: 15,
		},
		{
			name: "Example 03",
			nums: []int{10, 8, 18, 9},
			k:    3,
			dist: 1,
			want: 36,
		},
		{
			name: "Example 1",
			nums: []int{10, 1, 2, 3, 4, 5, 6},
			k:    3,
			dist: 2,
			want: 13,
		},
		{
			name: "Example 2",
			nums: []int{10, 2, 8, 3, 4, 9, 1},
			k:    4,
			dist: 2,
			want: 23,
		},
		{
			name: "Single element, k=1",
			nums: []int{5},
			k:    1,
			dist: 0,
			want: 5,
		},
		{
			name: "k=2, small dist",
			nums: []int{1, 10, 2, 9, 3, 8},
			k:    2,
			dist: 1,
			want: 3,
		},
		{
			name: "k-1 > dist+1: impossible",
			nums: []int{1, 2, 3, 4, 5},
			k:    5,
			dist: 1,
			want: math.MaxInt64,
		},
		{
			name: "All same values",
			nums: []int{7, 7, 7, 7, 7, 7},
			k:    3,
			dist: 2,
			want: 21,
		},
		{
			name: "Longer array, small k",
			nums: []int{10, 1, 100, 2, 50, 3, 200, 4},
			k:    3,
			dist: 3,
			want: 13,
		},
		{
			name: "k-1 equals window size",
			nums: []int{1, 2, 3, 4, 5, 6},
			k:    4,
			dist: 2,
			want: 10,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minimumCost3013(tt.nums, tt.k, tt.dist)
			if got != tt.want {
				t.Errorf("minimumCost3013() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMinimumCost(t *testing.T) {
	tests := []struct {
		name     string
		source   string
		target   string
		original []string
		changed  []string
		cost     []int
		want     int64
	}{
		{
			name:     "Example 1",
			source:   "abcd",
			target:   "acbd",
			original: []string{"bc", "cb"},
			changed:  []string{"cb", "bc"},
			cost:     []int{1, 2},
			want:     1, // a->a (0), bc->cb (1), d->d (0)
		},
		{
			name:     "Example 2",
			source:   "aaaa",
			target:   "bbbb",
			original: []string{"a", "aa"},
			changed:  []string{"b", "bb"},
			cost:     []int{5, 1},
			want:     2, // aa->bb (1) + aa->bb (1) = 2
		},
		{
			name:     "Impossible",
			source:   "abcd",
			target:   "abce",
			original: []string{"a"},
			changed:  []string{"e"},
			cost:     []int{10},
			want:     -1,
		},
		{
			name:     "Intermediate transformation",
			source:   "a",
			target:   "c",
			original: []string{"a", "b"},
			changed:  []string{"b", "c"},
			cost:     []int{10, 20},
			want:     30, // a->b (10) + b->c (20) = 30
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MinimumCost(tt.source, tt.target, tt.original, tt.changed, tt.cost)
			if got != tt.want {
				t.Errorf("MinimumCost() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExampleMinimumCost() {
	fmt.Println(MinimumCost("abcd", "acbd", []string{"bc", "cb"}, []string{"cb", "bc"}, []int{1, 2}))
	fmt.Println(MinimumCost("aaaa", "bbbb", []string{"a", "aa"}, []string{"b", "bb"}, []int{5, 1}))
	fmt.Println(MinimumCost("abcd", "abce", []string{"a"}, []string{"e"}, []int{10}))
	// Output:
	// 1
	// 2
	// -1
}

func ExampleIsTrionic() {
	fmt.Println(IsTrionic([]int{1, 3, 5, 4, 2, 6}))
	fmt.Println(IsTrionic([]int{2, 1, 3}))
	// Output:
	// true
	// false
}

func ExampleMaxSumTrionic() {
	fmt.Println(MaxSumTrionic([]int{-754, 167, -172, 202, 735, -941, 992}))
	fmt.Println(MaxSumTrionic([]int{2, 993, -791, -635, -569}))
	fmt.Println(MaxSumTrionic([]int{0, -2, -1, -3, 0, 2, -1}))
	fmt.Println(MaxSumTrionic([]int{1, 2, 8, 5, 3, 2, 9, 10}))
	fmt.Println(MaxSumTrionic([]int{10, 20, 15, 12, 25, 30}))
	fmt.Println(MaxSumTrionic([]int{1, 2, 3, 2, 1, 2, 3}))
	fmt.Println(MaxSumTrionic([]int{1, 2, 1, 2, 1, 2, 1}))
	fmt.Println(MaxSumTrionic([]int{5, 4, 3, 2, 1}))
	// Output:
	// 988
	// -431
	// -4
	// 40
	// 112
	// 14
	// 6
	// 0
}

func ExampleConstructTransformedArray() {
	fmt.Println(ConstructTransformedArray([]int{-10}))
	fmt.Println(ConstructTransformedArray([]int{3, -2, 1, 1}))
	fmt.Println(ConstructTransformedArray([]int{-1, 4, -1}))
	// Output:
	// [-10]
	// [1 1 1 3]
	// [-1 -1 4]
}

func ExampleMinRemoval() {
	fmt.Println(MinRemoval([]int{2, 1, 5}, 2))
	fmt.Println(MinRemoval([]int{1, 6, 2, 9}, 3))
	fmt.Println(MinRemoval([]int{4, 6}, 2))
	// Output:
	// 1
	// 2
	// 0
}

func ExampleBinaryGap() {
	fmt.Println(BinaryGap(22))
	fmt.Println(BinaryGap(8))
	fmt.Println(BinaryGap(5))
	// Output:
	// 2
	// 0
	// 2
}

func ExampleIsBalanced() {
	root := &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}}
	fmt.Println(IsBalanced(root))
	root2 := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}
	fmt.Println(IsBalanced(root2))
	// Output:
	// false
	// true
}

func ExampleBalanceBST() {
	root := &TreeNode{Val: 1, Right: &TreeNode{Val: 2, Right: &TreeNode{Val: 3, Right: &TreeNode{Val: 4}}}}
	balanced := BalanceBST(root)
	fmt.Println(IsBalanced(balanced))
	// Output:
	// true
}

func ExampleLongestBalanced() {
	fmt.Println(LongestBalanced([]int{1, 2, 1, 2}))
	fmt.Println(LongestBalanced([]int{1, 2, 3, 4}))
	// Output:
	// 4
	// 4
}

func ExampleLongestBalancedSubstring() {
	fmt.Println(LongestBalancedSubstring("aaabbb"))
	fmt.Println(LongestBalancedSubstring("ab"))
	// Output:
	// 6
	// 2
}

func ExampleLongestBalancedSubstring2() {
	fmt.Println(LongestBalancedSubstring2("aaabbbccc"))
	fmt.Println(LongestBalancedSubstring2("abc"))
	// Output:
	// 9
	// 3
}

func ExampleReverseBits() {
	fmt.Printf("%032b\n", ReverseBits(0b00000010100101000001111010011100))
	// Output:
	// 00111001011110000010100101000000
}

func ExampleHasAlternatingBits() {
	fmt.Println(HasAlternatingBits(5))
	fmt.Println(HasAlternatingBits(7))
	// Output:
	// true
	// false
}

func ExampleCountBinarySubstrings() {
	fmt.Println(CountBinarySubstrings("00110011"))
	fmt.Println(CountBinarySubstrings("10101"))
	// Output:
	// 6
	// 4
}

func ExampleCountPrimeSetBits() {
	fmt.Println(CountPrimeSetBits(6, 10))
	fmt.Println(CountPrimeSetBits(10, 15))
	// Output:
	// 4
	// 5
}

func ExampleHasAllCodes() {
	fmt.Println(HasAllCodes("00110110", 2))
	fmt.Println(HasAllCodes("0110", 1))
	fmt.Println(HasAllCodes("0110", 2))
	// Output:
	// true
	// true
	// false
}

func ExampleSumRootToLeaf() {
	// Example 1: [1,0,1,0,1,0,1]
	//       1
	//     /   \
	//    0     1
	//   / \   / \
	//  0   1 0   1
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   0,
			Left:  &TreeNode{Val: 0},
			Right: &TreeNode{Val: 1},
		},
		Right: &TreeNode{
			Val:   1,
			Left:  &TreeNode{Val: 0},
			Right: &TreeNode{Val: 1},
		},
	}
	fmt.Println(SumRootToLeaf(root))

	// Example 2: [0]
	root2 := &TreeNode{Val: 0}
	fmt.Println(SumRootToLeaf(root2))

	// Output:
	// 22
	// 0
}

func ExampleSortByBits() {
	fmt.Println(SortByBits([]int{0,1,2,3,4,5,6,7,8}))
	fmt.Println(SortByBits([]int{1024,512,256,128,64,32,16,8,4,2,1}))
	// Output:
	// [0 1 2 4 8 3 5 6 7]
	// [1 2 4 8 16 32 64 128 256 512 1024]
}