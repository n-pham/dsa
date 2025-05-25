package main

import (
	"testing"
)

func TestAll(t *testing.T) {
	if 22 != longestPalindrome([]string{"dd", "aa", "bb", "dd", "aa", "dd", "bb", "dd", "aa", "cc", "bb", "cc", "dd", "cc"}) {
		t.Error("longestPalindrome")
	}
	if 2 != longestPalindrome([]string{"cc", "ll", "xx"}) {
		t.Error("longestPalindrome")
	}
	if 14 != longestPalindrome([]string{"ab", "ty", "yt", "lc", "cl", "ab", "aa", "aa", "bb"}) {
		t.Error("longestPalindrome")
	}
	// if !slices.Equal([]int{0, 2}, findWordsContaining([]string{"abc", "bcd", "aaaa", "cbc"}, 'a')) {
	// 	t.Error("findWordsContaining")
	// }
	// if 1 != maxRemoval([]int{2, 0, 2}, [][]int{{0, 2}, {0, 2}, {1, 1}}) {
	// 	t.Errorf(`maxRemoval`)
	// }
	// sortColorsNums := []int{2, 0, 2, 1, 1, 0}
	// sortColors(sortColorsNums)
	// if !slices.Equal(sortColorsNums, []int{0, 0, 1, 1, 2, 2}) {
	// 	t.Errorf(`sortColors`)
	// }
	// if !slices.Equal([]string{"a", "c", "d"}, getLongestSubsequence([]string{"a", "b", "c", "d"}, []int{1, 1, 0, 1})) {
	// 	t.Errorf(`getLongestSubsequence`)
	// }
	// if 7 != lengthAfterTransformations3337("abcyy", 2, []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2}) {
	// 	t.Errorf(`lengthAfterTransformations3337`)
	// }
	// if 7 != lengthAfterTransformations("abcyy", 2) {
	// 	t.Errorf(`lengthAfterTransformations`)
	// }
	// if 2 != lengthAfterTransformations("v", 7) {
	// 	t.Errorf(`lengthAfterTransformations`)
	// }
	// actual := lengthAfterTransformations("abcyy", 2)
	// expected := 7
	// if actual != expected {
	// 	t.Errorf(`lengthAfterTransformations("abcyy", 2) = %v, expected %v`, actual, expected)
	// }
}
