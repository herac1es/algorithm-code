package jzoffer

// 时间: O(mxn)
// 空间: O(1)
func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return nil
	}
	m := len(matrix)
	n := len(matrix[0])
	ret := make([]int, 0, m*n)
	top, bottom, left, right := 0, m-1, 0, n-1
	for {
		for i := left; i <= right; i++ {
			ret = append(ret, matrix[top][i])
		}
		top++
		if top > bottom {
			break
		}
		for i := top; i <= bottom; i++ {
			ret = append(ret, matrix[i][right])
		}
		right--
		if right < left {
			break
		}
		for i := right; i >= left; i-- {
			ret = append(ret, matrix[bottom][i])
		}
		bottom--
		if bottom < top {
			break
		}
		for i := bottom; i >= top; i-- {
			ret = append(ret, matrix[i][left])
		}
		left++
		if left > right {
			break
		}
	}
	return ret
}
