package leetcode

// 347.前K个高频元素
// 时间复杂度 O(n)
// 空间：O(n)
func topKFrequent(nums []int, k int) []int {
	cnt := make(map[int]int, len(nums))
	for _, v := range nums {
		cnt[v]++
	}
	bucket := make([][]int, len(nums)+1)
	for k, v := range cnt {
		bucket[v] = append(bucket[v], k)
	}
	ret := make([]int, 0, k)
	for i := len(nums); i >= 0; i-- {
		if len(ret) == k {
			break
		}
		if bucket[i] != nil {
			ret = append(ret, bucket[i]...)
		}
	}
	return ret
}
