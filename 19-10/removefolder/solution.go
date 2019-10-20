package main

import (
	"fmt"
	"sort"
)

func removeSubfolders(folder []string) []string {
	sort.Slice(folder, func(i, j int) bool { return len(folder[i]) < len(folder[j]) })
	result := []string{}
	mapFolder := make(map[string]bool)
	for _, f := range folder {
		remove := false
		for i := 1; i < len(f); i++ {
			if f[i] == '/' && mapFolder[f[:i]] {
				remove = true
				break
			}
		}
		if !remove {
			mapFolder[f] = true
			result = append(result, f)
		}

	}
	return result
}

func main() {
	folders := []string{"/ah/al/am", "/ah/al"}
	fmt.Println(removeSubfolders(folders))
}
