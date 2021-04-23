package classic

// 快速排序
// 时间: O(nlogN)
// 空间: O(1)
func QuickSort(nums []int) []int {
	// quickSort(nums, 0, len(nums)-1)
	quickSort(nums, 0, len(nums)-1)
	return nums
}

func quickSort(nums []int, l, r int) {
	if l >= r {
		return
	}
	j := partition(nums, l, r)
	quickSort(nums, l, j-1)
	quickSort(nums, j+1, r)
}

func partition(nums []int, l, r int) int {
	i := l + 1
	j := r
	mid := l + (r-l)/2
	nums[l], nums[mid] = nums[mid], nums[l]
	v := nums[l]
	for {
		for i < r && nums[i] < v {
			i++
		}
		for nums[j] > v {
			j--
		}
		if i >= j {
			break
		}
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
	nums[l], nums[j] = nums[j], nums[l]
	return j
}

// 三路快排
func thirdPathQuickSort(nums []int, l, r int) {
	if l >= r {
		return
	}
	lt, i, rt := l-1, l, r+1 // [l:lt]都小于v，[rt:r]都大于v
	v := nums[l]
	for i <= rt && i <= r {
		cmp := nums[i] - v
		if cmp > 0 {
			rt--
			nums[i], nums[rt] = nums[rt], nums[i]
		} else if cmp < 0 {
			lt++
			nums[i], nums[lt] = nums[lt], nums[i]
			i++
		} else {
			i++
		}
	}
	thirdPathQuickSort(nums, l, lt)
	thirdPathQuickSort(nums, rt, r)
}
