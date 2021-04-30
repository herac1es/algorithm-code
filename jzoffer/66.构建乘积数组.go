package jzoffer

// 时间: O(n)
// 空间: O(1)
func constructArr(a []int) []int {
	ret := make([]int, len(a))

	multi := 1
	// 计算左乘积：矩阵下三角
	for i := 0; i < len(a); i++ {
		ret[i] = multi
		multi *= a[i]
	}
	multi = 1
	// 计算右乘积与之相乘：矩阵上三角
	for i := len(a) - 1; i >= 0; i-- {
		ret[i] *= multi
		multi *= a[i]
	}
	return ret
}
