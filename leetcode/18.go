package leetcode

import "sort"

// 给定一个包含 n 个整数的数组 nums 和一个目标值 target，判断 nums 中是否存在四个元素 a，b，c 和 d ，使得 a + b + c +
// d 的值与 target 相等？找出所有满足条件且不重复的四元组。
//
// 注意：答案中不可以包含重复的四元组。
//
//
//
// 示例 1：
//
//
// 输入：nums = [1,0,-1,0,-2,2], target = 0
// 输出：[[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]
//
//
// 示例 2：
//
//
// 输入：nums = [], target = 0
// 输出：[]
//
//
//
//
// 提示：
//
//
// 0 <= nums.length <= 200
// -109 <= nums[i] <= 109
// -109 <= target <= 109
//
// Related Topics 数组 哈希表 双指针
// 👍 788 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func fourSum(nums []int, target int) [][]int {
	if len(nums) < 4 {
		return [][]int{}
	}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	ret := make([][]int, 0)
	sum := 0
	for i := 0; i <= len(nums)-4; {
		for j := i + 1; j <= len(nums)-3; {
			l, r := j+1, len(nums)-1
			for l < r {
				sum = nums[i] + nums[j] + nums[l] + nums[r]
				if sum == target {
					item := []int{nums[i], nums[j], nums[l], nums[r]}
					ret = append(ret, item)
					lVal := nums[l]
					for l < len(nums) && nums[l] == lVal {
						l++
					}
					rVal := nums[j]
					for r > j && nums[r] == rVal {
						r--
					}
				} else if sum > target {
					r--
				} else {
					l++
				}
			}
			v := nums[j]
			for j < len(nums) && nums[j] == v {
				j++
			}
		}
		v := nums[i]
		for i < len(nums) && nums[i] == v {
			i++
		}
	}
	return ret
}

// leetcode submit region end(Prohibit modification and deletion)
