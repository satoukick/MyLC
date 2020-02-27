package main

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	index := -1

outer:
	for {
		if index+1 == len(strs[0]) {
			break
		}
		c := strs[0][index+1]
		for i := 1; i < len(strs); i++ {
			if index+1 == len(strs[i]) || c != strs[i][index+1] {
				break outer
			}
		}
		index++
	}

	if index < 0 {
		return ""
	}
	return strs[0][:index+1]
}
