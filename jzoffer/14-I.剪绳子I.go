package jzoffer

// 动态规划：
// 对于长度为n的绳子，最佳切分dp[n] = dp[1]*(n-1),...,dp[n-1]*1中的最大值
// 注意：每根绳子至少被切分成两根，所以对于小于n的j来说，至少存在一种切分方式：j*(i-j)
// 时间: O(n2)
// 空间: O(n)
func cuttingRope(n int) int {
	dp := make([]int, n+1)
	dp[2] = 1
	for i := 3; i <= n; i++ {
		for j := 1; j < i; j++ {
			dp[i] = max(max(dp[i], j*(i-j)), dp[j]*(i-j))
		}
	}
	return dp[n]
}
