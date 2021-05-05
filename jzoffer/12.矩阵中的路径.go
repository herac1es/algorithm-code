package jzoffer

// 时间: O(3^Kxmxn)
// 空尽: O(K) ,k: word长度
func exist(board [][]byte, word string) bool {
	m := len(board)
	n := len(board[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if existRecur(board, word, i, j, 0) {
				return true
			}
		}
	}
	return false
}

// idx: 需要对比的word位置
// 当前[i][j]位置被访问后，临时将board[i][j]置成一个无效字符表示已经访问，回溯后再还原
func existRecur(board [][]byte, word string, i, j, idx int) bool {
	if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) {
		return false
	}
	if board[i][j] != word[idx] {
		return false
	}
	if idx == len(word)-1 {
		return true
	}
	b := board[i][j]
	board[i][j] = 0
	ret := existRecur(board, word, i+1, j, idx+1) || existRecur(board, word, i-1, j, idx+1) || existRecur(board, word, i, j+1, idx+1) || existRecur(board, word, i, j-1, idx+1)
	board[i][j] = b
	return ret
}
