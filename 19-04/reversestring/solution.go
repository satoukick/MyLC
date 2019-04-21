package main

func reverseString(s []byte) {
	if len(s) <= 1 {
		return
	}

	start, end := 0, len(s)-1
	for start < end {
		tmp := s[start]
		s[start] = s[end]
		s[end] = tmp
		start++
		end--
	}
}
