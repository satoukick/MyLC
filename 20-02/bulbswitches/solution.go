package main

func numTimesAllBlue(light []int) int {
	prev := 0
	onMap := make(map[int]int)
	count := 0
	for i := 0; i < len(light); i++ {
		onMap[light[i]] = 1
		newPrev := prev
		for j := prev + 1; j <= i+1; j++ {
			if onMap[j] != 1 {
				break
			}
			newPrev++
		}
		prev = newPrev
		if prev == i+1 {
			count++
		}
	}
	return count
}
