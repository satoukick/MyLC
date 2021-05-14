package main

func intToRoman(num int) string {
	value := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	key := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	cur := 0
	result := ""
	for num > 0 {
		if value[cur] > num {
			cur++
			continue
		}
		result += key[cur]
		num -= value[cur]
	}
	return result
}
