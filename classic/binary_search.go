package classic

// 二分查找
// 给定一个有序递增数组，找出数组中目标值为target的元素索引，若不存在则返回-1
// func BinarySearch(nums []int, target int) int {
// 	// 定义操作区间为[l:r]
// 	l, r := 0, len(nums)-1
// 	for l <= r {
// 		mid := l + (r-l)/2
// 		if nums[mid] == target {
// 			return mid
// 		} else if nums[mid] < target {
// 			l = mid + 1
// 		} else {
// 			r = mid - 1
// 		}
// 	}
// 	return -1
// }

// 时间：O(lgN)
// 空间：O(1)
func BinarySearch(nums []int, target int) int {
	// 定义操作区间为[l:r)
	l, r := 0, len(nums)
	mid := 0
	for l < r {
		mid = l + (r-l)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return -1
}
