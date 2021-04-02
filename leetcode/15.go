package leetcode

import "sort"

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
func ThreeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return [][]int{}
	}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	ret := make([][]int, 0)
	for idx := 0; idx <= len(nums)-3; {
		i, j := idx+1, len(nums)-1
		target := nums[idx]
		for i < j {
			sum := target + nums[i] + nums[j]
			if sum == 0 {
				ret = append(ret, []int{nums[idx], nums[i], nums[j]})
				t := nums[i]
				for i < len(nums) && nums[i] == t {
					i++
				}
				j = len(nums) - 1
			} else if sum > 0 {
				t := nums[j]
				for j > i && nums[j] == t {
					j--
				}
			} else if sum < 0 {
				t := nums[i]
				for i < j && nums[i] == t {
					i++
				}
			}
		}
		for idx < len(nums) && nums[idx] == target {
			idx++
		}
	}
	return ret
}

//leetcode submit region end(Prohibit modification and deletion)
