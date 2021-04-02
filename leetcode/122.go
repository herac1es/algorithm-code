package leetcode

func MaxProfit122(prices []int) int {
	n := len(prices)
	dp1 := make([][2]int, n/2+1)
	dp := make([][][2]int, n)
	for i := 0; i < n; i++ {
		dp[i] = dp1
	}
	for i := 0; i < n; i++ {
		if i == 0 {
			for k := 1; k <= n/2; k++ {
				dp[i][k][0] = 0
				dp[i][k][1] = -prices[i]
			}
			continue
		}
		for k := 1; k <= n/2; k++ {
			dp[i][k][0] = max(dp[i-1][k][0], dp[i-1][k][1]+prices[i])
			dp[i][k][1] = max(dp[i-1][k][1], dp[i-1][k-1][0]-prices[i])
		}
	}
	return dp[n-1][n/2][0]
}
