package leetcode

func FindNumberOfLIS(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	res := make([]int, n+1)
	for i := range dp {
		dp[i] = 1
		res[i]++
	}
	max := dp[0]
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				val := dp[j] + 1
				if val > max {
					max = val
					dp[i] = val
					res[val] = res[val-1] + 1
				}
			}
		}
	}
	return res[n]
}
