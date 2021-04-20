package main

func strStr(haystack string, needle string) int {
	if len(haystack) == 0 && len(needle) == 0 {
		return 0
	}

	if len(haystack) == 0 {
		return -1
	}

	if len(needle) == 0 {
		return 0
	}

	lenN := len(needle)

	for i := 0; i <= len(haystack)-lenN; i++ {
		cur := haystack[i : i+lenN]
		if cur == needle {
			return i
		}
	}
	return -1
}
