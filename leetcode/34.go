package leetcode

// 给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。
//
// 如果数组中不存在目标值 target，返回 [-1, -1]。
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/find-first-and-last-position-of-element-in-sorted-array
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// 递归
// 时间O(lgN)
// 空间O(1)
func searchRange(nums []int, target int) []int {
	l, r := seach(nums, 0, len(nums)-1, target)
	return []int{l, r}
}

func seach(nums []int, l, r, target int) (int, int) {
	if l > r {
		return -1, -1
	}
	mid := l + (r-l)/2
	if nums[mid] == target {
		lIdx, rIdx := mid, mid
		if isValidIndex(nums, mid-1) && nums[mid-1] == target {
			lIdx, _ = seach(nums, l, mid-1, target)
		}
		if isValidIndex(nums, mid+1) && nums[mid+1] == target {
			_, rIdx = seach(nums, mid+1, r, target)
		}
		return lIdx, rIdx
	} else if nums[mid] > target {
		return seach(nums, l, mid-1, target)
	} else {
		return seach(nums, mid+1, r, target)
	}
}

func isValidIndex(nums []int, idx int) bool {
	return idx >= 0 && idx < len(nums)
}

// 遍历法
// 时间：O(n)
// 空间O:(1)
// func searchRange(nums []int, target int) []int {
// 	l, r := 0, len(nums)-1
// 	ret := []int{-1,-1}
// 	for l < len(nums)&& nums[l] != target{
// 		l++
// 	}
// 	for r >= 0 && nums[r] != target{
// 		r--
// 	}
// 	if l < len(nums){
// 		ret[0] = l
// 	}
// 	if r >= 0{
// 		ret[1] = r
// 	}
// 	return ret
// }
