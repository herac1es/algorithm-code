package classic

// 归并排序
// 时间：O(nlgN)
// 空间：O(n)

func MergeSort(nums []int) []int {
	// mergeSort(nums, 0, len(nums)-1)
	mergeSortII(nums, 0, len(nums)-1)
	return nums
}

// 自底向上
func mergeSort(nums []int) {
	n := len(nums)
	for size := 1; size < n; size *= 2 {
		for i := 0; i < n-size; i += size * 2 {
			merge(nums, i, i+size-1, min(i+size*2-1, n-1))
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 自顶向下
func mergeSortII(nums []int, l, r int) {
	if l >= r {
		return
	}
	mid := l + (r-l)/2
	mergeSortII(nums, l, mid)
	mergeSortII(nums, mid+1, r)
	merge(nums, l, mid+1, r)
}

func merge(nums []int, l, mid, r int) {
	i, j := l, mid // 对[l:mid-1]和[mid:r]归并
	aux := make([]int, 0, r-l+1)
	for i < mid || j <= r {
		if i >= mid {
			aux = append(aux, nums[j])
			j++
		} else if j > r {
			aux = append(aux, nums[i])
			i++
		} else if nums[i] <= nums[j] {
			aux = append(aux, nums[i])
			i++
		} else {
			aux = append(aux, nums[j])
			j++
		}
	}
	i = l
	for _, v := range aux {
		nums[i] = v
		i++
	}
}
