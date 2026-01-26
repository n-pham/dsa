package main

import (
	"fmt"
)

func ExampleMinimumDifference() {
	fmt.Println(MinimumDifference([]int{9,4,1,7}, 2))
	fmt.Println(MinimumDifference([]int{90}, 1))
	// Output:
	// 2
	// 0
}

func ExampleMinimumPairRemoval() {
	fmt.Println(MinimumPairRemoval([]int{5,2,3,1}))
	fmt.Println(MinimumPairRemoval([]int{1,2,2}))
	// Output:
	// 2
	// 0
}

func ExampleCountConsistentStrings() {
	fmt.Println(CountConsistentStrings("ab", []string{"ad","bd","aaab","baa","badab"}))
	fmt.Println(CountConsistentStrings("abc", []string{"a","b","c","ab","ac","bc","abc"}))
	fmt.Println(CountConsistentStrings("cad", []string{"cc","acd","b","ba","bac","bad","ac","d"}))
	// Output:
	// 2
	// 7
	// 4
}

func ExampleReverseDegree() {
	fmt.Println(ReverseDegree("abc"))
	fmt.Println(ReverseDegree("zaza"))
	// Output:
	// 148
	// 160
}

func ExampleGetConcatenation() {
	fmt.Println(GetConcatenation([]int{1, 2, 1}))
	fmt.Println(GetConcatenation([]int{1, 3, 2, 1}))
	// Output:
	// [1 2 1 1 2 1]
	// [1 3 2 1 1 3 2 1]
}

func ExampleRecoverOrder() {
	fmt.Println(RecoverOrder([]int{3, 1, 2, 5, 4}, []int{1, 3, 4}))
	fmt.Println(RecoverOrder([]int{1, 4, 5, 3, 2}, []int{2, 5}))
	// Output:
	// [3 1 4]
	// [5 2]
}

func ExampleTheMaximumAchievableX() {
	fmt.Println(TheMaximumAchievableX(4, 1))
	fmt.Println(TheMaximumAchievableX(3, 2))
	// Output:
	// 6
	// 7
}

func ExampleAddStrings() {
	fmt.Println(AddStrings("11", "123"))
	fmt.Println(AddStrings("456", "77"))
	fmt.Println(AddStrings("0", "0"))
	// Output:
	// 134
	// 533
	// 0
}

func ExampleThirdMax() {
	fmt.Println(ThirdMax([]int{3, 2, 1}))
	fmt.Println(ThirdMax([]int{1, 2}))
	fmt.Println(ThirdMax([]int{2, 2, 3, 1}))
	// Output:
	// 1
	// 2
	// 1
}

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

func ExampleMinBitwiseArray() {
	fmt.Println(MinBitwiseArray([]int{2}))
	fmt.Println(MinBitwiseArray([]int{3}))
	fmt.Println(MinBitwiseArray([]int{2, 3, 5, 7}))
	fmt.Println(MinBitwiseArray([]int{11, 13, 31}))
	// Output:
	// [-1]
	// [1]
	// [-1 1 4 3]
	// [9 12 15]
}

func ExampleMinPairSum() {
	fmt.Println(MinPairSum([]int{3,5,2,3}))
	fmt.Println(MinPairSum([]int{3,5,4,2,4,6}))
	// Output:
	// 7
	// 8
}

func ExampleMinimumAbsDifference() {
	fmt.Println(MinimumAbsDifference([]int{4, 2, 1, 3}))
	fmt.Println(MinimumAbsDifference([]int{1, 3, 6, 10, 15}))
	fmt.Println(MinimumAbsDifference([]int{3, 8, -10, 23, 19, -4, -14, 27}))
	// Output:
	// [[1 2] [2 3] [3 4]]
	// [[1 3]]
	// [[-14 -10] [19 23] [23 27]]
}
