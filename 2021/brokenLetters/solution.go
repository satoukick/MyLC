package main

import (
	"strings"
)

func canBeTypedWords(text string, brokenLetters string) int {
	strs := strings.Split(text, " ")

	broken := make([]bool, len(strs))

	for i := range brokenLetters {
		for j, str := range strs {
			if broken[j] {
				continue
			}

			if strings.Contains(str, string(brokenLetters[i])) {
				broken[j] = true
				// fmt.Println(string(brokenLetters[i]).)
			}
		}
	}

	brokenCount := 0
	for _, v := range broken {
		if v {
			brokenCount++
		}
	}
	return len(strs) - brokenCount
}
