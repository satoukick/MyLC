package main

import "fmt"

func findTheWinner(n int, k int) int {
	players := make([]int, n)
	for i := 0; i < n; i++ {
		players[i] = i + 1
	}

	start := 0
	for i := 1; i <= n-1; i++ {
		out := (start + k - 1) % (n - i + 1)
		fmt.Println(out, players[out])
		players = append(players[:out], players[out+1:]...)
		start = out
	}
	return players[0]
}
