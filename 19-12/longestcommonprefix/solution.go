package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}

	res := ""
	for i := 0; i < len(strs[0]); i++ {
		cur := strs[0][i]
		for j := 1; j < len(strs); j++ {
			if i == len(strs[j]) {
				return res
			}
			if strs[j][i] != cur {
				return res
			}
		}
		res += string(cur)
	}
	return res
}

func main() {
	strs := []string{"",""}
	fmt.Println(strs)
}