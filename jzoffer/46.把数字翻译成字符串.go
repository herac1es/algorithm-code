package jzoffer

// 递归
// 时间:O(length)
// 空间:O(length)
func translateNum(num int) int {
	ret := 0
	translate(num, &ret)
	return ret
}

func translate(num int, ret *int) {
	if num == 0 {
		*ret++
		return
	}
	// 可以取一位翻译
	translate(num/10, ret)
	// 可以取两位翻译
	if num%100 <= 25 && num%100 >= 10 {
		translate(num/100, ret)
	}
}
