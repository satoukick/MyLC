package main

func isAnagram(s string, t string) bool {
	strs := make(map[rune]int)
	for _, v := range s {
		strs[v]++
	}
	for _, v := range t {
		if _, ok := strs[v]; !ok {
			return false
		}
		strs[v]--
		if strs[v] == 0 {
			delete(strs, v)
		}
	}

	return len(strs) == 0
}
