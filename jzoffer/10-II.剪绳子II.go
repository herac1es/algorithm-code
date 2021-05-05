package jzoffer

// 贪心，尽量切成3段
func cuttingRopeII(n int) int {

	// n:2 -> 1x1 = 1
	// n:3 -> 1x2 = 2
	if n < 4 {
		return n - 1
	}
	res := 1
	for n > 4 {
		res = res * 3 % (1e9 + 7)
		n -= 3
	}
	return res * n % (1e9 + 7)
}
