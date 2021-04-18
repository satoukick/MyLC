package main

func checkIfPangram(sentence string) bool {
	if len(sentence) < 26 {
		return false
	}

	alphabetics := make(map[rune]bool)

	for _, v := range sentence {
		if _, ok := alphabetics[v]; !ok {
			alphabetics[v] = true
			if len(alphabetics) == 26 {
				return true
			}
		}
	}
	return false
}
