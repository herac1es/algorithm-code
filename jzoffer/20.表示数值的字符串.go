package jzoffer

var fsm = [...]map[byte]int{
	{' ': 0, 's': 1, 'd': 2, '.': 4}, // 0: 前缀空格
	{'d': 2, '.': 4},                 // 1:e之前的正负号
	{'d': 2, '.': 3, 'e': 5, ' ': 8}, // 2:小数点前的数字
	{'d': 3, 'e': 5, ' ': 8},         // 3:小数点后的数字
	{'d': 3},                         // 4:小数点
	{'s': 6, 'd': 7},                 // 5：幂符号
	{'d': 7},                         // 6:幂符号后的正负号
	{'d': 7, ' ': 8},                 // 7:幂符号后的数字
	{' ': 8},                         // 8:后缀空格
}

func isNumber(s string) bool {
	stat := 0
	var label byte
	for _, v := range []byte(s) {

		if v >= '0' && v <= '9' {
			label = 'd'
		} else if v == '+' || v == '-' {
			label = 's'
		} else if v == '.' {
			label = '.'
		} else if v == 'e' || v == 'E' {
			label = 'e'
		} else if v == ' ' {
			label = ' '
		} else {
			return false
		}
		if nextStat, ok := fsm[stat][label]; ok {
			stat = nextStat
		} else {
			return false
		}
	}
	return stat == 2 || stat == 3 || stat == 7 || stat == 8
}
