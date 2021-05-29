package main

func countGoodSubstrings(s string) int {
	count := 0
	charMap := make(map[byte]int)

	for i := 0; i < len(s); i++ {
		charMap[s[i]]++
		if i < 2 {
			continue
		}

		if len(charMap) == 3 {
			count++
		}

		charMap[s[i-2]]--
		if charMap[s[i-2]] == 0 {
			delete(charMap, s[i-2])
		}

	}
	return count
}
