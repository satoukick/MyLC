package main

import (
	"strconv"
	"strings"
)

func sortSentence(s string) string {
	strs := strings.Split(s, " ")
	strSli := make([]string, len(strs))

	for _, s := range strs {
		num, _ := strconv.Atoi(string(s[len(s)-1]))
		strSli[num-1] = s[:len(s)-1]
	}

	return strings.Join(strSli, " ")
}
