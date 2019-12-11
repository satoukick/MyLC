package main

import "fmt"

//Implement strStr() from explore

func strStr(haystack string, needle string) int {
	if len(haystack) < len(needle) {
		return -1
	}
	for i := 0; i <= len(haystack)-len(needle); i++ {
		if haystack[i:i+len(needle)] == needle {
			return i
		}
	}
	return -1
}

func main() {
	haystack := "aaaa"
	needle := "bba"
	fmt.Println(strStr(haystack, needle))
}
