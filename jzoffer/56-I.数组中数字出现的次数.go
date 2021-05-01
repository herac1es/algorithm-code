package jzoffer

// 时间: O(n)
// 空间: O(1)
func singleNumbers(nums []int) []int {
	xor := 0
	for i := 0; i < len(nums); i++ {
		xor ^= nums[i]
	}
	// 找到第一位为1的二进制位
	bit := 0
	for xor > 0 {
		if xor&1 == 1 {
			break
		}
		bit++
		xor >>= 1
	}
	// 根据第bit位是否为1将原数组分成两组进行异或，所得即是两个只出现一次的数
	val := 1 << bit
	a := 0
	b := 0
	for i := 0; i < len(nums); i++ {
		if nums[i]&val > 0 {
			a ^= nums[i]
		} else {
			b ^= nums[i]
		}
	}
	return []int{a, b}
}
