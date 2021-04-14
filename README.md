# algorithm-code
my algorithm code repo in Go

## Leetcode Solution

- [x] 1143.最长公共子序列

<https://leetcode-cn.com/problems/longest-common-subsequence/>

定义 dp 数组: dp\[i][j] 为 text1[0: i-1] 和 text2[0: j-1] 的最长公共子序列
base: 当 i = 0 或者 j = 0 时，区间里没有字符，所以 dp\[i][j] = 0
当 text1\[i-1] = text2\[j-1] 时, dp\[i]\[j] = dp\[i-i][j-1]+ 1
否则等于 dp\[i-1]\[j] 和 dp\[i][j-1] 中较大的一个

时间：O(n^2)
空间: O(n^2)

```go
func longestCommonSubsequence(text1 string, text2 string) int {
	len1 := len(text1)
	len2 := len(text2)

	dp := make([][]int, len1+1)
	for i := range dp {
		dp[i] = make([]int, len2+1)
	}

	for i := 1; i <= len1; i++ {
		for j := 1; j <= len2; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[len1][len2]
}
```

- [ ]  