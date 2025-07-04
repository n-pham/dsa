package main

func PossibleStringCount_fail(word string, k int) (cnt int) {
	// 3333
	const MOD = 1e9 + 7
	n := len(word)

	if n == 0 {
		if k == 0 {
			return 1
		}
		return 0
	}

	var groupLengths []int
	count := 1
	for i := 1; i < n; i++ {
		if word[i] == word[i-1] {
			count++
		} else {
			groupLengths = append(groupLengths, count)
			count = 1
		}
	}
	groupLengths = append(groupLengths, count)

	dp := make([]int64, n+1)
	dp[0] = 1

	for _, l := range groupLengths {
		nextDp := make([]int64, n+1)
		currentWindowSum := int64(0)

		// Calculate nextDp[1]
		// nextDp[1] = dp[0] (if l >= 1)
		if l >= 1 && dp[0] > 0 {
			currentWindowSum = (currentWindowSum + dp[0]) % MOD
		}
		nextDp[1] = currentWindowSum

		// Calculate nextDp for i from 2 to n
		for i := 2; i <= n; i++ {
			// Add dp[i-1] to window
			currentWindowSum = (currentWindowSum + dp[i-1]) % MOD
			// Remove dp[i-l-1] from window if it's outside the window
			if i-l-1 >= 0 {
				currentWindowSum = (currentWindowSum - dp[i-l-1] + MOD) % MOD
			}
			nextDp[i] = currentWindowSum
		}
		dp = nextDp
	}

	var totalCnt int64
	for j := k; j <= n; j++ {
		totalCnt = (totalCnt + dp[j]) % MOD
	}

	return int(totalCnt)
}

func KthCharacter(k int) byte {
	// 3304
	// b  ab  abbc  abbcbccd
	word := []int{0}
	for len(word) < k {
		m := len(word)
		for i := 0; i < m; i++ {
			word = append(word, (word[i]+1)%26)
		}
	}
	return 'a' + byte(word[k-1])
}

func KthCharacterII(k int64, operations []int) byte {
	// 3307
	k--
	inc := 0
	for i := len(operations) - 1; i >= 0; i-- {
		if i >= 63 {
			// k is at most 10^18, which is less than 2^63.
			// 2^i for i>=63 will be larger than k, so k >= 2^i is false.
			// We can skip these iterations.
			continue
		}
		length := int64(1) << i
		if k >= length {
			k -= length
			inc = (inc + operations[i]) % 26
		}
	}
	return 'a' + byte(inc)
}
