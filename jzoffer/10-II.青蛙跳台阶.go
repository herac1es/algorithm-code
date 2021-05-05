package jzoffer

func numWays(n int) int {
	if n < 2 {
		return 1
	}
	p1, p2 := 1, 1
	c := 0
	for i := 2; i <= n; i++ {
		c = p1
		p1 = p2
		p2 = (c%mod + p1%mod) % mod
	}
	return p2
}
