package main

import (
    "cmp"
    "fmt"
    "math"
    "slices"
)

func maxCount2554_1(banned []int, n int, maxSum int) int {
    chosen_count := 0
    current_sum := 0
    for number := 1; number <= min(n, maxSum); number++ {
        if !slices.Contains(banned, number) {
            if current_sum + number <= maxSum {
                chosen_count += 1
                current_sum += number
                // fmt.Println(number)
            }
        }
    }
    return chosen_count
}

// 66ms
func maxCount2554(banned []int, n int, maxSum int) int {
    chosen_count := 0
    current_sum := 0
    banned_set := make(map[int]struct{})
    for _,banned_number := range banned {
        banned_set[banned_number] = struct{}{}
    }
    for number := 1; number <= min(n, maxSum); number++ {
        if _, ok := banned_set[number]; !ok {
            if current_sum + number > maxSum {
                break
            }
            chosen_count += 1
            current_sum += number
        }
    }
    return chosen_count
}

func maxCount2554_3(banned []int, n int, maxSum int) int {
    // TODO maxSum vs sum(banned)
    chosen_count := 0
    current_sum := 0
    for number := 1; number <= min(n, maxSum); number++ {
        if !slices.Contains(banned, number) {
            if current_sum + number > maxSum {
                break
            }
            chosen_count += 1
            current_sum += number
        }
    }
    return chosen_count
}

func maxCount2554_2(banned []int, n int, maxSum int) int {
    //      1  2  3  4  5
    // sum  1  3  6 10 15    
    //      n*(n+1)/2 <= maxSum  -->  from sqrt(2*maxSum) up
    // ban     2     4     |
    // sum        4     9  | 
    from_number := min(n, int(math.Sqrt(float64(2*maxSum)))-1)
    from_sum:= from_number * (from_number+1) / 2
    chosen_count := from_number
    current_sum := from_sum
    fmt.Println(from_number, current_sum, chosen_count)
    banned_set := make(map[int]struct{})
    for _, banned_number := range banned {
        if banned_number <= from_number {
            if _, ok := banned_set[banned_number]; !ok { // dup in banned
                chosen_count -= 1
                current_sum -= banned_number
            }
        }
        banned_set[banned_number] = struct{}{}
    }
    fmt.Println(from_number, current_sum, chosen_count)
    number := from_number
    for number < n && current_sum <= maxSum {
        number += 1
        if _, ok := banned_set[number]; !ok {
            current_sum += number
            if current_sum <= maxSum {
                chosen_count += 1
            }
        }
        fmt.Println(number, current_sum, chosen_count)
    }
    return chosen_count
}

func maxCount2554_4(banned []int, n int, maxSum int) int {
    //      1  2  3  4  5
    // sum  1  3  6 10 15    
    //      n*(n+1)/2 <= maxSum  -->  from sqrt(2*maxSum) up
    // ban     2     4     |
    // sum        4     9  | 
    from_number := max(min(n, int(math.Sqrt(float64(2*maxSum)))-1), 1)
    from_sum:= from_number * (from_number+1) / 2
    chosen_count := from_number
    current_sum := from_sum
    fmt.Println(from_number, current_sum, chosen_count)
    banned_set := make(map[int]struct{})
    for _, banned_number := range banned {
        if banned_number <= from_number {
            if _, ok := banned_set[banned_number]; !ok { // dup in banned
                chosen_count -= 1
                current_sum -= banned_number
            }
        }
        banned_set[banned_number] = struct{}{}
    }
    fmt.Println(from_number, current_sum, chosen_count)
    for number := from_number+1; number <= min(n, maxSum); number++ {
        if _, ok := banned_set[number]; !ok {
            if current_sum + number > maxSum {
                break
            }
            chosen_count += 1
            current_sum += number
        }
        fmt.Println(number, current_sum, chosen_count)
    }
    return chosen_count
}

func minimumSize1760_fail_1(nums []int, maxOperations int) int {
    // 8,1 4,1 2,2
    // 6,1 4,1 2,3
    // 4,2 2,4
    // 4,1 2,6
    // 2,8
    maxBag := len(nums) + maxOperations
    total := 0
    var numSlice, countSlice []int
    for _, number := range nums {
        total += number
        i, found := slices.BinarySearch(numSlice, number)
        if found {
            countSlice[i] += 1
        } else {
            numSlice = slices.Insert(numSlice, i, number)
            countSlice = slices.Insert(countSlice, i, 1)
        }
    }
    fmt.Println(numSlice, countSlice)
    optimalSize := total / maxBag
    fmt.Println(total, maxBag, "optimalSize", optimalSize)
    var maxSize, maxSizeCount int
    for i := 0; i < maxOperations; i++ {
        fmt.Println("begin", i, numSlice, countSlice)
        maxSize = numSlice[len(numSlice)-1]
        maxSizeCount = countSlice[len(numSlice)-1]
        if maxSizeCount == 1 {
            numSlice = numSlice[:len(numSlice)-1]
            maxSize -= optimalSize
            countSlice = countSlice[:len(countSlice)-1]
            iNum, found := slices.BinarySearch(numSlice, maxSize)
            if found {
                countSlice[iNum] += 1
            } else {
                numSlice = slices.Insert(numSlice, iNum, maxSize)
                countSlice = slices.Insert(countSlice, iNum, 1)
            }
        } else {
            countSlice[len(numSlice)-1] -= 1
        }
        iNum, found := slices.BinarySearch(numSlice, optimalSize)
        if found {
            countSlice[iNum] += 1
        } else {
            numSlice = slices.Insert(numSlice, iNum, optimalSize)
            countSlice = slices.Insert(countSlice, iNum, 1)
        }
    }
    fmt.Println("end", numSlice, countSlice)
    return numSlice[len(numSlice)-1]
}

func minimumSize1760(nums []int, maxOperations int) int {
    panic("minimumSize1760 not implemented")
}

func isArraySpecial3152_1(nums []int, queries []int) bool {
    // 0 1 1 0 1 0
    // 0   1   1     nums1
    //   1   0   0   nums2
    // 0   1         nums1[0:2]
    //   1           nums1[0:2]
    //     1   1     nums1[2:5]
    //       0   0   nums2[2:5]
    return true
}

func isArraySpecial3152_time(nums []int, queries [][]int) []bool {
    evenMap := make(map[int]struct{})
    oddMap := make(map[int]struct{})
    for _, number := range nums {
        if number % 2 == 0 {
            evenMap[number] = struct{}{}
        } else {
            oddMap[number] = struct{}{}
        }
    }
    results := make([]bool, len(queries))
    for queryIndex, query := range queries {
        results[queryIndex] = true
        start := query[0]
        end := query[1]
        _, previousEven := evenMap[nums[start]]
        for index := start+1; index <= end; index++ {
            _, currentOdd := oddMap[nums[index]]
            fmt.Println(nums[index-1], nums[index], "previousEven", previousEven, "currentOdd", currentOdd)
            if previousEven != currentOdd {
                results[queryIndex] = false
                break
            }
            previousEven = !previousEven
        }
    }
    return results
}

func isArraySpecial3152(nums []int, queries [][]int) []bool {
    previousEven := nums[0] % 2 == 0
    falseIndexSlice := []int {}
    for i := 1; i < len(nums); i++ {
        currentOdd := nums[i] % 2 == 1
        // fmt.Println(nums[i-1], nums[i], "previousEven", previousEven, "currentOdd", currentOdd)
        if previousEven != currentOdd {
           falseIndexSlice = append(falseIndexSlice, i)
        }
        previousEven = !currentOdd
    }
    fmt.Println(falseIndexSlice)
    results := make([]bool, len(queries))
    for queryIndex, query := range queries {
        results[queryIndex] = true
        if query[0] != query[1] {
            i2, found2 := slices.BinarySearch(falseIndexSlice, query[1])
            if found2 {
                results[queryIndex] = false
            } else {
                fmt.Println(query, i2)
                if len(falseIndexSlice) > 0 && i2 >= len(falseIndexSlice) { // "to" is right of falseIndexSlice, get last index to compare
                    leftFalseIndex := falseIndexSlice[len(falseIndexSlice)-1]
                    results[queryIndex] = leftFalseIndex <= query[0]
                } else if i2 > 0 { // "to" is within falseIndexSlice, get previous index to compare
                    leftFalseIndex := falseIndexSlice[i2-1]
                    results[queryIndex] = leftFalseIndex <= query[0]
                }
            }
        }
    }
    return results
}

func maximumLength2981(s string) int {
    // 1    ..a..a..a..
    // 2    ..aaa.. ..aa..aa..aa..
    // n-1  ..aaaaa..
    specialStringMap := make(map[byte][]int)
    specialStringLength := 1
    previousChar := s[0]
    var currentChar byte
    for i := 1; i <= len(s); i++ {
        if i == len(s) {
            currentChar = 0
        } else {
            currentChar = s[i]
        }
        if currentChar == previousChar {
            specialStringLength += 1
        } else {
            slice, sliceFound := specialStringMap[previousChar]
            if sliceFound {
                index, _ := slices.BinarySearch(slice, specialStringLength)
                slice = slices.Insert(slice, index, specialStringLength)
                if len(slice) > 3 {
                    slice = slice[len(slice)-3:]
                }
                specialStringMap[previousChar] = slice
            } else {
                specialStringMap[previousChar] = []int {specialStringLength}
            }
            previousChar = currentChar
            specialStringLength = 1
        }
    }
    fmt.Println(specialStringMap)
    maxLength := 0
    var lengthMax, lengthNext int
    for _, lengthSlice := range specialStringMap {
        var possibleLengths []int
        if len(lengthSlice) == 3 && lengthSlice[0] == lengthSlice[1] && lengthSlice[1] == lengthSlice[2] {
            maxLength = max(maxLength, lengthSlice[0])
            continue
        }
        if len(lengthSlice) == 3 {
            lengthMax = lengthSlice[2]
            lengthNext = lengthSlice[1]
        } else if len(lengthSlice) == 2 {
            lengthMax = lengthSlice[1]
            lengthNext = lengthSlice[0]
        }
        if len(lengthSlice) == 3 || len(lengthSlice) == 2 {
            fmt.Println(lengthSlice, lengthNext, lengthMax)
            if lengthMax > lengthNext {
                if lengthMax == lengthNext + 1 {
                    possibleLengths = append(possibleLengths, lengthNext)
                }
                if lengthMax >= 3 {
                    possibleLengths = append(possibleLengths, lengthMax - 2)
                }
            } else {
                possibleLengths = append(possibleLengths, lengthMax - 1)
            }
            maxLength = max(maxLength, slices.Max(possibleLengths))
        } else {
            lengthMax := lengthSlice[0]
            if lengthMax >= 3 {
                lengthMax = lengthMax - 2
            } else {
                lengthMax = 0
            }
            maxLength = max(maxLength, lengthMax)
        }
    }
    if maxLength == 0 {
        return -1
    }
    return maxLength
}

func maximumBeauty2779_fail_1(nums []int, k int) int {
    // 2
    // 4       2,3,4,5,6
    // 6           4,5,6,7,8
    // 3     1,2,3,4,5
    // 1   0,1,2,3
    // 2   0,1,2,3,4
    // 5  0..20
    // 57       45..72
    // 46    31..51
    finalStart := max(0, nums[0] - k)
    finalEnd := nums[0] + k
    finalCommonCount := len(nums)
    finalCount := 1
    fmt.Println(nums[0], finalStart, finalEnd, finalCount)
    for _, number := range nums[1:] {
        start := max(max(0, number - k), finalStart)
        end := min(number + k, finalEnd)
        commonCount := max(0, end - start + 1)
        if commonCount > 0 {
            finalStart = start
            finalEnd = end
            finalCommonCount = min(finalCommonCount, commonCount)
            finalCount += 1
        }
        fmt.Println(number, "commonCount", commonCount, "finalStart", finalStart, "finalEnd", finalEnd, "finalCount", finalCount)   
    }
    return finalCount
}

func maximumBeauty2779_time(nums []int, k int) int {
    // 2
    // 4       2,3,4,5,6
    // 6           4,5,6,7,8
    // 3     1,2,3,4,5
    // 1   0,1,2,3
    // 2   0,1,2,3,4
    // SUM 2 3 4 5 4 3 2 1 1
    // 5  0..20
    // 57       45..72
    // 46    31..51
    numberMap := make(map[int]int)
    for _, number := range nums {
        for i := max(0, number - k); i <= number + k; i++ {
            numberMap[i] = numberMap[i] + 1 // default 0
        }
    }
    fmt.Println(numberMap)
    maxCount := 0
    for _, count := range numberMap {
        maxCount = max(maxCount, count)
    }
    return maxCount
}

func maximumBeauty2779_time_2(nums []int, k int) int {
    // 5   0 .. 20
    // 57             42 .. 72
    // 46          31 .. 61
    //       0  20 31 42 61 72
    // for
    // 5   .......
    // 57                +1
    // 46             +1
    //-------------------------
    //     0  2 3 4  6  8
    // 4      2  ..  6
    // 6          4 ..  8
    // 1   0 .. 3
    // 2   0 ..   4
    // for
    // 4       +1+1
    // 6            +1
    // 1     +1
    // 2     +1+1
    numberMap := make(map[int]int)
    var numberSlice []int
    for _, number := range nums {
        start := max(0, number - k)
        end := number + k
        if index, found := slices.BinarySearch(numberSlice, start); !found {
            numberSlice = slices.Insert(numberSlice, index, start)
        }
        if index, found := slices.BinarySearch(numberSlice, end); !found {
            numberSlice = slices.Insert(numberSlice, index, end)
        }
        numberMap[start] = numberMap[start] + 1 // default 0
        if k!= 0 {
            numberMap[end] = numberMap[end] + 1 // default 0
        }
    }
    // fmt.Println(numberMap)
    fmt.Println(numberSlice)
    for _, number := range nums {
        start, _ := slices.BinarySearch(numberSlice, max(0, number - k))
        end, _ := slices.BinarySearch(numberSlice, number + k)
        for index := start + 1; index < end; index++ {
            numberMap[numberSlice[index]] = numberMap[numberSlice[index]] + 1
        }
    }
    fmt.Println(numberMap)
    maxCount := 0
    for _, count := range numberMap {
        maxCount = max(maxCount, count)
    }
    return maxCount
}

func maximumBeauty2779(nums []int, k int) int {
    // 4          2   ..  6
    // 6               4 ..  8
    // 1       0 ..  3
    // 2       0 ..    4
    //         0  2  3 4  6  8
    // ends  3,4  6    8
    //       3,4
    //        3,4,6
    //           3,4,6
    //             4,6,8
    //                  6,8
    //                       8
    //         ^  ^  ^ ^  ^  ^
    //         |  |  | |  |  |
    //         2  3  3 3  2  1
    numberToEndsMap := make(map[int][]int)
    var numberSlice []int
    for _, number := range nums {
        start := max(0, number - k)
        end := number + k
        if index, found := slices.BinarySearch(numberSlice, start); !found {
            numberSlice = slices.Insert(numberSlice, index, start)
        }
        if index, found := slices.BinarySearch(numberSlice, end); !found {
            numberSlice = slices.Insert(numberSlice, index, end)
        }
        index, _ := slices.BinarySearch(numberToEndsMap[start], end)
        numberToEndsMap[start] = slices.Insert(numberToEndsMap[start], index, end)
    }
    fmt.Println(nums, numberSlice, numberToEndsMap)
    if k == 0 {
        finalCount := 1
        for _, v := range(numberToEndsMap) {
            finalCount = max(finalCount, len(v))
        }
        return finalCount
    }
    var endSlice = numberToEndsMap[numberSlice[0]]
    finalCount := len(endSlice)
    fmt.Println(endSlice)
    for _, number := range numberSlice[1:] {
        thisEndSlice, found := numberToEndsMap[number]
        if found {
            for _, thisEnd := range thisEndSlice {
                index, _ := slices.BinarySearch(endSlice, thisEnd)
                endSlice = slices.Insert(endSlice, index, thisEnd)
            }
        }
        finalCount = max(finalCount, len(endSlice))
        fmt.Println("number", number, "endSlice", endSlice, "finalCount", finalCount)
        for len(endSlice) > 0 && endSlice[0] == number { // for next loop
            endSlice = endSlice[1:]
        }
    }
    return finalCount
}

func pickGifts2558(gifts []int, k int) int64 {
    // 25 64 9  4 100
    // 25 64 9  4  10
    // 25 8  9  4  10
    // 5  8  9  4  10 
    // 5  8  9  4   3
    slices.Sort(gifts)
    // fmt.Println(gifts)
    for i := 0; i < k; i++ {
        remain := int(math.Sqrt(float64(gifts[len(gifts)-1])))
        index, _ := slices.BinarySearch(gifts, remain)
        gifts = slices.Insert(gifts[:len(gifts)-1], index, remain)
    }
    fmt.Println(gifts)
    var sum int64
    for _, val := range gifts { sum += int64(val) }
    return sum
}

func findScore2593_fail_1(nums []int) int64 {
    // 2,1,3,4,5,2
    //       4 5 2  1
    //       4      2
    //              4
    var sum int64
    for len(nums) > 0 {
        minIndex := -1
        minNumber := math.MaxUint32
        for index, number := range nums {
            if number < minNumber {
                minIndex = index
                minNumber = number
            }
        }
        sum += int64(minNumber)
        var before, after []int
        if minIndex > 1 { before = nums[:minIndex-1] }
        if minIndex < len(nums) - 2 { after = nums[minIndex+2:] }
        nums = append(before, after...)
        fmt.Println(minIndex, minNumber, nums, sum)
    }
    return sum
}

func findScore2593_time(nums []int) int64 {
    // 2,2,1,3,1,5,2     0 1 2 3 4 5 6  remove adjacents if -1 or +1
    // 2       1 5 2  1  0       4 5 6
    // 2           2  1  0           6
    //             2  2              6
    //                2   
    indexSlice := make([]int, len(nums))
    for i := range indexSlice { indexSlice[i] = i }
    fmt.Println(nums)
    var sum int64
    for len(indexSlice) > 0 {
        minIndex := -1
        minNumber := math.MaxUint32
        for index, indexInNums := range indexSlice {
            number := nums[indexInNums]
            if number < minNumber {
                minIndex = index
                minNumber = number
            }
        }
        sum += int64(minNumber)
        var before, after []int
        fmt.Println(indexSlice, minIndex, minNumber)
        if minIndex >= 1 {
            if indexSlice[minIndex-1] == indexSlice[minIndex]-1 {
                before = indexSlice[:minIndex-1]
            } else {
                before = indexSlice[:minIndex]
            }
        }
        if minIndex <= len(indexSlice) - 2 {
            if indexSlice[minIndex+1] == indexSlice[minIndex]+1 {
                after = indexSlice[minIndex+2:]
            } else {
                after = indexSlice[minIndex+1:]
            }
        }
        indexSlice = append(before, after...)
        fmt.Println(minIndex, minNumber, indexSlice, sum)
    }
    return sum
}

func findScore2593_time_2(nums []int) int64 {
    //    2,2,1,3,1,5,2
    //    1 1 2 2 2 3 5  sortedNums
    //
    //    2 4 0 1 6 3 5  oriInds slice
    // 1    4 0   6   5              not in markInds map
    // 2      0   6
    // 3          6
    // 4
    var sortedNums, oriInds []int
    markedInds := make(map[int]struct{})
    for i := len(nums)-1; i >= 0; i-- {  // backwards to insert same values at higher index first
        index, _ := slices.BinarySearch(sortedNums, nums[i])
        sortedNums = slices.Insert(sortedNums, index, nums[i])
        oriInds = slices.Insert(oriInds, index, i)
    }
    fmt.Println(oriInds, sortedNums)
    var sum int64
    for i := 0; len(markedInds) <= len(sortedNums) && i < len(sortedNums); i++ {
        fmt.Println(i, sortedNums[i], oriInds)
        if _, marked := markedInds[oriInds[i]]; marked { continue }
        sum += int64(sortedNums[i])
        if oriInds[i] >= 1 { markedInds[oriInds[i]-1] = struct{}{} }
        if oriInds[i] <= len(sortedNums) - 2 { markedInds[oriInds[i]+1] = struct{}{} }
        markedInds[oriInds[i]] = struct{}{} // to check length
    }
    return sum
}

func findScore2593_time_3(nums []int) int64 {
    //    2,2,1,3,1,5,2
    //    1 1 2 2 2 3 5  sortedNums
    type VI struct {
        Value int
        Index int
    }
    var sortedNums []VI
    for i := len(nums)-1; i >= 0; i-- {  // backwards to insert same values at higher index first
        index, _ := slices.BinarySearchFunc(sortedNums, VI{nums[i], i}, func(a, b VI) int {
            return cmp.Compare(a.Value, b.Value)
            // return cmp.Or(
            //     cmp.Compare(a.Value, b.Value),
            //     cmp.Compare(a.Index, b.Index),
            // )
        })
        sortedNums = slices.Insert(sortedNums, index, VI{nums[i], i})
    }
    // fmt.Println(oriInds, sortedNums)
    markedInds := make(map[int]struct{})
    var total int64
    for _, vi := range sortedNums {
        if _, found := markedInds[vi.Index]; found { continue }
        total += int64(vi.Value)
        markedInds[vi.Index-1] = struct{}{}
        markedInds[vi.Index+1] = struct{}{}
    }
    return total
}

func findScore2593(nums []int) int64 {
    panic("not implemented")
}

func lengthOfLongestSubstring3(s string) int {
    maxLength := 0
    if len(s) == 1 {
        maxLength = 1
    }
    for i := 0; i < len(s)-1; i++ {
        charMap := make(map[byte]struct{})
        charMap[s[i]] = struct{}{}
        length := 1
        for j := i+1; j < len(s); j++ {
            fmt.Println(i, j, charMap, s[j])
            if _, found := charMap[s[j]]; found { break }
            charMap[s[j]] = struct{}{}
            length += 1
        }
        maxLength = max(maxLength, length)
    }
    return maxLength
}

func continuousSubarrays2762_fail_1(nums []int) int64 {
    // 5 4 2 4
    // 5 4
    // 5 4 2     5 - 2 > 2 --> stop
    //   4 2
    //   4 2 4
    //     2 4
    m := make(map[int]int)
    for i := 0; i < len(nums)-1; i++ {
        for j := i+1; j < len(nums); j++ {
            diff := nums[i] - nums[j]
            if diff < -2 || diff > 2 { break }
            fmt.Println(i, j, nums[i], nums[j])
            m[j - i + 1] = m[j - i + 1] + 1 // default 0
        }
    }
    var total int64
    for _, v := range m { total += int64(v) }
    return total + int64(len(nums))
}

func continuousSubarrays2762(nums []int) int64 {
    panic("not implemented")
}

func main() {
    fmt.Println(continuousSubarrays2762([]int {65,66,67,66,66,65,64,65,65,64})) // 43
    fmt.Println(continuousSubarrays2762([]int {5,4,2,4})) // 8
    // fmt.Println(lengthOfLongestSubstring3(" ")) // 1
    // fmt.Println(lengthOfLongestSubstring3("abcabcbb")) // abc 3
    // fmt.Println(lengthOfLongestSubstring3("pwwkew")) // wke 3
    // fmt.Println(findScore2593([]int {46,777,1916,780,1857,1523,1016,389,117,934,1121,191,471,399,949,763,517,928,463,438,1496,1490,1552,211,280,122,200,1980,1437,1496,879,866,609,1923,836,1482,460,1080,1135,756,1870,30,1841,1860,1812,1121,1715,1930,1997,1531,1939,1674,346})) // 17122
    // fmt.Println(findScore2593([]int {2,2,1,3,1,5,2})) // 6
    // fmt.Println(findScore2593([]int {10,44,10,8,48,30,17,38,41,27,16,33,45,45,34,30,22,3,42,42})) // 212
    // fmt.Println(findScore2593([]int {2,1,3,4,5,2})) // 7
    // fmt.Println(pickGifts2558([] int {25,64,9,4,100}, 4))  // 29
    // fmt.Println(maximumBeauty2779([] int {13, 46, 71}, 29))  // 3
    // fmt.Println(maximumBeauty2779([] int {72,95,53,58,12,93,9,12,95,65}, 24))  // 7
    // fmt.Println(maximumBeauty2779([] int {13,68,81,61,13,70,23,46,4}, 5))  // 3
    // fmt.Println(maximumBeauty2779([] int {38,11,31,15,50,15}, 0))  // 2
    // fmt.Println(maximumBeauty2779([] int {45,33,34,35,70,35}, 21))  // 6
    // fmt.Println(maximumBeauty2779([]int {100000}, 0)) // 1
    //fmt.Println(maximumBeauty2779([]int {83,10,99,99}, 18)) // 3
    // fmt.Println(maximumBeauty2779([]int {5,57,46}, 15)) // 2
    // fmt.Println(maximumBeauty2779([]int {4,6,1,2}, 2)) // 3
    // fmt.Println(maximumLength2981("acc")) // -1
    // fmt.Println(maximumLength2981("abcaba")) // 1
    // fmt.Println(maximumLength2981("abcdef")) // -1
    // fmt.Println(maximumLength2981("aaabcaabaaacaa")) // 2
    // fmt.Println(maximumLength2981("eccdnmcnkl")) // 1
    // fmt.Println(maximumLength2981("eeeyyyybbbbbbbbssppb")) // 6
    // fmt.Println(maximumLength2981("ceeeeeeeeeeeebmmmfffeeeeeeeeeeeewww")) // 11
    // fmt.Println(maximumLength2981("hejlgbsjpedppppppdddddpphpppiiiiiga")) // 4
    // fmt.Println(isArraySpecial3152([]int {1,4}, [][]int {{0,1}}))
    // fmt.Println(isArraySpecial3152([]int {6,6,4,6,9,2,2}, [][]int {{2,4}}))
    // fmt.Println(isArraySpecial3152([]int {2,8,3,8,10}, [][]int {{0,3}}))
    // fmt.Println(isArraySpecial3152([]int {10,2,10,9,7}, [][]int {{2,3}}))
    // fmt.Println(isArraySpecial3152([]int {7,7}, [][]int {{1,1}}))
    // fmt.Println(isArraySpecial3152([]int {1,2,4,6,7,8,10,11}, [][]int {{0,2},{2,3},{4,7},{6,7},{4,5}}))
    // fmt.Println(isArraySpecial3152([]int {4,3,1,6}, [][]int {{0,2},{2,3}}))
    // fmt.Println(minimumSize1760([]int {431,922,158,60,192,14,788,146,788,775,772,792,68,143,376,375,877,516,595,82,56,704,160,403,713,504,67,332,26}, 80)) // 129
    // fmt.Println(minimumSize1760([]int {7,17}, 2)) // 7
    // fmt.Println(minimumSize1760([]int {2,4,8,2}, 4))
    // fmt.Println(maxCount2554([]int {8155, 8108}, 2431, 7821))
    // fmt.Println(maxCount2554([]int {179,266,77,196,59,313,286,41,21,201,57,237,74,333,101,281,227,25,138,10,304,55,50,72,244,113,159,330,154,156,311,170,283,9,224,46,197,2,325,237,54,168,275,166,236,30,250,48,274,331,240,153,312,63,303,342,79,37,165,20,79,293,103,152,215,44,56,196,29,251,264,210,212,135,296,123,289,257,208,309,67,114,170,119,337,163,242,162,109,318,51,105,272,240,107,226,224,188,224,317,27,102,63,128,3,133,27,134,186,220,198,24,274,287,267,8,13,322,278,166,304,165,342,89,184,300,312,339,163,307,123,137,293,227,229,57,66,13,71,233,260,79,228,301,4,4,89,196,193,337,205,51,144,99,104,73,10,311,240,168,77,244,114,217,186,134,229,241,46,89,54,127}, 4085, 109718563))
    // fmt.Println(maxCount2554([]int {176,36,104,125,188,152,101,47,51,65,39,174,29,55,13,138,79,81,175,178,42,108,24,80,183,190,123,20,139,22,140,62,58,137,68,148,172,76,173,189,151,186,153,57,142,105,133,114,165,118,56,59,124,82,49,94,8,146,109,14,85,44,60,181,95,23,150,97,28,182,157,46,160,155,12,67,135,117,2,25,74,91,71,98,127,120,130,107,168,18,69,110,61,147,145,38}, 3016, 240))
    // fmt.Println(maxCount2554([]int {87,193,85,55,14,69,26,133,171,180,4,8,29,121,182,78,157,53,26,7,117,138,57,167,8,103,32,110,15,190,139,16,49,138,68,69,92,89,140,149,107,104,2,135,193,87,21,194,192,9,161,188,73,84,83,31,86,33,138,63,127,73,114,32,66,64,19,175,108,80,176,52,124,94,33,55,130,147,39,76,22,112,113,136,100,134,155,40,170,144,37,43,151,137,82,127,73}, 1079, 87)) // 9
	// fmt.Println(maxCount2554([]int {1,6,5}, 5, 6)) // 2
    // fmt.Println(maxCount2554([]int {1,2,3,4,5,6,7}, 8, 1)) // 0
    // fmt.Println(maxCount2554([]int {11}, 7, 50)) // 7
}
