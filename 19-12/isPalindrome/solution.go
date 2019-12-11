package main

import "strings"

import "fmt"

func isPalindrome(s string) bool {
	if s == "" {
		return true
	}
	s = strings.ToLower(s)

	start, end := 0, len(s)-1
	for start < end {
		if start > end {
			break
		}
		// fmt.Println(s[start])
		// fmt.Println(s[end])
		// fmt.Println("")
		if (s[start] < '0' || s[start] > '9') && (s[start] < 'a' || s[start] > 'z') {
			start++
			continue
		}
		if (s[end] < '0' || s[end] > '9') && (s[end] < 'a' || s[end] > 'z') {
			end--
			continue
		}

		if s[start] != s[end] {
			return false
		}
		start++
		end--
	}
	return true
}

func main() {
	s := "race a car"
	fmt.Println(isPalindrome(s))
}
