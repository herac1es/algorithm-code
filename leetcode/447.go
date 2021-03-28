package leetcode

// 给定平面上 n 对 互不相同 的点 points ，其中 points[i] = [xi, yi] 。回旋镖 是由点 (i, j, k) 表示的元组 ，其中
// i 和 j 之间的距离和 i 和 k 之间的距离相等（需要考虑元组的顺序）。
//
// 返回平面上所有回旋镖的数量。
//
//
// 示例 1：
//
//
// 输入：points = [[0,0],[1,0],[2,0]]
// 输出：2
// 解释：两个回旋镖为 [[1,0],[0,0],[2,0]] 和 [[1,0],[2,0],[0,0]]
//
//
// 示例 2：
//
//
// 输入：points = [[1,1],[2,2],[3,3]]
// 输出：2
//
//
// 示例 3：
//
//
// 输入：points = [[1,1]]
// 输出：0
//
//
//
//
// 提示：
//
//
// n == points.length
// 1 <= n <= 500
// points[i].length == 2
// -104 <= xi, yi <= 104
// 所有点都 互不相同
//
// Related Topics 哈希表 数学
// 👍 129 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func numberOfBoomerangs(points [][]int) int {
	disCount := make(map[int]int)
	ret := 0
	for i := 0; i < len(points); i++ {
		for j := 0; j < len(points); j++ {
			if i == j {
				continue
			}
			disCount[pow(points[j][1]-points[i][1], 2)+pow(points[j][0]-points[i][0], 2)]++
		}
		for k, v := range disCount {
			if v > 1 {
				ret += calC(v, 2) * 2
			}
			delete(disCount, k)
		}
	}
	return ret
}
func calC(n, m int) int {
	if n == m {
		return 1
	}
	if n == 1 {
		return 1
	}
	if m == 1 {
		return n
	}
	return calC(n-1, m-1) + calC(n-1, m)
}

// leetcode submit region end(Prohibit modification and deletion)
