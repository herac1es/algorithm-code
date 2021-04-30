package jzoffer

// 利用 逻辑运算的短路逻辑 做递归终止控制
// 时间: O(n)
// 空间：
var res int

func sumNums(n int) int {
	res = 0
	var _ bool = (n > 1) && sumNums(n-1) > 1
	res += n
	return res
}
