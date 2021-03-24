package leetcode

// 给定两个数组，编写一个函数来计算它们的交集。
// 两个数组的交集
func intersection(nums1 []int, nums2 []int) []int {
	existed := make(map[int]struct{}, len(nums1))
	for i := range nums1 {
		existed[nums1[i]] = struct{}{}
	}
	m := make(map[int]struct{}, len(nums2))
	ret := make([]int, 0, len(nums2))
	for i := range nums2 {
		if _, ok := existed[nums2[i]]; ok {
			if _, ok := m[nums2[i]]; !ok {
				ret = append(ret, nums2[i])
				m[nums2[i]] = struct{}{}
			}
		}
	}
	return ret
}
