package jzoffer

// 对于长度为n的nums数组，元素取值范围是 0~n
// 从缺少元素的位置开始往右，索引值都小于元素值
// 所以找第一个索引值小于元素值的位置，索引值就是缺失元素
// 时间：O(logn)
// 空间: O(1)
func missingNumber(nums []int) int {
	l, r := 0, len(nums)-1
	mid := 0
	for l <= r {
		mid = l + (r-l)/2
		if nums[mid] == mid {
			l = mid + 1
		} else if mid-1 >= l && nums[mid-1] > mid-1 {
			r = mid - 1
		} else {
			return mid
		}
	}
	if l == len(nums) {
		return l
	}
	return -1
}
