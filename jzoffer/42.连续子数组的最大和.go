package jzoffer

// 时间: O(n)
// 空间: O(n)
func maxSubArray(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	ret := dp[0]
	for i := 1; i < len(nums); i++ {
		dp[i] = max(nums[i], dp[i-1]+nums[i])
		if dp[i] > ret {
			ret = dp[i]
		}
	}
	return ret
}
