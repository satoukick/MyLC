package main

func nearestExit(maze [][]byte, entrance []int) int {
	m := len(maze)
	n := len(maze[0])

	count := 0
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}

	q := [][]int{entrance}
	for len(q) > 0 {
		newQ := make([][]int, 0)

		for i := 0; i < len(q); i++ {
			item := q[i]
			if maze[item[0]][item[1]] == '+' {
				continue
			}

			if visited[item[0]][item[1]] {
				continue
			}

			visited[item[0]][item[1]] = true

			if item[0] == 0 || item[1] == 0 || item[0] == m-1 || item[1] == n-1 {
				if count != 0 {
					return count
				}
			}

			if item[0] > 0 {
				newItem := []int{item[0] - 1, item[1]}
				newQ = append(newQ, newItem)
			}
			if item[0] < m-1 {
				newItem := []int{item[0] + 1, item[1]}
				newQ = append(newQ, newItem)
			}
			if item[1] > 0 {
				newItem := []int{item[0], item[1] - 1}
				newQ = append(newQ, newItem)
			}
			if item[1] < n-1 {
				newItem := []int{item[0], item[1] + 1}
				newQ = append(newQ, newItem)
			}
		}

		q = newQ
		count++
	}
	return -1
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
