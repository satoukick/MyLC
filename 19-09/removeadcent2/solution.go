package main

import "fmt"

func removeDuplicates(s string, k int) string {
	if len(s) <= 1 {
		return s
	}
	start := 0
	curStr := s
	for {
		new := ""
		for i := 0; i < len(curStr)-1; i++ {
			if curStr[i] == curStr[i+1] {
				if i-start+1 == k {
					start = i + 1
				}
				continue
			}
			if i-start+1 < k {
				new += curStr[start : i+1]
			}
			start = i + 1
		}

		if start == len(curStr)-1 {
			new += string(curStr[start])
		} else if len(curStr)-start < k {
			new += curStr[start:len(curStr)]
		}

		if new == curStr {
			break
		}
		curStr = new
		start = 0
	}
	return curStr
}

func main() {
	st := "deeedbbcccbdaa"
	fmt.Println(removeDuplicates(st, 3))
}
