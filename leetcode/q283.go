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
		for i < n {
			if nums[i] != 0 {
				i++
			}
		}
		for j < n {
			if nums[j] == 0 {
				j++
			}
		}
		if j > i && j < n {
			nums[i], nums[j] = nums[j], nums[i]
		} else if i >= j {
			j = i + 1
		}
	}
}

//leetcode submit region end(Prohibit modification and deletion)
