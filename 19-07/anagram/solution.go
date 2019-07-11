package main

func isAnagram(s string, t string) bool {
	ana := make(map[rune]int)
	for _, c := range t {
		ana[c]++
	}

	for _, a := range s {
		v, ok := ana[a]
		if !ok || v == 0 {
			return false
		}

		ana[a]--
	}

	for _, v := range ana {
		if v != 0 {
			return false
		}
	}
	return true
}
