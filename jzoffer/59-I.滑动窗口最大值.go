package jzoffer

// 时间:(n)
// 空间:(k)
func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 {
		return nil
	}
	maxQueue := make([]int, 0, len(nums)/k)
	// 维护一个最长为k的滑动窗口
	l, r := 0, -1
	// 栈底的元素 nums[stack[0]] 是滑动窗口当前的最大值
	// 窗口右边纳入新元素，将栈顶小于等于此值的元素都出栈 （存索引）
	// 当窗口左边删除元素，检查是否是否为栈底元素，是的话，删除栈底元素
	maxStack := make([]int, 0, len(nums))
	for r < len(nums) && l < len(nums) {
		if r-l+1 == k {
			maxQueue = append(maxQueue, nums[maxStack[0]])
			if l == maxStack[0] {
				maxStack = maxStack[1:]
			}
			l++
		} else {
			r++
			for r < len(nums) && len(maxStack) > 0 && nums[maxStack[len(maxStack)-1]] <= nums[r] {
				maxStack = maxStack[:len(maxStack)-1]
			}
			maxStack = append(maxStack, r)
		}
	}
	return maxQueue
}
