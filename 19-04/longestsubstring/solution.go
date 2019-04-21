package main

import "fmt"

func longestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}

	length := len(s)

	res := ""

	dp := make([][]bool, len(s))
	for i := range dp {
		dp[i] = make([]bool, len(s))
	}

	for i := length - 1; i >= 0; i-- {
		for j := i; j < length; j++ {
			dp[i][j] = s[i] == s[j] && (j-i <= 2 || dp[i+1][j-1])

			if dp[i][j] && (len(res) == 0 || j-i+1 > len(res)) {
				res = s[i : j+1]
			}
		}

	}
	return res
}

func main() {
	fmt.Println(longestPalindrome("babad"))
	fmt.Println(longestPalindrome("aaaa"))
}
