package main

import "strconv"

func replaceDigits(s string) string {
	result := ""
	for i := 0; i < len(s); i += 2 {
		result += string(s[i])

		if i+1 < len(s) {
			add, _ := strconv.Atoi(string(s[i+1]))
			r := int(rune(s[i])) + add
			result += string(r)
		}

	}
	return result
}
