package main

// 387. First Unique Character in a String

func firstUniqChar(s string) int {
	smap := make(map[rune]int)
	for _, v := range s {
		smap[v]++
	}
	for i, v := range s {
		if smap[v] == 1 {
			return i
		}
	}
	return -1
}
