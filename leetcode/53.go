package leetcode

//给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
//
//
//
// 示例 1：
//
//
//输入：nums = [-2,1,-3,4,-1,2,1,-5,4]
//输出：6
//解释：连续子数组 [4,-1,2,1] 的和最大，为 6 。
//
//
// 示例 2：
//
//
//输入：nums = [1]
//输出：1
//
//
// 示例 3：
//
//
//输入：nums = [0]
//输出：0
//
//
// 示例 4：
//
//
//输入：nums = [-1]
//输出：-1
//
//
// 示例 5：
//
//
//输入：nums = [-100000]
//输出：-100000
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 3 * 104
// -105 <= nums[i] <= 105
//
//
//
//
// 进阶：如果你已经实现复杂度为 O(n) 的解法，尝试使用更为精妙的 分治法 求解。
// Related Topics 数组 分治算法 动态规划
// 👍 3143 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
// 动态规划：dp[i] 以nums[i-1]结尾的子数组最大和
// 时间：O(n)
// 空间：O(n)
func maxSubArrayDP(nums []int) int {
	dp := make([]int, len(nums)+1)
	ret := 0
	if len(nums) > 0 {
		ret = nums[0]
	}
	for i := 1; i <= len(nums); i++ {
		if dp[i-1]+nums[i-1] < nums[i-1] {
			dp[i] = nums[i-1]
		} else {
			dp[i] = dp[i-1] + nums[i-1]
		}
		if dp[i] > ret {
			ret = dp[i]
		}
	}
	return ret
}

//leetcode submit region end(Prohibit modification and deletion)

type Sum struct {
	LSum   int // 以最后一个元素结尾的最长连续子数组和
	RSum   int // 以最一个元素开头的最长连续子数组和
	Sum    int // 数组和
	MaxSum int // 最大连续子数组和
}

// 二分法
// 时间 ：O(n)
// 空间：O（logN）递归栈的空间（递归logN次）
func maxSubArray(nums []int) int {
	return maxSubArrayRecur(nums, 0, len(nums)-1).MaxSum
}

func maxSubArrayRecur(nums []int, l, r int) Sum {
	if l == r {
		return Sum{
			LSum:   nums[l],
			RSum:   nums[l],
			MaxSum: nums[l],
			Sum:    nums[l],
		}
	}

	mid := l + (r-l)/2
	ls := maxSubArrayRecur(nums, l, mid)
	rs := maxSubArrayRecur(nums, mid+1, r)
	sum := Sum{
		LSum:   max(ls.LSum, ls.Sum+rs.LSum),
		RSum:   max(rs.RSum, ls.RSum+rs.Sum),
		Sum:    ls.Sum + rs.Sum,
		MaxSum: max(max(ls.MaxSum, rs.MaxSum), ls.RSum+rs.LSum),
	}
	return sum
}
