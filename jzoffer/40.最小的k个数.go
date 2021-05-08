package jzoffer

// 时间：O(n)
// 空间: O(logn)
func getLeastNumbers(arr []int, k int) []int {
	quickSortSentinel(arr, 0, len(arr)-1, k)
	return arr[:k]
}

func quickSortSentinel(nums []int, l, r int, k int) {
	if l >= r {
		return
	}
	j := partitionSentinel(nums, l, r)
	if j == k-1 {
		return
	} else if j < k-1 {
		quickSortSentinel(nums, j+1, r, k)
	} else {
		quickSortSentinel(nums, l, j-1, k)
	}
}

func partitionSentinel(nums []int, l, r int) int {
	i := l + 1
	j := r
	mid := l + (r-l)/2
	nums[mid], nums[l] = nums[l], nums[mid]
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
