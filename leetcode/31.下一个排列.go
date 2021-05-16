package leetcode

// 时间: O(n)
// 空间: O(1)
func nextPermutation(nums []int) {

	// 从右往左找到一个升序对 nums[j-1], nums[j]
	// 此时 nums[j,) 是降序的
	n := len(nums)
	j := n - 1
	for j > 0 && (nums[j-1] >= nums[j]) {
		j--
	}
	// i:=j-1 当 i 有效时，在[j,)从右往左找第一个大于 nums[i]的数
	if j > 0 {
		k := n - 1
		for nums[k] <= nums[j-1] {
			k--
		}
		nums[j-1], nums[k] = nums[k], nums[j-1]
	}

	// 将 [j:）部分的翻转，形成升序
	k := n - 1
	for j < k {
		nums[j], nums[k] = nums[k], nums[j]
		j++
		k--
	}
}
