package jzoffer

const mod int = 1e9 + 7

// (a+b)%c = (a%c+b%c)%c
// 时间: O(n)
// 空间: O(1)
func fib(n int) int {
	if n < 2 {
		return n
	}
	a, b := 0, 1
	c := 0
	for i := 2; i <= n; i++ {
		c = a
		a = b
		b = (c%mod + b%mod) % mod
	}
	return b
}
