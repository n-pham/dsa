package main

import "testing"

func TestLengthAfterTransformations(t *testing.T) {
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
