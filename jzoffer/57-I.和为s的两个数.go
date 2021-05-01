package jzoffer

// 时间: O(n)
// 空间: O(1)
func twoSum(nums []int, target int) []int {
	ret := make([]int, 0, 2)
	i, j := 0, len(nums)-1
	for i < j {
		sum := nums[i] + nums[j]
		if sum == target {
			ret = append(ret, nums[i], nums[j])
			return ret
		} else if sum > target {
			j--
		} else {
			i++
		}
	}
	return ret
}
