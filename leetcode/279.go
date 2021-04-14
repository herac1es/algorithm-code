package leetcode

//给定正整数 n，找到若干个完全平方数（比如 1, 4, 9, 16, ...）使得它们的和等于 n。你需要让组成和的完全平方数的个数最少。
//
// 给你一个整数 n ，返回和为 n 的完全平方数的 最少数量 。
//
// 完全平方数 是一个整数，其值等于另一个整数的平方；换句话说，其值等于一个整数自乘的积。例如，1、4、9 和 16 都是完全平方数，而 3 和 11 不是。
//
//
//
//
// 示例 1：
//
//
//输入：n = 12
//输出：3
//解释：12 = 4 + 4 + 4
//
// 示例 2：
//
//
//输入：n = 13
//输出：2
//解释：13 = 4 + 9
//
//
// 提示：
//
//
// 1 <= n <= 104
//
// Related Topics 广度优先搜索 数学 动态规划
// 👍 830 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func numSquares(n int) int {

	// [2]{x,y} x:当前的节点 y:走的步数
	queue := make([][2]int, 1)
	queue[0] = [2]int{n, 0}

	visited := make([]int, n+1)
	visited[n] = 1

	for len(queue) > 0 {
		num := queue[0][0]
		step := queue[0][1]
		queue = queue[1:]

		for i := 1; ; i++ {
			next := num - i*i
			if next < 0 {
				break
			}
			if next == 0 {
				return step + 1
			}
			if visited[next] == 1 {
				continue
			}
			queue = append(queue, [2]int{next, step + 1})
			visited[next] = 1
		}
	}
	return -1
}

//leetcode submit region end(Prohibit modification and deletion)
