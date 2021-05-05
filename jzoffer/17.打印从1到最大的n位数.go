package jzoffer

import "math"

// 时间: O(10^n)
// 空间: O(1)
func printNumbers(n int) []int {
	maxNum := int(math.Pow(10, float64(n)) - 1)
	ret := make([]int, maxNum)
	for i := 0; i < maxNum; i++ {
		ret[i] = i + 1
	}
	return ret
}
