package jzoffer

// 时间: O(n)
// 空间: O(1)
func majorityElement(nums []int) int {
	ret := 0
	votes := 0
	for i := 0; i < len(nums); i++ {
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
