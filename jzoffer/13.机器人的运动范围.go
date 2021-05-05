package jzoffer

// 因为机器人可以上下左右移动，此题是求从[0][0]位置开始，机器人能到达的相连的格子数量，类似于岛屿面积
// 时间: O(mxn)
// 空间: O(mxn)
func movingCount(m int, n int, k int) int {
	depth := 0
	visited := make(map[[2]int]struct{})
	movingRecur(m, n, k, 0, 0, visited, &depth)
	return depth
}

func movingRecur(m, n, k, i, j int, visited map[[2]int]struct{}, depth *int) {
	if i < 0 || i >= m || j < 0 || j >= n {
		return
	}
	if _, ok := visited[[2]int{i, j}]; ok || sum10(i)+sum10(j) > k {
		return
	}
	visited[[2]int{i, j}] = struct{}{}
	*depth++
	movingRecur(m, n, k, i+1, j, visited, depth)
	movingRecur(m, n, k, i-1, j, visited, depth)
	movingRecur(m, n, k, i, j+1, visited, depth)
	movingRecur(m, n, k, i, j-1, visited, depth)
}

func sum10(a int) int {
	ret := 0
	for a > 0 {
		ret += a % 10
		a /= 10
	}
	return ret
}
