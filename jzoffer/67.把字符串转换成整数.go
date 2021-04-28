package jzoffer

import "math"

// 时间: O(n)
// 空间: O(1)
func strToInt(str string) int {
	isNegtive := false
	num := int64(0)
	idx := 0
	// 判断正负 & 首个非空字符合法校验
	for idx < len(str) {
		if str[idx] == ' ' {
			idx++
		} else if str[idx] == '+' {
			idx++
			break
		} else if str[idx] == '-' {
			idx++
			isNegtive = true
			break
		} else if str[idx] >= '0' && str[idx] <= '9' {
			break
		} else {
			return 0
		}
	}
	// 遇到非法字符终止转换
	// int32 范围越界检查
	for i := idx; i < len(str); i++ {
		if str[idx] < '0' || str[idx] > '9' {
			break
		}
		num = num*10 + int64(str[idx]-'0')
		if !isNegtive && num > math.MaxInt32 {
			num = math.MaxInt32
			break
		} else if isNegtive && -num < math.MinInt32 {
			num = -math.MinInt32
			break
		}
		idx++
	}
	if isNegtive {
		num = -num
	}
	return int(num)
}
