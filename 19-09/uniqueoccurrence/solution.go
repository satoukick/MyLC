package logic

import "sort"

func uniqueOccurrences(arr []int) bool {
	times := make(map[int]int)
	for _, v := range arr {
		times[v]++
	}
	
	t := []int{}
	for _, v := range times {
		t = append(t, v)
	}
	sort.Slice(t, func(i,j int) bool {return t[i] < t[j]})
	for i := 0; i < len(t)- 1; i++ {
		if t[i] == t[i+1] {
			return false
		}
	}
	return true
}