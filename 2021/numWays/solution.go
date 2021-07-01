package main

func numWays(n int, relation [][]int, k int) int {
	var dfs func(k int, index int)

	visited := make(map[int]bool)

	count := 0
	dfs = func(round, index int) {
		if visited[index] {
			return
		}

		if k == round {
			if relation[index][1] == n-1 {
				count++
			}
			return
		}

		visited[index] = true

		for i := 0; i < len(relation); i++ {
			if i != index && relation[i][0] == relation[index][1] {
				dfs(round+1, i)
			}
		}

		visited[index] = false
	}

	for i := 0; i < len(relation); i++ {
		if relation[i][0] == 0 {
			dfs(0, i)
		}
	}
	return count
}

func main() {

}
