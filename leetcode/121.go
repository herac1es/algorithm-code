package leetcode

// dp[i] : 在i对应的这一天卖出，所能取得的最大收益
// dp[i] = prices[i] - prices[0, i) 的最小值
// 时间: O(n)
// 空间: O(1)
func maxProfit(prices []int) int {
	min := prices[0]
	ret := 0
	for i := 1; i < len(prices); i++ {
		if prices[i]-min > ret {
			ret = prices[i] - min
		}
		if prices[i] < min {
			min = prices[i]
		}
	}
	return ret
}
