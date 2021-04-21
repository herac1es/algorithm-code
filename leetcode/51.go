package leetcode

// 51.N皇后
// 回溯
// 时间: O(n2)
// 空间：O(n)
var (
	leftUpLine  []int
	rightUpLine []int
	c           []int
)

func solveNQueens(n int) [][]string {

	res := make([][]string, 0)
	cols := make([]int, 0)
	leftUpLine = make([]int, 2*n-1)
	rightUpLine = make([]int, 2*n-1)
	c = make([]int, n)
	find(n, 0, &res, cols)
	return res
}

func find(n int, row int, res *[][]string, cols []int) {
	if row == n {
		item := makeRes(n, cols)
		*res = append(*res, item)
		return
	}

	for i := 0; i < n; i++ {
		if c[i] == 1 || leftUpLine[row-i+n-1] == 1 || rightUpLine[row+i] == 1 {
			continue
		}
		cols = append(cols, i)
		leftUpLine[row-i+n-1] = 1
		rightUpLine[row+i] = 1
		c[i] = 1
		find(n, row+1, res, cols)
		leftUpLine[row-i+n-1] = 0
		rightUpLine[row+i] = 0
		c[i] = 0

		cols = cols[:len(cols)-1]
	}

}

func makeRes(n int, cols []int) (ret []string) {
	if n != len(cols) {
		panic("invalid cols")
	}
	initial := make([]byte, len(cols))
	for i := range initial {
		initial[i] = '.'
	}
	for i := range cols {
		a := append([]byte{}, initial...)
		a[cols[i]] = 'Q'
		ret = append(ret, string(a))
	}
	return
}
