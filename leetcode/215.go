package leetcode

//在未排序的数组中找到第 k 个最大的元素。请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。
//
// 示例 1:
//
// 输入: [3,2,1,5,6,4] 和 k = 2
//输出: 5
//
//
// 示例 2:
//
// 输入: [3,2,3,1,2,4,5,5,6] 和 k = 4
//输出: 4
//
// 说明:
//
// 你可以假设 k 总是有效的，且 1 ≤ k ≤ 数组的长度。
// Related Topics 堆 分治算法
// 👍 1044 👎 0

// 堆排序法：
// 时间：O(nlogn)
// 空间: O(1)
func findKthLargest(nums []int, k int) int {
	for i := len(nums); i >= 1; i-- {
		sortSink(nums, len(nums), i)
	}
	i := 0
	size := len(nums)
	for {
		v := nums[0]
		nums[0], nums[size-1] = nums[size-1], nums[0]
		size--
		sortSink(nums, size, 1)
		i++
		if i == k {
			return v
		}
	}
}

func sortSink(nums []int, size int, idx int) {
	for idx*2 <= size {
		k := idx * 2
		if k+1 <= size && less(nums, k-1, k) {
			k = k + 1
		}
		if !less(nums, idx-1, k-1) {
			return
		}
		nums[idx-1], nums[k-1] = nums[k-1], nums[idx-1]
		idx = k
	}
}

func less(nums []int, i, j int) bool {
	return nums[i] < nums[j]
}

// 排序法：
// 时间：O(nlogn)
// 空间: O(1)
// func findKthLargest(nums []int, k int) int {
// 	sort.Slice(nums, func(i, j int) bool {
// 		return nums[i] > nums[j]
// 	})
// 	return nums[k-1]
// }
