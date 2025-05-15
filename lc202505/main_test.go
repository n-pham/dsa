package main

import (
	"slices"
	"testing"
)

func TestLengthAfterTransformations(t *testing.T) {
	if !slices.Equal([]string{"a", "c", "d"}, getLongestSubsequence([]string{"a", "b", "c", "d"}, []int{1, 1, 0, 1})) {
		t.Errorf(`getLongestSubsequence`)
	}
	if 7 != lengthAfterTransformations3337("abcyy", 2, []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2}) {
		t.Errorf(`lengthAfterTransformations3337`)
	}
	if 7 != lengthAfterTransformations("abcyy", 2) {
		t.Errorf(`lengthAfterTransformations`)
	}
	if 2 != lengthAfterTransformations("v", 7) {
		t.Errorf(`lengthAfterTransformations`)
	}
	// actual := lengthAfterTransformations("abcyy", 2)
	// expected := 7
	// if actual != expected {
	// 	t.Errorf(`lengthAfterTransformations("abcyy", 2) = %v, expected %v`, actual, expected)
	// }
}
