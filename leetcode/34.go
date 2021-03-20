package leetcode

// 给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。
//
// 如果数组中不存在目标值 target，返回 [-1, -1]。
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/find-first-and-last-position-of-element-in-sorted-array
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// 时间O(n)
// 空间O(1)
func searchRange(nums []int, target int) []int {
	l, r := 0, len(nums)-1
	ret := []int{-1, -1}
	for l < len(nums) && nums[l] != target {
		l++
	}
	for r >= 0 && nums[r] != target {
		r--
	}
	if l < len(nums) {
		ret[0] = l
	}
	if r >= 0 {
		ret[1] = r
	}
	return ret
}
