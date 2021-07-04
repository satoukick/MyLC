package main

const mod = 1e9 + 7

func countGoodNumbers(n int64) int {
	var odd int64 = n / 2
	even := n - odd
	return fastPowering(5, int(even)) * fastPowering(4, int(odd)) % mod
}

func fastPowering(n, pow int) int {
	if pow == 0 {
		return 1
	}

	ans, base := 1, n
	for pow > 0 {
		if pow&1 == 1 {
			ans = ans * base % mod
		}
		base = base * base % mod
		pow >>= 1
	}
	return ans
}
