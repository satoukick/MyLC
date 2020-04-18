package main

func numWays(n int, relation [][]int, k int) int {
	ways := []int{}
	for i := 0; i < len(relation); i++ {
		if relation[i][0] == 0 {
			ways = append(ways, relation[i][1])
		}
	}
	for i := 1; i < k; i++ {
		newWays := []int{}
		for j := 0; j < len(ways); j++ {
			for a := 0; a < len(relation); a++ {
				if relation[a][0] == ways[j] {
					newWays = append(newWays, relation[a][1])
				}
			}
		}
		ways = newWays
	}
	count := 0
	for i := 0; i < len(ways); i++ {
		if ways[i] == n-1 {
			count++
		}
	}
	return count
}
