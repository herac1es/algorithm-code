package leetcode

func MaxProfit123(prices []int) int {
	n := len(prices)
	if n < 2 {
		return 0
	}
	dp := make([][3][2]int, n)
	for i := 0; i < n; i++ {
		if i == 0 {
			dp[i][1][0], dp[i][2][0] = 0, 0
			dp[i][1][1], dp[i][2][1] = -prices[i], -prices[i]
			continue
		}
		for k := 1; k <= 2; k++ {
			dp[i][k][0] = max(dp[i-1][k][0], dp[i-1][k][1]+prices[i])
			dp[i][k][1] = max(dp[i-1][k][1], dp[i-1][k-1][0]-prices[i])
		}
	}
	return dp[n-1][2][0]
}
