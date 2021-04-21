package leetcode

// 最短无序子数组
// 时间：O(n)
// 空间：O(1)
func findUnsortedSubarray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	minv, maxv := int(1e5+1), int(-1e5-1)
	needContinue := false
	for i := 0; i < len(nums)-1; i++ {
		// 找到逆序对
		if nums[i] > nums[i+1] {
			needContinue = true
			minv = min(nums[i+1], minv)
			maxv = max(nums[i], maxv)
		}
	}
	if !needContinue {
		return 0
	}
	l, r := 0, -1
	ldone, rdone := 0, 0
	for i := 0; i < len(nums); i++ {
		if nums[i] > minv && ldone != 1 {
			l = i
			ldone = 1
		}
		tail := len(nums) - 1 - i
		if nums[tail] < maxv && rdone != 1 {
			r = tail
			rdone = 1
		}
		if ldone == 1 && rdone == 1 {
			break
		}
	}
	if r-l < 0 {
		return 0
	}
	return r - l + 1
}
