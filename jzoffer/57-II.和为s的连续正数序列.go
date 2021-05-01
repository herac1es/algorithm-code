package jzoffer

// 时间: O(n)
// 空间： O(n)
func findContinuousSequence(target int) [][]int {
	ret := make([][]int, 0)
	nums := make([]int, 0, target-1)
	for i := 1; i <= target-1; i++ {
		nums = append(nums, i)
	}
	l, r := 0, -1
	sum := 0
	for r < target {
		if sum == target {
			ret = append(ret, nums[l:r+1])
			sum -= nums[l]
			l++
		} else if sum > target {
			sum -= nums[l]
			l++
		} else {
			r++
			if r < len(nums) {
				sum += nums[r]
			}
		}
	}
	return ret
}
