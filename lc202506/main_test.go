package main

import (
	"testing"
)

func TestAll(t *testing.T) {
	if "" != clearStars("d*o*") {
		t.Error("clearStars")
	}
	// if "addb" != robotWithString("bdda") {
	// 	t.Error("robotWithString")
	// }
	// if "abc" != robotWithString("bac") {
	// 	t.Error("robotWithString")
	// }
}
