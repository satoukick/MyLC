package main

import (
	"fmt"
	"strconv"
)

func letterCombinations(digits string) []string {
	nums := map[int]string{
		2: "abc",
		3: "def",
		4: "ghi",
		5: "jkl",
		6: "mno",
		7: "pqrs",
		8: "tuv",
		9: "wxyz",
	}

	result := []string{}
	phone(digits, nums, &result, "")
	return result
}

func phone(digits string, nums map[int]string, result *[]string, cur string) {
	if len(digits) <= 0 {
		*result = append(*result, cur)
		return
	}

	n, _ := strconv.Atoi(string(digits[0]))
	for _, v := range nums[n] {
		phone(digits[1:], nums, result, cur+string(v))
	}
}

func main() {
	fmt.Println(letterCombinations("23"))
}
