package main

import "fmt"

func ExampleWordPattern() {
	fmt.Println(WordPattern("abba", "dog cat cat fish"))
	// Unordered output:
	// false
}

func ExampleKthCharacter() {
	fmt.Println(KthCharacter(5))
	// Unordered output:
	// 98
}

func ExampleKthCharacterII() {
	fmt.Println(KthCharacterII(10, []int{0, 1, 0, 1}))
	fmt.Println(KthCharacterII(33354182522397, []int{0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 0, 1, 0, 0, 1, 1, 0, 1, 0, 1, 1, 0, 1, 1, 1, 0, 1, 0, 1, 0, 1, 0, 0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 1, 0, 0, 0, 1, 0, 1, 1, 0, 1, 0, 0, 0, 1, 0, 1, 0, 1, 1, 0, 0, 0, 0, 1, 0, 1, 1, 0, 0, 1, 0, 0, 1, 1, 0, 1, 1, 1, 1, 1, 1, 0, 0, 0}))
	// Unordered output:
	// 98
	// 107
}

// func ExamplePossibleStringCount() {
// "aabbccdd", 7                                                                     --> 5 ("aabbccdd", "aabbccd", "aabbcdd", "aabccdd", and "abbccdd")
// "aaabbb", 3                                                                       --> 8
// "ggggggggaaaaallsssssaaaaaaaaaiiqqqqqqqqqqbbbbbbbvvfffffjjjjeeeeeefffmmiiiix", 34 --> 834168507
// fmt.Println(PossibleStringCount("aabbccdd", 7))
// fmt.Println(PossibleStringCount("aaabbb", 3))
// fmt.Println(PossibleStringCount("ggggggggaaaaallsssssaaaaaaaaaiiqqqqqqqqqqbbbbbbbvvfffffjjjjeeeeeefffmmiiiix", 34))
// Unordered output:
// 5
// 8
// 834168507
// }
