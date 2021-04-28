package jzoffer

// 二进制和跟异或运算规律相同，进位和跟与运算并且左移一位的规律相同
// 当进位和为0时，二进制和为所求结果
// 时间: O(1)
// 空间: O(1)
func add(a int, b int) int {
	sum := a ^ b
	surplus := (a & b) << 1
	for surplus != 0 {
		preSum := sum
		sum = sum ^ surplus
		surplus = (preSum & surplus) << 1
	}
	return sum
}
