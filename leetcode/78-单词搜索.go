package leetcode

func exist(board [][]byte, word string) bool {
	var ret bool
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			existRecur(board, word, 0, i, j, &ret)
			if ret {
				return true
			}
		}
	}
	return false
}

func existRecur(board [][]byte, word string, idx, i, j int, ret *bool) {
	if idx >= len(word) {
		*ret = true
		return
	}
	if !valid(board, i, j) || board[i][j] == '-' {
		return
	}
	if board[i][j] == word[idx] {
		tmp := board[i][j]
		board[i][j] = '-'
		existRecur(board, word, idx+1, i+1, j, ret)
		existRecur(board, word, idx+1, i-1, j, ret)
		existRecur(board, word, idx+1, i, j+1, ret)
		existRecur(board, word, idx+1, i, j-1, ret)
		board[i][j] = tmp
	}
}

func valid(board [][]byte, i, j int) bool {
	m := len(board)
	n := 0
	if m > 0 {
		n = len(board[0])
	}
	return i >= 0 && i < m && j >= 0 && j < n
}
