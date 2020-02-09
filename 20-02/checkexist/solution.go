package main

import (
	"fmt"
)

func checkIfExist(arr []int) bool {
	aMap := make(map[int]bool)
	for i := 0; i < len(arr); i++ {
		if aMap[arr[i]*2] {
			return true
		}
		if arr[i]%2 == 0 && aMap[arr[i]/2] {
			return true
		}

		aMap[arr[i]] = true
	}
	return false
}

func main() {
	arr := []int{-10, 12, -20, -8, 15}
	fmt.Println(checkIfExist(arr))
}
