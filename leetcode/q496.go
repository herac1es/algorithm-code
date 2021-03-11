package leetcode

import "container/list"

//给你两个 没有重复元素 的数组 nums1 和 nums2 ，其中nums1 是 nums2 的子集。
//
// 请你找出 nums1 中每个元素在 nums2 中的下一个比其大的值。
//
// nums1 中数字 x 的下一个更大元素是指 x 在 nums2 中对应位置的右边的第一个比 x 大的元素。如果不存在，对应位置输出 -1 。
//
//
//
// 示例 1:
//
//
//输入: nums1 = [4,1,2], nums2 = [1,3,4,2].
//输出: [-1,3,-1]
//解释:
//    对于 num1 中的数字 4 ，你无法在第二个数组中找到下一个更大的数字，因此输出 -1 。
//    对于 num1 中的数字 1 ，第二个数组中数字1右边的下一个较大数字是 3 。
//    对于 num1 中的数字 2 ，第二个数组中没有下一个更大的数字，因此输出 -1 。
//
// 示例 2:
//
//
//输入: nums1 = [2,4], nums2 = [1,2,3,4].
//输出: [3,-1]
//解释:
//    对于 num1 中的数字 2 ，第二个数组中的下一个较大数字是 3 。
//    对于 num1 中的数字 4 ，第二个数组中没有下一个更大的数字，因此输出 -1 。
//
//
//
//
// 提示：
//
//
// 1 <= nums1.length <= nums2.length <= 1000
// 0 <= nums1[i], nums2[i] <= 104
// nums1和nums2中所有整数 互不相同
// nums1 中的所有整数同样出现在 nums2 中
//
//
//
//
// 进阶：你可以设计一个时间复杂度为 O(nums1.length + nums2.length) 的解决方案吗？
// Related Topics 栈
// 👍 381 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func NextGreaterElement(nums1 []int, nums2 []int) []int {
	stack := list.New()
	ng := make([]int, len(nums2))
	idxMap := make(map[int]int, len(nums1))
	for i := len(nums2) - 1; i >= 0; i-- {
		idxMap[nums2[i]] = i
		for stack.Len() != 0 && stack.Back().Value.(int) < nums2[i] {
			stack.Remove(stack.Back())
		}
		ng[i] = -1
		if stack.Len() > 0 {
			ng[i] = stack.Back().Value.(int)
		}
		stack.PushBack(nums2[i])
	}
	ret := make([]int, len(nums1))
	for i := range nums1 {
		ret[i] = ng[idxMap[nums1[i]]]
	}
	return ret
}

//leetcode submit region end(Prohibit modification and deletion)
