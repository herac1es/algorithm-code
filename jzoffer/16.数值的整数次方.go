package jzoffer

// 时间：O(logn)
// 空间: O(logn)
func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	t := n
	if t < 0 {
		t = -t
	}
	r := powRecur(x, t)
	if n < 0 {
		return 1 / r
	}
	return r
}

func powRecur(x float64, n int) float64 {
	if n == 1 {
		return x
	}
	half := powRecur(x, n/2)
	if n%2 == 1 {
		return half * half * x
	}
	return half * half
}
