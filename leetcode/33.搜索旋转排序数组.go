package leetcode

func search(nums []int, target int) int {
	return binarySearchReverse(nums, 0, len(nums)-1, target)
}

// 时间: O(logn)
// 空间: O(logn)
func binarySearchReverse(nums []int, l, r, target int) int {
	if l > r {
		return -1
	}
	lo, hi := nums[l], nums[r]
	mid := l + (r-l)/2
	// 常规的升序数组
	if lo < hi {
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			return binarySearchReverse(nums, l, mid-1, target)
		} else {
			return binarySearchReverse(nums, mid+1, r, target)
		}
	}

	// 旋转数组
	if nums[mid] == target {
		return mid
	} else if (nums[mid] <= hi && target <= hi) || (nums[mid] > hi && target > hi) {
		// 在旋转的同一侧
		if nums[mid] > target {
			return binarySearchReverse(nums, l, mid-1, target)
		} else {
			return binarySearchReverse(nums, mid+1, r, target)
		}
	} else {
		// 在不同的侧
		if nums[mid] > target {
			return binarySearchReverse(nums, mid+1, r, target)
		} else {
			return binarySearchReverse(nums, l, mid-1, target)
		}
	}
}
