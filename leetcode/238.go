package leetcode

// 238.除自身以外数的乘积
// 进行两轮遍历，第一轮遍历计算i左边的元素乘积，第二轮计算i右边的元素乘积
// 时间: O(n)
// 空间: O(1)  结果数组的额外空间不算
func productExceptSelf(nums []int) []int {
	ret := make([]int, len(nums))
	k := 1 // 记录累乘
	// 计算nums[i] 左边数的乘积
	for i := 0; i < len(nums); i++ {
		ret[i] = k // 在乘上自己之前先赋值，所以是不包含自己的左边元素乘积
		k *= nums[i]
	}
	k = 1
	// 一边计算nums[i] 右边数的乘积 ，同时与已计算出的左乘积相乘，则得出结果
	for i := len(nums) - 1; i >= 0; i-- {
		ret[i] *= k
		k *= nums[i]
	}
	return ret
}
