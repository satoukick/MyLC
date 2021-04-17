package main

func countPoints(points [][]int, queries [][]int) []int {
	result := []int{}
	for _, circle := range queries {
		count := 0

		for _, p := range points {
			x := circle[0] - p[0]
			y := circle[1] - p[1]

			if x*x+y*y <= circle[2]*circle[2] {
				count++
			}
		}

		result = append(result, count)
	}
	return result
}
