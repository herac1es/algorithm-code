package leetcode

// 给定一个包含红色、白色和蓝色，一共 n 个元素的数组，原地对它们进行排序，使得相同颜色的元素相邻，并按照红色、白色、蓝色顺序排列。
//
// 此题中，我们使用整数 0、 1 和 2 分别表示红色、白色和蓝色。
//
//
//
//
//
//
// 示例 1：
//
//
// 输入：nums = [2,0,2,1,1,0]
// 输出：[0,0,1,1,2,2]
//
//
// 示例 2：
//
//
// 输入：nums = [2,0,1]
// 输出：[0,1,2]
//
//
// 示例 3：
//
//
// 输入：nums = [0]
// 输出：[0]
//
//
// 示例 4：
//
//
// 输入： nums = [1]
// 输出: [1]
//
//
//
//
// 提示：
//
//
// n == nums.length
// 1 <= n <= 300
// nums[i] 为 0、1 或 2
//
//
//
//
// 进阶：
//
//
// 你可以不使用代码库中的排序函数来解决这道题吗？
// 你能想出一个仅使用常数空间的一趟扫描算法吗？
//
// Related Topics 排序 数组 双指针
// 👍 822 👎 0

// leetcode submit region begin(Prohibit modification and deletion)

// Solution 2
// 三指针法:
// 指针zero:代表[0:zero]区间的数全为0
// 指针cursor: 代表未排好序区间的当前游标,(zero:cursor)都是1
// 指针second: 代表[second:len(nums)-1]区间的数全为2
// 核心思路：假设cursor左边已经是全部为0，cursor右边已经全部为2,现在cursor需要走完[zero+1:second-1]区间
// base case： zero=-1,cursor=0,second=len(nums)
func sortColors(nums []int) {
	zero, cursor, second := -1, 0, len(nums)
	for cursor < second {
		if nums[cursor] == 0 {
			nums[cursor], nums[zero+1] = nums[zero+1], nums[cursor]
			zero++
			cursor++
		} else if nums[cursor] == 1 {
			cursor++
		} else if nums[cursor] == 2 {
			nums[cursor], nums[second-1] = nums[second-1], nums[cursor]
			second--
		} else {
			panic("input is not valid")
		}
	}
}

// Solution 1
// 使用一个额外的数组记录0，1，2分别出现的次数，然后依次将响应个数的0，1，2放回原数组
// 时间复杂度: O(n)
// 空间复杂度: O(1)
// func sortColors(nums []int) {
// 	count := [3]int{}
// 	for i := 0; i < len(nums); i++ {
// 		// 防止不合法输入导致数组越界
// 		if nums[i] < 0 || nums[i] > 2 {
// 			return
// 		}
// 		count[nums[i]]++
// 	}
// 	idx := 0
// 	for i := 0; i < len(count) && idx < len(nums); i++ {
// 		for j := 0; j < count[i]; j++ {
// 			nums[idx] = i
// 			idx++
// 		}
// 	}
// }

// leetcode submit region end(Prohibit modification and deletion)
