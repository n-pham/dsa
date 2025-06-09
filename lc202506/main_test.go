package main

import (
	"slices"
	"testing"
)

func TestAll(t *testing.T) {
	if !slices.Equal([]int{1, 0}, plusOne([]int{9})) {
		t.Error("plusOne")
	}
	// if "" != clearStars("d*o*") {
	// 	t.Error("clearStars")
	// }
	// if "addb" != robotWithString("bdda") {
	// 	t.Error("robotWithString")
	// }
	// if "abc" != robotWithString("bac") {
	// 	t.Error("robotWithString")
	// }
}
