package main

import (
	"fmt"
	"strings"
)

func printVertically(s string) []string {
	result := []string{}
	if len(s) == 0 {
		return result
	}
	strs := strings.Split(s, " ")
	length := 0
	for i := 0; i < len(strs); i++ {
		length = max(length, len(strs[i]))
	}
	result = make([]string, length)
	for i := 1; i <= length; i++ {
		for j := 0; j < len(strs); j++ {
			if len(strs[j]) < i {
				result[i-1] += " "
			} else {
				result[i-1] += string(strs[j][i-1])
			}
		}
		result[i-1] = strings.TrimRight(result[i-1], " ")
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	s := printVertically("TO BE OR NOT TO BE")
	fmt.Println(s)
}
