package leetcode

import (
	"github.com/herac1es/algorithm-code/classic"
)

//给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有和为 0 且不重
//复的三元组。
//
// 注意：答案中不可以包含重复的三元组。
//
//
//
// 示例 1：
//
//
//输入：nums = [-1,0,1,2,-1,-4]
//输出：[[-1,-1,2],[-1,0,1]]
//
//
// 示例 2：
//
//
//输入：nums = []
//输出：[]
//
//
// 示例 3：
//
//
//输入：nums = [0]
//输出：[]
//
//
//
//
// 提示：
//
//
// 0 <= nums.length <= 3000
// -105 <= nums[i] <= 105
//
// Related Topics 数组 双指针
// 👍 3058 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
// 时间: O(n2)
// 空间: O(1)
func threeSum(nums []int) [][]int {
	classic.QuickSort(nums)
	ret := make([][]int, 0)

	for i := 0; i < len(nums); {
		l, r := i+1, len(nums)-1
		for l < r {
			sum := nums[l] + nums[r] + nums[i]
			if sum == 0 {
				ret = append(ret, []int{nums[i], nums[l], nums[r]})
			}
			// 右指针左移
			rv := nums[r]
			if sum == 0 || sum > 0 {
				for r > l && nums[r] == rv {
					r--
				}
			}
			// 左指针右移
			lv := nums[l]
			if sum == 0 || sum < 0 {
				for l < r && nums[l] == lv {
					l++
				}
			}
		}

		iv := nums[i]
		for i < len(nums) && nums[i] == iv {
			i++
		}
	}
	return ret
}

//leetcode submit region end(Prohibit modification and deletion)
