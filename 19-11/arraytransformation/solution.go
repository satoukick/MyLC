package main

import "fmt"

func transformArray(arr []int) []int {
	cur := arr
	changed := true
	var new []int
	for changed {
		new = make([]int, len(arr))
		changed = false
		for i := 0; i < len(cur); i++ {
			if i == 0 || i == len(cur)-1 {
				new[i] = cur[i]
				continue
			}
			if cur[i] < cur[i-1] && cur[i] < cur[i+1] {
				changed = true
				new[i] = cur[i] + 1
			} else if cur[i] > cur[i-1] && cur[i] > cur[i+1] {
				changed = true
				new[i] = cur[i] - 1
			} else {
				new[i] = cur[i]
			}
		}
		cur = new
	}
	return new
}

func main() {
	arr := []int{1, 6, 3, 4, 3, 5}
	fmt.Println(transformArray(arr))
}
