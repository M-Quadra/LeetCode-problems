package main

func validSubstringCount(word1 string, word2 string) int64 {
	charToCnt, find := [26]int32{}, 0
	for _, char := range word2 {
		if char -= 'a'; charToCnt[char] == 0 {
			find++
		}
		charToCnt[char]++
	}

	res, i := int64(0), 0
	for _, char := range word1 {
		for ; find > 0 && i < len(word1); i++ {
			char := word1[i] - 'a'
			if charToCnt[char] == 1 {
				find--
			}
			charToCnt[char]--
		}
		if find > 0 {
			break
		}
		res += int64(len(word1) - i + 1)

		if char -= 'a'; charToCnt[char] == 0 {
			find++
		}
		charToCnt[char]++
	}
	return res
}
