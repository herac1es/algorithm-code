package jzoffer

// 时间:O(n)
// 空间:O(1)
func singleNumber(nums []int) int {
	bitMap := [32]int{}
	for i := 0; i < len(nums); i++ {

		for j := 0; j < 32; j++ {
			bitMap[j] += nums[i] & (1 << j)
		}
	}
	ret := 0
	// 不能被3整除的位，所求值的对应位为1
	for i := 0; i < 32; i++ {
		if bitMap[i]%3 != 0 {
			ret |= 1 << i
		}
	}
	return ret
}
