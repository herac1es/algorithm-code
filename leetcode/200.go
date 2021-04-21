package leetcode

// 200.岛屿数量
// 递归填充
// 时间：O(mn)
// 空间: O(mn)
func numIslands(grid [][]byte) int {
	m := len(grid)
	if m == 0 {
		return 0
	}
	n := len(grid[0])
	ret := 0

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] != '1' {
				continue
			}

			ret++
			fillIslands(grid, i, j, m, n)
		}
	}
	return ret
}

func fillIslands(grid [][]byte, i, j, m, n int) {
	if !checkBorder(m, n, i, j) || grid[i][j] != '1' {
		return
	}
	grid[i][j] += 1
	fillIslands(grid, i+1, j, m, n)
	fillIslands(grid, i-1, j, m, n)
	fillIslands(grid, i, j+1, m, n)
	fillIslands(grid, i, j-1, m, n)
}

func checkBorder(m, n, i, j int) bool {
	return i >= 0 && i < m && j >= 0 && j < n
}
