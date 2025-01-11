package main

import (
	"fmt"
	// "math"
	"slices"
	"strconv"
	"strings"
)

func maxScore1422_fail(s string) int {
	// 011101
	// iZCount = -1,0,4
	// zCount 2: oCount = len-1 - iZCount[zCount]
	// zCount 1: oCount = len-1 - iZCount[zCount] - 1 // zCount == 2
	panic("not implemented")
	iZCount := []int{-1}
	zTotal := 0
	for i, c := range s {
		if c == '0' {
			zTotal += 1
			iZCount = append(iZCount, i)
		}
	}
	fmt.Println(iZCount)
	if len(iZCount) == 1 {
		return len(s) - 1
	}
	maxS := max(0, iZCount[1]) // "1"s before first 0
	for zCount := len(iZCount) - 1; zCount > 0; zCount-- {
		// ignore if end is 0
		if iZCount[zCount] == len(s)-1 {
			continue
		}
		oCount := len(s) - 1 - iZCount[zCount] - (len(iZCount) - 1 - zCount)
		fmt.Println("maxS", maxS, "zCount", zCount, "oCount", oCount)
		maxS = max(maxS, zCount+oCount)
	}
	return maxS
}

func maxScore1422(s string) int {
	// 1422
	maxS, forwardSums, backwardSums := 0, make([]int, len(s)+1), make([]int, len(s)+1)
	for i, c := range s {
		forwardSums[i+1] = forwardSums[i]
		if c == '0' {
			forwardSums[i+1] += 1
		}
	}
	for i := len(s) - 1; i > 0; i-- {
		backwardSums[i] = backwardSums[i+1]
		if s[i] == '1' {
			backwardSums[i] += 1
		}
	}
	for i := 1; i <= len(s)-1; i++ {
		maxS = max(maxS, forwardSums[i]+backwardSums[i])
	}
	return maxS
}

func vowelStrings2559_time(words []string, queries [][]int) []int {
	queriesByStart := make(map[int]map[int]struct{})
	queriesByEnd := make(map[int]map[int]struct{})
	for i, q := range queries {
		if m, found := queriesByStart[q[0]]; found {
			m[i] = struct{}{}
			// queriesByStart[q[0]] = m
		} else {
			m := make(map[int]struct{})
			m[i] = struct{}{}
			queriesByStart[q[0]] = m
		}
		if m, found := queriesByEnd[q[1]]; found {
			m[i] = struct{}{}
			// queriesByStart[q[0]] = m
		} else {
			m := make(map[int]struct{})
			m[i] = struct{}{}
			queriesByEnd[q[1]] = m
		}
	}
	// fmt.Println(queriesByStart, queriesByEnd)
	rs := make([]int, len(queries))
	currentQueryByIdx := make(map[int]struct{})
	for i, w := range words {
		// fmt.Println("i", i, "w", w, "currentQueryByIdx", currentQueryByIdx)
		if qIs, found := queriesByStart[i]; found {
			// fmt.Println("Start qIs", qIs)
			for qi := range qIs {
				currentQueryByIdx[qi] = struct{}{}
			}
		}
		if isVowel(w[0]) && isVowel(w[len(w)-1]) {
			for qi := range currentQueryByIdx {
				rs[qi] += 1
			}
		}
		if qIs, found := queriesByEnd[i]; found {
			// fmt.Println("End qIs", qIs)
			for qi := range qIs {
				delete(currentQueryByIdx, qi)
			}
		}
	}
	return rs
}

func vowelStrings2559_time_2(words []string, queries [][]int) []int {
	ones := make([]int, len(words))
	for i, w := range words {
		if isVowel(w[0]) && isVowel(w[len(w)-1]) {
			ones[i] = int(1)
		}
	}
	fmt.Println(ones)
	rs := make([]int, len(queries))
	for i, q := range queries {
		fmt.Println(ones[q[0]:q[1]])
		for _, val := range ones[q[0] : q[1]+1] {
			rs[i] += val
		}
	}
	return rs
}

func vowelStrings2559(words []string, queries [][]int) []int {
	//     aba bcb ece  aa   e
	//       1   0   1   1   1
	//   0   1   1   2   3   4
	// 2559
	prefixSums := make([]int, len(words)+1)
	for i, w := range words {
		prefixSums[i+1] = prefixSums[i]
		if isVowel(w[0]) && isVowel(w[len(w)-1]) {
			prefixSums[i+1] += 1
		}
	}
	rs := make([]int, len(queries))
	for i, q := range queries {
		rs[i] = prefixSums[q[1]+1] - prefixSums[q[0]]
	}
	return rs
}

func isVowel(c byte) bool {
	switch c {
	case
		'a',
		'e',
		'i',
		'o',
		'u',
		'A',
		'E',
		'I',
		'O',
		'U':
		return true
	}
	return false
}

func waysToSplitArray220_12ms(nums []int) int {
	//     10  4 -8  7
	// --> 10 14  6 13
	//     13  3 -1  7 <--
	// 2270
	forwardSums := make([]int, len(nums)+1)
	for i, num := range nums {
		forwardSums[i+1] = forwardSums[i] + num
	}
	backwardSums := make([]int, len(nums)+1)
	for i := len(nums) - 1; i >= 0; i-- {
		backwardSums[i] = backwardSums[i+1] + nums[i]
	}
	// fmt.Println(forwardSums, backwardSums)
	cnt := 0
	for i := 0; i < len(nums)-1; i++ {
		// fmt.Println(i, forwardSums[i], backwardSums[i+1])
		if forwardSums[i+1] >= backwardSums[i+1] {
			cnt += 1
		}
	}
	return cnt
}

func waysToSplitArray220(nums []int) int {
	//     10  4 -8  7  --> sum 13
	//     10  3            cur, sum-cur
	//        14 -1         cur, sum-cur
	//            6  7      cur, sum-cur
	// 2270
	var cnt, curSum, total int
	for _, num := range nums {
		total += num
	}
	for i := 0; i < len(nums)-1; i++ {
		curSum += nums[i]
		if curSum >= total-curSum {
			cnt++
		}
	}
	return cnt
}

func executeInstructions2120(n int, startPos []int, s string) []int {
	//    R  R  D  D  L  U
	// dx 1  2  2  2  1  1
	// dy 0  0 -1 -2 -1  0
	// dx -1  0 -1 -1 -1  0 <--|
	// dy  1  1  1  0 -1 -1 <--|
	// 2120
	panic("not implemented")
}

func minOperations1551(n int) int {
	// 1551
	// 1 3 5        --> n 3
	// 2 3 4
	// 3 3 3
	// 2+0+2
	// 1 3 5 7 9 11 --> n 6
	// 5+3+1
	rs := 0
	for d := n - 1; d > 0; d -= 2 {
		rs += d
	}
	return rs
}

func minSteps1347_115ms(s string, t string) int {
	// 1347
	// bab map aba
	// b     a  b?
	//  a        ?
	//   b       a
	// leetcode     map     practice
	// l          practice
	//  e         practic
	//   e
	//    t       pracic
	//     c      praci
	//      o     praci
	//       d    praci
	//        e   praci
	tMap := make(map[byte]int)
	tI := 0
	for sI := 0; sI < len(s); sI++ {
		fmt.Println(sI, s[sI], tI, tMap)
		if cnt, found := tMap[s[sI]]; found {
			if cnt == 1 {
				delete(tMap, s[sI])
			} else {
				tMap[s[sI]] = cnt - 1
			}
			continue
		}
		for (tI < len(t)) && (s[sI] != t[tI]) {
			cnt, found := tMap[t[tI]]
			tMap[t[tI]] = 1
			if found {
				tMap[t[tI]] += cnt
			}
			tI++
		}
		if tI < len(t) {
			tI++
		}
	}
	rs := 0
	for _, cnt := range tMap {
		rs += cnt
	}
	return rs
}

func minSteps1347(s string, t string) int {
	// 1347
	rs, charCnt := 0, [26]int{}
	for i, c := range s {
		charCnt[c-97]++
		charCnt[t[i]-97]--
	}
	for _, cnt := range charCnt {
		rs += max(0, cnt)
	}
	return rs
}

func countPalindromicSubsequence1930(s string) int {
	// 1930
	//         a a b c a
	//         1 1 2 3 3      first a:0 b:2 c:3  last a:4 b:2 c:3
	//         b b c b a b a
	// uniqSum 1 1 2 2 3 3 3  first b:0 c:2 a:4  last b:5 c:2 a:6
	panic("not implemented")
	uniqPrefixSum, firsts, lasts, rs := make([]int, len(s)), make(map[rune]int), [26]int{}, 0
	for i, c := range s {
		// fmt.Println(i, c, firsts, lasts)
		lasts[c-'a'] = i
		if _, found := firsts[c-'a']; !found {
			firsts[c-'a'] = i
		}
		uniqPrefixSum[i] = len(firsts)
	}
	fmt.Println(uniqPrefixSum, firsts, lasts)
	for i, _ := range firsts {
		rs += uniqPrefixSum[lasts[i]] - uniqPrefixSum[firsts[i]]
	}
	return rs
}

func findTheWinner1823(n int, k int) int {
	// 1823
	panic("not implemented")
}

func minPairSum1877(nums []int) int {
	// 1877
	// 3 5 2 3
	// 3+3 5+2
	// 3+5 2+3
	panic("not implemented")
}

func shiftingLetters2381_time(s string, shifts [][]int) string {
	// 2381
	ds := make([]int, len(s))
	m := map[int]int{1: 1, 0: -1}
	for _, s := range shifts {
		for i := s[0]; i <= s[1]; i++ {
			ds[i] += m[s[2]]
		}
	}
	fmt.Print(ds)
	rs := []byte(s)
	for i, d := range ds {
		if d != 0 {
			if d < 0 {
				d = -(-d) % 26
			}
			num := int(rs[i]-'a') + d
			fmt.Println(i, num)
			if num < 0 {
				num = 26 + num
			}
			rs[i] = 'a' + byte(num%26)
		} else {
			rs[i] = s[i]
		}
	}
	return string(rs)
}

func shiftingLetters2381(s string, shifts [][]int) string {
	// 2381
	// 0 1 2 3 4 5 6 7 8
	//         1 1 1 1 1
	//         1
	//     1 1 1
	//            -1-1
	//    -1
	//-1-1-1
	//                 1
	//  -1-1-1
	panic("not implemented")
	ds := make([]int, len(s))
	fmt.Print(ds)
	rs := []byte(s)
	for i, d := range ds {
		if d != 0 {
			if d < 0 {
				d = -(-d) % 26
			}
			num := int(rs[i]-'a') + d
			fmt.Println(i, num)
			if num < 0 {
				num = 26 + num
			}
			rs[i] = 'a' + byte(num%26)
		} else {
			rs[i] = s[i]
		}
	}
	return string(rs)
}

func sortVowels2785_709ms(s string) string {
	// 2785
	vs, is, rs := []rune{}, []int{}, []rune(s)
	for i, c := range rs {
		if isVowel2785(c) {
			pos, _ := slices.BinarySearch(vs, c)
			vs = slices.Insert(vs, pos, c)
			iPos, _ := slices.BinarySearch(is, i)
			is = slices.Insert(is, iPos, i)
		}
	}
	fmt.Println(rs, vs, is)
	for i, c := range vs {
		rs[is[i]] = c
	}
	fmt.Println(rs)
	return string(rs)
}

func isVowel2785(c rune) bool {
	switch c {
	case
		'a',
		'e',
		'i',
		'o',
		'u',
		'A',
		'E',
		'I',
		'O',
		'U':
		return true
	}
	return false
}

func sortVowels2785(s string) string {
	// 2785
	//                             A E I O U a e i o u
	// lEetcOde is [1 2 5 7] cnts [0 1 0 1 0 0 2 0 0 0]
	vs := []rune{'A', 'E', 'I', 'O', 'U', 'a', 'e', 'i', 'o', 'u'}
	rs, cnts, is := []rune(s), [10]int{}, []int{}
	for i, c := range rs {
		if pos, found := slices.BinarySearch(vs, c); found {
			cnts[pos] += 1
			is = append(is, i)
		}
	}
	fmt.Println(is, cnts)
	currentI := 0
	for vI, cnt := range cnts {
		for j := currentI; j < currentI+cnt; j++ {
			rs[is[j]] = vs[vI]
		}
		currentI += cnt
	}
	return string(rs)
}

func findingUsersActiveMinutes1817(logs [][]int, k int) []int {
	// 1817
	// time 1 2 3 4 5 cnt
	// user   0     0   2
	// user   1 1       2
	// 0,5  0:5
	// 1,2  1:2
	// 0,2  0:5,2
	// 0,5  0:5,2
	// 1,3  1:2,3
	timesByUser := make(map[int]map[int]struct{})
	for _, log := range logs {
		if m, found := timesByUser[log[0]]; found {
			m[log[1]] = struct{}{}
		} else {
			timesByUser[log[0]] = map[int]struct{}{log[1]: struct{}{}}
		}
	}
	fmt.Println(timesByUser)
	rs := make([]int, k)
	for _, m := range timesByUser {
		rs[len(m)-1] += 1
	}
	return rs
}

func stringSequence3324(target string) []string {
	// 3324
	rs, prefix := []string{}, ""
	for cCnt := 0; cCnt < len(target); cCnt++ {
		for i := 'a'; i < rune(target[cCnt]); i++ {
			rs = append(rs, prefix, string(i))
		}
		prefix += string(target[cCnt])
	}
	return rs
}

func stringMatching1408(words []string) []string {
	// 1408
	rs := []string{}
	for i, w := range words {
		for j, otherWord := range words {
			if i == j {
				continue
			}
			if strings.Contains(otherWord, w) {
				rs = append(rs, w)
				break
			}
		}
	}
	return rs
}

func minimumPushes3016(word string) int {
	// 3016
	// 2-9 = 8 slots
	// a2 b2 c2 d2 e2 f2 g2 h2 i6
	rs, clickCnt, times, cntByChar := 0, 1, 0, make([]int, 26)
	for _, c := range word {
		cntByChar[c-'a'] += 1
	}
	slices.Sort(cntByChar)
	fmt.Println(cntByChar)
	for i := len(cntByChar) - 1; i >= 0 && cntByChar[i] > 0; i-- {
		rs += clickCnt * cntByChar[i]
		fmt.Println(i, cntByChar[i], clickCnt, times, rs)
		times += 1
		if times == 8 {
			times = 0
			clickCnt += 1
		}
	}
	return rs
}

func partitionLabels763(s string) []int {
	// 763
	// ababcbacadefegdehijhklij
	// a       a
	//  b   b
	//     c  c
	//          d    d
	//           e    e
	//            f
	//              g
	//                 h  h
	//                  i    i
	//                   j    j
	lefts, rights, rs := make(map[rune]int), [26]int{}, []int{}
	for i, c := range s {
		// fmt.Println(i, c, lefts, rights)
		rights[c-'a'] = i
		if _, found := lefts[c-'a']; !found {
			lefts[c-'a'] = i
		}
	}
	fmt.Println(lefts, rights)
	q, nextQ := lefts, make(map[rune]int)
	// TODO use rights, []int{}, q[0] is not random key
	for len(q) > 0 {
		c := rune(q[0])
		l, r := q[c], rights[c]
		delete(q, c)
		for len(q) > 0 {
			otherC := rune(q[0])
			if !(r < q[otherC] || l > rights[otherC]) {
				l, r = min(l, q[otherC]), max(r, rights[otherC])
			} else {
				nextQ[otherC] = q[otherC]
			}
			fmt.Println(q, otherC, nextQ)
			delete(q, otherC)
			fmt.Println(q, otherC, nextQ)
		}
		rs = append(rs, r-l+1)
		q = nextQ
	}
	return rs
}

func buildArray1441(target []int, n int) []string {
	// 1441
	rs, nextNum := []string{}, 1
	for _, num := range target {
		for i := nextNum; i < num; i++ {
			fmt.Println(i, nextNum, num)
			rs = append(rs, "Push", "Pop")
			nextNum += 1
		}
		rs = append(rs, "Push")
		nextNum += 1
	}
	return rs
}

func countDistinctIntegers2442(nums []int) int {
	// 2442
	m := make(map[int]struct{}, len(nums)*2)
	for _, num := range nums {
		m[num] = struct{}{}
		rev := 0
		for num > 0 {
			rev = (rev * 10) + (num % 10)
			num /= 10
		}
		m[rev] = struct{}{}
	}
	return len(m)
}

func reverse_int_111ms(value int) int {
	intString := strconv.Itoa(value)
	newSlice := make([]rune, len(intString))
	for i, c := range intString {
		newSlice[len(intString)-i-1] = c
	}
	newInt, _ := strconv.Atoi(string(newSlice))
	return newInt
}

func countPrefixSuffixPairs3042(words []string) int {
	// 3042
	rs := 0
	for i, s1 := range words {
		for _, s2 := range words[i+1:] {
			if len(s2) >= len(s1) {
				fmt.Println(s1, s2, s2[:len(s1)], s2[len(s2)-len(s1):])
			}
			if len(s2) >= len(s1) &&
				s1 == s2[:len(s1)] &&
				s1 == s2[len(s2)-len(s1):] {
				rs += 1
			}
		}
	}
	return rs
}

func subsets78(nums []int) [][]int {
	panic("not implemented")
}

func wateringPlants2079(plants []int, capacity int) int {
	// 2079
	// 2  2  3  3
	// 2  4 (7)
	//       3 (6)
	//          3
	stepCnt, amountSum := 0, 0
	for i, amount := range plants {
		amountSum += amount
		if amountSum > capacity {
			stepCnt += i+i
			amountSum = amount
		}
		stepCnt += 1
	}
	return stepCnt
}

func findThePrefixCommonArray2657_8ms(A []int, B []int) []int {
	// 2657
	// 1 3 2 4
	// 3 1 2 4
	// 0 2 3 4
	// 2 3 1
	// 3 1 2
	// 0 1 3
	n := len(A)
	rs, mA, mB, mAB := make([]int, n), make(map[int]struct{}, n), make(map[int]struct{}, n), make(map[int]struct{}, n)
	for i := 0; i < n; i++ {
		mA[A[i]] = struct{}{}
		mB[B[i]] = struct{}{}
		if _, found := mA[B[i]]; found {
			mAB[B[i]] = struct{}{}
		}
		if _, found := mB[A[i]]; found {
			mAB[A[i]] = struct{}{}
		}
		rs[i] = len(mAB)
	}
	return rs
}

func findThePrefixCommonArray2657(A []int, B []int) []int {
	// 2657
	n := len(A)
	m := make(map[int]int, n) // A 1 B 2 AB 3
	rs, mAB := make([]int, n), make(map[int]struct{}, n)
	for i := 0; i < n; i++ {
		if vA, foundA := m[A[i]]; foundA {
			if vA == 2 {
				m[A[i]] = 3
				mAB[A[i]] = struct{}{}
			}
		} else {
			m[A[i]] = 1
		}
		if vB, foundB := m[B[i]]; foundB {
			if vB == 1 {
				m[B[i]] = 3
				mAB[B[i]] = struct{}{}
			}
		} else {
			m[B[i]] = 2
		}
		rs[i] = len(mAB)
	}
	return rs
}

func prefixCount2185(words []string, pref string) int {
	// 2185
	cnt := 0
	for _, w := range words {
		if len(w) >= len(pref) && pref == w[:len(pref)] {
			cnt += 1
		}
	}
	return cnt
}

func maximumXOR2317(nums []int) int {
	// 2317
	//  11
	//  10
	// 100
	// 110
	panic("not implemented")
}

func wordSubsets916_109ms(words1 []string, words2 []string) []string {
	// 916
	// eo oo
	// amazon
	// facebook
	rs, m := []string{}, make(map[rune]int, 26)
	for _, w := range words2 {
		mw := make(map[rune]int, 26)
		for _, c := range w {
			mw[c-'a'] += 1
		}
		for c, cnt := range mw {
			m[c] = max(m[c], cnt)
		}
	}
	for _, w := range words1 {
		mw := make(map[rune]int, len(w))
		for c, cnt := range m {
			mw[c] = cnt
		}
		for _, c := range w {
			if cnt, found := mw[c-'a']; found {
				if cnt > 1 {
					mw[c-'a'] = cnt-1
				}  else {
					delete(mw, c-'a')
				}
			}
		}
		if len(mw) == 0 {
			rs = append(rs, w)
		}
	}
	return rs
}

func wordSubsets916(words1 []string, words2 []string) []string {
	// 916    [26]int is faster than map
	// eo oo
	// amazon
	// facebook
	rs, m := []string{}, [26]int{}
	for _, w := range words2 {
		mw := [26]int{}
		for _, c := range w {
			mw[c-'a'] += 1
		}
		for c, cnt := range mw {
			m[c] = max(m[c], cnt)
		}
	}
	for _, w := range words1 {
		isSub, mw := true, [26]int{}
		for _, c := range w {
			mw[c-'a'] += 1
		}
		for c, cnt := range m {
			fmt.Println(w, mw, mw[c], c, cnt)
			if cnt > 0 && mw[c] < cnt {
				isSub = false
				break
			}
		}
		if isSub {
			rs = append(rs, w)
		}
	}
	return rs
}

func rotateTheBox1861(boxGrid [][]byte) [][]byte {
	// 1861
	// ##*.*.   ##*.*.
	// ###*..   ###*..
	// ###.#.   ..####
	panic("not implemented")
}

func canConstruct1400(s string, k int) bool {
	// 1400 k palindromes
	// annabelle 2 --> "anna" + "elble", "anbna" + "elle", "anellena" + "b"
	// a2b1e2l2n2
	// e3c1d1l1o1t1 3
    if len(s) < k {
        return false
    }
	oddM := [26]int{} // default 0
	for _, c := range s {
		oddM[c-'a'] = oddM[c-'a'] ^ 1
	}
	oddCnt := 0
	for _, v := range oddM {
		oddCnt += v
	}
	if oddCnt > k {
		return false
	}
	return true
}

func canConstruct1400_39ms(s string, k int) bool {
	// 1400 k palindromes
	// annabelle 2 --> "anna" + "elble", "anbna" + "elle", "anellena" + "b"
	// a2b1e2l2n2
	// e3c1d1l1o1t1 3
    if len(s) < k {
        return false
    }
	oddM := make(map[rune]struct{}, 26)
	for _, c := range s {
		if _, found := oddM[c]; found {
			delete(oddM,c)
		} else {
			oddM[c] = struct{}{}
		}
	}
	if len(oddM) > k {
		return false
	}
	return true
}

func main() {
	fmt.Println(canConstruct1400("annabelle",2))
	fmt.Println(canConstruct1400("cr",7))
	// fmt.Println(wordSubsets916([]string {"amazon","apple","facebook","google","leetcode"}, []string {"e","oo"}))
	// fmt.Println(maximumXOR2317([]int{3,2,4,6})) // 7
	// fmt.Println(maximumXOR2317([]int{3,2,4,6})) // 11
	// fmt.Println(prefixCount2185([]string{"pay","attention","practice","attend"}, "at"))
	// fmt.Println(findThePrefixCommonArray2657([]int{1,3,2,4}, []int{3,1,2,4}))
	// fmt.Println(findThePrefixCommonArray2657([]int{2,3,1}, []int{3,1,2}))
	// fmt.Println(wateringPlants2079([]int{2, 2, 3, 3}, 5))
	// fmt.Println(subsets78([]int{1,2,3}))
	// fmt.Println(countPrefixSuffixPairs3042([]string{"a", "aba", "ababa", "aa"}))
	// fmt.Println(countDistinctIntegers2442([]int{1,13,10,12,31}))
	// fmt.Println(buildArray1441([]int {1,3}, 3))
	// fmt.Println(buildArray1441([]int {2,3,4}, 4))
	// fmt.Println(partitionLabels763("ababcbacadefegdehijhklij"))
	// fmt.Println(minimumPushes3016("hiknogatpyjzcdbe")) // 24
	// fmt.Println(minimumPushes3016("aabbccddeeffgghhiiiiii")) // 24
	// fmt.Println(stringMatching1408([]string {"mass","as","hero","superhero"}))
	// fmt.Println(stringSequence3324("abc"))
	// fmt.Println(findingUsersActiveMinutes1817([][]int {{0,5},{1,2},{0,2},{0,5},{1,3}}, 5))
	// fmt.Println(sortVowels2785("lEetcOde"))
	// fmt.Println(shiftingLetters2381("abc", [][]int{{0, 1, 0}, {1, 2, 1}, {0, 2, 1}}))
	// fmt.Println(shiftingLetters2381("dztz", [][]int{{0, 0, 0}, {1, 1, 1}})) // catz
	// fmt.Println(shiftingLetters2381("xuwdbdqik", [][]int{{4,8,0},{4,4,0},{2,4,0},{2,4,0},{6,7,1},{2,2,1},{0,2,1},{8,8,0},{1,3,1}})) // ywxcxcqii
	// fmt.Println(minPairSum1877([]int {3,5,2,3}))
	// fmt.Println(countPalindromicSubsequence1930("aabca")) // aba aaa aca
	// fmt.Println(countPalindromicSubsequence1930("bbcbaba")) // bbb bcb bab aba
	// fmt.Println(minSteps1347("gctcxyuluxjuxnsvmomavutrrfb", "qijrjrhqqjxjtprybrzpyfyqtzf")) // 18
	// fmt.Println(minSteps1347("leetcode", "practice")) // 5
	// fmt.Println(minSteps1347("bab", "aba")) // 1
	// fmt.Println(executeInstructions2120(3, []int {0,1}, "RRDDLU")) //
	// fmt.Println(waysToSplitArray220([]int {10,4,-8,7})) // 2
	// fmt.Println(waysToSplitArray220([]int {-2,-1})) // 0
	// fmt.Println(vowelStrings2559([]string{"aba", "bcb", "ece", "aa", "e"}, [][]int{{0, 2}, {1, 4}, {1, 1}})) //
	// fmt.Println(maxScore1422("00"))     // 1
	// fmt.Println(maxScore1422("010"))    // 2
	// fmt.Println(maxScore1422("011101")) // 5
	// fmt.Println(maxScore1422("1111"))   // 3
	// fmt.Println(maxScore1422("11100"))  // 2
}
