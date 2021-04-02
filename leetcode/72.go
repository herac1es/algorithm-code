package leetcode

func MinDistance(word1 string, word2 string) int {
	w1 := []rune(word1)
	w2 := []rune(word2)
	n1 := len(w1)
	n2 := len(w2)
	dp := make([][]int, n1+1)
	for i := range dp {
		dp1 := make([]int, n2+1)
		dp[i] = dp1
	}
	for j := range dp[0] {
		dp[0][j] = j
	}
	for i := range dp {
		dp[i][0] = i
	}
	dp[0][0] = 0
	for i := 1; i <= n1; i++ {
		for j := 1; j <= n2; j++ {
			if w1[i-1] == w2[j-1] {
				dp[i][j] = dp[i-1][j-1]
				continue
			}
			dp[i][j] = min(min(dp[i-1][j]+1, dp[i][j-1]+1), dp[i-1][j-1]+1)
		}
	}
	return dp[n1][n2]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
