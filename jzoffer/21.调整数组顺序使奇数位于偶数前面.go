package jzoffer

// 时间: O(n)
// 空间: O(1)
func exchange(nums []int) []int {
	i, j := 0, len(nums)-1
	for {
		for i < len(nums) && nums[i]%2 != 0 {
			i++
		}
		for j >= 0 && nums[j]%2 == 0 {
			j--
		}
		if i >= j {
			break
		}
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
	return nums
}
