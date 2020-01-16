package main

import (
	"fmt"
	"math"
)

// Input: "eceba"
// Output: 3
// Explanation: t is "ece" which its length is 3.
// Example 2:

// Input: "ccaabbb"
// Output: 5
// Explanation: t is "aabbb" which its length is 5.

func lengthOfLongestSubstringTwoDistinct(s string) int {
	if len(s) <= 2 {
		return len(s)
	}
	curStr := s[:1]
	n := make(map[byte]int) // save count per byte
	for _, r := range curStr {
		n[byte(r)]++
	}
	maxN := math.MinInt32
	for i := 1; i < len(s); i++ {
		curStr += string(s[i])
		n[s[i]]++

		if len(n) > 2 {
			n[curStr[0]]--
			if num, ok := n[curStr[0]]; ok && num == 0 {
				delete(n, curStr[0])
			}
			curStr = curStr[1:]
		}
		// fmt.Println(i, curStr)
		maxN = max(maxN, len(curStr))
	}
	return maxN
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	s := "bacc"
	fmt.Println(lengthOfLongestSubstringTwoDistinct(s))
}
