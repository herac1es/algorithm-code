package leetcode

//给你一个整数数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位
//。
//
// 返回滑动窗口中的最大值。
//
//
//
// 示例 1：
//
//
//输入：nums = [1,3,-1,-3,5,3,6,7], k = 3
//输出：[3,3,5,5,6,7]
//解释：
//滑动窗口的位置                最大值
//---------------               -----
//[1  3  -1] -3  5  3  6  7       3
// 1 [3  -1  -3] 5  3  6  7       3
// 1  3 [-1  -3  5] 3  6  7       5
// 1  3  -1 [-3  5  3] 6  7       5
// 1  3  -1  -3 [5  3  6] 7       6
// 1  3  -1  -3  5 [3  6  7]      7
//
//
// 示例 2：
//
//
//输入：nums = [1], k = 1
//输出：[1]
//
//
// 示例 3：
//
//
//输入：nums = [1,-1], k = 1
//输出：[1,-1]
//
//
// 示例 4：
//
//
//输入：nums = [9,11], k = 2
//输出：[11]
//
//
// 示例 5：
//
//
//输入：nums = [4,-2], k = 2
//输出：[4]
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 105
// -104 <= nums[i] <= 104
// 1 <= k <= nums.length
//
// Related Topics 堆 Sliding Window
// 👍 885 👎 0
// 时间:(n)
// 空间:(k)
func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 {
		return nil
	}
	maxQueue := make([]int, 0, len(nums)/k)
	l, r := 0, -1
	// 栈底的元素 nums[stack[0]] 是滑动窗口当前的最大值
	// 窗口右边纳入新元素，将栈顶小于此值的元素都出栈
	// 当窗口左边删除元素，检查大小是否等于栈底元素，是的话，删除栈底元素
	maxStack := make([]int, 0, len(nums))
	for r < len(nums) && l < len(nums) {
		if r-l+1 == k {
			maxQueue = append(maxQueue, maxStack[0])
			if nums[l] == maxStack[0] {
				maxStack = maxStack[1:]
			}
			l++
		} else {
			r++
			if r < len(nums) {
				for len(maxStack) > 0 && maxStack[len(maxStack)-1] < nums[r] {
					maxStack = maxStack[:len(maxStack)-1]
				}
				maxStack = append(maxStack, nums[r])
			}
		}
	}
	return maxQueue
}
