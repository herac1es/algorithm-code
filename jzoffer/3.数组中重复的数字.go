package jzoffer

// 时间: O(n)
// 空间: O(n)
func findRepeatNumber(nums []int) int {
	ht := make(map[int]struct{}, len(nums))
	for _, v := range nums {
		if _, ok := ht[v]; ok {
			return v
		}
		ht[v] = struct{}{}
	}
	return -1
}
