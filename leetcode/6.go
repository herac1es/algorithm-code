package leetcode

// 6.Z字形变换

// 时间复杂度：O(K),K为字符串的长度
// 空间复杂度: O(K)
func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	queue := make([][]byte, numRows)
	idx := 0
	direct := up
	for _, v := range []byte(s) {
		queue[idx] = append(queue[idx], v)
		if direct == up && idx == numRows-1 {
			direct = down
			idx--
		} else if direct == down && idx == 0 {
			direct = up
			idx++
		} else if direct == up {
			idx++
		} else {
			idx--
		}
	}
	ret := make([]byte, 0)
	for i := range queue {
		ret = append(ret, queue[i]...)
	}
	return string(ret)
}

const (
	up   = 0
	down = 1
)
