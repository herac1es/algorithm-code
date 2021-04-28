package leetcode

//给你 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为 (i,
//ai) 和 (i, 0) 。找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。
//
// 说明：你不能倾斜容器。
//
//
//
// 示例 1：
//
//
//
//
//输入：[1,8,6,2,5,4,8,3,7]
//输出：49
//解释：图中垂直线代表输入数组 [1,8,6,2,5,4,8,3,7]。在此情况下，容器能够容纳水（表示为蓝色部分）的最大值为 49。
//
// 示例 2：
//
//
//输入：height = [1,1]
//输出：1
//
//
// 示例 3：
//
//
//输入：height = [4,3,2,1,4]
//输出：16
//
//
// 示例 4：
//
//
//输入：height = [1,2,1]
//输出：2
//
//
//
//
// 提示：
//
//
// n = height.length
// 2 <= n <= 3 * 104
// 0 <= height[i] <= 3 * 104
//
// Related Topics 数组 双指针
// 👍 2247 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
// 时间：O(n)
// 空间: O(1)
func maxArea(height []int) int {
	i, j := 0, len(height)-1
	volume := -1
	for i < j {
		volume = max(volume, getVolume(height, i, j))
		// 容器的体积由较低的侧高决定
		// 从最外侧向里收缩，容器的底一定变小，若收缩高侧，（上限被低测限制）容积一定是减少或者不变，而收缩低测才有可能容积变大，所以每次选择高度低的一侧向里收缩
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}
	return volume
}

// 计算容积
func getVolume(height []int, i, j int) int {
	return min(height[i], height[j]) * (j - i)
}

//leetcode submit region end(Prohibit modification and deletion)
