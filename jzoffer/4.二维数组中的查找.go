package jzoffer

// 时间: O(m+n)
// 时间: O(1)
func findNumberIn2DArray(matrix [][]int, target int) bool {
	m := len(matrix)
	if m == 0 || len(matrix[0]) == 0 {
		return false
	}
	n := len(matrix[0])
	i, j := m-1, 0
	for i >= 0 && j < n {
		if target > matrix[i][j] {
			j++
		} else if target < matrix[i][j] {
			i--
		} else {
			return true
		}
	}
	return false
}
