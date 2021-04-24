package leetcode

// 343.整数拆分
// 时间:O(n2)
// 空间:O(n)
func integerBreak(n int) int {
	dp := make([]int, n+1)
	dp[1] = 1
	for i := 2; i <= n; i++ {
		for j := 1; j <= i-1; j++ {
			dp[i] = max(max(dp[i], j*(i-j)), j*dp[i-j])
		}
	}
	return dp[n]
}
