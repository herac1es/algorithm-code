package leetcode

// 给定一个含有 n 个正整数的数组和一个正整数 target 。
//
// 找出该数组中满足其和 ≥ target 的长度最小的 连续子数组 [numsl, numsl+1, ..., numsr-1, numsr] ，并返回其长度。如果不存在符合条件的子数组，返回 0 。
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/minimum-size-subarray-sum
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func minSubArrayLen(target int, nums []int) int {
	ret := len(nums) + 1
	l, r := 0, -1
	sum := 0
	for l < len(nums) {
		if r+1 < len(nums) && sum < target {
			r++
			sum += nums[r]
		} else {
			sum -= nums[l]
			l++
		}
		if sum >= target && r-l+1 < ret {
			ret = r - l + 1
		}
	}
	if ret > len(nums) {
		ret = 0
	}
	return ret
}
