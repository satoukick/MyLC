package main

func generateTheString(n int) string {
	if n == 0 {
		return ""
	}
	s := ""
	if n%2 == 0 {
		s += "a"
		for i := 1; i <= n-1; i++ {
			s += "b"
		}
	} else {
		for i := 1; i <= n; i++ {
			s += "b"
		}
	}
	return s
}
