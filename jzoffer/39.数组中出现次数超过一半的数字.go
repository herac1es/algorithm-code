package jzoffer

// 时间: O(n)
// 空间: O(1)
func majorityElement(nums []int) int {
	ret := nums[0]
	votes := 1
	for i := 1; i < len(nums); i++ {
		if votes == 0 {
			ret = nums[i]
			votes = 1
			continue
		}
		if nums[i] != ret {
			votes--
		} else {
			votes++
		}
	}
	return ret
}
