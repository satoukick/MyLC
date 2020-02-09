package main

func minSteps(s string, t string) int {
	sMap, tMap := make(map[byte]int), make(map[byte]int)
	for i := 0; i < len(s); i++ {
		sMap[s[i]]++
		tMap[t[i]]++
	}
	for k, v := range tMap {
		if sMap[k] > 0 {
			sMap[k] -= v
			if sMap[k] <= 0 {
				delete(sMap, k)
			}
		}
	}
	count := 0
	for _, v := range sMap {
		count += v
	}
	return count
}
