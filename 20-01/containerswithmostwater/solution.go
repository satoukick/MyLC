package main

func maxArea(height []int) int {
	left, right := 0, len(height)-1
	area := -1
	for left < right {
		area = max(area, (right-left)*min(height[left], height[right]))

		if height[left] > height[right] {
			right--
		} else {
			left++
		}
	}
	return area
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
