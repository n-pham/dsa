package main

import "fmt"

func ExampleKthCharacter() {
	fmt.Println(KthCharacter(5))
	// Unordered output:
	// 98
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
