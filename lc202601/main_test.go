package main

import (
	"fmt"
)

func ExampleReverseVowels() {
	fmt.Println(ReverseVowels("IceCreAm"))
	fmt.Println(ReverseVowels("leetcode"))
	// Unordered output:
	// AceCreIm
	// leotcede
}
