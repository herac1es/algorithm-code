package leetcode

// 时间：O(2^n)
// 空间: O(n)
func generateParenthesis(n int) []string {
	ret := make([]string, 0)
	generateRecur(n, 0, 0, "", &ret)
	return ret
}

func generateRecur(n int, lc, rc int, item string, ret *[]string) {

	if lc == n && rc == n {
		*ret = append(*ret, item)
		return
	}

	// 只走合法的路径
	// 当前位置可选
	// 1.左括号
	// 2.当右括号数小于左括号时，才可选右括号
	if lc < n {
		generateRecur(n, lc+1, rc, item+"(", ret)
	}
	if rc < lc {
		generateRecur(n, lc, rc+1, item+")", ret)
	}
}
