package main

func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	var ind int
	neeLength := len(needle)
	for ind+neeLength <= len(haystack) {
		for j := 0; j < len(needle); j++ {
			if needle[j] != haystack[ind+j] {
				break
			}
			if j == len(needle)-1 {
				return ind
			}
		}
		ind++
	}
	return -1
}
