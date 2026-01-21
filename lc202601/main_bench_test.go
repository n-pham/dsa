package main

import (
	"testing"
)

func BenchmarkMinBitwiseArray(b *testing.B) {
	nums := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MinBitwiseArray(nums)
	}
}
