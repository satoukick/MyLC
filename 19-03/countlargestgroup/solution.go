package main

import (
	"fmt"
	"sort"
)

func countLargestGroup(n int) int {
	tmpN := n
	var space int
	for tmpN != 0 {
		space += 9
		tmpN = tmpN / 10
	}
	groups := make([]int, space+1)
	for i := 1; i <= n; i++ {
		var sum int
		num := i
		for num != 0 {
			sum += num % 10
			num = num / 10
		}
		groups[sum]++
	}
	sort.Slice(groups, func(i, j int) bool { return groups[i] > groups[j] })
	fmt.Println(groups)
	largest := groups[0]
	count := 1
	for i := 1; i <= space; i++ {
		if groups[i] != largest {
			break
		}
		count++
	}
	return count
}
