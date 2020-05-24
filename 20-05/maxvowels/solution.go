func maxVowels(s string, k int) int {
	vowels := map[rune]bool{
		'a': true,
		'e': true,
		'i': true,
		'o': true,
		'u': true,
	}
	var max, count int
	cur := s[:k]
	for _, r := range cur {
		if vowels[r] {
			count++
		}
	}
	if count == k {
		return count
	}
	max = count
	for i := k; i <= len(s)-1; i++ {
		newCount := count
		if vowels[rune(s[i-k])] {
			newCount--
		}
		if vowels[rune(s[i])] {
			newCount++
		}
		if newCount == k {
			return k
		}
		if newCount > max {
			max = newCount
		}
		count = newCount
	}
	return max
}