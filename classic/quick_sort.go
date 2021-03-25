package classic

// 快速排序
// 时间: O(nlogN)
// 空间: O(1)
func QuickSort(nums []int) []int {
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
	v := nums[l]
	for {
		for nums[i] < v {
			i++
			if i == r {
				break
			}
		}
		for nums[j] > v {
			j--
			if j == l {
				break
			}
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
