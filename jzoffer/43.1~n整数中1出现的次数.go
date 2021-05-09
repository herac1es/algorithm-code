package jzoffer

// 时间: O(logn)
// 空间: O(1)
func countDigitOne(n int) int {
	ret := 0

	cur := n % 10
	digit := 1
	high := n / 10
	low := 0

	for cur > 0 || high > 0 {
		if cur == 0 {
			ret += high * digit
		} else if cur == 1 {
			ret += high*digit + low + 1
		} else {
			ret += (high + 1) * digit
		}
		low += cur * digit
		cur = high % 10
		high /= 10
		digit *= 10
	}
	return ret
}
