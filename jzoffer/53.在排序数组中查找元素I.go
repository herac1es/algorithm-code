package jzoffer

func search(nums []int, target int) int {

	// 右边界
	right := -1
	l, r := 0, len(nums)-1
	mid := 0
	for l <= r {
		mid = l + (r-l)/2
		if nums[mid] == target {
			if mid+1 <= r && nums[mid+1] == target {
				l = mid + 1
			} else {
				right = mid
				break
			}
		} else if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	if right == -1 {
		return 0
	}
	// 左边界
	l, r = 0, right
	left := right
	for l <= r {

		mid = l + (r-l)/2
		if nums[mid] == target {
			if mid-1 >= l && nums[mid-1] == target {
				r = mid - 1
			} else {
				left = mid
				break
			}
		} else if nums[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return right - left + 1
}

// 二分
// 时间: O(logn)
// 空间: O(k)
// func search(nums []int, target int) int {
// 	cnt := 0
// 	binaryCount(nums, target, 0, len(nums)-1, &cnt)
// 	return cnt
// }
//
// func binaryCount(nums []int, target, l, r int, ret *int) {
// 	if l > r {
// 		return
// 	}
// 	mid := l + (r-l)/2
// 	if nums[mid] == target {
// 		*ret++
// 		if mid-1 >= l && nums[mid-1] == target {
// 			binaryCount(nums, target, l, mid-1, ret)
// 		}
// 		if mid+1 <= r && nums[mid+1] == target {
// 			binaryCount(nums, target, mid+1, r, ret)
// 		}
//
// 	} else if nums[mid] > target {
// 		binaryCount(nums, target, l, mid-1, ret)
// 	} else {
// 		binaryCount(nums, target, mid+1, r, ret)
// 	}
// }

// func search(nums []int, target int) int {
//     left, right := findBorder(nums, target, 0, len(nums)-1)
//     if left == -1 && right == -1 {
//         return 0
//     }
//     return right-left+1
// }

// func findBorder(nums []int, target, l, r int) (int, int) {

//     if l > r {
//         return -1, -1
//     }

//     mid := l+(r-l)/2
//     if nums[mid] < target {
//         return findBorder(nums, target, mid+1, r)
//     } else if nums[mid] > target {
//         return findBorder(nums, target, l, mid-1)
//     } else {
//         left, right := mid, mid
//         if mid-1 >= l && nums[mid-1] == target {
//             left, _ = findBorder(nums, target, l, mid-1)
//         }
//         if mid+1 <= r && nums[mid+1] == target {
//             _, right = findBorder(nums, target, mid+1, r)
//         }
//         return left, right
//     }
// }
