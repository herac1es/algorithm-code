package leetcode

func MaxProfit188(k int, prices []int) int {
	n := len(prices)
	if n == 0 || k == 0 {
		return 0
	}
	if k > n/2 {
		k = n / 2
	}
	dp1 := make([][2]int, k+1)
	dp := make([][][2]int, n)
	for i := 0; i < n; i++ {
		dp[i] = dp1
	}
	for i := 0; i < n; i++ {
		if i == 0 {
			for kk := 1; kk <= k; kk++ {
				dp[i][kk][0] = 0
				dp[i][kk][1] = -prices[i]
			}
			continue
		}
		for kk := 1; kk <= k; kk++ {
			dp[i][kk][0] = max(dp[i-1][kk][0], dp[i-1][kk][1]+prices[i])
			dp[i][kk][1] = max(dp[i-1][kk][1], dp[i-1][kk-1][0]-prices[i])
		}
	}
	return dp[n-1][k][0]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
