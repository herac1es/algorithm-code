package leetcode

// 152.乘积最大子数组
// 记录以当前元素结尾的最大最小元素，当此元素为负数时，交换最大最小值
// 时间:O(n)
// 空间:O(1)
func maxProduct(nums []int) int {
	min, max := nums[0], nums[0]
	ret := max
	for i := 1; i < len(nums); i++ {

		if nums[i] < 0 {
			max, min = min, max
		}
		imax := max * nums[i]
		imin := min * nums[i]
		max = nums[i]
		min = nums[i]
		if imax > max {
			max = imax
		}
		if imin < min {
			min = imin
		}
		if max > ret {
			ret = max
		}
	}
	return ret
}
