package main

func reverseWords(s []byte) {
	reverse(s, 0, len(s)-1)
	i := 0
	for j := 0; j < len(s); j++ {
		if s[j] == ' ' {
			reverse(s, i, j-1)
			i = j + 1
		}
	}
	reverse(s, i, len(s)-1)
}

func reverse(s []byte, i, j int) {
	for i < j {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}
}
