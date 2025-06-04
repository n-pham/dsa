package main

//lint:file-ignore U1000 Ignore all unused code, it's generated

func answerString(word string, numFriends int) string {
	// 3403 1208ms
	stringLen := len(word) - numFriends + 1
	maxSubstr := ""
	for length := 1; length <= stringLen; length++ {
		for start := 0; start <= len(word)-length; start++ {
			curr := word[start : start+length]
			if curr > maxSubstr {
				maxSubstr = curr
			}
		}
	}
	return maxSubstr
}
