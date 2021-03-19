package leetcode

//给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
//
// 示例:
//
// 输入: [0,1,0,3,12]
//输出: [1,3,12,0,0]
//
// 说明:
//
//
// 必须在原数组上操作，不能拷贝额外的数组。
// 尽量减少操作次数。
//
// Related Topics 数组 双指针
// 👍 980 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func moveZeroes(nums []int) {
	n := len(nums)
	i, j := 0, 0
	for i < n && j < n {
		for i < n { // 从左往右找第一个为0的数
			if nums[i] != 0 {
				i++
			}
			break
		}
		for j < n { // 从左往右找第一个不为0的数
			if nums[j] == 0 {
				j++
			}
			break
		}
		if i < j && j < n { // 如果没有索引有效并且0在非0值前，则交换这两个数
			nums[i], nums[j] = nums[j], nums[i]
		} else if i > j { // 如果 0在非0之后，将非零指针直接移动到0指针之后
			j = i + 1
		}
	}
}

//leetcode submit region end(Prohibit modification and deletion)
