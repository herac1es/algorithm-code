package leetcode

import "math"

// 时间: O(n)
// 空间： O(1)
func myAtoi(s string) int {

	isNegtive := false

	idx := 0
	// 先找到数字开始部分，确认符号
	for idx < len(s) {

		if s[idx] == ' ' {
			idx++
		} else if s[idx] == '+' {
			idx++
			break
		} else if s[idx] == '-' {
			isNegtive = true
			idx++
			break
		} else if s[idx] >= '0' && s[idx] <= '9' {
			break
		} else {
			return 0
		}
	}

	ret := int64(0)

	// 进行转换
	for ; idx < len(s); idx++ {
		if s[idx] >= '0' && s[idx] <= '9' {
			ret = ret*10 + int64(s[idx]-'0')
			if !isNegtive && ret > math.MaxInt32 {
				return math.MaxInt32
			} else if isNegtive && ret > math.MaxInt32+1 {
				return math.MinInt32
			}
		} else {
			break
		}
	}
	if isNegtive {
		ret = -ret
	}
	return int(ret)
}
