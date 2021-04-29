package leetcode

// 由于不限交易次数， 只要当天较前一天是涨，则买
// 时间: O(n)
// 空间: O(1)
func maxProfitII(prices []int) int {
	profit := 0
	tmp := 0
	for i := 1; i < len(prices); i++ {
		tmp = prices[i] - prices[i-1]
		if tmp > 0 {
			profit += tmp
		}
	}
	return profit
}
