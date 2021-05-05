package jzoffer

// 时间: O(1)
// 空间: O(1)
func hammingWeight(num uint32) int {
	ret := 0
	for num > 0 {
		ret += int(num & 1)
		num >>= 1
	}
	return ret
}
