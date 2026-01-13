package main

import (
	"fmt"
)

func ExampleFizzBuzz() {
	fmt.Println(FizzBuzz(3))
	fmt.Println(FizzBuzz(5))
	fmt.Println(FizzBuzz(15))
	// Output:
	// [1 2 Fizz]
	// [1 2 Fizz 4 Buzz]
	// [1 2 Fizz 4 Buzz Fizz 7 8 Fizz Buzz 11 Fizz 13 14 FizzBuzz]
}

func ExampleMinTimeToVisitAllPoints() {
	fmt.Println(MinTimeToVisitAllPoints([][]int{{1, 1}, {3, 4}, {-1, 0}}))
	fmt.Println(MinTimeToVisitAllPoints([][]int{{3, 2}, {-2, 2}}))
	// Output:
	// 7
	// 5
}

func ExampleLongestPalindrome() {
	fmt.Println(LongestPalindrome("abccccdd"))
	fmt.Println(LongestPalindrome("a"))
	// Output:
	// 7
	// 1
}

func ExampleMinimumDeleteSum() {
	fmt.Println(MinimumDeleteSum("sea", "eat"))
	fmt.Println(MinimumDeleteSum("delete", "leet"))
	fmt.Println(MinimumDeleteSum("", ""))
	fmt.Println(MinimumDeleteSum("a", ""))
	fmt.Println(MinimumDeleteSum("a", "b"))
	fmt.Println(MinimumDeleteSum("abc", "abc"))
	// Output:
	// 231
	// 403
	// 0
	// 97
	// 195
	// 0
}

func ExampleFindTheDifference() {
	fmt.Println(string(FindTheDifference("abcd", "abcde")))
	fmt.Println(string(FindTheDifference("", "y")))
	// Output:
	// e
	// y
}

func ExampleFirstUniqChar() {
	fmt.Println(FirstUniqChar("leetcode"))
	fmt.Println(FirstUniqChar("loveleetcode"))
	fmt.Println(FirstUniqChar("aabb"))
	// Output:
	// 0
	// 2
	// -1
}

func ExampleSumFourDivisors() {
	fmt.Println(SumFourDivisors([]int{21, 4, 7}))
	fmt.Println(SumFourDivisors([]int{21, 21}))
	fmt.Println(SumFourDivisors([]int{1, 2, 3, 4, 5}))
	fmt.Println(SumFourDivisors([]int{1, 6, 8, 10, 12})) // 6 (1,2,3,6), 8 (1,2,4,8), 10 (1,2,5,10), 12 (1,2,3,4,6,12)
	fmt.Println(SumFourDivisors([]int{}))                // Empty slice
	fmt.Println(SumFourDivisors([]int{30}))              // Has more than 4 divisors (1,2,3,5,6,10,15,30)
	fmt.Println(SumFourDivisors([]int{13}))              // Prime, only 2 divisors
	// Output:
	// 32
	// 64
	// 0
	// 45
	// 0
	// 0
	// 0
}

func ExampleReverseVowels() {
	fmt.Println(ReverseVowels("IceCreAm"))
	fmt.Println(ReverseVowels("leetcode"))
	fmt.Println(ReverseVowels("hello"))
	fmt.Println(ReverseVowels("aeiou"))
	fmt.Println(ReverseVowels("Aa"))
	fmt.Println(ReverseVowels("rhythm"))
	fmt.Println(ReverseVowels("a"))
	fmt.Println(ReverseVowels("b"))
	fmt.Println(ReverseVowels(""))
	fmt.Println(ReverseVowels("!@#$"))
	// Output:
	// AceCreIm
	// leotcede
	// holle
	// uoiea
	// aA
	// rhythm
	// a
	// b
	//
	// !@#$
}

func ExampleRepeatedNTimes() {
	fmt.Println(RepeatedNTimes([]int{1, 2, 3, 3}))
	fmt.Println(RepeatedNTimes([]int{2, 1, 2, 5, 3, 2}))
	fmt.Println(RepeatedNTimes([]int{5, 1, 5, 2, 5, 3, 5, 4}))
	// Output:
	// 3
	// 2
	// 5
}

func ExampleIsSubsequence() {
	fmt.Println(IsSubsequence("abc", "ahbgdc"))
	fmt.Println(IsSubsequence("axc", "ahbgdc"))
	fmt.Println(IsSubsequence("", "ahbgdc"))
	fmt.Println(IsSubsequence("abc", ""))
	fmt.Println(IsSubsequence("aa", "a"))
	// Output:
	// true
	// false
	// true
	// false
	// false
}

func ExampleCanConstruct() {
	fmt.Println(CanConstruct("a", "b"))
	fmt.Println(CanConstruct("aa", "ab"))
	fmt.Println(CanConstruct("aa", "aab"))
	// Output:
	// false
	// false
	// true
}
