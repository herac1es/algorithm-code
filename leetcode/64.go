package leetcode

// 64.最小路径和
// 动态规划
// 时间:O (mn)
// 空间：O(mn)
func minPathSum(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	dp[0][0] = grid[0][0]
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				dp[i][j] = grid[0][0]
			} else if j == 0 {
				dp[i][j] = dp[i-1][j] + grid[i][j]
			} else if i == 0 {
				dp[i][j] = dp[i][j-1] + grid[i][j]
			} else {
				dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
			}
		}
	}
	return dp[m-1][n-1]
}

// 递归回溯
// 时间：O(mn)
// 空间: O(mn)
func minPathSumII(grid [][]int) int {
	ret := make([][]int, len(grid))
	row := make([]int, len(grid[0]))
	for i := range row {
		row[i] = -1
	}
	for i := range ret {
		ret[i] = append(make([]int, 0, len(grid[0])), row...)
	}
	return minPathSumRecur(grid, ret, 0, 0)

}
func minPathSumRecur(grid [][]int, ret [][]int, i, j int) (res int) {
	if i == len(grid)-1 && j == len(grid[0])-1 {
		res = grid[i][j]
		ret[i][j] = res
		return
	}

	if ret[i][j] != -1 {
		res = ret[i][j]
		return
	}

	if i+1 < len(grid) && j+1 < len(grid[0]) {
		res = min(minPathSumRecur(grid, ret, i, j+1), minPathSumRecur(grid, ret, i+1, j)) + grid[i][j]
	} else if i+1 < len(grid) {
		res = minPathSumRecur(grid, ret, i+1, j) + grid[i][j]
	} else {
		res = minPathSumRecur(grid, ret, i, j+1) + grid[i][j]
	}
	ret[i][j] = res
	return
}
