package jzoffer

import "math"

const MaxPoint = 6

// 时间: O(n2)
// 空间: O(n)
func dicesProbability(n int) []float64 {

	total := int64(math.Pow(float64(MaxPoint), float64(n))) // 排列总数
	cnt := [2][]int64{}
	// 数组元素: cnt[slot][n] 表示当前和为n出现的次数
	cnt[0] = make([]int64, MaxPoint*n+1)
	cnt[1] = make([]int64, MaxPoint*n+1)
	// 一个骰子的情况
	for i := 1; i <= MaxPoint; i++ {
		cnt[0][i] = 1
	}
	slot := 0
	for ni := 2; ni <= n; ni++ {

		for i := 1; i < ni; i++ {
			cnt[1-slot][i] = 0
		}
		for s := ni; s <= ni*MaxPoint; s++ {
			sum := int64(0)
			for i := 1; i <= MaxPoint && i <= s; i++ {
				sum += cnt[slot][s-i]
			}
			cnt[1-slot][s] = sum
		}
		slot = 1 - slot
	}
	ret := make([]float64, 0, MaxPoint*n-n)
	for i := n; i <= MaxPoint*n; i++ {
		ret = append(ret, float64(cnt[slot][i])/float64(total))
	}
	return ret
}
