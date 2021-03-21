package main

import (
	"fmt"
	"strconv"
)

func secondHighest(s string) int {
	second, first := -1, -1
	// nums := make(map[int]int)
	for _, bn := range s {
		n, err := strconv.Atoi(string(bn))
		if err != nil {
			continue
		}
		if first == n {
			continue
		}

		if first < n {
			second = first
			first = n
			// fmt.Println("1", second, n)
			continue
		}
		if second < n {
			// fmt.Println("2", second, n)
			second = n
		}

	}
	return second
}

func main() {
	s := secondHighest("abc1111")
	fmt.Println(s)
}
