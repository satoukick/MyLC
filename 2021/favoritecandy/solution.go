package main

// Time Complexity: O(n)
// Space Complexity: O(n)
func canEat(candiesCount []int, queries [][]int) []bool {
	pre := make([]int, len(candiesCount))
	pre[0] = candiesCount[0]
	for i := 1; i < len(candiesCount); i++ {
		pre[i] = pre[i-1] + candiesCount[i]
	}

	result := make([]bool, len(queries))

	for i, q := range queries {
		leastConsume := q[1] + 1
		maxConsume := min(q[2]*(q[1]+1), pre[len(candiesCount)-1])

		// how many candies until favorite type?
		var leastNeed int
		if q[0] != 0 {
			leastNeed = pre[q[0]-1]
		}
		maxNeed := pre[q[0]]

		// fmt.Println(leastConsume, maxConsume, leastNeed, maxNeed)

		if leastNeed > maxConsume {
			result[i] = false
			continue
		}
		if maxNeed < leastConsume {
			result[i] = false
			continue
		}
		result[i] = true
	}
	return result
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
