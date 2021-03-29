package leetcode

// 直线上最多的点
//leetcode submit region begin(Prohibit modification and deletion)
func maxPoints(points [][]int) int {
	slopeCount := make(map[float64]int) // 记录有斜率的点个数
	infCount := 0                       // 记录垂直于x轴（无斜率的点个数）
	ret := 0
	for i := 0; i < len(points); i++ {
		for j := 0; j < len(points); j++ {
			if i == j || points[i][0] == points[j][0] {
				infCount++
				continue
			}
			slopeCount[float64(points[j][1]-points[i][1])/float64(points[j][0]-points[i][0])]++
		}
		if infCount > ret {
			ret = infCount
		}
		infCount = 0
		for k, v := range slopeCount {
			if v+1 > ret {
				ret = v + 1
			}
			delete(slopeCount, k)
		}
	}
	return ret
}

//leetcode submit region end(Prohibit modification and deletion)
