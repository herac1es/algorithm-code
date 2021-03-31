package leetcode

//在整数数组 nums 中，是否存在两个下标 i 和 j，使得 nums [i] 和 nums [j] 的差的绝对值小于等于 t ，且满足 i 和 j 的差的
//绝对值也小于等于 ķ 。
//
// 如果存在则返回 true，不存在返回 false。
//
//
//
// 示例 1:
//
// 输入: nums = [1,2,3,1], k = 3, t = 0
//输出: true
//
// 示例 2:
//
// 输入: nums = [1,0,1,1], k = 1, t = 2
//输出: true
//
// 示例 3:
//
// 输入: nums = [1,5,9,1,5,9], k = 2, t = 3
//输出: false
// Related Topics 排序 Ordered Map
// 👍 308 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	mark := make(map[int64]int64, len(nums))
	l, r := 0, 0
	bkt := int64(0)
	for r < len(nums) {
		bkt = getBucket(nums[r], t)
		if _, ok := mark[bkt]; ok {
			return true
		} else if left, ok := mark[bkt-1]; ok && int64(nums[r])-left <= int64(t) {
			return true
		} else if right, ok := mark[bkt+1]; ok && right-int64(nums[r]) <= int64(t) {
			return true
		} else {
			mark[bkt] = int64(nums[r])
		}
		mark[bkt] = int64(nums[r])
		r++
		if len(mark) == k+1 {
			delete(mark, getBucket(nums[l], t))
			l++
		}
	}
	return false
}

func getBucket(num int, t int) int64 {
	if num < 0 {
		return int64(num/(t+1)) - 1
	}
	return int64(num / (t + 1))
}

//leetcode submit region end(Prohibit modification and deletion)
